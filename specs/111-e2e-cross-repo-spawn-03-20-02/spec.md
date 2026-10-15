# Feature Specification: e2e cross-repo spawn (03:20:02)

**Feature Branch**: `fabrik/issue-111`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610150720020653() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261015-072002-0653\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610150720020653() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It exercises two earlier fixes:

- #803: on-demand initialisation of a spawn-target repo.
- #1308: a per-run-unique fixture token, so each run's final assertion can tell its own output from earlier runs.

The work spans two repositories in a strict dependency order. The beta repo (`handarbeit/fabrik-test-delta`) must gain a new exported greeting function before the alpha repo (`handarbeit/fabrik-test-gamma`, this repo) can call it. Plan is therefore expected to decompose the work by spawning a child issue in the beta repo. The scenario's final assertion checks that the exact function name and string literal below land in the beta repo, so they must be used verbatim and never substituted.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns the per-run-unique literal string, covered by a test. This work is delivered by a child issue spawned in the beta repo, not by this issue's own PR.

**Why this priority**: Everything else depends on it. The alpha repo cannot reference a function that does not exist, and the scenario's final assertion inspects the beta repo.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610150720020653()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261015-072002-0653`.
2. **Given** the beta repo's test suite, **When** it is run, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Alpha repo calls and prints the new function (Priority: P2)

After the beta change has landed, this repo's entry point calls the new greeting function in addition to its existing behaviour and prints the result.

**Why this priority**: It depends on User Story 1 and completes the cross-repo flow, showing the dependency order is respected.

**Independent Test**: Build and run this repo's program. Its output contains the literal `e2e-cross-repo-spawn-20261015-072002-0653`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta change is available to this repo, **When** the program runs, **Then** it prints the string returned by `greeting.HelloE2E202610150720020653()`.
2. **Given** the program's behaviour before this change, **When** it runs after this change, **Then** its previous output is still produced.
3. **Given** the dependency set after the change, **When** the module is built, **Then** it resolves without errors (running `go mod tidy` if needed).

---

### User Story 3 - Decomposition happens across repos in order (Priority: P1)

Plan decomposes the work by spawning a child issue targeting the beta repo, and the beta work lands before this repo's change.

**Why this priority**: This is the behaviour under test, cross-repo spawn with strict ordering. The two code changes are only the vehicle.

**Independent Test**: After Plan completes, a child issue exists in the beta repo describing the beta change. It is registered as blocking this issue, and this issue's implementation does not proceed until the child is done.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the child issue exists in the beta repo, **When** the pipeline proceeds, **Then** this issue waits on the child, and the beta change is merged before this repo's change.
3. **Given** the child spec, **When** it is read, **Then** it carries the exact function name and string literal verbatim.

---

### Edge Cases

- The beta repo has not been initialised for spawning (#803): the spawn must still succeed through on-demand initialisation.
- The function name or literal gets altered or "improved" in any stage: the final assertion fails, so the values must be copied exactly.
- The alpha change is attempted before the beta change has landed: the build fails on the missing function, so ordering must be enforced.
- This repo's existing output must not regress when the new call is added.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST expose an exported function named exactly `HelloE2E202610150720020653`, taking no arguments and returning a `string`, in its greeting package (`pkg/greeting/greeting.go`).
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261015-072002-0653`.
- **FR-003**: The beta repo MUST include a test for that function in `pkg/greeting/greeting_test.go` that asserts the returned literal.
- **FR-004**: This repo's entry point (`main.go`) MUST call `greeting.HelloE2E202610150720020653()`, in addition to its existing behaviour, and print the returned value.
- **FR-005**: This repo MUST build cleanly after the change, running `go mod tidy` if needed.
- **FR-006**: The beta repo's change MUST land before this repo's change.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the beta repo, `handarbeit/fabrik-test-delta`.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, in the spawned child issue, the code, and the tests. No alternative names or tokens may be invented.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo, the consumer of the new function.
- **Beta repo**: `handarbeit/fabrik-test-delta`, the repo that provides the greeting function and receives the spawned child issue.
- **Per-run-unique token**: the value embedded in the function name and literal (`20261015-072002-0653`), which distinguishes this run's output from other runs.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the beta repo contains the function `HelloE2E202610150720020653` returning `e2e-cross-repo-spawn-20261015-072002-0653`, with a passing test.
- **SC-002**: A child issue targeting the beta repo was created by Plan through a spawn block, with no manual issue creation.
- **SC-003**: The beta change was merged before the alpha change, and this repo's build succeeds with the new call.
- **SC-004**: Running this repo's program prints the literal `e2e-cross-repo-spawn-20261015-072002-0653` and still produces its prior output.

## Assumptions

- The beta repo already has a greeting package at the path named in the issue, and this repo already imports it (hence "in addition to whatever it already does"). If it does not, creating the package is in scope for the beta child issue.
- "Alpha" means this repo (`handarbeit/fabrik-test-gamma`) and "beta" means `handarbeit/fabrik-test-delta`, as the issue's own headings state.
- Printing the result means writing it to standard output, the same way the program prints its existing output.
- The specific paths and identifiers in this spec are stated by the issue author as fixed fixture values for this e2e scenario, not design choices made here.

## Out of Scope *(optional)*

- Any changes to the greeting package beyond the single new function and its test.
- Changes to the engine's spawn, ordering, or repo-initialisation logic. This scenario only verifies the existing behaviour.
- Any additional repos, or additional functions or tokens.
- Refactoring this repo's existing entry-point behaviour.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
