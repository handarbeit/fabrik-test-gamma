# Feature Specification: e2e cross-repo spawn (decompose across fabrik-test-gamma and fabrik-test-delta)

**Feature Branch**: `fabrik/issue-134`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description:

> ## Goal
>
> Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.
>
> ## What needs to change
>
> ### In handarbeit/fabrik-test-delta (the beta repo)
> - Add an exported function HelloE2E202610171410524438() string to pkg/greeting/greeting.go returning the literal string "e2e-cross-repo-spawn-20261017-141052-4438".
> - Add a test for it in pkg/greeting/greeting_test.go.
>
> ### In handarbeit/fabrik-test-gamma (this repo)
> - Update main.go to call greeting.HelloE2E202610171410524438() (in addition to whatever it already does) and print the result.
> - Run go mod tidy if needed.
>
> ## Constraint
>
> Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.
>
> This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token).
>
> Labels: fabrik:yolo

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. The feature under test is that a single issue whose work spans two repositories in a strict dependency order is decomposed by the Plan stage into a child issue in the other repository, and that the child lands before the parent's own change.

Two earlier fixes are exercised: #803 (on-demand initialisation of the spawn-target repository) and #1308 (a per-run-unique fixture token, so the final assertion cannot be satisfied by leftovers from an earlier run). The names and string literal below embed that per-run token, which is why they must be used exactly as given.

The ordering constraint is real, not incidental: this repo (gamma, referred to in the request as "alpha") cannot import a function that does not yet exist in the delta repo (referred to as "beta"), so the delta change must land first.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Delta repo gains the per-run greeting function (Priority: P1)

The delta repo exposes a new exported greeting function that returns the per-run-unique token string, with a test covering it. This is the dependency the gamma change needs, and it is also what the scenario's final assertion inspects.

**Why this priority**: The gamma change cannot compile without it, and the scenario's final assertion depends on this text landing in the delta repo.

**Independent Test**: In the delta repo, run the Go tests for the greeting package. The new test passes, and the function returns exactly `e2e-cross-repo-spawn-20261017-141052-4438`.

**Acceptance Scenarios**:

1. **Given** the delta repo's greeting package, **When** `HelloE2E202610171410524438()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261017-141052-4438`.
2. **Given** the delta repo's greeting test file, **When** the package tests are run, **Then** a test for `HelloE2E202610171410524438` is present and passes.
3. **Given** the change is merged in the delta repo, **When** its default branch is inspected, **Then** the function name and string literal appear there verbatim.

---

### User Story 2 - Gamma repo prints the delta greeting (Priority: P2)

This repo's `main.go` calls the new delta function in addition to its existing behaviour and prints the result. The change is made only after the delta change has landed.

**Why this priority**: It consumes the delta change, so it is blocked by User Story 1. It is the visible end of the cross-repo chain.

**Independent Test**: With the delta change available, build and run this repo's program. The output still contains its previous output and also contains `e2e-cross-repo-spawn-20261017-141052-4438`.

**Acceptance Scenarios**:

1. **Given** the delta function exists, **When** this repo's program runs, **Then** it prints the string returned by `greeting.HelloE2E202610171410524438()`.
2. **Given** the program's existing behaviour, **When** the new call is added, **Then** the existing behaviour is unchanged.
3. **Given** the dependency set changed, **When** `go mod tidy` is run, **Then** the module files are consistent and the program builds.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The Plan stage recognises that the work spans two repos with a strict order, and emits a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the delta repo. The engine creates the child issue there and makes the parent wait on it, so the delta change lands first.

**Why this priority**: This is the behaviour under test. If decomposition does not happen, the scenario fails regardless of the code changes.

**Independent Test**: After Plan, a child issue exists in `handarbeit/fabrik-test-delta` describing the delta work, and the parent is blocked on it until it is resolved.

**Acceptance Scenarios**:

1. **Given** the parent issue spans both repos, **When** Plan completes, **Then** its output contains a spawn block targeting `handarbeit/fabrik-test-delta`.
2. **Given** the spawn block, **When** the engine processes it, **Then** a child issue is created in the delta repo and the parent is blocked on it.
3. **Given** the child is not yet resolved, **When** the pipeline continues, **Then** the parent's gamma change is not implemented ahead of the delta change.
4. **Given** the child issue's description, **When** it is read, **Then** it carries the exact function name and string literal from this spec.

---

### Edge Cases

- The delta repo may not yet be initialised for spawning (#803): the first spawn must still succeed.
- A fixture token from a previous run may be present in the delta repo (#1308): the final assertion must be satisfied only by this run's exact token, so no other function name or literal may be substituted.
- The spawn target may not be served by this instance's board configuration: the spawn must fail loudly rather than silently creating an unassigned child, per existing engine behaviour.
- If the gamma change would be attempted before the delta change lands, it cannot compile; ordering must prevent that.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The delta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610171410524438() string` in `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610171410524438()` MUST return exactly the string `e2e-cross-repo-spawn-20261017-141052-4438`.
- **FR-003**: The delta repo MUST have a test for `HelloE2E202610171410524438` in `pkg/greeting/greeting_test.go`, and it MUST pass.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610171410524438()` in addition to its existing behaviour and print the result.
- **FR-005**: This repo's module files MUST be consistent after the change (`go mod tidy` run if needed) and the program MUST build.
- **FR-006**: The delta change MUST land before the gamma change is implemented or merged.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the delta repo, so the delta work is carried out as a child issue there.
- **FR-008**: The function name and string literal given in this spec MUST be used verbatim in every artefact (child issue, code, test, printed output); no other name or literal may be invented or substituted.

### Key Entities

- **Parent issue (#134, gamma repo)**: Owns the `main.go` change and the decomposition; blocked on the child.
- **Child issue (delta repo)**: Owns the greeting function and its test; must land first.
- **Per-run fixture token**: `20261017-141052-4438`, embedded in the function name and string literal, making the final assertion specific to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A child issue exists in `handarbeit/fabrik-test-delta`, created by the Plan stage's spawn block, and the parent was blocked on it until it was resolved.
- **SC-002**: The default branch of the delta repo contains the text `HelloE2E202610171410524438` and the literal `e2e-cross-repo-spawn-20261017-141052-4438`, exactly once each in the code and in the test as specified.
- **SC-003**: The delta repo's greeting package tests pass, including the new test.
- **SC-004**: This repo builds, and its program's output includes `e2e-cross-repo-spawn-20261017-141052-4438` alongside its prior output.
- **SC-005**: The delta change merged before the gamma change, as shown by merge order.

## Assumptions

- The delta repo already has a `pkg/greeting` package with `greeting.go` and `greeting_test.go` (or the child issue may create them); this spec does not depend on which.
- The delta repo is reachable as a Go module by this repo, so a normal dependency update (`go mod tidy`) resolves the new function once the delta change has landed.
- "Alpha" and "beta" in the original request refer to this repo (gamma) and the delta repo respectively; the repo names given in the request are authoritative.
- The scenario runs under `fabrik:yolo`, so stages advance without waiting for human approval of this spec.
- The exact function name and string literal appearing in this spec are mandated by the request as test fixtures, so they are stated here despite this spec otherwise staying at the what/why level.

## Out of Scope

- Any change to Fabrik's engine, the spawn mechanism or the Plan stage itself; this issue only exercises them.
- Other behaviour changes to `main.go` or to the greeting package beyond what is described.
- Fixes for #803 or #1308; they are covered by this scenario but not part of this work.
- Any use of a different function name or string literal.

## Source References

- #803 (on-demand spawn-target repo init), mentioned for context
- #1308 (per-run-unique fixture token), mentioned for context
