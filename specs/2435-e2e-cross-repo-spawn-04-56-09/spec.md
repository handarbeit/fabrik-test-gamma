# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2435`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611170956090573() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261117-095609-0573\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611170956090573() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. Some work spans two repositories with a strict ordering dependency: the library repo (`handarbeit/fabrik-test-delta`, called "beta" in the request) must gain a new function before the application repo (`handarbeit/fabrik-test-gamma`, this repo, called "alpha" in the request) can call it.

The scenario exists to confirm two behaviours:

- On-demand initialisation of a spawn-target repo (#803).
- Per-run-unique fixture tokens (#1308). The function name and string literal embed a token unique to this run, so the final assertion can prove the exact text landed in the beta repo rather than coming from a previous run.

The pipeline must decompose the work by spawning a child issue in the beta repo that lands first. This issue (the parent) then lands its own change in this repo, which depends on the child's change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the unique greeting function (Priority: P1)

A child work item in the beta repo adds the new exported greeting function and a test for it, so the function exists for the application repo to consume.

**Why this priority**: The application-repo change cannot compile without it, and the scenario's final assertion checks this exact text in the beta repo.

**Independent Test**: In the beta repo, call `HelloE2E202611170956090573()` from its test and confirm it returns `e2e-cross-repo-spawn-20261117-095609-0573`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** the child change lands, **Then** an exported function `HelloE2E202611170956090573() string` exists in `pkg/greeting/greeting.go` and returns exactly `e2e-cross-repo-spawn-20261117-095609-0573`.
2. **Given** the child change, **When** the beta repo's tests run, **Then** a test in `pkg/greeting/greeting_test.go` covers the new function and passes.

---

### User Story 2 - This repo calls the new function and prints the result (Priority: P1)

After the beta change has landed, this repo's `main.go` calls `greeting.HelloE2E202611170956090573()` and prints the result, in addition to its existing behaviour.

**Why this priority**: This is the parent's own deliverable and proves the ordering constraint was respected: it can only build once the beta function exists.

**Independent Test**: Build and run this repo's program with the beta change available and confirm the output contains `e2e-cross-repo-spawn-20261117-095609-0573` alongside its previous output.

**Acceptance Scenarios**:

1. **Given** the beta function is available, **When** this repo's program runs, **Then** it prints the string returned by `greeting.HelloE2E202611170956090573()`.
2. **Given** the change, **When** the program runs, **Then** its pre-existing output is still produced.
3. **Given** the dependency set changed, **When** the module is tidied, **Then** the module files are consistent (`go mod tidy` leaves no further changes).

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage recognises that the work spans two repos with a strict ordering and decomposes it by spawning a child issue targeting the beta repo, blocking this parent until the child is done.

**Why this priority**: Verifying decomposition is the purpose of the scenario.

**Independent Test**: After Plan runs, confirm a child issue exists in `handarbeit/fabrik-test-delta`, is on the board, and blocks this issue.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output declares a spawn of one child issue targeting `handarbeit/fabrik-test-delta` that carries the beta-repo work.
2. **Given** the child exists, **When** this issue's Implement stage would start, **Then** it waits until the child has landed.
3. **Given** the beta repo has not been used as a spawn target before, **When** the child is spawned, **Then** the target repo is initialised on demand (#803).

---

### Edge Cases

- The function name and string literal must be used exactly as given; an invented variant must not be accepted, since the final assertion depends on the exact text.
- If the beta repo is not yet set up as a spawn target, spawning must still succeed (on-demand initialisation).
- This repo's change must not be merged ahead of the beta change; it would not compile without it.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST gain an exported function `HelloE2E202611170956090573() string` in `pkg/greeting/greeting.go` returning the literal `e2e-cross-repo-spawn-20261117-095609-0573`.
- **FR-002**: The beta repo MUST have a test for that function in `pkg/greeting/greeting_test.go` that passes.
- **FR-003**: This repo's `main.go` MUST call `greeting.HelloE2E202611170956090573()` and print the result, in addition to its existing behaviour.
- **FR-004**: This repo's module files MUST be left consistent (run `go mod tidy` if needed).
- **FR-005**: The beta-repo change MUST land before this repo's change.
- **FR-006**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
- **FR-007**: The function name and string literal MUST be used verbatim as given in this spec, with no variants.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the text `e2e-cross-repo-spawn-20261117-095609-0573` in the greeting package, which is the scenario's final assertion.
- **SC-002**: Exactly one child issue is spawned in the beta repo, and it lands before this issue's change does.
- **SC-003**: This repo builds and its program output includes the unique string after both changes have landed.
- **SC-004**: Tests pass in both repos.

## Assumptions

- The beta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`; if not, the child creates them.
- This repo's `main.go` already imports or can import the beta repo's greeting package; adding that dependency is part of the change.
- "Alpha" in the original request means this repo (`handarbeit/fabrik-test-gamma`), and "beta" means `handarbeit/fabrik-test-delta`.
- The unique token is fixed for this run, so no variation is expected.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the one new function and its test.
- Any change to this repo's existing behaviour beyond calling and printing the new function.
- Changes to the pipeline, the engine, or the e2e harness itself.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation.
- #1308: per-run-unique fixture token.
