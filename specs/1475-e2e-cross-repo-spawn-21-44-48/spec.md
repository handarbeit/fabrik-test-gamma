# Feature Specification: E2E cross-repo spawn

**Feature Branch**: `fabrik/issue-1475`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610310144483284() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261031-014448-3284\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610310144483284() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It checks that when work spans two repositories with a strict ordering dependency, the Plan stage can split it. The part that lands in the dependency repository (`handarbeit/fabrik-test-delta`, "beta") must be delivered as a spawned child issue. The part in this repository (`handarbeit/fabrik-test-gamma`, "alpha") must wait until the beta part has landed, because alpha consumes a function that does not exist in beta yet.

The scenario covers two earlier changes:
- **#803**: on-demand initialization of the spawn-target repository.
- **#1308**: a per-run-unique fixture token, so the final assertion identifies exactly this run's output.

The function name and string literal contain a per-run-unique token. They must be used exactly as written, because the scenario's final assertion checks that this exact text landed in the beta repository.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repository gains the greeting function (Priority: P1)

The beta repository exposes a new exported greeting function that returns the run-unique string. A test in that repository covers it. This work is delivered through a child issue spawned by Plan, not done directly from this issue.

**Why this priority**: Alpha depends on it, and the scenario's final assertion checks for this text in the beta repo. Nothing else can land until it does.

**Independent Test**: In the beta repository, build and run the tests. The new function must exist, be exported, and return the exact literal. The new test must pass.

**Acceptance Scenarios**:

1. **Given** the beta repository's greeting package, **When** `HelloE2E202610310144483284()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261031-014448-3284`.
2. **Given** the beta repository's greeting test file, **When** the test suite runs, **Then** a test for the new function exists and passes.
3. **Given** this issue's Plan output, **When** it decomposes the work, **Then** it contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` that carries the exact function name and literal.

---

### User Story 2 - Alpha repository calls the new function (Priority: P2)

After the beta change has landed, this repository's `main.go` calls the new greeting function in addition to its existing behavior and prints the result. Dependency metadata is tidied if needed.

**Why this priority**: It completes the cross-repo chain, but it can only be done after User Story 1.

**Independent Test**: Once beta has landed, build and run this repository. The program must compile and its output must include the returned literal.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** this repository is built, **Then** the build succeeds.
2. **Given** the built program, **When** it runs, **Then** it prints `e2e-cross-repo-spawn-20261031-014448-3284` and its prior output is unchanged.

---

### User Story 3 - Strict ordering is enforced (Priority: P1)

The alpha change must not be implemented, merged or built before the beta change has landed. The decomposition must express this dependency.

**Why this priority**: Ordering is the point of the scenario. If it is violated, alpha fails to compile.

**Independent Test**: Inspect the pipeline state. The parent issue must be blocked on the spawned beta child until the child completes.

**Acceptance Scenarios**:

1. **Given** Plan has spawned the beta child, **When** the child is not yet done, **Then** the parent does not proceed to Implement.
2. **Given** the beta child has landed, **When** the parent resumes, **Then** it proceeds with the alpha change.

---

### Edge Cases

- The function name and literal must not be paraphrased, re-generated or altered anywhere. Any deviation fails the final assertion.
- If the beta repository was never initialized for spawn targets, on-demand initialization (#803) must handle it. This is part of what the scenario exercises.
- If `go mod tidy` produces no changes, none are required.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repository MUST gain an exported function `HelloE2E202610310144483284() string` in its greeting package that returns exactly `e2e-cross-repo-spawn-20261031-014448-3284`.
- **FR-002**: The beta repository MUST gain a test for that function in the greeting package's test file, and it MUST pass.
- **FR-003**: The Plan stage MUST decompose the work with a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` for the beta change.
- **FR-004**: The spawned beta work MUST use the exact function name and string literal given in this spec, verbatim.
- **FR-005**: This repository's `main.go` MUST call `greeting.HelloE2E202610310144483284()` in addition to its existing behavior and print the result.
- **FR-006**: Existing behavior and output of `main.go` MUST be preserved.
- **FR-007**: Dependency metadata (`go mod tidy`) MUST be updated if, and only if, needed for the build to succeed.
- **FR-008**: The beta change MUST land before the alpha change is implemented or merged.

### Key Entities

- **Alpha repository** (`handarbeit/fabrik-test-gamma`): this repository, the consumer of the new function.
- **Beta repository** (`handarbeit/fabrik-test-delta`): the dependency repository, the provider of the new function and the spawn target.
- **Spawned child issue**: the beta-repository work item created by Plan, which this issue is blocked on.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repository contains the exact function name and exact literal given in this spec, and its tests pass.
- **SC-002**: After the run, this repository builds and its program output includes the exact literal.
- **SC-003**: The Plan output contains at least one spawn block targeting the beta repository, and exactly one child issue is created for it.
- **SC-004**: The beta change lands before the alpha change, and this repository never has a state on its default branch that fails to build for lack of the beta function.

## Assumptions

- The beta repository already has a greeting package with a source file and test file at the stated paths. If not, on-demand initialization (#803) is expected to cover it.
- The per-run-unique token in the function name and literal is final. It must not be regenerated.
- The existing output of `main.go` is not otherwise specified and must simply be preserved.
- The beta repository is reachable by this pipeline (the board and credentials can serve it).

## Out of Scope *(optional)*

- Any changes to the engine, the spawn logic, or the pipeline itself.
- Any functionality in either repository beyond the one function, its test, and the single call and print.
- Refactoring or cleanup of the greeting package or of `main.go`.

## Source References *(optional)*

- #803: on-demand spawn-target repo init (mentioned for context).
- #1308: per-run-unique fixture token (mentioned for context).
