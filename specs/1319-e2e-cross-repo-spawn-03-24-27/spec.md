# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1319`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610300724272450() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261030-072427-2450".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610300724272450() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).

## Background

This is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It guards two earlier fixes: #803 (on-demand initialisation of a spawn-target repo) and #1308 (a per-run-unique fixture token, so that runs cannot be confused with each other).

The work spans two repositories with a hard ordering. Two repos are involved:

- `handarbeit/fabrik-test-delta` (the "beta" repo) must gain a new greeting function, with a test.
- `handarbeit/fabrik-test-gamma` (this repo, the "alpha" repo) must then call that function from `main.go`.

Gamma cannot compile against a function that does not yet exist in delta, so delta's change must land first. That dependency is what should push the Plan stage to decompose the work, spawning a child issue in the delta repo.

The function name and string literal embed a per-run-unique token (`20261030-072427-2450`). The scenario's final assertion checks that this exact text lands in the delta repo, so the values must be used exactly as given.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Delta repo gains the unique greeting function (Priority: P1)

The delta repo exposes a new exported function that returns this run's unique token string, and it has a test.

**Why this priority**: It is the prerequisite for everything else. The scenario's final assertion checks that this text exists in the delta repo.

**Independent Test**: In the delta repo, run the package tests for `pkg/greeting`. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261030-072427-2450`.

**Acceptance Scenarios**:

1. **Given** the delta repo before the change, **When** the work lands, **Then** `pkg/greeting/greeting.go` contains an exported function `HelloE2E202610300724272450() string` returning the literal `"e2e-cross-repo-spawn-20261030-072427-2450"`.
2. **Given** the new function, **When** the delta tests run, **Then** a test in `pkg/greeting/greeting_test.go` asserts that return value and passes.

---

### User Story 2 - Gamma repo calls the new function and prints the result (Priority: P2)

After delta's change has landed, this repo's `main.go` calls `greeting.HelloE2E202610300724272450()` and prints the result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1. It is the half of the work that proves the cross-repo dependency was honoured in order.

**Independent Test**: Once the delta change is available, build and run this repo. The program still does what it did before and also prints `e2e-cross-repo-spawn-20261030-072427-2450`.

**Acceptance Scenarios**:

1. **Given** the delta function exists, **When** `main.go` is run, **Then** the output includes the string `e2e-cross-repo-spawn-20261030-072427-2450` alongside the program's pre-existing output.
2. **Given** the module dependencies after the change, **When** the build runs, **Then** it succeeds, running `go mod tidy` if that is needed to make it so.

---

### User Story 3 - Plan decomposes the work across repos in dependency order (Priority: P1)

The Plan stage recognises that the work needs ordered changes in two repositories and decomposes it by spawning a child issue in the delta repo. The gamma change is then sequenced after it.

**Why this priority**: Exercising cross-repo decomposition is the entire purpose of the scenario. Without a spawn, #803 is not tested.

**Independent Test**: After Plan completes, a child issue exists in `handarbeit/fabrik-test-delta` that carries the delta-side work, and this issue is blocked by it.

**Acceptance Scenarios**:

1. **Given** this spec, **When** Plan completes, **Then** its output contains a `FABRIK_SPAWN_CHILD_BEGIN` block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawned child, **When** its description is read, **Then** it uses the exact function name and string literal from this spec, with no invented substitutes.
3. **Given** the child is not yet complete, **When** the pipeline continues, **Then** the gamma-side change does not land ahead of the delta-side change.

---

### Edge Cases

- The delta repo may not have been initialised for spawning yet. This is the situation #803 covers, so on-demand initialisation of the spawn target must succeed.
- A function name or literal that differs even slightly from the specified text would defeat the scenario's final assertion. There must be no normalisation, renaming or paraphrasing.
- `main.go` may already print other output. That output must be preserved, and the new call is additive.
- If a `go mod tidy` run changes nothing, that is acceptable. It is only to be run if needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST expose an exported function `HelloE2E202610300724272450() string` in `pkg/greeting/greeting.go`.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261030-072427-2450`.
- **FR-003**: The delta repo MUST include a test for the function in `pkg/greeting/greeting_test.go` that passes.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610300724272450()` and print the result, while keeping its existing behaviour.
- **FR-005**: This repo's module files MUST be consistent after the change (`go mod tidy` run if needed), so the project builds.
- **FR-006**: The delta-side change MUST land before the gamma-side change.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block that targets the delta repo.
- **FR-008**: The exact function name and string literal given in this spec MUST be used verbatim everywhere they appear. No alternative names or values may be introduced.

### Key Entities

- **Alpha repo**: `handarbeit/fabrik-test-gamma`, this repo. It consumes the new function.
- **Beta repo**: `handarbeit/fabrik-test-delta`. It provides the new function and is the spawn target.
- **Per-run-unique token**: `20261030-072427-2450`. It is embedded in the function name and the returned literal, and the scenario's final assertion keys on it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the delta repo's default branch contains the exact text `e2e-cross-repo-spawn-20261030-072427-2450` returned by `HelloE2E202610300724272450()`.
- **SC-002**: The delta repo's `pkg/greeting` tests pass, including the new test.
- **SC-003**: Running this repo's program prints `e2e-cross-repo-spawn-20261030-072427-2450` in addition to its prior output, and the project builds cleanly.
- **SC-004**: A child issue exists in the delta repo, spawned from this issue's Plan, and it merged before the gamma change did.

## Assumptions

- "Alpha" in the original request refers to this repo (`handarbeit/fabrik-test-gamma`) and "beta" to `handarbeit/fabrik-test-delta`. The names are used interchangeably in the original text.
- `pkg/greeting` already exists in the delta repo, as the file paths in the request imply, and this repo can import it. If the import path or module requirement is missing, wiring it up is part of the gamma-side work and `go mod tidy` is how it gets resolved.
- The `fabrik:yolo` label means stages advance automatically, so no human approval of this spec is expected.
- The token `20261030-072427-2450` is final. This spec treats it as a given and does not generate or alter it.

## Out of Scope

- Any greeting-related behaviour in either repo beyond the single function, its test and the one call in `main.go`.
- Changes to Fabrik's engine, to its spawn mechanism or to its board configuration. This scenario only exercises the existing mechanism.
- Refactoring or cleaning up existing code in either repo.
- Cleaning up the spawned child issue or the delta-side artifacts after the run.

## Source References

- #803: on-demand spawn-target repo initialisation (regression guarded by this scenario).
- #1308: per-run-unique fixture token (regression guarded by this scenario).
