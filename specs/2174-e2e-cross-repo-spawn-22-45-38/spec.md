# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2174`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611170345385247() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261117-034538-5247\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611170345385247() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two behaviours:

- **#803**: on-demand initialisation of a spawn-target repo.
- **#1308**: a per-run-unique fixture token, so each run's final assertion checks text that only this run could have produced.

The work spans two repos with a strict ordering dependency. The consuming repo (this one, "alpha", `handarbeit/fabrik-test-gamma`) can only call a greeting function once the providing repo ("beta", `handarbeit/fabrik-test-delta`) has it. Planning must therefore split the work into a child issue in the beta repo, which lands first, and the remaining change in this repo.

The function name and string literal embed a per-run-unique token. They must be used exactly as written. The scenario's final assertion checks that this exact text appears in the beta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo gains an exported function that returns the per-run-unique string, with a test, delivered through a child issue spawned from this one.

**Why this priority**: It is the prerequisite for everything else and the artifact the scenario's final assertion inspects.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes, and the greeting source contains the exact function name and literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611170345385247()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261117-034538-5247`.
2. **Given** the new function, **When** the beta repo's test suite runs, **Then** a test covering the function passes.
3. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** a child issue targeting the beta repo is created, and that child's change is what lands in the beta repo.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P1)

After the beta change has landed, this repo's entry point calls the new function and prints the result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo outcome and proves the ordering constraint (beta first, then alpha).

**Independent Test**: Build and run this repo's program. Its output includes `e2e-cross-repo-spawn-20261117-034538-5247` along with its prior output.

**Acceptance Scenarios**:

1. **Given** the beta change is available, **When** this repo's program runs, **Then** it prints the string returned by `greeting.HelloE2E202611170345385247()`.
2. **Given** the program's previous behaviour, **When** the change is applied, **Then** that behaviour is preserved (the new output is additive).
3. **Given** the change, **When** the module's dependencies are checked, **Then** they are consistent (`go mod tidy` run if needed) and the project builds.

---

### Edge Cases

- The beta repo must be ready before alpha depends on it: alpha's change must not be built or merged against a beta that lacks the function.
- The function name and literal must match the issue text exactly. Any variation (renamed function, altered token) breaks the final assertion.
- The spawn target repo may not yet be initialised for Fabrik (the #803 on-demand init case). Spawning must still succeed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The work is decomposed so the beta repo change lands before the alpha repo change. Plan creates a child issue in `handarbeit/fabrik-test-delta` for the beta-side work.
- **FR-002**: The beta repo's greeting package exports `HelloE2E202611170345385247() string`, returning exactly `e2e-cross-repo-spawn-20261117-034538-5247`.
- **FR-003**: The beta repo has a test for that function in the greeting package's test file.
- **FR-004**: This repo's `main.go` calls `greeting.HelloE2E202611170345385247()` and prints the result, in addition to its existing behaviour.
- **FR-005**: This repo's module dependencies are tidied (`go mod tidy`) if the change requires it.
- **FR-006**: The function name and string literal are used verbatim as given. No alternate token is introduced.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo. It hosts the parent issue and consumes the function.
- **Beta repo**: `handarbeit/fabrik-test-delta`. It provides the function and receives the spawned child issue.
- **Per-run-unique token**: the substring embedded in the function name and literal that makes the final assertion specific to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting the beta repo exists as a result of decomposing this issue.
- **SC-002**: The beta repo's default branch contains the exact function name and the exact literal `e2e-cross-repo-spawn-20261117-034538-5247`, with a passing test.
- **SC-003**: Running this repo's program prints the literal above in addition to its previous output.
- **SC-004**: The beta change is merged before the alpha change.

## Assumptions

- The beta repo already has a greeting package with a source file and a test file, as the issue describes. Only the addition of the function and its test is needed there.
- Tool and dependency resolution for alpha to see the beta change (for example a version bump or local replacement) is a planning and implementation concern, not part of this spec.
- The existing behaviour of this repo's `main.go` is unspecified here and must simply be preserved.

## Out of Scope *(optional)*

- Changes to Fabrik's engine, spawn logic or e2e harness. This issue is a fixture for them, not a change to them.
- Any other change to either repo beyond the function, its test and the call/print in `main.go`.
- Alternative function names or tokens.
