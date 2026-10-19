# Feature Specification: E2E cross-repo spawn (beta-first decomposition)

**Feature Branch**: `fabrik/issue-442`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610190709125993() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-070912-5993\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610190709125993() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier changes:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: a per-run-unique fixture token, so each run's final assertion can tell its own output apart from earlier runs.

The work spans two repositories with a strict ordering dependency. The shared greeting package lives in the beta repo (`handarbeit/fabrik-test-delta`). The consuming program lives in this repo (`handarbeit/fabrik-test-gamma`, the "alpha" side). This repo cannot call a function that does not yet exist in beta, so the beta change must land first. Because of that, the Plan stage is expected to decompose the work by spawning a child issue targeting the beta repo, rather than doing everything in one repo.

The function name and string literal embed a per-run-unique token. They must be used exactly as written, because the scenario's final assertion checks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the per-run greeting function (Priority: P1)

The beta repo's greeting package gains a new exported function that returns the per-run-unique string, with a test covering it.

**Why this priority**: It is the prerequisite for everything else, and it is what the scenario's final assertion inspects.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261019-070912-5993`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610190709125993()` is called, **Then** it returns `e2e-cross-repo-spawn-20261019-070912-5993`.
2. **Given** the beta repo's greeting test file, **When** the package tests run, **Then** a test for this function exists and passes.

---

### User Story 2 - This repo's program prints the beta greeting (Priority: P2)

This repo's main program calls the new beta function in addition to its existing behaviour and prints the result.

**Why this priority**: It depends on User Story 1 and proves the two-repo change composes.

**Independent Test**: Once the beta change is available, build and run this repo's program. Its output includes `e2e-cross-repo-spawn-20261019-070912-5993`, and the module builds cleanly.

**Acceptance Scenarios**:

1. **Given** the beta function is available, **When** this repo's program runs, **Then** it prints the string returned by `greeting.HelloE2E202610190709125993()`.
2. **Given** the program's pre-existing behaviour, **When** the change is applied, **Then** that behaviour is unchanged.
3. **Given** the dependency set changed, **When** the module is tidied, **Then** `go mod tidy` leaves no pending changes and the build succeeds.

---

### User Story 3 - Plan decomposes the work across repos in beta-first order (Priority: P1)

The Plan stage emits a spawn block targeting the beta repo. That creates a child issue for the beta work, and this issue is blocked on it.

**Why this priority**: Demonstrating this decomposition is the purpose of the scenario (#803, #1308).

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` for the beta work. This issue is blocked by that child until it lands.

**Acceptance Scenarios**:

1. **Given** the requirements span two repos with a strict order, **When** Plan completes, **Then** a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` has been emitted and a child issue created there.
2. **Given** the child issue exists, **When** this issue is processed further, **Then** the alpha-side change is not completed before the beta-side change has landed.
3. **Given** the beta repo needs on-demand initialisation for spawning (#803), **When** the child is spawned, **Then** it succeeds without manual repo setup.

---

### Edge Cases

- The function name and string literal must not be altered, abbreviated or regenerated. A different token would fail the final assertion.
- If the alpha-side change is attempted before the beta function exists, the build fails. The ordering constraint exists to prevent this.
- Re-running the scenario uses a different per-run token, so the token must not be assumed to match earlier runs.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's `pkg/greeting/greeting.go` MUST export a function `HelloE2E202610190709125993() string`.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261019-070912-5993`.
- **FR-003**: The beta repo's `pkg/greeting/greeting_test.go` MUST include a test for this function.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610190709125993()` and print the result, in addition to its existing behaviour.
- **FR-005**: The module in this repo MUST remain tidy and buildable after the change. `go mod tidy` is run if needed.
- **FR-006**: The beta change MUST land before the alpha change. This repo's change MUST NOT land first.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-008**: The exact function name and string literal given in this spec MUST be used verbatim, with no substitute token.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The beta repo's default branch contains the text `e2e-cross-repo-spawn-20261019-070912-5993` in its greeting package, which is the scenario's final assertion.
- **SC-002**: The greeting package tests in the beta repo pass with the new test included.
- **SC-003**: This repo builds and its program prints the per-run string.
- **SC-004**: A child issue was created in the beta repo from this issue, without manual repo setup, and the beta change landed before this repo's change.

## Assumptions

- The beta repo already has the `pkg/greeting` package; the issue says to add to the existing `greeting.go` and `greeting_test.go`.
- This repo already imports the beta greeting package, or `go mod tidy` can resolve it once the function exists.
- "Whatever it already does" in `main.go` means existing output is preserved and the new line is additive.

## Out of Scope

- Any changes to the beta repo beyond the single function and its test.
- Changing Fabrik's engine, spawn mechanism or fixture-token generation. This issue only exercises them.
- Refactoring the greeting package or `main.go` beyond the additions above.

## Source References

- #803 (on-demand spawn-target repo init)
- #1308 (per-run-unique fixture token)
