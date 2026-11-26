# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2933`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202611262151585332() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261126-215158-5332".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202611262151585332() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a strict ordering: the greeting function must exist in the beta repo (`handarbeit/fabrik-test-delta`) before this repo (`handarbeit/fabrik-test-gamma`) can call it. The scenario exists to prove that Plan decomposes such work by spawning a child issue in the beta repo, that the spawn-target repo is initialised on demand (#803), and that a per-run-unique fixture token (#1308) lands verbatim in the beta repo. The function name and string literal embed that per-run token and must be used exactly as given; the scenario's final assertion depends on the exact text.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function returning the run-unique literal, with a test, delivered through a child issue spawned from this one.

**Why this priority**: Everything else depends on it, and the scenario's final assertion checks that this text exists in the beta repo.

**Independent Test**: In the beta repo, run its test suite and confirm `HelloE2E202611262151585332()` exists in `pkg/greeting/greeting.go` returning `"e2e-cross-repo-spawn-20261126-215158-5332"`, with a passing test in `pkg/greeting/greeting_test.go`.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`, and a child issue is created there.
2. **Given** the child issue is implemented and merged, **When** the beta repo is inspected, **Then** the function and its test exist with the exact name and literal given.

---

### User Story 2 - This repo calls the new function (Priority: P2)

After the beta change has landed, this repo's `main.go` calls the new greeting function and prints its result, in addition to existing behaviour.

**Why this priority**: It completes the scenario, but it cannot be done until User Story 1 lands.

**Independent Test**: Build and run this repo; the output includes `e2e-cross-repo-spawn-20261126-215158-5332`, and the build succeeds.

**Acceptance Scenarios**:

1. **Given** the beta repo change has landed, **When** this repo's work runs, **Then** `main.go` calls `greeting.HelloE2E202611262151585332()` and prints the result, and existing output is preserved.
2. **Given** the dependency requires it, **When** the work completes, **Then** module files are tidy (`go mod tidy` applied if needed) and the build passes.

---

### Edge Cases

- The beta repo has not been initialised for spawning: the spawn target must be initialised on demand (#803) rather than failing.
- A different function name or literal is invented in place of the given one: the scenario fails; the exact values must be used.
- This repo's work starts before the beta change lands: the child issue must block this issue so ordering is enforced.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-002**: The beta repo MUST gain an exported function `HelloE2E202611262151585332() string` in `pkg/greeting/greeting.go` returning exactly `"e2e-cross-repo-spawn-20261126-215158-5332"`.
- **FR-003**: The beta repo MUST gain a test for that function in `pkg/greeting/greeting_test.go`.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202611262151585332()` in addition to its existing behaviour and print the result.
- **FR-005**: Module dependency files in this repo MUST be tidy (`go mod tidy` run if needed) so the build succeeds.
- **FR-006**: The beta change MUST land before this repo's change; this issue's work MUST NOT depend on a function that does not yet exist.
- **FR-007**: The function name and string literal MUST be used verbatim as given and MUST NOT be replaced by invented values.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in `handarbeit/fabrik-test-delta` as a result of Plan's spawn block.
- **SC-002**: The exact text `HelloE2E202611262151585332` and `e2e-cross-repo-spawn-20261126-215158-5332` appear in the beta repo's greeting package after the child lands.
- **SC-003**: This repo builds and its program output contains the literal `e2e-cross-repo-spawn-20261126-215158-5332`.
- **SC-004**: The beta repo's tests pass, including the new one.

## Assumptions

- The beta repo already has a `pkg/greeting` package; if it does not, creating it is part of the child work.
- This repo already depends on the beta repo as a Go module, or can be made to via `go mod tidy`.
- Printing the result to standard output satisfies "print the result".
- `fabrik:yolo` is present, so the issue advances through stages without waiting for spec approval.

## Out of Scope *(optional)*

- Any greeting behaviour beyond the single function and its test.
- Changes to Fabrik's engine, spawn logic, or other repositories.
- Refactoring existing code in either repo.

## Source References *(optional)*

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
