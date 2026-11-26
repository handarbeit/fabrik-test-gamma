# Feature Specification: E2E cross-repo spawn (beta change lands before alpha change)

**Feature Branch**: `fabrik/issue-2758`
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
> - Add an exported function HelloE2E202611261652371503() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261126-165237-1503".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202611261652371503() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. One unit of work needs changes in two repositories with a hard ordering dependency: the helper function must exist in the beta repo (`handarbeit/fabrik-test-delta`) before this repo (`handarbeit/fabrik-test-gamma`) can call it. The scenario checks that the pipeline splits the work, so the beta-side change is carried by a child issue in the beta repo and the dependent change here waits until the child has landed.

It is a regression check for two earlier fixes:

- **#803**: on-demand initialization of a spawn-target repo.
- **#1308**: a per-run-unique fixture token, so a stale artifact from an earlier run cannot satisfy this run's assertion.

The function name and string literal embed a per-run-unique token. They must be used exactly as given. The scenario's final assertion looks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta-side change is decomposed into a child issue in the beta repo (Priority: P1)

The work that belongs in the beta repo (the new greeting function and its test) is split out as a child issue in `handarbeit/fabrik-test-delta`. It is not done as part of this repo's issue.

**Why this priority**: This is the behavior under test. Without the decomposition there is no cross-repo spawn to verify.

**Independent Test**: After the planning stage, check that a child issue exists in the beta repo, is linked to this issue as a blocker, and describes the exact function and literal.

**Acceptance Scenarios**:

1. **Given** this issue requires a function in the beta repo that does not yet exist, **When** the plan is produced, **Then** a child issue targeting `handarbeit/fabrik-test-delta` is created and this issue is blocked by it.
2. **Given** the beta repo has not been used as a spawn target before, **When** the child is created, **Then** the beta repo is initialized on demand as needed (#803) and the child issue is created without manual setup.

---

### User Story 2 - Beta function lands with the exact per-run token (Priority: P1)

The beta repo gains the exported function `HelloE2E202611261652371503() string` in `pkg/greeting/greeting.go`. It returns exactly `"e2e-cross-repo-spawn-20261126-165237-1503"`. A test in `pkg/greeting/greeting_test.go` covers it.

**Why this priority**: The scenario's final assertion checks this exact text in the beta repo, and the alpha-side change depends on it.

**Independent Test**: In the beta repo, run the tests for `pkg/greeting` and confirm that the function exists with that name and returns that literal.

**Acceptance Scenarios**:

1. **Given** the child issue is implemented, **When** its change is merged in the beta repo, **Then** `HelloE2E202611261652371503()` exists and returns `"e2e-cross-repo-spawn-20261126-165237-1503"`.
2. **Given** the new test in `pkg/greeting/greeting_test.go`, **When** the beta repo's tests are run, **Then** the test passes and fails if the returned literal changes.

---

### User Story 3 - This repo consumes the beta function only after beta has landed (Priority: P1)

This repo's `main.go` additionally calls `greeting.HelloE2E202611261652371503()` and prints the result. Existing behavior is retained. This work does not start or merge before the beta change has landed.

**Why this priority**: It completes the end-to-end flow and proves the ordering constraint is honored.

**Independent Test**: Build and run this repo after the beta change is available. The output contains `e2e-cross-repo-spawn-20261126-165237-1503` in addition to the previous output, and the build succeeds.

**Acceptance Scenarios**:

1. **Given** the beta change has landed, **When** this repo's change is implemented, **Then** `main.go` calls `greeting.HelloE2E202611261652371503()` and prints its result, and the module builds.
2. **Given** the beta change has not landed, **When** this issue would otherwise proceed to implementation, **Then** it is held until the beta child is resolved.

---

### Edge Cases

- The function name and literal must not be altered, abbreviated or regenerated. Any difference fails the scenario.
- If dependency metadata must be updated for the beta package to resolve (for example `go mod tidy`), it is done as part of this repo's change.
- A leftover function or literal from an earlier run (a different token) must not satisfy this run's check (#1308).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The work is decomposed so that the beta-repo change is carried by a child issue created in `handarbeit/fabrik-test-delta`, and this issue is blocked by that child.
- **FR-002**: The beta repo exports `HelloE2E202611261652371503() string` in `pkg/greeting/greeting.go`, returning exactly `"e2e-cross-repo-spawn-20261126-165237-1503"`.
- **FR-003**: The beta repo has a test for that function in `pkg/greeting/greeting_test.go` that asserts the returned literal.
- **FR-004**: This repo's `main.go` calls `greeting.HelloE2E202611261652371503()` and prints the result, retaining all existing behavior.
- **FR-005**: This repo's module metadata is updated (for example via `go mod tidy`) if needed so the build succeeds with the new import.
- **FR-006**: The beta change lands before this repo's change; this repo's change does not merge while the beta function is absent.
- **FR-007**: The function name and string literal are used verbatim as given in this spec, including in the child issue's description.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, where this issue lives and where `main.go` is updated.
- **Beta repo**: `handarbeit/fabrik-test-delta`, the spawn target that receives the new greeting function and its test.
- **Child issue**: the issue in the beta repo that carries the beta-side change and blocks this issue.
- **Per-run token**: the unique suffix baked into the function name and literal, identifying this run's artifacts.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Exactly one child issue exists in `handarbeit/fabrik-test-delta` for the beta-side change, and this issue is blocked by it.
- **SC-002**: The beta repo's default branch contains `HelloE2E202611261652371503` and the literal `e2e-cross-repo-spawn-20261126-165237-1503` after the run, and its `pkg/greeting` tests pass.
- **SC-003**: This repo builds, and running it prints a line containing `e2e-cross-repo-spawn-20261126-165237-1503`.
- **SC-004**: The beta change merges before this repo's change.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, as the issue states. If not, the child issue creates the minimum needed.
- This repo already imports the beta repo's `pkg/greeting` package, or can add that import. The issue's mention of `go mod tidy` implies a dependency update may be needed.
- The existing output of `main.go` should keep working, and the new line is purely additive.
- The `fabrik:yolo` label is present, so stages auto-advance without a human approving the spec.

## Out of Scope

- Changing any other behavior of either repo's greeting code or `main.go`.
- Cleanup of the per-run artifacts after the scenario.
- Changes to Fabrik's engine, spawn logic or the e2e harness itself.

## Source References

- #803: on-demand spawn-target repo initialization.
- #1308: per-run-unique fixture token.
