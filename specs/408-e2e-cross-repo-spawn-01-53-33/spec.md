# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-408`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610190553330025() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261019-055333-0025".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610190553330025() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. Some work spans two repositories with a strict ordering: a dependency repo (`handarbeit/fabrik-test-delta`, the "beta" repo) must gain a capability before the repo being worked on (`handarbeit/fabrik-test-gamma`, this repo, the "alpha" repo) can use it. The pipeline is expected to recognise this, split the work, and spawn a child issue in the beta repo that lands first.

The scenario guards two earlier fixes:
- **#803**: on-demand initialisation of a spawn-target repo that Fabrik has not yet worked in.
- **#1308**: the per-run-unique fixture token. The function name and string literal in this issue carry a token unique to this run, and the scenario's final assertion checks that exact text in the beta repo. Any deviation from the given names and literals fails the scenario.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

A child issue in the beta repo adds an exported greeting function returning a fixed, run-unique string, together with a test, so the alpha repo has something to import.

**Why this priority**: Everything else depends on it; the alpha change cannot compile without it, and the scenario's final assertion inspects this text in the beta repo.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting` and confirm that the new function exists, is exported and returns exactly `e2e-cross-repo-spawn-20261019-055333-0025`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the function, **When** the child issue's work is merged, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202610190553330025() string` returning exactly `e2e-cross-repo-spawn-20261019-055333-0025`.
2. **Given** the merged change, **When** the beta repo's tests run, **Then** a test in `pkg/greeting/greeting_test.go` covers the new function and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

Once the beta change has landed, this repo's `main.go` calls the new function and prints its result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo change, but it can only land after User Story 1.

**Independent Test**: Build and run this repo and confirm the output includes `e2e-cross-repo-spawn-20261019-055333-0025` alongside the previous output.

**Acceptance Scenarios**:

1. **Given** the beta function is available to this repo, **When** `main.go` runs, **Then** it prints the string returned by `greeting.HelloE2E202610190553330025()` and still does everything it did before.
2. **Given** the dependency was updated, **When** the module is tidied, **Then** the module files are consistent and the repo builds.

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The Plan stage decomposes the work by spawning a child issue in the beta repo, and the parent's own changes are sequenced after the child's.

**Why this priority**: This is the behaviour under test (#803). Without the spawn, the scenario does not exercise cross-repo decomposition.

**Independent Test**: After Plan, confirm that a child issue exists in `handarbeit/fabrik-test-delta`, that it is registered as a blocker of this issue, and that this issue does not proceed to implementation until the child is resolved.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` is emitted and a child issue is created there.
2. **Given** the beta repo has not been worked on by Fabrik before, **When** the spawn happens, **Then** the target repo is initialised on demand and the child proceeds.
3. **Given** the child is unresolved, **When** the pipeline advances this issue, **Then** the alpha-side change does not land before the beta-side change.

---

### Edge Cases

- The function name and string literal must be used verbatim; no renamed, altered or regenerated token is acceptable.
- If the beta repo has not been initialised for Fabrik, the spawn must still succeed (#803).
- If no tidy step is needed for the module files, none is required.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610190553330025() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261019-055333-0025`.
- **FR-003**: The beta repo MUST include a test for the function in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610190553330025()` and print the result, in addition to its existing behaviour.
- **FR-005**: This repo's module files MUST be left consistent (tidied if needed) so the repo builds with the new dependency.
- **FR-006**: The beta-repo change MUST land before the alpha-repo change.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-008**: The exact function name and string literal in FR-001 and FR-002 MUST be used verbatim in all artifacts (child issue, code, tests).

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, the repo this issue lives in; it consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-delta`, the dependency repo that provides the greeting function.
- **Child issue**: The issue spawned in the beta repo, blocking this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in the beta repo and was created by the pipeline's Plan-stage decomposition rather than manually.
- **SC-002**: The beta repo's default branch contains `HelloE2E202610190553330025` returning exactly `e2e-cross-repo-spawn-20261019-055333-0025`, and its tests pass.
- **SC-003**: The alpha repo's default branch builds and its program output includes `e2e-cross-repo-spawn-20261019-055333-0025`.
- **SC-004**: The beta-repo change is merged before the alpha-repo change.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, or the child issue creates them if absent.
- This repo's `main.go` already exists and already imports, or can import, the beta repo's greeting package.
- The `fabrik:yolo` label means the pipeline advances through stages without waiting for human approval.
- The token in the function name and literal is final for this run and must not be regenerated.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single function and its test.
- Any change to this repo's behaviour beyond calling and printing the new function.
- Changes to Fabrik's engine, spawn logic or e2e harness; this issue only exercises them.

## Source References *(optional)*

- #803 (on-demand spawn-target repo init)
- #1308 (per-run-unique fixture token)
