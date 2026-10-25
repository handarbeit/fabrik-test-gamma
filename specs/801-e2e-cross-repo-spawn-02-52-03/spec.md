# Feature Specification: Cross-repo spawn e2e scenario (greeting function in delta, consumed by gamma)

**Feature Branch**: `fabrik/issue-801`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "e2e cross-repo spawn (02:52:03)" — the original request, verbatim:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610250652037932() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261025-065203-7932".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610250652037932() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a strict ordering dependency: `handarbeit/fabrik-test-gamma` (this repo, called "alpha" in the original request) cannot call a greeting function until `handarbeit/fabrik-test-delta` (called "beta") defines it. The scenario verifies that the Plan stage decomposes the work by spawning a child issue in the delta repo, that the child lands before this issue's change, and that the exact per-run-unique token ends up in the delta repo. It is a regression test for #803 (on-demand initialization of a spawn-target repo) and #1308 (per-run-unique fixture token).

The function name and string literal embed a per-run-unique token. They must be used exactly as given, because the scenario's final assertion checks for that exact text in the delta repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Delta repo exposes the unique greeting function (Priority: P1)

The delta repo gains an exported greeting function that returns the run-unique string, with a test, delivered through a child issue spawned from this issue.

**Why this priority**: The delta change is the prerequisite for everything else and is what the scenario's final assertion inspects.

**Independent Test**: In the delta repo, the greeting package's tests pass, and the function `HelloE2E202610250652037932()` returns exactly `e2e-cross-repo-spawn-20261025-065203-7932`.

**Acceptance Scenarios**:

1. **Given** the delta repo's greeting package, **When** `HelloE2E202610250652037932()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261025-065203-7932`.
2. **Given** the delta repo's greeting test file, **When** the delta repo's tests run, **Then** a test covering `HelloE2E202610250652037932()` exists and passes.
3. **Given** this issue's Plan stage, **When** it decomposes the work, **Then** it emits a spawn block targeting the delta repo and a child issue is created there carrying the delta-side work.

---

### User Story 2 - Gamma repo consumes the function after delta lands (Priority: P2)

This repo's main program calls the new delta function and prints its result, in addition to its existing behavior, only after the delta change has landed.

**Why this priority**: It depends on Story 1 and demonstrates that the strict ordering (delta first) is respected.

**Independent Test**: With the delta change available, this repo builds, and running the program prints the string `e2e-cross-repo-spawn-20261025-065203-7932` in addition to its prior output.

**Acceptance Scenarios**:

1. **Given** the delta function exists, **When** this repo's main program runs, **Then** its prior behavior is preserved and the returned string is printed.
2. **Given** the delta function does not yet exist, **When** this issue's implementation would run, **Then** it does not proceed ahead of the child issue (the parent is blocked by the child).
3. **Given** the dependency graph changed, **When** module housekeeping is needed, **Then** `go mod tidy` has been run so the module files are consistent.

---

### Edge Cases

- The spawn target repo may not be initialized yet (the #803 case): on-demand initialization must make the spawn succeed.
- The function name and literal must not be altered or "improved"; a different token would fail the final assertion (the #1308 case).
- If the child issue cannot be created or placed on the board, the parent must not proceed to import a non-existent function.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo MUST export `HelloE2E202610250652037932() string` in its greeting package (`pkg/greeting/greeting.go`), returning exactly `e2e-cross-repo-spawn-20261025-065203-7932`.
- **FR-002**: The delta repo MUST include a test for that function in its greeting package tests (`pkg/greeting/greeting_test.go`), and the test MUST pass.
- **FR-003**: This repo's `main.go` MUST call `greeting.HelloE2E202610250652037932()` and print the result, while preserving its existing behavior.
- **FR-004**: Module files in this repo MUST be consistent with the new dependency use (`go mod tidy` run if needed).
- **FR-005**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the delta repo, so the delta-side work is carried by a child issue there.
- **FR-006**: The delta-side change MUST land before the gamma-side change; this issue's implementation MUST NOT depend on a function that does not yet exist.
- **FR-007**: The exact function name and string literal given in this spec MUST be used verbatim; no alternative token may be substituted.

### Key Entities

- **Parent issue**: This issue (#801) in the gamma repo, owning the `main.go` change.
- **Child issue**: The issue spawned in the delta repo, owning the greeting function and its test.
- **Run-unique token**: `20261025-065203-7932`, embedded in the function name (`E2E202610250652037932`) and the returned literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the delta repo's default branch contains `HelloE2E202610250652037932` returning `e2e-cross-repo-spawn-20261025-065203-7932`, and the delta repo's tests pass.
- **SC-002**: A child issue exists in the delta repo, created from this issue's Plan, and was completed before this issue's change merged.
- **SC-003**: After the run, this repo's program output includes `e2e-cross-repo-spawn-20261025-065203-7932`, and this repo builds and tests cleanly.
- **SC-004**: The scenario's final assertion (the exact token present in the delta repo) passes.

## Assumptions

- "Alpha" in the original request refers to this repo, `handarbeit/fabrik-test-gamma`, and "beta" refers to `handarbeit/fabrik-test-delta`.
- The delta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go`, and this repo's `main.go` already imports or can import that package; if not, adding them is within scope.
- The existing behavior of this repo's `main.go` is kept unchanged other than the added call and print.
- The `fabrik:yolo` label means the pipeline advances without waiting for manual approval between stages.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the single function and its test.
- Changes to Fabrik engine behavior; this issue only exercises the existing spawn mechanism.
- Cleanup or reuse of the per-run token for other scenarios.

## Source References *(optional)*

- #803 — on-demand spawn-target repo initialization (regression covered here)
- #1308 — per-run-unique fixture token (regression covered here)
