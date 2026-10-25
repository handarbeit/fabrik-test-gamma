# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-788`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610250615015471() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261025-061501-5471\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610250615015471() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work spans two repositories with a strict ordering dependency: the greeting function must exist in the beta repo (`handarbeit/fabrik-test-delta`) before this repo (`handarbeit/fabrik-test-gamma`, the "alpha" repo) can call it. A correct pipeline run must therefore split the work, with the beta-repo change spawned as a sub-issue by Plan and landed first, then the alpha-repo change completed on top of it.

The scenario guards two earlier fixes:

- **#803** — on-demand initialisation of a spawn-target repo.
- **#1308** — a per-run-unique fixture token, so the final assertion can tell this run's output from any other run's.

The function name and string literal contain that per-run-unique token. They must be used exactly as written, because the scenario's final assertion checks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique string, with a test covering it. This is the artifact the whole scenario's final assertion checks.

**Why this priority**: The scenario's pass/fail assertion depends on this exact text landing in the beta repo, and the alpha-repo change cannot compile without it.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes, and the function returns `e2e-cross-repo-spawn-20261025-061501-5471`.

**Acceptance Scenarios**:

1. **Given** the beta repo without the new function, **When** the spawned sub-issue is implemented and merged, **Then** the beta repo exports `HelloE2E202610250615015471() string` returning exactly `e2e-cross-repo-spawn-20261025-061501-5471`.
2. **Given** the new function exists, **When** the beta repo's tests run, **Then** a test for that function exists and passes.

---

### User Story 2 - Alpha repo consumes the function and prints the result (Priority: P2)

The alpha repo's entry point calls the new beta-repo function in addition to its existing behaviour and prints the returned string.

**Why this priority**: It completes the cross-repo change, but it depends on User Story 1 landing first.

**Independent Test**: With the beta change available, build and run this repo's program. It succeeds and its output includes `e2e-cross-repo-spawn-20261025-061501-5471`.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** the alpha repo's entry point is updated, **Then** it calls `greeting.HelloE2E202610250615015471()` and prints the result, while still doing everything it did before.
2. **Given** the dependency set changed, **When** the module's dependencies are tidied if needed, **Then** the module builds cleanly.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

Plan recognises that the work crosses two repos with an ordering constraint and decomposes it by spawning a sub-issue in the beta repo, rather than attempting both changes in one repo.

**Why this priority**: This is the behaviour the scenario exists to verify. Without it, the alpha change would reference a function that does not exist.

**Independent Test**: After Plan completes, a sub-issue exists in `handarbeit/fabrik-test-delta` describing the beta-repo change, and this issue is blocked until that sub-issue lands.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo carrying the beta-repo requirements.
2. **Given** the sub-issue was spawned, **When** the pipeline proceeds, **Then** the beta change lands before the alpha change is implemented.

---

### Edge Cases

- The beta repo may not yet be set up for spawning; on-demand initialisation of the spawn target (#803) must handle this.
- The function name or literal must not be altered, abbreviated, or replaced with a freshly generated token; any variation fails the final assertion.
- If the alpha change were attempted before the beta change landed, the alpha repo would fail to build; ordering must prevent this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function named exactly `HelloE2E202610250615015471` that takes no arguments and returns a `string`, located in its `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261025-061501-5471`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: The alpha repo (`handarbeit/fabrik-test-gamma`) MUST update `main.go` to call `greeting.HelloE2E202610250615015471()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha repo's dependency metadata MUST be tidied (`go mod tidy`) if the change requires it, so the module builds.
- **FR-006**: The beta-repo change MUST land before the alpha-repo change is implemented or merged.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this spec. No substitute names or tokens may be introduced.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo; the consumer of the new function.
- **Beta repo**: `handarbeit/fabrik-test-delta`; the provider of the new function and the target of the spawned sub-issue.
- **Run-unique token**: the `20261025-061501-5471` fragment (in the literal) and its compact form in the function name; it identifies this run's output.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the exact function name and exact string literal given in this spec.
- **SC-002**: A sub-issue targeting the beta repo was created by Plan's spawn mechanism, and the beta change landed before the alpha change.
- **SC-003**: The alpha repo builds and runs, and its output includes `e2e-cross-repo-spawn-20261025-061501-5471`.
- **SC-004**: The beta repo's greeting tests pass, including the new test.

## Assumptions

- The beta repo has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, and the alpha repo's module already depends on, or can be tidied to depend on, the beta repo's module.
- The spawned sub-issue is created, tracked and landed through Fabrik's normal cross-repo spawn mechanism, so no manual cross-repo coordination is part of this issue.
- The alpha repo's existing `main.go` behaviour is preserved; the new call is purely additive.

## Out of Scope *(optional)*

- Changes to Fabrik's engine, spawn mechanism or on-demand repo initialisation (#803) and fixture-token handling (#1308); this issue only exercises them.
- Any other greeting functions, or any change to the beta repo beyond the single function and its test.
- Any other behaviour change to the alpha repo's `main.go`.
- Handling of failure modes in the spawn mechanism itself.

## Source References *(optional)*

- #803 — on-demand spawn-target repo init
- #1308 — per-run-unique fixture token
