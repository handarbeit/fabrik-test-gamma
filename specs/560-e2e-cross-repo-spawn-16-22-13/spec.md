# Feature Specification: E2E cross-repo spawn (decompose across fabrik-test-delta and fabrik-test-gamma)

**Feature Branch**: `fabrik/issue-560`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610192022130633() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-202213-0633\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610192022130633() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a strict ordering dependency. The greeting function must exist in `handarbeit/fabrik-test-delta` (the "beta" repo) before `handarbeit/fabrik-test-gamma` (this repo, referred to in the original text as "alpha") can call it. Fabrik should therefore split the work. The Plan stage is expected to emit a child-spawn block targeting the delta repo, so the delta change is delivered by its own child issue and this issue's change in gamma follows once that lands.

The scenario guards two behaviours:

- **#803**: on-demand initialisation of a spawn-target repo (the child's repo is brought into service when the spawn happens).
- **#1308**: a per-run-unique fixture token. The function name and returned string below embed a token unique to this run (`202610192022130633` in the name, `e2e-cross-repo-spawn-20261019-202213-0633` in the value). The scenario's final assertion checks that exactly this text landed in the delta repo, so the names and literal must be used verbatim and never substituted.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Delta repo gains the greeting function via a spawned child (Priority: P1)

The work for `handarbeit/fabrik-test-delta` is carried out by a child issue spawned from this issue during Plan, not by this issue's own branch.

**Why this priority**: This is the point of the scenario. Cross-repo spawn is what is under test, and gamma's change cannot be done without it.

**Independent Test**: After the pipeline runs, inspect `handarbeit/fabrik-test-delta`. The greeting package has the new exported function returning the exact literal, and a test covering it, delivered through an issue spawned by this one.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a child-spawn block targeting `handarbeit/fabrik-test-delta` describing the delta-side change.
2. **Given** the child issue is created in the delta repo, **When** it is implemented and merged, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202610192022130633() string` returning `"e2e-cross-repo-spawn-20261019-202213-0633"`.
3. **Given** the delta change, **When** its tests run, **Then** `pkg/greeting/greeting_test.go` contains a test asserting that exact return value, and it passes.

---

### User Story 2 - Gamma consumes the new function only after delta has landed (Priority: P1)

This repo's entry point calls the new function and prints its result, and this only happens after the delta change is available.

**Why this priority**: It verifies the strict-ordering constraint, since the parent depends on the child. It also produces the visible end result of the cross-repo change.

**Independent Test**: Build and run this repo's program after both changes have landed. Its output includes the string `e2e-cross-repo-spawn-20261019-202213-0633`, and its existing output is preserved.

**Acceptance Scenarios**:

1. **Given** the delta child has not yet landed, **When** this issue's pipeline reaches implementation, **Then** it does not proceed ahead of the child. The parent is blocked on the child, so it never references a function that does not exist.
2. **Given** the delta change has landed, **When** this issue is implemented, **Then** `main.go` calls `greeting.HelloE2E202610192022130633()` in addition to its existing behaviour and prints the returned value.
3. **Given** the updated `main.go`, **When** the module is built, **Then** dependencies resolve (running `go mod tidy` if needed) and the build and tests succeed.

---

### Edge Cases

- The delta repo has not yet been initialised for Fabrik use. The spawn must still succeed (#803).
- A re-run or retry of Plan must not create a duplicate child issue for the same decomposition block.
- The function name or literal altered, shortened or regenerated in any stage. This would make the final assertion fail and must not happen.
- The dependency version of the delta repo that gamma resolves does not yet contain the new function. The gamma build would then fail, which is why the ordering constraint exists.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST export a function named exactly `HelloE2E202610192022130633` with signature `() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610192022130633()` MUST return exactly the string `e2e-cross-repo-spawn-20261019-202213-0633`.
- **FR-003**: The delta repo MUST contain a test in `pkg/greeting/greeting_test.go` that verifies the return value in FR-002 and passes.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610192022130633()` and print the result, while keeping all of its existing behaviour.
- **FR-005**: The delta-side work MUST be delivered through a child issue spawned from this issue's Plan stage, using a child-spawn block that targets the delta repo.
- **FR-006**: The delta change MUST land before this repo's change. This issue must not be implemented against a delta version that lacks the function.
- **FR-007**: The module metadata in this repo MUST be left consistent (`go mod tidy` run if needed) so the program builds.
- **FR-008**: The exact function name and string literal given in this issue MUST be used verbatim in every stage and artifact, and never replaced by generated alternatives.

### Key Entities *(if applicable)*

- **Parent issue**: This issue (#560) in `handarbeit/fabrik-test-gamma`. It owns the `main.go` change and is blocked on the child.
- **Child issue**: The issue spawned in `handarbeit/fabrik-test-delta` that owns the greeting function and its test.
- **Fixture token**: The per-run-unique value embedded in the function name and returned literal, used by the final assertion.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Exactly one child issue exists in `handarbeit/fabrik-test-delta` that was spawned from this issue, and it reaches a merged or completed state.
- **SC-002**: A search of the delta repo's default branch finds `HelloE2E202610192022130633` and `e2e-cross-repo-spawn-20261019-202213-0633` exactly as given. The scenario's final assertion passes.
- **SC-003**: The delta repo's tests pass, including the new test, on the default branch after the change lands.
- **SC-004**: This repo builds and its tests pass after the `main.go` change, and running the program prints `e2e-cross-repo-spawn-20261019-202213-0633`.
- **SC-005**: Chronologically, the delta change is merged before this repo's change is merged.

## Assumptions

- "Beta" in the original text means `handarbeit/fabrik-test-delta`, and "alpha" means this repo, `handarbeit/fabrik-test-gamma`. The names reflect an earlier naming of the test repos, and the explicit repo names in the issue are authoritative.
- This repo's `main.go` already imports or can import the delta repo's `pkg/greeting` package. If not, adding that dependency is part of the `go mod tidy` step.
- The pipeline's existing parent-blocked-on-child behaviour provides the ordering guarantee in FR-006, so no additional ordering mechanism is required.
- The existing behaviour of `main.go` that must be preserved is whatever it does today. The spec does not enumerate it.

## Out of Scope *(optional)*

- Any change to Fabrik's engine, spawn logic or stage configuration. This issue only exercises the existing behaviour.
- Additional functions, refactors or unrelated edits in either test repo.
- Changes to the greeting package beyond the single function and its test.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation (regression covered by this scenario)
- #1308: per-run-unique fixture token (regression covered by this scenario)
