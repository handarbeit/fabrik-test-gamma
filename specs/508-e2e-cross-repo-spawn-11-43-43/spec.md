# Feature Specification: e2e cross-repo spawn (strict-order decomposition across two repos)

**Feature Branch**: `fabrik/issue-508`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610191543431133() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-154343-1133\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610191543431133() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a hard ordering dependency: the function must exist in the shared greeting package (`handarbeit/fabrik-test-delta`) before this repository (`handarbeit/fabrik-test-gamma`), which imports that package, can call it. The Plan stage is therefore expected to decompose the work by emitting a child-spawn block that targets the delta repository, rather than trying to do everything in one repository.

The scenario regression-tests two earlier fixes: on-demand initialisation of a spawn-target repository (#803) and the per-run-unique fixture token (#1308). The function name and string literal embed a per-run-unique token, and the scenario's final assertion checks that exact text in the delta repository, so they must be used verbatim.

Note: the issue text refers to the two repositories as "alpha"/"beta"; the concrete repositories are this one (`fabrik-test-gamma`, the importing consumer) and `fabrik-test-delta` (the dependency that must land first).

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Dependency repo gains the new greeting function (Priority: P1)

The delta repository's greeting package gains a new exported function with a fixed name and return value, plus a test, delivered as a spawned child issue that completes before the consumer change.

**Why this priority**: It is the prerequisite for everything else and is the artifact the scenario's final assertion inspects.

**Independent Test**: In the delta repository, call `HelloE2E202610191543431133()` and compare with the expected literal; run the delta repository's tests.

**Acceptance Scenarios**:

1. **Given** the delta repository's greeting package, **When** `HelloE2E202610191543431133()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-154343-1133`.
2. **Given** the delta repository's test suite, **When** it is run, **Then** a test covering the new function exists in the greeting package's tests and passes.

---

### User Story 2 - Consumer repo prints the new function's result (Priority: P2)

This repository's program calls the new function and prints its result, in addition to its existing output, using a dependency version that includes the new function.

**Why this priority**: It completes the end-to-end change, but can only be done after Story 1 has landed.

**Independent Test**: Build and run this repository's program (and its tests) and confirm the new line appears.

**Acceptance Scenarios**:

1. **Given** the delta change has landed and is available to this repository, **When** the program runs, **Then** its output includes `e2e-cross-repo-spawn-20261019-154343-1133` in addition to all previous output lines.
2. **Given** the updated program, **When** the module is built and tested, **Then** the build and tests succeed with module metadata consistent (tidy).

---

### User Story 3 - Plan decomposes across repositories in strict order (Priority: P1)

The Plan stage splits the work into a child issue in the delta repository and the remaining consumer work in this repository, with the child required to land first.

**Why this priority**: Verifying this decomposition behaviour is the purpose of the scenario.

**Independent Test**: Inspect the Plan output and the board after Plan: a child issue exists in the delta repository, is linked as a blocker of this issue, and this issue does not proceed to implementation until it is done.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** it has emitted a child-spawn block targeting the delta repository describing the delta work with the exact function name and literal.
2. **Given** the child exists, **When** the pipeline continues, **Then** this issue's implementation does not begin until the child's change has landed.

---

### Edge Cases

- The spawn-target repository has not been initialised for Fabrik yet (the #803 scenario): the spawn must still succeed on demand.
- The function name or literal is altered, shortened or regenerated by any stage: the final assertion fails; they must appear verbatim.
- This repository's dependency pin does not yet include the new function when the consumer change is attempted: the consumer change must not be considered complete until it does.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repository's greeting package MUST export `HelloE2E202610191543431133() string` returning exactly `e2e-cross-repo-spawn-20261019-154343-1133`.
- **FR-002**: The delta repository MUST include a test of `HelloE2E202610191543431133` in the greeting package's tests, and it MUST pass.
- **FR-003**: This repository's program MUST call `greeting.HelloE2E202610191543431133()` and print the result, in addition to (not replacing) its existing output.
- **FR-004**: This repository's dependency on the delta repository MUST resolve to a version that contains the new function, and module metadata MUST be tidy.
- **FR-005**: The delta-repository work MUST be decomposed by the Plan stage into a spawned child issue targeting `handarbeit/fabrik-test-delta`, via the child-spawn marker block.
- **FR-006**: The delta change MUST land before the consumer change in this repository is completed.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this issue, in every stage and both repositories.

### Key Entities *(if applicable)*

- **Consumer repository**: `handarbeit/fabrik-test-gamma` (this repo); imports the greeting package and prints its results.
- **Dependency repository**: `handarbeit/fabrik-test-delta`; owns the greeting package and the new function.
- **Spawned child issue**: The delta-repository work item created by Plan and blocking this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The exact text `e2e-cross-repo-spawn-20261019-154343-1133` is present in the delta repository's greeting package on its default branch after the run.
- **SC-002**: Running this repository's program prints that literal exactly once, in addition to all previously printed lines.
- **SC-003**: A child issue in the delta repository exists, was created by the Plan stage, and was completed before this issue's implementation merged.
- **SC-004**: Build and tests pass in both repositories.

## Assumptions

- "alpha" in the issue text means this repository (`fabrik-test-gamma`) and "beta" means `fabrik-test-delta`.
- Printing the result as an additional line alongside the existing printed lines satisfies "print the result".
- Updating the consumer's dependency pin to the delta version that contains the new function is part of making the consumer change work (as in earlier runs of this scenario).
- The test in the delta repository follows the conventions of that repository's existing greeting tests.

## Out of Scope *(optional)*

- Any change to existing greeting functions or existing output lines.
- Changes to Fabrik's engine, spawn logic or the e2e harness itself.
- Further decomposition beyond one child issue in the delta repository.
