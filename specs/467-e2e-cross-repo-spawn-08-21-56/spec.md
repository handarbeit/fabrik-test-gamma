# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-467`
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
> - Add an exported function HelloE2E202610191221560514() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261019-122156-0514".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610191221560514() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. A single piece of work needs changes in two repositories with a hard ordering dependency: the greeting function must exist in `handarbeit/fabrik-test-delta` (the "beta" repo) before `handarbeit/fabrik-test-gamma` (this repo, called "alpha" in the original request) can import and call it. The scenario checks that planning decomposes the work by spawning a child issue in the beta repo, that the child lands first, and that the dependent change here then lands on top of it.

It is a regression test for two earlier pieces of work: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token). The function name and string literal contain that per-run-unique token. The scenario's final assertion checks that exact text in the beta repo, so the identifiers must be used verbatim and never altered or replaced.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo gains an exported greeting function returning the per-run-unique literal, with a test covering it. This work is delivered through a child issue spawned in the beta repo.

**Why this priority**: It is the prerequisite for everything else. This repo cannot compile against a function that does not exist, and the scenario's final assertion is made against this repo.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes and confirms the function returns the expected literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610191221560514()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-122156-0514`.
2. **Given** the beta repo after the change, **When** its test suite runs, **Then** the new test for this function exists and passes.
3. **Given** the change is delivered, **When** the beta repo's `pkg/greeting/greeting.go` is inspected, **Then** it contains the function name and string literal verbatim as specified.

---

### User Story 2 - This repo consumes the new function (Priority: P2)

This repo's `main.go` calls `greeting.HelloE2E202610191221560514()` in addition to its existing behaviour, and prints the result. It depends on User Story 1 having landed.

**Why this priority**: It completes the cross-repo change, but it can only be done and verified after the beta repo change exists.

**Independent Test**: Build and run this repo after the beta change is available. The program output includes the literal `e2e-cross-repo-spawn-20261019-122156-0514`, and existing output is preserved.

**Acceptance Scenarios**:

1. **Given** the beta repo change has landed, **When** this repo is built, **Then** it compiles with the new import usage and module dependencies are consistent.
2. **Given** this repo's program runs, **When** it executes, **Then** it prints the result of `greeting.HelloE2E202610191221560514()` and still does whatever it did before.

---

### User Story 3 - Decomposition happens via a cross-repo spawn (Priority: P1)

The beta-repo work is split out from this issue by Plan as a spawned child issue targeting the beta repo, rather than being done inline or by hand. The spawn ordering makes the beta child land before this issue's change.

**Why this priority**: Exercising the spawn path (including on-demand initialisation of the spawn-target repo, #803) is the purpose of the scenario.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` carrying the beta-repo requirements, and this issue is blocked by it until it completes.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it declares a child spawn block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the child is spawned, **When** this issue proceeds, **Then** its implementation is ordered after the child's completion.
3. **Given** the child's requirements, **When** inspected, **Then** they carry the exact function name and string literal from this spec.

---

### Edge Cases

- The spawn-target repo has not been initialised on demand yet (the #803 situation): the spawn must still succeed.
- The function name or literal is altered during planning, spawning or implementation: the scenario's final assertion would fail, so exact preservation is required at every step.
- This repo's change is attempted before the beta change exists: it cannot compile, which is why ordering is enforced.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610191221560514() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261019-122156-0514`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610191221560514()` in addition to its existing behaviour and print the result.
- **FR-005**: This repo's module dependencies MUST be consistent after the change (`go mod tidy` run if needed).
- **FR-006**: The beta repo change MUST land before this repo's change.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-008**: The function name and string literal MUST be used verbatim as given; no substitute identifier or token may be invented.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the scenario completes, `pkg/greeting/greeting.go` in the beta repo contains `HelloE2E202610191221560514` and the literal `e2e-cross-repo-spawn-20261019-122156-0514`.
- **SC-002**: The beta repo's tests, including the new one, pass.
- **SC-003**: This repo builds and its program output includes `e2e-cross-repo-spawn-20261019-122156-0514`.
- **SC-004**: Exactly one child issue exists in `handarbeit/fabrik-test-delta` for the beta-repo work, and it completed before this issue's change landed.

## Assumptions

- "Alpha" in the original request refers to this repo (`handarbeit/fabrik-test-gamma`); "beta" refers to `handarbeit/fabrik-test-delta`.
- Existing behaviour of `main.go` is retained; only the new call and print are added.
- The beta repo already has `pkg/greeting/greeting.go` and `greeting_test.go` (or the spawn-target repo is initialised on demand per #803).
- The `fabrik:yolo` label means the pipeline advances without human approval of this spec.

## Out of Scope

- Any change to the greeting package other than the one new function and its test.
- Changes to Fabrik engine behaviour; this issue only exercises it.
- Additional repos or additional spawned children beyond the single beta-repo child.

## Source References

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
