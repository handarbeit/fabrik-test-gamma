# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2096`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611161548245157() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261116-154824-5157\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611161548245157() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. Work that spans two repositories must be split so that the dependency lands first. Here, the repo this issue lives in (`handarbeit/fabrik-test-gamma`, "alpha") needs a function that does not yet exist in the other repo (`handarbeit/fabrik-test-delta`, "beta"). The scenario checks two behaviours:

- **#803**: the spawn-target repo is initialised on demand when a child issue is spawned into it.
- **#1308**: a per-run-unique fixture token, already baked into the function name and string literal below, lets the final assertion tell this run's output apart from any earlier run's.

The function name and string literal are fixed by the scenario and must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new exported greeting function that returns the run-unique string. This is the dependency, so it must land first.

**Why this priority**: The alpha change cannot compile without it, and the scenario's final assertion checks this exact text in the beta repo.

**Independent Test**: In the beta repo, run the unit test for the new function and confirm it passes. Confirm that `HelloE2E202611161548245157()` returns `"e2e-cross-repo-spawn-20261116-154824-5157"`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611161548245157()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261116-154824-5157`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test for the new function exists and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

The alpha repo's entry point calls the new beta function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on Story 1 and demonstrates the cross-repo dependency being honoured.

**Independent Test**: Build and run alpha's entry point once beta's change is available, and confirm the printed output includes the run-unique string.

**Acceptance Scenarios**:

1. **Given** beta's function is available, **When** alpha's entry point runs, **Then** it prints the result of `greeting.HelloE2E202611161548245157()` alongside its existing output.
2. **Given** alpha's module dependencies, **When** they are tidied, **Then** the module builds without error.

---

### User Story 3 - Work is decomposed across repos in order (Priority: P1)

The planning stage decomposes the work into a child issue in the beta repo that must land before the alpha change.

**Why this priority**: This is the behaviour the scenario exists to verify (#803, #1308).

**Independent Test**: After planning, a child issue exists in the beta repo, is registered as a blocker of this issue, and covers the beta-side change. The beta change merges before the alpha change.

**Acceptance Scenarios**:

1. **Given** the work requires changes in both repos, **When** Plan runs, **Then** it emits a spawn block targeting the beta repo for the beta-side change.
2. **Given** the child issue exists, **When** the pipeline proceeds, **Then** the alpha change is not implemented until the beta change has landed.

---

### Edge Cases

- The beta repo has not been initialised for spawning before this run: it must be initialised on demand (#803).
- A previous run's artifacts are present in either repo: the run-unique token must keep this run's assertion unambiguous (#1308).
- The alpha entry point already prints other output: that existing output must be preserved.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function named exactly `HelloE2E202611161548245157` in its greeting package, taking no arguments and returning a string.
- **FR-002**: `HelloE2E202611161548245157()` MUST return exactly the literal `e2e-cross-repo-spawn-20261116-154824-5157`.
- **FR-003**: The beta repo's greeting tests MUST include a test for `HelloE2E202611161548245157` that passes.
- **FR-004**: The alpha repo's (`handarbeit/fabrik-test-gamma`) entry point MUST call `greeting.HelloE2E202611161548245157()` and print the result, in addition to its existing behaviour.
- **FR-005**: The alpha repo's module dependencies MUST be tidied if needed so it builds.
- **FR-006**: The beta-side change MUST land before the alpha-side change.
- **FR-007**: The Plan stage MUST decompose the work by spawning a child issue targeting the beta repo (using the spawn block mechanism), rather than handling both repos in a single issue.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this spec; no substitute names or tokens may be introduced.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo contains the exact text `e2e-cross-repo-spawn-20261116-154824-5157` returned by `HelloE2E202611161548245157`, which is the scenario's final assertion.
- **SC-002**: A child issue exists in the beta repo as a result of Plan's decomposition, and this issue is blocked on it until it lands.
- **SC-003**: The beta repo's tests pass, including the new test.
- **SC-004**: The alpha repo builds, and its entry point's output includes the run-unique string.
- **SC-005**: The beta change merges before the alpha change.

## Assumptions

- The existing `pkg/greeting` package in the beta repo is the right home for the new function and test, as stated in the request.
- The alpha repo already imports (or can import) the beta repo's greeting package.
- "In addition to whatever it already does" means alpha's existing output is unchanged and the new line is added.
- The run-unique token is already final and no regeneration is needed.

## Out of Scope

- Any change to the greeting package beyond the one function and its test.
- Any other repos beyond the alpha and beta repos named above.
- Changes to the engine's spawn or on-demand initialisation behaviour; this scenario only exercises them.
