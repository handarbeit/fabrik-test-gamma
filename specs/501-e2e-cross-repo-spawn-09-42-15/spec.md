# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-501`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610191342157557() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-134215-7557\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610191342157557() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two behaviours: on-demand initialisation of a spawn-target repository (#803) and a per-run-unique fixture token (#1308).

The work is deliberately split across two repositories with a strict ordering dependency. This repository (`handarbeit/fabrik-test-gamma`, "alpha" in the original text) consumes a function that does not exist yet in `handarbeit/fabrik-test-delta` ("beta"). Beta must therefore land first. The planning stage is expected to recognise this and decompose the work by spawning a child issue in the beta repo, rather than doing everything in one repo.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks that exactly this text lands in the beta repo, so these values are fixed inputs. They must not be altered, shortened or re-invented.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Cross-repo work is decomposed into a beta-first child issue (Priority: P1)

An operator files an issue in the alpha repo whose work needs a prerequisite change in the beta repo. The pipeline decomposes the work by creating a child issue in the beta repo, and the alpha work waits for it.

**Why this priority**: Exercising decomposition across repos is the whole purpose of the scenario. Without it, #803 and #1308 are not covered.

**Independent Test**: Run the issue through the pipeline and confirm that a child issue exists in `handarbeit/fabrik-test-delta` and that this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this issue is in the pipeline with its scope spanning both repos, **When** the planning stage completes, **Then** a child issue targeting `handarbeit/fabrik-test-delta` has been created and this issue is blocked by it.
2. **Given** the beta repo has not been initialised for spawn use, **When** the child issue is spawned, **Then** the beta repo is prepared on demand (#803) and the spawn succeeds.

---

### User Story 2 - Beta repo gains the greeting function and its test (Priority: P1)

The beta repo receives an exported function and a test that covers it.

**Why this priority**: The alpha change cannot compile or be validated until this exists, and the scenario's final assertion checks for it.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting` and confirm the new function returns the exact required literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting/greeting.go`, **When** the child issue is implemented and merged, **Then** it exports `HelloE2E202610191342157557() string` returning exactly `e2e-cross-repo-spawn-20261019-134215-7557`.
2. **Given** `pkg/greeting/greeting_test.go`, **When** the tests run, **Then** a test for the new function passes.

---

### User Story 3 - Alpha repo uses the new function (Priority: P1)

After the beta change has landed, this repo's `main.go` calls the new function and prints its result.

**Why this priority**: It completes the end-to-end flow and proves the ordering constraint was honoured.

**Independent Test**: Build and run this repo's program and confirm the output includes the required literal string.

**Acceptance Scenarios**:

1. **Given** the beta change is available, **When** this repo's `main.go` runs, **Then** it calls `greeting.HelloE2E202610191342157557()` in addition to its existing behaviour and prints the returned value.
2. **Given** the dependency set changed, **When** the module is built, **Then** the module files are consistent (`go mod tidy` run if needed) and the build succeeds.

---

### Edge Cases

- The alpha work is attempted before the beta change has landed: it must not proceed, since the import would not resolve.
- The function name or literal is altered (shortened, re-generated, or given a different token): the final assertion fails, so the values must be used verbatim.
- The spawn-target repo is not yet initialised: on-demand initialisation (#803) must handle it rather than failing the spawn.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The work MUST be decomposed so that a child issue targeting `handarbeit/fabrik-test-delta` is spawned from this issue, with this issue blocked until the child's work has landed.
- **FR-002**: In `handarbeit/fabrik-test-delta`, `pkg/greeting/greeting.go` MUST export a function named exactly `HelloE2E202610191342157557` with signature `() string`.
- **FR-003**: That function MUST return exactly the string `e2e-cross-repo-spawn-20261019-134215-7557`.
- **FR-004**: `pkg/greeting/greeting_test.go` in the beta repo MUST contain a test of the new function that asserts the exact returned value.
- **FR-005**: In this repo, `main.go` MUST call `greeting.HelloE2E202610191342157557()` and print the result, while retaining its existing behaviour.
- **FR-006**: This repo's module files MUST be consistent with the new dependency (`go mod tidy` run if needed) and the project MUST build.
- **FR-007**: The alpha change MUST NOT be merged until the beta change has landed.
- **FR-008**: The function name and string literal MUST be used verbatim as given in this issue; no alternative token may be substituted.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repository; consumes the new function from `main.go`.
- **Beta repo**: `handarbeit/fabrik-test-delta`; owns `pkg/greeting` and must receive the function first.
- **Child issue**: the decomposed work item in the beta repo, spawned from this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in `handarbeit/fabrik-test-delta` that was spawned from this issue, and this issue was blocked on it until it landed.
- **SC-002**: The text `e2e-cross-repo-spawn-20261019-134215-7557` appears in the beta repo's `pkg/greeting/greeting.go` on its default branch.
- **SC-003**: The beta repo's `pkg/greeting` tests pass, including the new test.
- **SC-004**: This repo builds, and running it prints `e2e-cross-repo-spawn-20261019-134215-7557`.
- **SC-005**: The beta change landed before the alpha change in every run.

## Assumptions

- The existing `main.go` in this repo already imports or can import the `pkg/greeting` package from the beta repo's module, so only the call and print are added.
- The beta repo's module path is the one this repo already depends on; adding a new dependency edge, if needed, is covered by `go mod tidy`.
- "Alpha" in the original text means this repo, `handarbeit/fabrik-test-gamma`, and "beta" means `handarbeit/fabrik-test-delta`.
- The existing output of `main.go` is kept as is; the new output is additional.

## Out of Scope

- Any change to other functions or behaviour in `pkg/greeting`.
- Changes to the pipeline or engine themselves; this issue only exercises existing spawn behaviour.
- Releasing or versioning the beta repo beyond what is needed for this repo to consume the new function.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
