# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-481`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610191305159466() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-130515-9466\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610191305159466() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier changes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token).

The work spans two repositories with a hard ordering dependency. The new greeting function lives in `handarbeit/fabrik-test-delta` (the "beta" repo). The caller lives in `handarbeit/fabrik-test-gamma` (this repo, the "alpha" repo). Alpha cannot compile against a function that does not yet exist in beta, so beta's change must land first.

The scenario is meant to prove that Plan decomposes the work by spawning a child issue in the beta repo, that the child lands first, and that the parent then completes using the new function. The function name and string literal embed a per-run-unique token. The scenario's final assertion checks that exact text in the beta repo, so the values must be used verbatim and never substituted.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the uniquely-tokened greeting function (Priority: P1)

A sub-issue in the beta repo (`handarbeit/fabrik-test-delta`) adds the exported greeting function and a test for it. This is the foundation that the alpha change depends on.

**Why this priority**: The alpha change cannot compile without it, and the scenario's final assertion checks this exact text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes and asserts the exact literal string.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610191305159466()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-130515-9466`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Alpha repo calls the new function and prints the result (Priority: P1)

After the beta change has landed, this repo's `main.go` calls the new greeting function in addition to its existing behaviour and prints the returned value.

**Why this priority**: It completes the cross-repo change. It demonstrates that the ordering constraint was respected: alpha builds only once beta has the function.

**Independent Test**: Build and run the alpha program with the beta dependency updated to include the new function. The output contains `e2e-cross-repo-spawn-20261019-130515-9466`, and the existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta repo's new function is available to alpha, **When** alpha is built, **Then** the build succeeds (running `go mod tidy` if needed).
2. **Given** a built alpha program, **When** it runs, **Then** it prints the string returned by `greeting.HelloE2E202610191305159466()` in addition to its existing output.

---

### User Story 3 - Plan decomposes the work via a cross-repo spawn (Priority: P1)

Plan recognises that the work needs ordered changes in two repos. It emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo, so the beta change becomes a child issue that blocks this one.

**Why this priority**: This is the behaviour under test (#803). Without a spawned child in the beta repo, the scenario does not exercise cross-repo decomposition.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta`. The parent is blocked on it, and the parent does not proceed to implement alpha until the child is done.

**Acceptance Scenarios**:

1. **Given** the spec above, **When** Plan runs, **Then** its output contains a spawn block whose target repo is `handarbeit/fabrik-test-delta`.
2. **Given** the child was spawned, **When** the pipeline continues, **Then** the beta change lands before the alpha change is implemented.

---

### Edge Cases

- The function name or string literal is altered, shortened or regenerated by any stage. This breaks the final assertion, so the values must be used verbatim.
- The alpha change is implemented before the beta change has landed. The build fails because the symbol does not exist.
- The beta repo needs on-demand initialisation as a spawn target (#803). Spawning must still succeed.
- `go mod tidy` produces no changes. That is acceptable, and it is only run if needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610191305159466() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610191305159466()` MUST return exactly the literal string `e2e-cross-repo-spawn-20261019-130515-9466`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610191305159466` in `pkg/greeting/greeting_test.go` that passes.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610191305159466()` and print the result, in addition to its existing behaviour.
- **FR-005**: Module dependencies MUST be tidied with `go mod tidy` if needed so the alpha repo builds.
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`, so the beta change is a separate child issue.
- **FR-007**: The beta change MUST land before the alpha change is implemented or merged.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this spec, and no stage may substitute its own values.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting `handarbeit/fabrik-test-delta` exists as a result of Plan's spawn block, and the parent is blocked on it until the child completes.
- **SC-002**: The beta repo's default branch contains `HelloE2E202610191305159466` returning `e2e-cross-repo-spawn-20261019-130515-9466`, plus a passing test for it.
- **SC-003**: The alpha repo builds successfully and its program output includes `e2e-cross-repo-spawn-20261019-130515-9466`.
- **SC-004**: The scenario's final assertion finds the exact token text in the beta repo.

## Assumptions

- `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go` already exist in the beta repo, and the alpha repo already imports the beta greeting package. This spec adds one function and its use, not the package itself.
- `main.go` has existing output that must be kept unchanged.
- The beta change reaches alpha through the normal Go module dependency mechanism (for example, an updated dependency pin), which is why `go mod tidy` may be needed.
- The `fabrik:yolo` label means stages auto-advance, so no human approval gate is expected between stages.

## Out of Scope

- Any greeting-related changes beyond the single new function, its test, and its call and print in `main.go`.
- Changing the spawn mechanism, on-demand spawn-target repo init (#803) or the fixture-token mechanism (#1308) themselves. This issue only tests them.
- Any additional repos beyond alpha (this repo) and beta (`handarbeit/fabrik-test-delta`).

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
