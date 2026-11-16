# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1989`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order. In handarbeit/fabrik-test-delta (the beta repo): add an exported function HelloE2E202611161354261489() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261116-135426-1489", and add a test for it in pkg/greeting/greeting_test.go. In handarbeit/fabrik-test-gamma (this repo): update main.go to call greeting.HelloE2E202611161354261489() (in addition to whatever it already does) and print the result, and run go mod tidy if needed. Constraint: beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo. This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. It checks that when work spans two repositories with a strict ordering dependency, the pipeline splits the work into a sub-issue in the second repository and sequences the two correctly.

It covers two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: a per-run-unique fixture token, so that each run's final assertion can identify the text that run produced.

The change is split across two repositories:

- `handarbeit/fabrik-test-delta` ("beta") supplies a new greeting function.
- `handarbeit/fabrik-test-gamma` ("alpha", this repo) consumes it.

The consumer cannot compile or import the function until it exists in beta. Beta's change must therefore land first, and the work must be decomposed in that order.

The function name and string literal contain a per-run-unique token. They must be used exactly as given, because the scenario's final assertion checks for that exact text in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the unique greeting (Priority: P1)

The beta repo (`handarbeit/fabrik-test-delta`) gains an exported greeting function with the run-unique name, returning the run-unique literal, and a test covering it.

**Why this priority**: The consumer change depends on it, and the scenario's final assertion checks that this text lands in beta.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes. The function returns exactly `e2e-cross-repo-spawn-20261116-135426-1489`.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting` package, **When** `HelloE2E202611161354261489()` is called, **Then** it returns the string `e2e-cross-repo-spawn-20261116-135426-1489`.
2. **Given** the beta repo's test suite, **When** it runs, **Then** a test in `pkg/greeting/greeting_test.go` asserts that return value and passes.

---

### User Story 2 - Alpha repo consumes and prints the greeting (Priority: P2)

This repo's `main.go` calls `greeting.HelloE2E202611161354261489()` from the beta repo's package, in addition to its existing behaviour, and prints the result.

**Why this priority**: It depends on Story 1 and exercises the cross-repo ordering constraint.

**Independent Test**: With beta's change available, build and run this repo. The output includes `e2e-cross-repo-spawn-20261116-135426-1489`, and the module's dependencies are consistent (`go mod tidy` leaves no diff).

**Acceptance Scenarios**:

1. **Given** beta's function exists, **When** this repo's program runs, **Then** it prints the returned literal in addition to its prior output.
2. **Given** the dependency on the beta repo's package, **When** `go mod tidy` is run, **Then** the module files are consistent and the build succeeds.

---

### User Story 3 - Work is decomposed across repos in the required order (Priority: P1)

The Plan stage decomposes the work so that the beta-repo change is a spawned sub-issue in `handarbeit/fabrik-test-delta`, and this issue's own work waits for it.

**Why this priority**: Verifying this decomposition is the purpose of the scenario.

**Independent Test**: After Plan runs, a sub-issue exists in the beta repo for the beta change, and this issue's implementation does not proceed until that sub-issue has landed.

**Acceptance Scenarios**:

1. **Given** the work needs changes in two repos in strict order, **When** Plan completes, **Then** a spawn block targeting `handarbeit/fabrik-test-delta` is emitted and a sub-issue is created there.
2. **Given** the sub-issue has not yet landed, **When** this issue would proceed to implementation, **Then** it is held until the beta change has landed.
3. **Given** the beta sub-issue has landed, **When** this issue completes, **Then** the exact token text is present in the beta repo.

---

### Edge Cases

- The spawn-target repo may not yet be initialised for Fabrik (the #803 scenario). On-demand initialisation must allow the spawn to proceed.
- The function name and literal must be copied verbatim. Any variation fails the scenario's final assertion (the #1308 scenario).
- If `go mod tidy` has nothing to change, the step is a no-op and not a failure.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export `HelloE2E202611161354261489() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202611161354261489()` MUST return exactly the literal `e2e-cross-repo-spawn-20261116-135426-1489`.
- **FR-003**: The beta repo MUST include a test in `pkg/greeting/greeting_test.go` that verifies FR-002.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202611161354261489()` and print the result, while keeping its existing behaviour.
- **FR-005**: This repo's module files MUST be consistent after the change (run `go mod tidy` if needed).
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, so the beta change is created as a sub-issue there.
- **FR-007**: The beta change MUST land before this repo's change that depends on it.
- **FR-008**: The function name and string literal MUST be used verbatim, with no substitution of other tokens.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo contains both the exact function name `HelloE2E202611161354261489` and the exact literal `e2e-cross-repo-spawn-20261116-135426-1489`.
- **SC-002**: The beta repo's greeting tests pass, including the new one.
- **SC-003**: This repo builds and its program output includes the literal.
- **SC-004**: Exactly one sub-issue is spawned in the beta repo for this work, and it lands before this repo's change merges.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, or on-demand initialisation provides them.
- This repo already has a `main.go` that can import the beta repo's `pkg/greeting` package.
- Because the issue carries `fabrik:yolo`, the pipeline advances automatically once the spec is complete.
- The per-run token in the issue text is already unique for this run and needs no regeneration.

## Out of Scope

- Any greeting behaviour beyond the single function and its test.
- Changes to other packages or files in either repo, beyond what the import and `go mod tidy` require.
- Changes to the engine's spawn or decomposition logic. This scenario verifies existing behaviour only.
