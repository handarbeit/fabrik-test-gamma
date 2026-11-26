# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-2657`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202611261505087429() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261126-150508-7429".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202611261505087429() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This issue is a live end-to-end regression scenario for Fabrik's cross-repo decomposition. The work genuinely spans two repositories with a hard ordering dependency. The consumer repo (`handarbeit/fabrik-test-gamma`, "alpha" in the original request) cannot call a function until the provider repo (`handarbeit/fabrik-test-delta`, "beta") has it.

The scenario checks that Plan decomposes the work by spawning a child issue in the provider repo. It is a regression test for two earlier changes:

- #803: on-demand initialisation of a spawn-target repo.
- #1308: a per-run-unique fixture token, so a stale result from an earlier run cannot satisfy this run's assertion.

The function name and string literal embed a per-run-unique token. They must be used exactly as written, because the scenario's final assertion looks for that exact text in the provider repo.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provider repo gains the greeting function (Priority: P1)

The provider repo (`handarbeit/fabrik-test-delta`) exposes a new greeting function that returns a fixed per-run token string, with a test covering it.

**Why this priority**: Everything else depends on it. The consumer repo cannot compile against a function that does not exist, and the scenario's final assertion checks this text in the provider repo.

**Independent Test**: In the provider repo, run its Go tests. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261126-150508-7429`.

**Acceptance Scenarios**:

1. **Given** the provider repo's greeting package, **When** the exported function `HelloE2E202611261505087429()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261126-150508-7429`.
2. **Given** the provider repo's greeting test file, **When** the package tests run, **Then** a test for `HelloE2E202611261505087429` is present and passes.

---

### User Story 2 - Consumer repo uses the new function (Priority: P2)

Once the provider change has landed, the consumer repo (`handarbeit/fabrik-test-gamma`) calls the new function from its entry point and prints the result, in addition to its existing behaviour.

**Why this priority**: It completes the cross-repo change. It can only be done after User Story 1 has landed.

**Independent Test**: With the provider change available, build and run the consumer program. Its output contains `e2e-cross-repo-spawn-20261126-150508-7429` and its previous behaviour is unchanged.

**Acceptance Scenarios**:

1. **Given** the provider function exists, **When** the consumer's `main.go` runs, **Then** it calls `greeting.HelloE2E202611261505087429()` and prints the returned string.
2. **Given** the consumer's previous output, **When** the program runs after the change, **Then** the previous behaviour is still present.
3. **Given** the module dependency set changes, **When** the change is prepared, **Then** `go mod tidy` has been run if needed and the module builds.

---

### User Story 3 - Plan decomposes the work across repos (Priority: P1)

The Plan stage splits the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the provider repo, so the provider change is made as a separate child issue that lands before the consumer change.

**Why this priority**: This is the behaviour the scenario exists to verify. The ordering constraint is what justifies the decomposition.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` carrying the provider work, and this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** work needing changes in both repos in strict order, **When** Plan runs, **Then** it emits a spawn block targeting the provider repo for the provider-side work.
2. **Given** the child issue was spawned, **When** the pipeline continues, **Then** the consumer-side change is not completed before the provider-side change has landed.

---

### Edge Cases

- The function name and string literal must not be altered, normalised or regenerated, since the final assertion depends on the exact text.
- The provider repo may need on-demand initialisation as a spawn target (#803); the scenario passes only if that works without manual setup.
- If the consumer change were attempted before the provider change lands, it would not compile. The ordering constraint must prevent this.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The provider repo's `pkg/greeting/greeting.go` MUST define an exported function `HelloE2E202611261505087429() string` returning exactly `e2e-cross-repo-spawn-20261126-150508-7429`.
- **FR-002**: The provider repo's `pkg/greeting/greeting_test.go` MUST contain a test that verifies `HelloE2E202611261505087429` returns that exact string.
- **FR-003**: The consumer repo's `main.go` MUST call `greeting.HelloE2E202611261505087429()` and print the result, while keeping its existing behaviour.
- **FR-004**: `go mod tidy` MUST be run in the consumer repo if the dependency set requires it, and the consumer module MUST build.
- **FR-005**: The provider-side change MUST land before the consumer-side change.
- **FR-006**: Plan MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta` for the provider-side work.
- **FR-007**: The exact function name and string literal given in this spec MUST be used verbatim, with no substitute identifiers or tokens.

### Key Entities

- **Provider repo ("beta")**: `handarbeit/fabrik-test-delta`. It hosts the greeting function and is the spawn target.
- **Consumer repo ("alpha")**: `handarbeit/fabrik-test-gamma`. It calls the greeting function from its entry point.
- **Per-run token**: the string `20261126-150508-7429` embedded in the function name and returned literal, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue targeting `handarbeit/fabrik-test-delta` is created by Plan's decomposition without manual repo setup.
- **SC-002**: After the provider change lands, the exact text `e2e-cross-repo-spawn-20261126-150508-7429` is present in the provider repo's greeting package, and the provider repo's tests pass.
- **SC-003**: The consumer repo builds and its program output contains `e2e-cross-repo-spawn-20261126-150508-7429` after the provider change has landed.
- **SC-004**: The provider change lands before the consumer change; at no point does the consumer repo reference a function that does not yet exist in the provider repo.

## Assumptions

- "Alpha" in the original request refers to this repo, `handarbeit/fabrik-test-gamma`, and "beta" refers to `handarbeit/fabrik-test-delta`.
- The provider repo already has a `pkg/greeting` package and the consumer's module already depends on it. If not, the provider-side work includes adding what is needed for the stated function to be importable.
- The existing behaviour of the consumer's `main.go` is preserved and only extended.
- File paths and identifiers are included because the requester fixed them as the scenario's contract, not as design choices made here.

## Out of Scope *(optional)*

- Any behaviour in either repo beyond the single greeting function, its test, and the consumer call-and-print.
- Changes to Fabrik's engine, spawn mechanism or Plan stage; this issue only exercises them.
- Generating or choosing a different per-run token.

## Source References *(optional)*

- Mentioned for context: #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
