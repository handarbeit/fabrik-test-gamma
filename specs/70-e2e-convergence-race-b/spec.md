# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-70`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "e2e cross-repo spawn — Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order. In handarbeit/fabrik-test-delta (the beta repo): add an exported function HelloE2E202610142327011219() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261014-232701-1219", and add a test for it in pkg/greeting/greeting_test.go. In handarbeit/fabrik-test-gamma (this repo): update main.go to call greeting.HelloE2E202610142327011219() (in addition to whatever it already does) and print the result, and run go mod tidy if needed. Constraint: beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo. This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work spans two repositories with a strict ordering dependency: `handarbeit/fabrik-test-delta` (the "beta" repo, which owns the `greeting` package) must gain a new exported function first, and only then can `handarbeit/fabrik-test-gamma` (this repo, whose `main.go` already imports and prints several `greeting` functions) call it.

The scenario exists to verify:
- **#803** — on-demand initialization of a spawn-target repo, when Plan decomposes work into a child issue in another repository.
- **#1308** — a per-run-unique fixture token. The function name and string literal embed a run-unique token (`20261014-232701-1219`); the scenario's final assertion checks that exactly this text lands in the beta repo, so it must be used verbatim and never substituted.

Plan is expected to decompose the work by emitting a spawn block for a child issue targeting the beta repo, with this issue (gamma) depending on that child.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo exposes the new greeting function (Priority: P1)

A child issue in `handarbeit/fabrik-test-delta` adds the new exported function and a test for it. This is the prerequisite slice and must land before anything in this repo changes.

**Why this priority**: This repo cannot compile a call to a function that does not exist upstream, and the scenario's final assertion is made against the beta repo.

**Independent Test**: In the beta repo, run the package tests for `pkg/greeting`; the new test passes and asserts the exact literal. The function is exported and returns that literal.

**Acceptance Scenarios**:

1. **Given** the beta repo's `pkg/greeting` package, **When** `HelloE2E202610142327011219()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261014-232701-1219`.
2. **Given** the new function, **When** the beta repo's test suite runs, **Then** a test for it exists in `pkg/greeting/greeting_test.go` and passes.

---

### User Story 2 - Gamma repo calls and prints the new function (Priority: P2)

Once the beta change has landed, this repo's `main.go` calls `greeting.HelloE2E202610142327011219()` and prints the result, in addition to everything it already prints.

**Why this priority**: It depends on User Story 1 and completes the end-to-end flow.

**Independent Test**: Build and run this repo's program (with the updated beta dependency available); its output includes the line `e2e-cross-repo-spawn-20261014-232701-1219` alongside all prior output.

**Acceptance Scenarios**:

1. **Given** the beta function has landed and this repo depends on a version containing it, **When** the program runs, **Then** its output contains `e2e-cross-repo-spawn-20261014-232701-1219`.
2. **Given** the change, **When** the program runs, **Then** all previously printed lines are still printed.
3. **Given** the change, **When** the module is built, **Then** it builds cleanly, with `go mod tidy` applied if it was needed.

---

### User Story 3 - Plan decomposes across repos in strict order (Priority: P1)

Plan recognises that the work cannot be done in one repo and decomposes it: it emits a spawn block targeting the beta repo, and this issue is blocked until the child is done.

**Why this priority**: Verifying this decomposition is the purpose of the scenario (#803, #1308).

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` carrying the beta-side requirements, and this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan runs, **Then** it emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawned child, **When** its body is read, **Then** it contains the exact function name and exact string literal from this spec.
3. **Given** the child is not yet complete, **When** the pipeline continues for this issue, **Then** no gamma-side change is implemented ahead of it.

---

### Edge Cases

- The function name and literal contain a run-unique token; any altered, shortened or re-generated token breaks the scenario's final assertion.
- The beta repo may need to be initialized on demand when the child is spawned (the subject of #803).
- This repo's dependency on the beta module must resolve to a version that contains the new function; a stale pinned version would make the call fail to compile.
- The beta function must not collide with or alter existing `greeting` functions.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST gain an exported function `HelloE2E202610142327011219() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261014-232701-1219`.
- **FR-003**: The beta repo MUST include a test for the function in `pkg/greeting/greeting_test.go`, asserting the exact literal.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610142327011219()` and print the result, while retaining all existing output.
- **FR-005**: This repo MUST build successfully after the change, running `go mod tidy` if needed.
- **FR-006**: The beta-side change MUST land before the gamma-side change.
- **FR-007**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
- **FR-008**: The function name and string literal MUST be used verbatim everywhere they appear (spawned child, code, tests); none may be invented, altered or regenerated.

### Key Entities

- **Beta repo**: `handarbeit/fabrik-test-delta`, owner of the `greeting` package.
- **Gamma repo (this repo)**: `handarbeit/fabrik-test-gamma`, a consumer of the `greeting` package.
- **Spawned child issue**: the beta-repo issue created by Plan to carry the beta-side changes.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the beta repo's default branch contains `HelloE2E202610142327011219` returning `e2e-cross-repo-spawn-20261014-232701-1219`, with a passing test.
- **SC-002**: Running this repo's program prints a line equal to `e2e-cross-repo-spawn-20261014-232701-1219` in addition to all prior output.
- **SC-003**: A child issue was created in the beta repo by Plan's spawn block, and the beta change merged before the gamma change.
- **SC-004**: The token `20261014-232701-1219` appears unmodified in every location it is expected.

## Assumptions

- The beta repo already has `pkg/greeting/greeting.go` and `pkg/greeting/greeting_test.go`, since this repo already imports from that package.
- The existing `HelloE2E…` functions in the beta repo are the established pattern for the new one.
- The `fabrik:yolo` label means the pipeline advances automatically through stages once each is complete.
- How this repo picks up the new beta version (dependency update) is a planning matter, not specified here.

## Out of Scope *(optional)*

- Any change to existing `greeting` functions in the beta repo.
- Any change to engine behaviour; this issue only exercises it.
- Any further repositories beyond beta and gamma.

## Source References *(optional)*

- Fabrik issues #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
