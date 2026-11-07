# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1739`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description (verbatim):

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202611071936352658() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261107-193635-2658".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202611071936352658() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It checks that when one piece of work needs changes in two repositories in a strict order, the pipeline splits it into a child issue in the second repository and a parent change in this one, and that the changes land in the right order.

The scenario covers two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: the per-run-unique fixture token, which lets the final assertion tell this run's output apart from any other run's.

The work has a real dependency. The greeting function must exist in the delta repo (`handarbeit/fabrik-test-delta`, "beta") before this repo (`handarbeit/fabrik-test-gamma`, "alpha") can import and call it. That ordering is why decomposition is needed, and the scenario exists to confirm the pipeline honours it.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the uniquely-tokenised greeting function (Priority: P1)

The delta repo gets a new exported greeting function that returns this run's unique token, along with a test for it. This is the child work the pipeline must spawn from this issue.

**Why this priority**: The final assertion of the scenario checks that this exact text lands in the beta repo, and the alpha change depends on it.

**Independent Test**: In the delta repo, the greeting package's tests pass, and calling `HelloE2E202611071936352658()` returns `"e2e-cross-repo-spawn-20261107-193635-2658"`.

**Acceptance Scenarios**:

1. **Given** the delta repo without the new function, **When** the child work is completed, **Then** `pkg/greeting/greeting.go` exports `HelloE2E202611071936352658() string` returning exactly `"e2e-cross-repo-spawn-20261107-193635-2658"`.
2. **Given** the new function exists, **When** the delta repo's tests are run, **Then** a test in `pkg/greeting/greeting_test.go` asserts the returned literal and passes.

---

### User Story 2 - Alpha repo consumes the beta function (Priority: P2)

This repo's `main.go` calls the new beta function and prints the result, after the beta change has landed.

**Why this priority**: It completes the end-to-end flow, but it depends on User Story 1 being merged first.

**Independent Test**: With the beta function available, building and running this repo prints the token returned by `greeting.HelloE2E202611071936352658()` and the build succeeds.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** `main.go` is updated, **Then** it calls `greeting.HelloE2E202611071936352658()` and prints the result in addition to its existing behaviour.
2. **Given** the updated `main.go`, **When** the module is built, **Then** the build succeeds, with `go mod tidy` run if needed to resolve the dependency.

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The Plan stage splits the work by emitting a spawn-child block (`FABRIK_SPAWN_CHILD_BEGIN`) that targets the delta repo. The parent (alpha) work stays blocked until the child (beta) work has landed.

**Why this priority**: Verifying cross-repo spawn and ordering is the purpose of the scenario.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` describing the beta change, and the parent is blocked on it.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it emits a spawn-child block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the child issue exists, **When** the pipeline runs, **Then** the beta change lands before the alpha change is implemented against it.

---

### Edge Cases

- The function name and string literal must be used verbatim. Inventing a different name or token would break the scenario's final assertion.
- The delta repo may not yet be initialised for spawning. The scenario exercises on-demand initialisation of the spawn target (#803).
- The alpha change must not be attempted before the beta change has landed, because the import cannot resolve until then.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202611071936352658() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202611071936352658()` MUST return exactly the literal `"e2e-cross-repo-spawn-20261107-193635-2658"`.
- **FR-003**: The delta repo MUST include a test for this function in `pkg/greeting/greeting_test.go`.
- **FR-004**: This repo's (`handarbeit/fabrik-test-gamma`) `main.go` MUST call `greeting.HelloE2E202611071936352658()`, in addition to its existing behaviour, and print the result.
- **FR-005**: The dependency changes needed for this repo to build against the new beta function MUST be applied (`go mod tidy` run if needed).
- **FR-006**: The delta change MUST land before the gamma change is implemented.
- **FR-007**: The Plan stage MUST decompose the work by emitting a spawn-child block (`FABRIK_SPAWN_CHILD_BEGIN`) that targets the delta repo.
- **FR-008**: The function name and string literal in FR-001 and FR-002 MUST be used exactly as given, with no substitution of a different name or token.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo, which consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-delta`, which provides the greeting function.
- **Per-run-unique token**: `20261107-193635-2658`, embedded in the function name and the returned literal so the scenario's final assertion can identify this run's output.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting `handarbeit/fabrik-test-delta` is created from this issue's Plan stage.
- **SC-002**: The delta repo's default branch contains `HelloE2E202611071936352658` returning exactly `"e2e-cross-repo-spawn-20261107-193635-2658"`, with a passing test.
- **SC-003**: This repo builds successfully with `main.go` calling and printing the result of the new function.
- **SC-004**: The delta change is merged before the alpha change is merged.

## Assumptions

- The delta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, as the issue implies. The new function and test are additions to those files.
- This repo already imports the delta repo's greeting package, or the dependency can be added through `go mod tidy`.
- The pipeline's existing behaviour (blocking the parent on the spawned child) provides the required ordering. No new ordering mechanism is needed.
- The token `20261107-193635-2658` is fixed for this run and must not be regenerated.

## Out of Scope *(optional)*

- Any change to Fabrik's spawn or decomposition logic. This issue only exercises it.
- Changes beyond the specified function, its test and the `main.go` call in either repo.
- Fixing #803 or #1308. This is a regression test for them.
