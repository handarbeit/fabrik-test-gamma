# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2253`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611170510559531() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261117-051055-9531\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611170510559531() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It covers two earlier changes:

- **#803**: on-demand initialisation of the repo a spawned child targets.
- **#1308**: a per-run-unique fixture token, so a run's final assertion cannot be satisfied by leftovers from an earlier run.

The work needs changes in two repositories, in a strict order. The new greeting function must exist in `handarbeit/fabrik-test-delta` (the "beta" repo) before `handarbeit/fabrik-test-gamma` (this repo, the "alpha" repo) can call it. The pipeline should therefore split the work, with the beta change as a spawned sub-issue that lands first. The parent issue then lands the change in this repo.

The function name and string literal contain a per-run-unique token. They must be used exactly as written, because the scenario's final assertion checks that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo (`handarbeit/fabrik-test-delta`) gains an exported greeting function that returns the run-unique literal. It has a test.

**Why this priority**: Everything else depends on it. The alpha change cannot compile without it, and the scenario's final assertion checks the text landing in this repo.

**Independent Test**: In the beta repo, run the Go tests for the greeting package. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261117-051055-9531`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611170510559531()` is called, **Then** it returns `"e2e-cross-repo-spawn-20261117-051055-9531"`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the new function (Priority: P1)

This repo's `main.go` calls the beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo change and shows the strict ordering holds. Alpha builds only once beta has landed.

**Independent Test**: After the beta change has landed, build and run this repo. The output includes `e2e-cross-repo-spawn-20261117-051055-9531`, and the previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta change has landed, **When** this repo is built, **Then** the build succeeds with module dependencies tidy.
2. **Given** the built program, **When** it runs, **Then** it prints the result of `greeting.HelloE2E202611170510559531()` in addition to its prior output.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

Plan emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`. This creates a sub-issue for the beta change, which the parent depends on.

**Why this priority**: Exercising the cross-repo spawn path (#803) is the purpose of this scenario.

**Independent Test**: After Plan, a child issue exists in the beta repo, is registered on the board, and blocks the parent. The beta change lands before the alpha change.

**Acceptance Scenarios**:

1. **Given** the Plan stage for this issue, **When** it completes, **Then** a child issue exists in the beta repo covering the greeting function and its test.
2. **Given** the child is unfinished, **When** the parent would proceed, **Then** it is held until the child's work has landed.

---

### Edge Cases

- The beta repo may need on-demand initialisation when the child is spawned (#803).
- A leftover function or literal from an earlier run must not satisfy this run's assertion. Only the exact token given here counts (#1308).
- If the alpha change is attempted before the beta change lands, it fails to build. The ordering constraint exists to prevent this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202611170510559531() string` in its greeting package that returns exactly `"e2e-cross-repo-spawn-20261117-051055-9531"`. The name and literal MUST be used verbatim, with no variation.
- **FR-002**: The beta repo MUST include a test for that function in its greeting package's test file.
- **FR-003**: This repo's `main.go` MUST call `greeting.HelloE2E202611170510559531()` and print the result, while keeping its existing behaviour.
- **FR-004**: This repo MUST remain buildable with tidy module dependencies once the beta function is available (`go mod tidy` if needed).
- **FR-005**: The beta change MUST land before the alpha change.
- **FR-006**: Plan MUST decompose the work with a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.

### Key Entities *(if applicable)*

- **Beta repo**: `handarbeit/fabrik-test-delta`. It hosts the greeting function and its test, and is the spawn target.
- **Alpha repo**: `handarbeit/fabrik-test-gamma`. It hosts this issue and the consuming `main.go`.
- **Run-unique token**: `20261117-051055-9531` (the function name carries the same digits without separators). It ties the final assertion to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The literal `e2e-cross-repo-spawn-20261117-051055-9531` and the function `HelloE2E202611170510559531` are present in the beta repo's default branch after the run.
- **SC-002**: The beta repo's greeting tests pass with the new test included.
- **SC-003**: This repo builds, and its program output includes the literal above alongside its prior output.
- **SC-004**: A child issue targeting the beta repo was created by Plan, and its change landed before the parent's.

## Assumptions

- The beta repo already has a greeting package whose source and test files are where the issue says. If it does not, the spawned child creates them, which is the on-demand initialisation behaviour (#803) under test.
- This repo already depends on the beta repo's greeting package, or the dependency can be added with the standard Go module tooling.
- The scenario's pass/fail is judged by the harness's assertion on the beta repo's content, not by manual inspection.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single function and its test.
- Any change to this repo's behaviour beyond calling and printing the new function.
- Changes to the Fabrik engine's spawn or board logic. This scenario only exercises it.
- Decomposing into more than one child or more than one other repo.
