# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-817`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610250723061285() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261025-072306-1285\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610250723061285() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two earlier fixes:

- #803: a spawn-target repo is initialised on demand.
- #1308: each run uses a unique fixture token, so a run's final assertion cannot be satisfied by leftovers from an earlier run.

The work needs changes in two repositories in a fixed order. This repo (`handarbeit/fabrik-test-gamma`, called "alpha" in the request) consumes a new greeting function that does not exist yet in the other repo (`handarbeit/fabrik-test-delta`, called "beta"). The function must therefore land in beta first. Planning is expected to recognise that dependency and decompose the work: it should create a sub-issue in beta for the beta-side change and keep the alpha-side change here, blocked on it.

The function name and string literal embed a per-run-unique token (`20261025-072306-1285`). They must be used exactly as written. The scenario's final assertion checks for that exact text in beta.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

The beta repo gains an exported greeting function that returns a fixed, per-run-unique string, with a test covering it. This is delivered through a sub-issue spawned in beta, not as a change in this repo.

**Why this priority**: Everything else depends on it. The alpha change cannot compile without it, and the scenario's final assertion looks for this exact text in beta.

**Independent Test**: In the beta repo, run the greeting package tests. The new test passes, and calling `HelloE2E202610250723061285()` returns `e2e-cross-repo-spawn-20261025-072306-1285`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610250723061285()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261025-072306-1285`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test for the new function exists and passes.
3. **Given** the work is decomposed, **When** the beta change is created, **Then** it exists as a sub-issue in `handarbeit/fabrik-test-delta`, spawned from this issue's Plan stage.

---

### User Story 2 - Alpha repo calls the new function and prints the result (Priority: P2)

This repo's `main.go` calls `greeting.HelloE2E202610250723061285()` from the beta repo, in addition to its existing behaviour, and prints the returned value. Dependency tidying is done if needed.

**Why this priority**: It completes the cross-repo change, but it can only be built and verified after the beta change has landed.

**Independent Test**: With the beta change available, build and run this repo's program. The output includes `e2e-cross-repo-spawn-20261025-072306-1285`, and the program's previous output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function has landed, **When** this repo builds, **Then** it compiles and its dependencies are consistent.
2. **Given** the program runs, **When** it executes, **Then** it prints the string returned by `greeting.HelloE2E202610250723061285()` and still performs its existing behaviour.

---

### User Story 3 - Decomposition happens in the required order (Priority: P1)

Planning decomposes the work so that the beta change is spawned as a sub-issue in the beta repo. This issue's own implementation is blocked on it, so beta lands before alpha.

**Why this priority**: Verifying cross-repo spawn and ordering is the purpose of this scenario. A pipeline that does not decompose has failed it, even if the code ends up correct.

**Independent Test**: After Plan, a spawned child issue exists in `handarbeit/fabrik-test-delta`. This issue is recorded as blocked by that child, and no alpha-side implementation is merged before the child's change has landed.

**Acceptance Scenarios**:

1. **Given** the plan for this issue, **When** Plan completes, **Then** it emits a spawn block targeting `handarbeit/fabrik-test-delta` that describes the beta-side change.
2. **Given** a spawned beta child, **When** it is still open, **Then** this issue does not proceed to implement the alpha-side change.
3. **Given** the child's change has landed in beta, **When** this issue resumes, **Then** the alpha-side change is implemented against the landed function.

---

### Edge Cases

- The beta repo has not been initialised or configured to receive spawned work. On-demand spawn-target initialisation (#803) is expected to handle this.
- A previous run's fixture exists in beta. The per-run-unique token (#1308) means the old function and string must not satisfy this run's assertion; only the exact names given here count.
- The alpha-side change is attempted before beta has landed. The build would fail on the missing function; ordering must prevent this.
- The spawn is retried after a partial failure. It must not create a duplicate beta child.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610250723061285() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610250723061285()` MUST return the literal string `e2e-cross-repo-spawn-20261025-072306-1285`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610250723061285` in `pkg/greeting/greeting_test.go`, and the test MUST pass.
- **FR-004**: The function name and string literal MUST be used verbatim as given; no alternative name or token may be substituted.
- **FR-005**: This repo's `main.go` MUST call `greeting.HelloE2E202610250723061285()` and print the result, in addition to its existing behaviour.
- **FR-006**: This repo's dependencies MUST be tidied (`go mod tidy`) where the change requires it.
- **FR-007**: The beta-side change MUST be delivered as a sub-issue spawned in `handarbeit/fabrik-test-delta` by this issue's Plan stage, using a spawn block that targets the beta repo.
- **FR-008**: The beta-side change MUST land before the alpha-side change is implemented, so that this repo never imports a function that does not yet exist.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo. It consumes the new function in `main.go`.
- **Beta repo**: `handarbeit/fabrik-test-delta`. It owns the `greeting` package and the new function.
- **Spawned child issue**: the sub-issue created in the beta repo for the beta-side change. It blocks this issue's implementation.
- **Per-run-unique token**: `20261025-072306-1285`, embedded in the function name and string literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains `HelloE2E202610250723061285` in `pkg/greeting/greeting.go` returning exactly `e2e-cross-repo-spawn-20261025-072306-1285`.
- **SC-002**: The beta repo's greeting tests pass with the new test included.
- **SC-003**: A child issue exists in `handarbeit/fabrik-test-delta`, created by this issue's Plan stage, and its change merged before this repo's change.
- **SC-004**: This repo builds and runs, and its output contains `e2e-cross-repo-spawn-20261025-072306-1285`.
- **SC-005**: No manual cross-repo intervention is needed to get from this issue to the final state.

## Assumptions

- "Alpha" in the request means this repo (`handarbeit/fabrik-test-gamma`) and "beta" means `handarbeit/fabrik-test-delta`, as the request's own headings state.
- This repo already depends on the beta repo's `greeting` package, or can add it with dependency tidying; the request implies this by saying to run `go mod tidy` if needed.
- "In addition to whatever it already does" means existing output and behaviour in `main.go` are preserved.
- The printing format (for example a trailing newline or a label) is not specified, so any print that includes the exact returned string satisfies the requirement.
- The pipeline is running in an automated mode (`fabrik:yolo`), so spawned work proceeds without extra manual approval.

## Out of Scope

- Any other change to the beta or alpha repos beyond the function, its test and the call site.
- Changing the spawn, board-servability or repo-initialisation machinery. This issue only exercises them (#803, #1308).
- Cleaning up fixtures from earlier runs.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
- `handarbeit/fabrik-test-delta`: `pkg/greeting/greeting.go`, `pkg/greeting/greeting_test.go`
