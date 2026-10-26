# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1296`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610261605192109() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261026-160519-2109\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610261605192109() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repo (#803), and a per-run-unique fixture token that lets the scenario's final assertion tell this run's output from any other run's (#1308).

The work spans two repositories with a strict ordering dependency. The consumer repo (`handarbeit/fabrik-test-gamma`, "alpha" in the original text) cannot import a function that does not yet exist in the provider repo (`handarbeit/fabrik-test-delta`, "beta" in the original text). The pipeline must therefore decompose the work. The provider-side change is spawned as a child issue in the provider repo, and the consumer-side change depends on it.

The function name and string literal embed a per-run-unique token. They are fixed by this issue and must be used exactly as written. The scenario's final assertion checks that this exact text landed in the provider repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provider repo exposes the unique greeting function (Priority: P1)

The provider repo (`handarbeit/fabrik-test-delta`) gains an exported greeting function that returns this run's unique literal, with a test covering it. This is the work the pipeline must delegate to a spawned child issue in that repo.

**Why this priority**: The consumer change cannot compile or be merged until this exists, and the scenario's final assertion checks this text in the provider repo.

**Independent Test**: In the provider repo, run the package tests for `pkg/greeting`. The new test passes and asserts the exact literal.

**Acceptance Scenarios**:

1. **Given** the provider repo's `pkg/greeting/greeting.go`, **When** the change lands, **Then** it exports `HelloE2E202610261605192109() string` returning exactly `"e2e-cross-repo-spawn-20261026-160519-2109"`.
2. **Given** `pkg/greeting/greeting_test.go`, **When** the package tests run, **Then** a test for `HelloE2E202610261605192109` exists and passes.

---

### User Story 2 - Consumer repo calls and prints the new function (Priority: P2)

The consumer repo (`handarbeit/fabrik-test-gamma`, this repo) updates its entry point to call the provider's new function and print the result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and completes the end-to-end flow, but it delivers nothing until the provider change is available.

**Independent Test**: Build and run the consumer program after the provider change is available. Its output includes `e2e-cross-repo-spawn-20261026-160519-2109`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the provider function is available, **When** `main.go` runs, **Then** it calls `greeting.HelloE2E202610261605192109()` and prints the returned string.
2. **Given** the consumer repo's existing behaviour, **When** the change lands, **Then** everything `main.go` did before still happens.
3. **Given** the dependency set changed, **When** the change is finalised, **Then** the module builds cleanly (running `go mod tidy` if needed).

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

Plan recognises that the work needs changes in both repos, with the provider first. It emits a spawn block targeting the provider repo, so a child issue is created there and the consumer issue is blocked on it.

**Why this priority**: Verifying this decomposition is the purpose of the scenario (#803, #1308).

**Independent Test**: After Plan completes, a child issue exists in `handarbeit/fabrik-test-delta` containing the provider-side requirements. The consumer issue is blocked by it.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawn block, **When** the child issue is created, **Then** its body carries the exact function name and string literal from this issue, unaltered.
3. **Given** the child issue exists, **When** the consumer work is scheduled, **Then** the consumer implementation does not proceed until the provider change has landed.

---

### Edge Cases

- The provider repo may not yet be initialised for Fabrik. The spawn must still succeed, which is the #803 on-demand init behaviour.
- A pipeline stage may be tempted to "improve" or regenerate the token. The function name and literal must be used verbatim, because a mismatch breaks the scenario's final assertion.
- If the provider change has not landed, the consumer change must not be merged. It would not compile.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The provider repo (`handarbeit/fabrik-test-delta`) MUST export `HelloE2E202610261605192109() string` from `pkg/greeting/greeting.go`, returning exactly `"e2e-cross-repo-spawn-20261026-160519-2109"`.
- **FR-002**: The provider repo MUST include a test for `HelloE2E202610261605192109` in `pkg/greeting/greeting_test.go`.
- **FR-003**: The consumer repo (`handarbeit/fabrik-test-gamma`) `main.go` MUST call `greeting.HelloE2E202610261605192109()` and print the result, while preserving its existing behaviour.
- **FR-004**: The consumer repo's module files MUST be left consistent (`go mod tidy` run if needed).
- **FR-005**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the provider repo, so the provider change is delivered via a spawned child issue.
- **FR-006**: The provider-side change MUST land before the consumer-side change is merged.
- **FR-007**: The function name `HelloE2E202610261605192109` and the literal `e2e-cross-repo-spawn-20261026-160519-2109` MUST be used verbatim everywhere. They MUST NOT be renamed, regenerated or altered.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the provider repo's default branch contains the exact literal `e2e-cross-repo-spawn-20261026-160519-2109` returned by `HelloE2E202610261605192109`.
- **SC-002**: A child issue targeting the provider repo was created by Plan's spawn block, and the consumer issue was blocked by it until the child completed.
- **SC-003**: The provider repo's `pkg/greeting` tests pass, including the new test.
- **SC-004**: The consumer program builds and its output includes the literal, with its prior output intact.

## Assumptions

- "Alpha" in the original text means this repo (`handarbeit/fabrik-test-gamma`). "Beta" means `handarbeit/fabrik-test-delta`.
- The provider repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, and the consumer repo already has a `main.go` that can import the provider's greeting package.
- The consumer repo already depends on the provider repo as a Go module. If it does not, adding that dependency is part of FR-004.
- The test for the new function asserts the literal return value. Its exact style follows the existing tests in the package.

## Out of Scope *(optional)*

- Changes to Fabrik engine behaviour. This issue only exercises existing spawn behaviour.
- Any changes in either repo beyond the function, its test and the consumer call.
- Generating or rotating the per-run token. It is supplied by the scenario.

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
