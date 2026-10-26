# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1220`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610260819152314() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261026-081915-2314\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610260819152314() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier fixes:

- #803: on-demand initialisation of a spawn-target repo.
- #1308: a per-run-unique fixture token, so each run's final assertion can tell its own output from an earlier run's.

The work spans two repos with a strict ordering dependency. `handarbeit/fabrik-test-gamma` (this repo, "alpha") imports a greeting function that must first exist in `handarbeit/fabrik-test-delta` ("beta"). Beta must therefore land first. The pipeline should decompose this by spawning a child issue in beta from the Plan stage, rather than handling both repos in one issue.

The function name and string literal embed a per-run-unique token. The scenario's final assertion checks for that exact text in the beta repo, so both must be used verbatim.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Plan decomposes cross-repo work into a beta child issue (Priority: P1)

The Plan stage recognises that the change needs a prior change in the beta repo. It declares a spawned child issue targeting `handarbeit/fabrik-test-delta` for the beta-side work. That child is created and linked as a blocker of this issue.

**Why this priority**: This is the behaviour under test (#803). Without the spawn, nothing else in the scenario is exercised.

**Independent Test**: Run the issue through Plan and verify a child issue exists in `handarbeit/fabrik-test-delta`, assigned and on the board, with this issue blocked by it.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawn block is processed, **When** the child is created, **Then** this issue is blocked by the child until the child is resolved.

---

### User Story 2 - Beta lands the greeting function with the exact per-run token (Priority: P1)

The beta child adds `HelloE2E202610260819152314() string` to `pkg/greeting/greeting.go`, returning `"e2e-cross-repo-spawn-20261026-081915-2314"`. It also adds a test in `pkg/greeting/greeting_test.go`.

**Why this priority**: The scenario's final assertion depends on this exact text landing in beta. Alpha also cannot compile its change until this exists.

**Independent Test**: Inspect beta's `pkg/greeting/greeting.go` after the child lands. It must contain the function and literal exactly as given, and the new test must pass.

**Acceptance Scenarios**:

1. **Given** the beta child is implemented and merged, **When** `HelloE2E202610260819152314()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-081915-2314`.
2. **Given** the beta test file, **When** beta's tests run, **Then** the new test for this function passes.

---

### User Story 3 - Alpha consumes the beta function after beta lands (Priority: P2)

Once beta has landed, alpha's `main.go` calls `greeting.HelloE2E202610260819152314()` in addition to its existing behaviour, and prints the result.

**Why this priority**: It completes the cross-repo flow. It depends on User Story 2 and exercises the ordering constraint.

**Independent Test**: With beta's function available, build and run alpha. The output includes the returned string and existing behaviour is unchanged.

**Acceptance Scenarios**:

1. **Given** beta's function has landed, **When** alpha is built and run, **Then** the program prints `e2e-cross-repo-spawn-20261026-081915-2314` in addition to its prior output.
2. **Given** the dependency change, **When** `go mod tidy` is needed, **Then** it has been run and the module builds cleanly.

---

### Edge Cases

- Alpha work must not start until beta has landed. If it did, the import would not compile.
- The function name and string literal must not be paraphrased, renamed or regenerated. Any drift breaks the final assertion.
- The spawn target repo may not yet be initialised for Fabrik. Spawn-target initialisation on demand (#803) must handle this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The Plan stage MUST emit a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` for the beta-side work.
- **FR-002**: The beta child MUST add an exported function `HelloE2E202610260819152314() string` to `pkg/greeting/greeting.go` in `handarbeit/fabrik-test-delta`.
- **FR-003**: `HelloE2E202610260819152314()` MUST return exactly the literal `e2e-cross-repo-spawn-20261026-081915-2314`.
- **FR-004**: The beta child MUST add a test for the function in `pkg/greeting/greeting_test.go`, and the test MUST pass.
- **FR-005**: This repo's `main.go` MUST call `greeting.HelloE2E202610260819152314()`, in addition to its existing behaviour, and print the result.
- **FR-006**: `go mod tidy` MUST be run in this repo if needed so that the module builds cleanly.
- **FR-007**: The beta change MUST land before the alpha change is implemented, and this issue MUST be blocked by the spawned child.
- **FR-008**: The function name and string literal MUST be used verbatim as given. No alternative name or token may be substituted.

### Key Entities *(if applicable)*

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo, which consumes the greeting function.
- **Beta repo**: `handarbeit/fabrik-test-delta`, which provides `pkg/greeting`.
- **Spawned child issue**: the beta-side work item created by Plan's spawn block and linked as a blocker of this issue.
- **Per-run-unique token**: the text baked into the function name and string literal, which the final assertion checks.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting `handarbeit/fabrik-test-delta` is created from Plan's spawn block, and this issue is blocked by it.
- **SC-002**: Beta's `pkg/greeting/greeting.go` contains `HelloE2E202610260819152314` returning `e2e-cross-repo-spawn-20261026-081915-2314`, verbatim.
- **SC-003**: Beta's tests, including the new one, pass.
- **SC-004**: Alpha builds and, when run, prints `e2e-cross-repo-spawn-20261026-081915-2314` alongside its existing output.
- **SC-005**: The beta change is merged before alpha's change merges, and alpha never has a state where the import is unresolvable on its base branch.

## Assumptions

- Alpha's `main.go` already imports or can import the beta module's `pkg/greeting` package. The import path and module wiring are a planning concern.
- Beta already has `pkg/greeting/greeting.go` and `greeting_test.go`. If it does not, adding them is part of the beta child's work.
- The spawn mechanism assigns the child, adds it to the board and links the blocking relationship, per existing Fabrik behaviour. This issue does not change it.
- The per-run token in the given names is final for this run and must not be regenerated.

## Out of Scope *(optional)*

- Any change to the spawn mechanism itself, to #803 or to #1308.
- Changes to beta beyond the one function and its test.
- Changes to alpha beyond calling and printing the function (and `go mod tidy` if needed).

## Source References *(optional)*

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
