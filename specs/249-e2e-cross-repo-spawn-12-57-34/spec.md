# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-249`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610171657344675() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261017-165734-4675\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610171657344675() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It is a regression test for #803 (on-demand initialisation of a spawn-target repo) and #1308 (per-run-unique fixture token).

The work spans two repositories with a hard ordering dependency. The shared greeting package lives in `handarbeit/fabrik-test-delta` (the "beta" repo). The consumer is this repo, `handarbeit/fabrik-test-gamma` (the "alpha" repo). The consumer cannot reference a function that does not yet exist in beta, so the beta change must land first. This forces the Plan stage to decompose the work and spawn a sub-issue in the beta repo.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks that exact text in the beta repo, so they must be used verbatim and never substituted.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Plan decomposes work across repos (Priority: P1)

When the Plan stage processes this issue, it recognises that part of the work belongs in `handarbeit/fabrik-test-delta` and must land first. It declares a child spawn targeting that repo instead of making the change in this repo.

**Why this priority**: This is the behaviour under test (#803). Without the spawn, nothing else in the scenario is exercised.

**Independent Test**: Run the Plan stage on this issue and confirm a child issue exists in `handarbeit/fabrik-test-delta`, registered as a blocker of this issue.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawn block was emitted, **When** the engine processes it, **Then** a child issue is created in the beta repo, even if that repo needed on-demand initialisation, and this issue is blocked by it.

---

### User Story 2 - Beta repo gains the greeting function (Priority: P1)

The child issue adds the new exported function and a test for it in the beta repo, and lands before any work on this repo proceeds.

**Why this priority**: The final assertion of the scenario checks that this exact text exists in the beta repo (#1308).

**Independent Test**: Inspect the beta repo's default branch after the child lands. It contains the function with the exact name and literal, and its test passes.

**Acceptance Scenarios**:

1. **Given** the child issue is implemented, **When** it lands, **Then** the beta greeting package exports `HelloE2E202610171657344675() string` returning exactly `"e2e-cross-repo-spawn-20261017-165734-4675"`.
2. **Given** the function exists, **When** the beta repo's tests run, **Then** a test covering that function passes.

---

### User Story 3 - This repo consumes the new function (Priority: P2)

Once the beta change has landed, this repo's entry point calls the new function and prints its result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 2 and completes the end-to-end flow. It is the downstream half of the ordering constraint.

**Independent Test**: Build and run this repo's program after the beta change is available. Its output includes `e2e-cross-repo-spawn-20261017-165734-4675` alongside its prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** this repo's main program runs, **Then** it prints the string returned by `greeting.HelloE2E202610171657344675()`.
2. **Given** the change is made, **When** the module is built, **Then** it compiles, with dependency metadata tidied if that was needed.

---

### Edge Cases

- The beta repo has not been initialised for spawning yet. The spawn must still succeed (#803).
- Implementation in this repo must not begin or merge before the beta child has landed, since the import would not compile.
- The function name and string must not be altered, normalised or regenerated. A different token would fail the final assertion.
- This repo's existing `main.go` behaviour must be preserved. The new call and print are additive.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The Plan stage MUST decompose this issue by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets `handarbeit/fabrik-test-delta`.
- **FR-002**: The beta repo's `pkg/greeting/greeting.go` MUST export a function `HelloE2E202610171657344675() string`.
- **FR-003**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261017-165734-4675`.
- **FR-004**: The beta repo's `pkg/greeting/greeting_test.go` MUST include a test that verifies the function's return value.
- **FR-005**: The beta change MUST land before the change in this repo is implemented or merged.
- **FR-006**: This repo's `main.go` MUST call `greeting.HelloE2E202610171657344675()` and print the result, in addition to its existing behaviour.
- **FR-007**: Dependency metadata for this repo (`go mod tidy`) MUST be updated if the build requires it.
- **FR-008**: The function name and string literal MUST be used exactly as given in this spec, with no substitution.

### Key Entities *(if applicable)*

- **Beta repo**: `handarbeit/fabrik-test-delta`. It owns the greeting package and is the target of the spawned child issue.
- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo. It consumes the greeting function.
- **Spawned child issue**: The beta-repo issue created by the Plan stage. It blocks this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in `handarbeit/fabrik-test-delta` as a result of this issue's Plan stage, and this issue is blocked by it until it is done.
- **SC-002**: The text `HelloE2E202610171657344675` and the string `e2e-cross-repo-spawn-20261017-165734-4675` are present in the beta repo's default branch after the child lands.
- **SC-003**: The beta repo's test suite passes with the new test included.
- **SC-004**: This repo builds, and running its program prints `e2e-cross-repo-spawn-20261017-165734-4675`.
- **SC-005**: The beta change lands strictly before the change in this repo.

## Assumptions

- `handarbeit/fabrik-test-delta` already exposes the `pkg/greeting` package, or the spawned child can create it. The issue names the files to change but does not say whether they already exist.
- This repo already imports the beta repo's greeting package, or the Plan stage will add the dependency as part of the change.
- The literal names given in this issue are the source of truth. The per-run token is generated before the issue is filed and is not regenerated.
- The issue is labelled `fabrik:yolo`, so stages advance automatically once complete.

## Out of Scope *(optional)*

- Changes to the engine's spawn logic itself. This issue only exercises it.
- Any behaviour in either repo beyond the single function, its test, and the call and print in `main.go`.
- Cleanup of fixture artefacts across runs.

## Source References *(optional)*

- #803: on-demand spawn-target repo initialisation
- #1308: per-run-unique fixture token
