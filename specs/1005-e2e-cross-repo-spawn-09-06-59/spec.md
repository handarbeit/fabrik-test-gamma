# Feature Specification: E2E cross-repo spawn (greeting function in delta, consumed by gamma)

**Feature Branch**: `fabrik/issue-1005`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610251306590520() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261025-130659-0520\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610251306590520() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a strict ordering: this repo (gamma) imports a greeting package from `handarbeit/fabrik-test-delta` (called "beta" in the original request), so gamma cannot call a new greeting function until delta provides it.

The scenario exercises two previously fixed behaviors:

- **#803**: on-demand initialization of the spawn-target repo, so a child issue can be created in a repo the instance has not yet set up.
- **#1308**: a per-run-unique fixture token. The function name and string literal below embed such a token. The scenario's final assertion checks that this exact text lands in the delta repo, so the values must be used verbatim and never substituted.

The expected pipeline behavior is that Plan decomposes the work by emitting a spawn block that targets the delta repo. The delta child lands first. The gamma change follows once the child is complete.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Delta provides the new greeting function (Priority: P1)

A maintainer needs the delta repo's greeting package to expose a new function that returns a per-run-unique token string, with a test covering it. This is the prerequisite for everything else in the scenario.

**Why this priority**: Gamma cannot compile against a function that does not exist. The final e2e assertion also checks for this text in the delta repo.

**Independent Test**: In the delta repo, run the package tests for `pkg/greeting`. The new test passes, and calling `HelloE2E202610251306590520()` returns exactly `e2e-cross-repo-spawn-20261025-130659-0520`.

**Acceptance Scenarios**:

1. **Given** the delta repo's `pkg/greeting` package, **When** `HelloE2E202610251306590520()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261025-130659-0520`.
2. **Given** the delta repo's test suite, **When** it is run, **Then** a test for `HelloE2E202610251306590520` exists and passes.

---

### User Story 2 - Gamma prints the delta-provided value (Priority: P1)

A maintainer needs gamma's `main.go` to call the new delta function and print its result, in addition to its existing output.

**Why this priority**: This is the half of the change that lives in this repo. It proves that the cross-repo dependency was consumed in the required order.

**Independent Test**: After the delta change has landed and gamma's dependency on delta is updated to include it, build and run gamma. The output includes the line `e2e-cross-repo-spawn-20261025-130659-0520`, and all previously printed output is unchanged.

**Acceptance Scenarios**:

1. **Given** the delta function exists and gamma depends on a delta revision that contains it, **When** gamma's `main` runs, **Then** it prints the value returned by `greeting.HelloE2E202610251306590520()`.
2. **Given** gamma's existing greeting calls, **When** the new call is added, **Then** all existing calls and their output are preserved.
3. **Given** the updated gamma module, **When** the build and tests are run, **Then** they succeed and the module files are consistent (tidy).

---

### User Story 3 - The work is decomposed across repos in order (Priority: P1)

An operator running the pipeline needs Plan to split this issue into a delta child issue and the gamma work, and to order them so that delta lands first.

**Why this priority**: Verifying this decomposition is the purpose of the scenario (regression coverage for #803 and #1308).

**Independent Test**: Run the issue through the pipeline. Confirm a child issue is created in the delta repo, that it is registered as a blocker of this issue, and that the gamma change is not merged until the delta change has landed.

**Acceptance Scenarios**:

1. **Given** Plan has analyzed this issue, **When** it declares the decomposition, **Then** it emits a spawn block targeting `handarbeit/fabrik-test-delta` covering the delta-side change.
2. **Given** the delta child exists, **When** this issue proceeds, **Then** the gamma change does not land before the delta child's change has landed.
3. **Given** the pipeline has completed, **When** the delta repo is inspected, **Then** it contains the exact function name and string literal given in this spec.

---

### Edge Cases

- The delta repo has not yet been initialized for spawn targets (#803): spawning the child must still succeed on demand.
- The function name or string literal is altered, shortened, or regenerated by any stage: the final assertion fails. They must be copied verbatim.
- Gamma's dependency on delta is pinned to a revision that predates the new function: gamma cannot build until the dependency is updated to a revision containing it.
- The delta child has not landed when the gamma work is ready: gamma must wait rather than land a change that cannot compile.
- Other `HelloE2E…` functions already exist in the delta repo and are already printed by gamma: the new function is purely additive and must not alter or remove them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST expose an exported function `HelloE2E202610251306590520() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610251306590520()` MUST return exactly the literal string `e2e-cross-repo-spawn-20261025-130659-0520`.
- **FR-003**: The delta repo MUST include a test for `HelloE2E202610251306590520` in `pkg/greeting/greeting_test.go` that passes.
- **FR-004**: Gamma's `main.go` MUST call `greeting.HelloE2E202610251306590520()` and print its result, in addition to its existing output.
- **FR-005**: Gamma's existing greeting calls and their output MUST remain unchanged.
- **FR-006**: Gamma's module files MUST be consistent after the change (`go mod tidy` run if needed), and gamma MUST build and pass its tests against a delta revision that contains the new function.
- **FR-007**: The delta-side change MUST land before the gamma-side change.
- **FR-008**: Plan MUST decompose this issue by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` for the delta-side work.
- **FR-009**: The exact function name and string literal in FR-001 and FR-002 MUST be used verbatim by every stage and in every repo. They MUST NOT be replaced by invented values.

### Key Entities

- **Delta repo (beta)**: `handarbeit/fabrik-test-delta`, which owns the `greeting` package and receives the spawned child issue.
- **Gamma repo (alpha)**: `handarbeit/fabrik-test-gamma`, this repo, which consumes the `greeting` package from delta.
- **Per-run-unique token**: The `20261025-130659-0520` value embedded in the function name and string literal, used by the e2e harness to assert the correct text landed in delta.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the delta repo's default branch contains `HelloE2E202610251306590520` returning `e2e-cross-repo-spawn-20261025-130659-0520`, with a passing test.
- **SC-002**: After the pipeline completes, running the gamma program prints `e2e-cross-repo-spawn-20261025-130659-0520` as part of its output, alongside all prior output.
- **SC-003**: A child issue targeting `handarbeit/fabrik-test-delta` exists as a result of Plan's spawn block, and is linked as a blocker of this issue.
- **SC-004**: The delta change is merged before the gamma change is merged, in 100% of runs.
- **SC-005**: Gamma's build and tests pass at the point of merge.

## Assumptions

- "Beta" and `handarbeit/fabrik-test-delta` are the same repo, and "alpha" is this repo (`handarbeit/fabrik-test-gamma`), as the original request's repo names indicate.
- Gamma already depends on the delta module, so the only dependency change needed is moving to a delta revision that contains the new function.
- The existing `HelloE2E…` functions in delta and their calls in gamma are unrelated to this scenario and stay as they are.
- The `fabrik:yolo` label means stages advance automatically. No human approval gate is expected between stages.

## Out of Scope *(optional)*

- Changing any existing greeting function or its test in either repo.
- Changes to the Fabrik engine's spawn behavior itself. This issue only exercises it.
- Any change to repos other than gamma and delta.

## Source References *(optional)*

- Fabrik issues #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token), which this scenario regression-tests.
