# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-100`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610150650002855() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261015-065000-2855\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610150650002855() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier behaviours:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: a per-run-unique fixture token, so that each run's final assertion can tell its own output from earlier runs.

The change is deliberately small but spans two repositories with a strict ordering dependency. This repo (gamma, "alpha" in the original wording) cannot call a greeting function that the delta repo ("beta") does not yet provide. The pipeline therefore has to split the work: the delta change is spawned as a sub-issue that must land first, and the change here follows once it has.

The function name and the returned string contain a per-run-unique token. They must be used exactly as given, because the scenario's final assertion checks for that exact text in the delta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provide the greeting function in the delta repo (Priority: P1)

The delta repo gains a new exported greeting function that returns a fixed, run-unique string, with a test that covers it. This work is delegated to a spawned sub-issue in `handarbeit/fabrik-test-delta`.

**Why this priority**: Everything else depends on it. The gamma change cannot compile until this function exists, and the scenario's final assertion checks that this text lands in the delta repo.

**Independent Test**: In the delta repo, the new function exists, returns the exact literal, and its unit test passes.

**Acceptance Scenarios**:

1. **Given** the delta repo's greeting package, **When** `HelloE2E202610150650002855()` is called, **Then** it returns `e2e-cross-repo-spawn-20261015-065000-2855`.
2. **Given** the delta repo's greeting tests, **When** the test suite runs, **Then** a test for the new function is present and passes.
3. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a spawn-child block targeting `handarbeit/fabrik-test-delta`, and a sub-issue is created there.

---

### User Story 2 - Use the new function from the gamma repo (Priority: P2)

After the delta change has landed, the gamma application calls the new greeting function in addition to its existing behaviour and prints the result.

**Why this priority**: It depends on Story 1, so it must come second. It completes the cross-repo flow the scenario exists to verify.

**Independent Test**: With the delta change available, the gamma application builds and its output includes the literal `e2e-cross-repo-spawn-20261015-065000-2855`.

**Acceptance Scenarios**:

1. **Given** the delta function is available to this repo, **When** the application runs, **Then** it prints the string returned by `HelloE2E202610150650002855()` and still performs its existing behaviour.
2. **Given** the module dependencies, **When** the project is built, **Then** the build succeeds, with dependencies tidied if that is needed.

---

### User Story 3 - Strict ordering between repos (Priority: P1)

The delta work is sequenced to complete before the gamma work begins, so the gamma change is never attempted against a delta repo that lacks the function.

**Why this priority**: The ordering constraint is the point of the scenario. Without it, cross-repo decomposition has not been verified.

**Independent Test**: The spawned delta sub-issue is linked as a blocker of this issue, and this issue's implementation does not proceed until the sub-issue is done.

**Acceptance Scenarios**:

1. **Given** the delta sub-issue has been spawned, **When** it is not yet complete, **Then** this issue is blocked on it.
2. **Given** the delta sub-issue has landed, **When** this issue proceeds, **Then** the gamma change builds against the delta function.

---

### Edge Cases

- The delta repo has not yet been initialised for spawning. On-demand spawn-target initialisation (#803) must handle this without manual setup.
- Any function name or string other than the given ones is wrong. The final assertion fails if the literal text is not found in the delta repo.
- If `go mod tidy` changes nothing, no module-file change is required.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST expose an exported function `HelloE2E202610150650002855() string` in its greeting package.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261015-065000-2855`.
- **FR-003**: The delta repo's greeting tests MUST include a test for that function.
- **FR-004**: The gamma application (this repo) MUST call `greeting.HelloE2E202610150650002855()` in addition to its existing behaviour and print the result.
- **FR-005**: Module dependencies in this repo MUST be tidied if the change requires it.
- **FR-006**: The Plan stage MUST decompose the work by emitting a spawn-child block targeting `handarbeit/fabrik-test-delta`, rather than implementing the delta change in this repo.
- **FR-007**: The delta change MUST land before the gamma change is implemented against it.
- **FR-008**: The exact function name and string literal above MUST be used verbatim; no substitute names or values may be introduced.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A spawned sub-issue exists in `handarbeit/fabrik-test-delta` for the delta change, created by Plan's cross-repo decomposition.
- **SC-002**: The delta repo contains the literal `e2e-cross-repo-spawn-20261015-065000-2855` and the function `HelloE2E202610150650002855` after the sub-issue lands, which is what the scenario's final assertion checks.
- **SC-003**: The delta repo's tests pass, including the new one.
- **SC-004**: The gamma application builds and prints the new greeting alongside its prior output.
- **SC-005**: The delta change is merged before the gamma change, with no failed or reverted gamma attempt caused by a missing function.

## Assumptions

- The per-run-unique token in the function name and string is already final and must not be regenerated or altered.
- The delta repo already has a greeting package with an existing test file, as the issue describes.
- The `fabrik:yolo` label means the pipeline advances automatically, so no human approval is expected between stages.
- "Alpha" and "beta" in the original wording mean this repo (gamma) and the delta repo respectively.

## Out of Scope

- Any change to the greeting package other than the one new function and its test.
- Changes to the gamma application's existing behaviour beyond calling and printing the new function.
- Changes to Fabrik's engine, spawn logic or stage configuration. This issue only exercises them.
- Decomposition into more than the single delta sub-issue.

## Source References

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
