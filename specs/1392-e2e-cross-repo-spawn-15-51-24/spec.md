# Feature Specification: E2E cross-repo spawn decomposition

**Feature Branch**: `fabrik/issue-1392`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: the original issue body, quoted verbatim below.

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610301951248021() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261030-195124-8021".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610301951248021() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two behaviours:

- **#803**: on-demand initialisation of a spawn-target repo.
- **#1308**: a per-run-unique fixture token, which lets the scenario's final assertion tell this run's output apart from earlier runs.

The work spans two repositories with a strict ordering dependency. The greeting function must exist in the beta repo (`handarbeit/fabrik-test-delta`) before this repo (`handarbeit/fabrik-test-gamma`) can import and call it. Because of that, the work cannot be done in one repository. The Plan stage should split it by spawning a sub-issue in the beta repo, and this issue's own work should be blocked until that sub-issue lands.

The function name and string literal contain a per-run-unique token that is already fixed. They must be used exactly as given, because the scenario's final assertion checks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the greeting function (Priority: P1)

A consumer of the beta repo's `pkg/greeting` package can call a new exported function that returns the run-unique string. This is the work the spawned sub-issue carries.

**Why this priority**: It is the prerequisite for everything else, and the scenario's final assertion checks that this exact text lands in the beta repo.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`. The new test passes, and calling `HelloE2E202610301951248021()` returns exactly `e2e-cross-repo-spawn-20261030-195124-8021`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting` package, **When** `HelloE2E202610301951248021()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261030-195124-8021`.
2. **Given** the beta repo's test suite, **When** it is run, **Then** a test in `pkg/greeting/greeting_test.go` covers the new function and passes.

---

### User Story 2 - This repo consumes and prints the beta function (Priority: P2)

Running this repo's program prints the string returned by the beta repo's new function, in addition to its existing output.

**Why this priority**: It completes the cross-repo change, but it depends on User Story 1 landing first.

**Independent Test**: After the beta change has landed, build and run this repo's program. Its output contains `e2e-cross-repo-spawn-20261030-195124-8021` alongside its previous output.

**Acceptance Scenarios**:

1. **Given** the beta function is available to this repo, **When** the program runs, **Then** it prints the result of `greeting.HelloE2E202610301951248021()` and keeps its existing behaviour.
2. **Given** the dependency set changed, **When** the module is tidied (if needed), **Then** the build succeeds with no missing or stale module requirements.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage recognises that the beta repo's change must land first and decomposes the work with a spawned sub-issue targeting the beta repo. This issue's own implementation waits for that sub-issue.

**Why this priority**: This is the behaviour under test. If the decomposition does not happen, the scenario fails.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` carrying the beta-side requirements, and this issue is blocked by it.

**Acceptance Scenarios**:

1. **Given** the spec describes changes in two repos with a strict order, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
2. **Given** the spawn block was emitted, **When** the engine processes it, **Then** a sub-issue is created in the beta repo, and this issue is blocked until that sub-issue is resolved.
3. **Given** the sub-issue's description, **When** it is read, **Then** it carries the exact function name and string literal from this spec.

---

### Edge Cases

- The beta repo may not yet be initialised for spawning (#803). The spawn must still succeed through on-demand initialisation.
- The function name and string literal must not be paraphrased, regenerated or substituted with a different token. Any deviation breaks the final assertion.
- If this repo's `go.mod` does not yet reference the beta module, or references an older version, the dependency must be resolved so the build succeeds once the beta change has landed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610301951248021() string` in its `pkg/greeting` package (`pkg/greeting/greeting.go`).
- **FR-002**: `HelloE2E202610301951248021()` MUST return exactly the string `e2e-cross-repo-spawn-20261030-195124-8021`.
- **FR-003**: The beta repo MUST include a test for the new function in `pkg/greeting/greeting_test.go` that asserts the exact returned string and passes.
- **FR-004**: This repo (`handarbeit/fabrik-test-gamma`) MUST update `main.go` to call `greeting.HelloE2E202610301951248021()` and print its result, while preserving the program's existing behaviour.
- **FR-005**: This repo's module dependencies MUST be tidied (`go mod tidy`) if the change requires it, so the build is clean.
- **FR-006**: The beta repo change MUST land before this repo's change, since this repo cannot import a function that does not yet exist.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo. This is the mechanism by which the beta change is split off.
- **FR-008**: The spawned sub-issue MUST reproduce the function name and string literal exactly as given in FR-001 and FR-002, with no invented substitute.

### Key Entities

- **Alpha repo** (`handarbeit/fabrik-test-gamma`): this repo, the consumer. It holds the parent issue and the `main.go` change.
- **Beta repo** (`handarbeit/fabrik-test-delta`): the provider repo, which receives the spawned sub-issue and the new function.
- **Run-unique token**: the fixed value `20261030-195124-8021`. It appears in the function name (as `20261030195124` followed by `8021`) and in the returned literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the beta repo's default branch contains `HelloE2E202610301951248021` returning `e2e-cross-repo-spawn-20261030-195124-8021`, and its test suite passes.
- **SC-002**: After the pipeline completes, this repo's default branch contains `main.go` calling `greeting.HelloE2E202610301951248021()`, the repo builds successfully, and the program's output includes the literal.
- **SC-003**: Exactly one sub-issue is spawned in the beta repo for the beta-side work, and it is created before this repo's implementation lands.
- **SC-004**: The scenario's final assertion finds the exact run-unique text in the beta repo.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, or the spawned sub-issue creates them. Either way the file locations in FR-001 and FR-003 apply.
- This repo's `main.go` already imports, or can import, the beta repo's `pkg/greeting` package, and the module wiring needed for that is part of FR-005. This spec does not decide how it is wired.
- The `fabrik:yolo` label stays on the issue, so the pipeline advances without manual approval.
- A single spawned sub-issue is enough to cover all the beta-side work.

## Out of Scope

- Changes to the engine's spawn logic itself. This issue only exercises it.
- Any change to the beta repo beyond the single function and its test.
- Any behaviour change to this repo's existing output beyond printing the additional line.
- Generating or rotating the run-unique token. It is fixed in this spec.

## Source References

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
