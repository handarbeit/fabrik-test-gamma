# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-26`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610141933573941() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261014-193357-3941".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610141933573941() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work is deliberately split across two repositories with a hard ordering dependency: the greeting function must exist in `handarbeit/fabrik-test-delta` (the "beta" repo) before `handarbeit/fabrik-test-gamma` (this repo, the "alpha" repo in the issue text) can call it. A single-repo plan cannot satisfy that, so the pipeline must decompose the work by spawning a sub-issue in the beta repo, which then lands first.

The scenario guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token). The function name `HelloE2E202610141933573941` and the string `e2e-cross-repo-spawn-20261014-193357-3941` are that per-run-unique token. The scenario's final assertion checks that this exact text landed in the beta repo, so it must be used verbatim and never substituted.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the greeting function (Priority: P1)

The beta repo (`handarbeit/fabrik-test-delta`) gains an exported greeting function that returns the per-run-unique string, with a test. This is delivered through a spawned sub-issue, not directly from this issue.

**Why this priority**: The alpha-side change cannot compile without it, and the scenario's final assertion checks for it in the beta repo.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261014-193357-3941`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting` package, **When** `HelloE2E202610141933573941()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261014-193357-3941`.
2. **Given** the beta repo's test suite, **When** the tests run, **Then** a test in `pkg/greeting/greeting_test.go` covers the new function and passes.

---

### User Story 2 - Alpha repo calls and prints the beta function's result (Priority: P2)

This repo's `main.go` calls `greeting.HelloE2E202610141933573941()` from the beta repo, in addition to its existing behaviour, and prints the result. This lands only after User Story 1 has landed.

**Why this priority**: It depends on User Story 1 and completes the end-to-end flow.

**Independent Test**: With the beta change available, build and run this repo's program. It still does what it did before and also prints `e2e-cross-repo-spawn-20261014-193357-3941`.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** this repo's program runs, **Then** its prior output is unchanged and the returned string is also printed.
2. **Given** the dependency on the beta repo, **When** `go mod tidy` is run if needed, **Then** the module files are consistent and the build succeeds.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

During Plan, the pipeline recognises that the work spans two repos in strict order and emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo. The beta sub-issue is created and blocks this issue until it resolves.

**Why this priority**: This is the behaviour under test (#803, #1308). Without the spawn, the scenario has no value.

**Independent Test**: After Plan, a sub-issue exists in `handarbeit/fabrik-test-delta`, linked as a blocker of this issue, carrying the function name and string verbatim.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output contains a spawn block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawn block, **When** the child is created, **Then** the child issue body states the exact function name and exact string literal given above.
3. **Given** the child is unresolved, **When** the pipeline continues, **Then** this issue does not implement the `main.go` change until the child has landed.

---

### Edge Cases

- The spawn-target repo may not yet be initialised for Fabrik; it must be initialised on demand (#803).
- The function name and string must not be re-invented, abbreviated or regenerated at any stage; the per-run token must appear unchanged in the child issue, the beta code and its test.
- The alpha change must not be implemented or merged before the beta change has landed, since it would not compile.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo `handarbeit/fabrik-test-delta` MUST export `HelloE2E202610141933573941() string` in `pkg/greeting/greeting.go`, returning exactly `e2e-cross-repo-spawn-20261014-193357-3941`.
- **FR-002**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-003**: This repo's `main.go` MUST call `greeting.HelloE2E202610141933573941()`, keep its existing behaviour, and print the returned value.
- **FR-004**: The module files in this repo MUST be tidy (`go mod tidy` run if needed) so the program builds.
- **FR-005**: The beta change MUST land before the alpha change.
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-007**: The function name and string literal MUST be used verbatim, with no substituted values, in the spawned child issue, the beta code and its test.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A sub-issue targeting `handarbeit/fabrik-test-delta` exists as a result of Plan, and blocks this issue until resolved.
- **SC-002**: The text `e2e-cross-repo-spawn-20261014-193357-3941` and the name `HelloE2E202610141933573941` are present in the beta repo on its default branch after the scenario.
- **SC-003**: The beta repo's tests pass, and this repo builds and prints the returned string when run.
- **SC-004**: The beta change merged before the alpha change.

## Assumptions

- "Alpha" in the issue text means this repo (`handarbeit/fabrik-test-gamma`); "beta" means `handarbeit/fabrik-test-delta`.
- `fabrik:yolo` is set, so stages auto-advance and this stage completes without waiting for human approval.
- This repo already has a `main.go` and module setup from which the beta package can be imported; no import-path questions block the spec (Research and Plan decide how).
- The `pkg/greeting` package and its test file exist in the beta repo or can be created there as part of the child issue.

## Out of Scope *(optional)*

- Any other changes to either repo beyond the function, its test, and the `main.go` call and print.
- Generating a new token or altering the given name and string.
- Changes to the pipeline's spawn mechanism itself; this issue only exercises it.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
