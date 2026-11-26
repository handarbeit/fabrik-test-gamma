# Feature Specification: Cross-repo spawn e2e scenario (beta lands before alpha)

**Feature Branch**: `fabrik/issue-2578`
**Created**: 2026-11-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611260445532995() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261126-044553-2995\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611260445532995() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. Some work spans two repositories with a strict ordering dependency: the consuming repo (this one, "alpha", `handarbeit/fabrik-test-gamma`) cannot import a function until the providing repo ("beta", `handarbeit/fabrik-test-delta`) has it. The pipeline's Plan stage is expected to recognise this and decompose the work by spawning a child issue in the beta repo, rather than attempting everything in one repo.

The scenario guards two earlier fixes:
- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: per-run-unique fixture tokens, so the final assertion cannot be satisfied by stale content from a previous run.

The function name `HelloE2E202611260445532995` and the string `e2e-cross-repo-spawn-20261126-044553-2995` carry that per-run-unique token. They must be used exactly as given, because the scenario's final assertion looks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo (`handarbeit/fabrik-test-delta`) exposes a new exported greeting function that returns the per-run-unique literal, with a test covering it. This work is delivered through a child issue spawned from this issue's Plan stage.

**Why this priority**: The alpha change depends on it, and the scenario's final assertion checks for this text in the beta repo.

**Independent Test**: In the beta repo, run the tests for `pkg/greeting` and confirm the new test passes. Confirm `HelloE2E202611260445532995()` returns exactly `e2e-cross-repo-spawn-20261126-044553-2995`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting/greeting.go`, **When** the child issue's work is merged, **Then** it contains an exported `HelloE2E202611260445532995() string` returning `"e2e-cross-repo-spawn-20261126-044553-2995"`.
2. **Given** the beta repo's `pkg/greeting/greeting_test.go`, **When** the child issue's work is merged, **Then** it contains a test for that function and the test passes.

---

### User Story 2 - Alpha repo consumes the new function (Priority: P1)

This repo's `main.go` calls `greeting.HelloE2E202611260445532995()` from the beta repo, in addition to its existing behaviour, and prints the result. This lands only after the beta change has landed.

**Why this priority**: It is the end-to-end proof that cross-repo ordering held. Alpha cannot compile against beta until the function exists.

**Independent Test**: Build and run the alpha program and confirm that its output includes `e2e-cross-repo-spawn-20261126-044553-2995` along with its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha change is implemented, **Then** `main.go` calls `greeting.HelloE2E202611260445532995()` and prints the returned value.
2. **Given** the alpha change, **When** the module is built, **Then** it builds without error, with `go mod tidy` run if needed.
3. **Given** the alpha change, **When** the program runs, **Then** its pre-existing behaviour is unchanged.

---

### User Story 3 - Plan decomposes across repos (Priority: P1)

The Plan stage for this issue decomposes the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, so that the beta work is a separate issue that is ordered before the alpha work.

**Why this priority**: Exercising the decomposition path (including #803's on-demand init of the spawn target) is the purpose of the scenario.

**Independent Test**: Inspect the Plan output and the board: a child issue exists in `handarbeit/fabrik-test-delta`, and the alpha issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** the spec's cross-repo constraint, **When** Plan runs, **Then** its output contains a spawn block targeting `handarbeit/fabrik-test-delta` for the beta work.
2. **Given** the spawned child, **When** the alpha issue proceeds, **Then** it does not implement the alpha change until the beta child has landed.

---

### Edge Cases

- The function name and string literal must not be varied or re-derived. Any difference defeats the final assertion.
- The beta repo may not yet be initialised as a spawn target (#803). Spawning must still succeed.
- A previous run's artifacts in the beta repo must not satisfy this run's assertion, since the token is unique per run (#1308).
- Existing content in `main.go` and in `pkg/greeting` must be preserved. The change is additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST contain an exported function `HelloE2E202611260445532995() string` in `pkg/greeting/greeting.go` that returns exactly `"e2e-cross-repo-spawn-20261126-044553-2995"`.
- **FR-002**: The beta repo MUST contain a test for that function in `pkg/greeting/greeting_test.go`, and the test MUST pass.
- **FR-003**: This repo's `main.go` MUST call `greeting.HelloE2E202611260445532995()` in addition to its existing behaviour and print the returned value.
- **FR-004**: This repo MUST build after the change, with `go mod tidy` run if needed.
- **FR-005**: The beta change MUST land before the alpha change that depends on it.
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo.
- **FR-007**: The exact function name and string literal given in this spec MUST be used verbatim. They MUST NOT be replaced with invented names or values.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the text `e2e-cross-repo-spawn-20261126-044553-2995` and `HelloE2E202611260445532995`, and its test for the function passes.
- **SC-002**: A child issue exists in the beta repo that was spawned from this issue, and it merged before this issue's alpha change.
- **SC-003**: Running the alpha program prints `e2e-cross-repo-spawn-20261126-044553-2995`.
- **SC-004**: The alpha module builds cleanly with no unresolved imports.

## Assumptions

- The beta repo's `pkg/greeting` package already exists, or the child work creates it. The spec asks only that the function and its test exist there.
- The alpha module can import the beta module's `pkg/greeting` package. Establishing that dependency is part of the alpha change (this is why `go mod tidy` is mentioned).
- "In addition to whatever it already does" means the alpha change is additive and removes no existing behaviour.
- The `fabrik:yolo` label means the stages advance without waiting for human approval.

## Out of Scope *(optional)*

- Any change in either repo beyond the function, its test, and the alpha call-and-print.
- Changes to the Fabrik engine's spawn or decomposition behaviour. This issue only exercises it.
- Cleanup of fixtures from earlier runs.

## Source References *(optional)*

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
