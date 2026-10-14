# Feature Specification: e2e cross-repo spawn (greeting.HelloE2E202610142252572433)

**Feature Branch**: `fabrik/issue-43`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610142252572433() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261014-225257-2433\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610142252572433() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. The work spans two repositories with a strict ordering dependency. The shared greeting package lives in `handarbeit/fabrik-test-delta` (the "beta" repo). This repository, `handarbeit/fabrik-test-gamma` (the "alpha" repo), imports it from its entry point. The entry point cannot reference a new greeting function until that function exists in the beta repo. Fabrik should therefore decompose the work by spawning a sub-issue in the beta repo, which must land first.

The scenario guards two behaviours:
- **#803**: on-demand initialisation of the spawn-target repository.
- **#1308**: per-run-unique fixture tokens. The function name and string literal embed a per-run token so the final assertion can prove that this run's change, and no earlier run's, landed in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo's greeting package gains a new exported function returning this run's unique string, covered by a unit test. This work is delivered by a spawned sub-issue in the beta repo.

**Why this priority**: It is the prerequisite for everything else. The alpha repo cannot compile against a function that does not exist, and the scenario's final assertion checks this text in the beta repo.

**Independent Test**: In the beta repo, run the package tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261014-225257-2433`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610142252572433()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261014-225257-2433`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo prints the new greeting (Priority: P2)

The alpha repo's entry point calls the new beta function and prints its result, in addition to its existing output. This work is done in this issue's own pull request, after the beta change has landed.

**Why this priority**: It depends on User Story 1 and demonstrates the ordered, cross-repo dependency being exercised.

**Independent Test**: Build and run this repo's program. Its output includes the line `e2e-cross-repo-spawn-20261014-225257-2433` alongside all prior output, and the module builds and tests cleanly.

**Acceptance Scenarios**:

1. **Given** the beta change has landed and this repo depends on a version containing it, **When** the program runs, **Then** it prints `e2e-cross-repo-spawn-20261014-225257-2433` in addition to its existing output.
2. **Given** the updated dependency, **When** the module is built and tested, **Then** both succeed with no missing or unresolved module requirements.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

During planning, Fabrik decomposes the work by emitting a child-spawn block that targets the beta repo, so the beta change is tracked and delivered as a separate sub-issue that blocks this one.

**Why this priority**: Proving this decomposition is the purpose of the scenario (#803, #1308).

**Independent Test**: After Plan, a sub-issue exists in `handarbeit/fabrik-test-delta` that blocks this issue, and its description carries the exact function name and literal from FR-001 and FR-002.

**Acceptance Scenarios**:

1. **Given** Plan completes for this issue, **When** its output is processed, **Then** a child issue is created in `handarbeit/fabrik-test-delta` and linked as a blocker of this issue.
2. **Given** the child issue exists, **When** its description is read, **Then** it names `HelloE2E202610142252572433` and the literal `e2e-cross-repo-spawn-20261014-225257-2433` verbatim.
3. **Given** the child issue is unresolved, **When** this issue is next evaluated, **Then** it does not proceed to implementation until the child is resolved.

---

### Edge Cases

- The beta change has not yet landed or been published when this repo is built. The dependency cannot resolve the new function, so this repo's change must not proceed until the beta change is available.
- A run reuses a name or literal from a previous run. The per-run token exists to prevent this, so the exact values given must be used and none invented.
- The spawn-target repo has not been initialised for Fabrik. Spawning must still succeed (#803).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's greeting package MUST export `HelloE2E202610142252572433() string`.
- **FR-002**: `HelloE2E202610142252572433()` MUST return exactly the literal `e2e-cross-repo-spawn-20261014-225257-2433`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610142252572433` in its greeting package tests.
- **FR-004**: This repo's entry point MUST call `greeting.HelloE2E202610142252572433()` and print the result, in addition to its existing output.
- **FR-005**: This repo's module requirements MUST resolve cleanly after the change, running dependency tidying if needed.
- **FR-006**: The beta change MUST land before this repo's change, since this repo cannot reference a function that does not yet exist.
- **FR-007**: Plan MUST decompose the work with a child-spawn block targeting `handarbeit/fabrik-test-delta`, and the child issue MUST be a blocker of this issue.
- **FR-008**: The exact function name and string literal above MUST be used verbatim everywhere. No alternative name or value may be introduced.
- **FR-009**: Existing output of this repo's program MUST be unchanged apart from the added line.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repository, which consumes the greeting package.
- **Beta repo**: `handarbeit/fabrik-test-delta`, which owns the greeting package and is the spawn target.
- **Per-run token**: The unique suffix baked into the function name and literal for this scenario run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo's default branch contains `HelloE2E202610142252572433` returning `e2e-cross-repo-spawn-20261014-225257-2433`, with a passing test.
- **SC-002**: This repo's program output includes `e2e-cross-repo-spawn-20261014-225257-2433` and the existing output lines.
- **SC-003**: A child issue in `handarbeit/fabrik-test-delta` was created by Plan and resolved before this issue's implementation merged.
- **SC-004**: Build and tests pass in both repos.

## Assumptions

- The beta repo's greeting package follows the same conventions as its existing functions, which are exported, parameterless and return a string.
- Making the new beta function visible to this repo (for example through a dependency version update) is part of this issue's work.
- The `fabrik:yolo` label means the stages auto-advance, so no human approval gate is expected.

## Out of Scope *(optional)*

- Changes to any other function in the greeting package.
- Cleanup or refactoring of the existing output lines in this repo's entry point.
- Changes to Fabrik's engine or spawn logic. This issue only exercises it.

## Source References *(optional)*

- Fabrik issues #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
