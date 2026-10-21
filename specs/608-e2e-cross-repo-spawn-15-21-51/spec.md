# Feature Specification: Cross-repo spawn e2e — HelloE2E202610211921513278

**Feature Branch**: `fabrik/issue-608`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610211921513278() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261021-192151-3278\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610211921513278() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work spans two repositories with a hard ordering dependency: the greeting library lives in `handarbeit/fabrik-test-delta`, and this repository (`handarbeit/fabrik-test-gamma`) imports it. The consumer cannot reference a function that does not yet exist in the library, so the library change must land first, and the consumer change follows once the library change is available.

The scenario exists to exercise two behaviours:

- **#803** — on-demand initialisation of a spawn-target repository when Plan decomposes work into a repo other than the issue's own.
- **#1308** — a per-run-unique fixture token, so that the scenario's final assertion can only be satisfied by this run's change landing in the library repo and not by a leftover from an earlier run. The token (`20261021192151` / `3278`) is already embedded in the function name and the returned literal, which must therefore be used exactly as given.

Plan is expected to decompose the work by spawning a sub-issue in the library repo for the library half. The issue is labelled `fabrik:yolo`, so stages advance automatically.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Library repo gains the unique greeting function (Priority: P1)

A new exported function exists in the greeting library (`handarbeit/fabrik-test-delta`) returning this run's unique literal, with a test, delivered through a sub-issue spawned from this issue.

**Why this priority**: It is the prerequisite for everything else and is the artefact the scenario's final assertion inspects.

**Independent Test**: In the library repo, the greeting package exposes `HelloE2E202610211921513278() string` returning exactly `e2e-cross-repo-spawn-20261021-192151-3278`, its test passes, and the library repo's test suite passes.

**Acceptance Scenarios**:

1. **Given** the library repo has no `HelloE2E202610211921513278`, **When** the spawned library sub-issue completes and merges, **Then** the function exists, is exported, and returns the exact literal `e2e-cross-repo-spawn-20261021-192151-3278`.
2. **Given** the function exists, **When** the library's tests run, **Then** a test asserts the exact return value and passes.

---

### User Story 2 - This repo consumes the new function (Priority: P2)

After the library change has landed, the program in this repo calls the new function and prints its result alongside its existing output.

**Why this priority**: It completes the cross-repo change and proves the ordering constraint held, but it is only possible once Story 1 is done.

**Independent Test**: Build and run the program in this repo; its output includes `e2e-cross-repo-spawn-20261021-192151-3278`, and the module builds and its tests pass.

**Acceptance Scenarios**:

1. **Given** the library change has landed and is resolvable as a dependency, **When** this repo's program is built and run, **Then** its output includes `e2e-cross-repo-spawn-20261021-192151-3278` in addition to all prior output.
2. **Given** the dependency is updated to a version containing the function, **When** the module is built, **Then** it builds without errors and module metadata is consistent (tidy).

---

### User Story 3 - Plan decomposes across repos in the right order (Priority: P1)

Plan recognises the two-repo, strictly ordered nature of the work and decomposes it by spawning a sub-issue targeting the library repo, with this issue blocked until that sub-issue is done.

**Why this priority**: Verifying this decomposition behaviour (#803, #1308) is the purpose of the scenario.

**Independent Test**: The Plan output for this issue contains a spawn block targeting `handarbeit/fabrik-test-delta`, a sub-issue is created in that repo, and this issue's implementation does not proceed until the sub-issue is complete.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan runs, **Then** it declares a child spawn targeting the library repo covering the library half of the work, using the exact function name and literal from this spec.
2. **Given** the spawn target repo has not been initialised for Fabrik, **When** the spawn occurs, **Then** the target repo is initialised on demand and the child is created and placed on the board.

---

### Edge Cases

- The function name or literal is altered, shortened, or regenerated by any stage — the final assertion fails; both must be used verbatim.
- The consumer change is attempted before the library change has landed — the build would fail on an undefined function; ordering must be enforced.
- Existing `main.go` output (all prior `Hello…` calls) must continue to print unchanged.
- The dependency version in this repo may need updating (and tidying) for the new function to be resolvable.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The library repo `handarbeit/fabrik-test-delta` MUST export a function `HelloE2E202610211921513278() string` in its greeting package, using this exact name.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261021-192151-3278`.
- **FR-003**: The library repo MUST include a test for that function in its greeting package tests that asserts the exact returned value.
- **FR-004**: The program in this repo (`main.go`) MUST call `greeting.HelloE2E202610211921513278()` and print the result, in addition to its existing output.
- **FR-005**: This repo's module metadata MUST be tidy and the module MUST resolve a library version containing the new function.
- **FR-006**: The library change MUST land before the consumer change in this repo is implemented or merged.
- **FR-007**: Plan MUST decompose the work by spawning a sub-issue in `handarbeit/fabrik-test-delta` for the library half of the work, and this issue MUST be blocked on it.
- **FR-008**: The exact function name and string literal given in this spec MUST be used verbatim in every artefact (spawned sub-issue, code, tests); no stage may invent substitutes.

### Key Entities

- **Library repo** (`handarbeit/fabrik-test-delta`): owns the greeting package and the new function; the spawn target.
- **Consumer repo** (`handarbeit/fabrik-test-gamma`, this repo): imports the greeting package and prints its output.
- **Spawned sub-issue**: the child issue in the library repo carrying the library half of the work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the library repo's default branch contains `HelloE2E202610211921513278` returning `e2e-cross-repo-spawn-20261021-192151-3278`, verifiable by text search of that repo.
- **SC-002**: A test in the library repo covering the function exists and passes.
- **SC-003**: Running this repo's program prints `e2e-cross-repo-spawn-20261021-192151-3278`, and building and testing the module succeeds.
- **SC-004**: Exactly one child sub-issue targeting the library repo is spawned from this issue, and this issue's implementation work starts only after that child is done.

## Assumptions

- The library repo is reachable by the engine and its board is servable by this Fabrik instance (cross-repo spawn board-servability), or it is initialised on demand as part of the scenario under test.
- The library repo's greeting package already exists and already follows the pattern of earlier `HelloE2E…` functions; this change adds one more in the same style.
- The consumer repo already imports the greeting package, so no new dependency is introduced — only a version update if required.
- The spec's literal values are authoritative and not to be reinterpreted (e.g. the date/time token in the name is not to be normalised).

## Out of Scope

- Changes to any function in the library other than adding the one specified.
- Removing or reordering existing output in this repo's `main.go`.
- Changes to Fabrik engine behaviour (this issue only exercises it).
- Releasing or tagging the library repo beyond what is needed for this repo to resolve the new function.
