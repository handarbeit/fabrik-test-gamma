# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-418`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610190622361759() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261019-062236-1759\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610190622361759() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for cross-repo decomposition. It exercises two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: a per-run-unique fixture token, so each run's final assertion can tell its own output from a previous run's.

The feature needs changes in two repositories, and they depend on each other in one direction. The consumer (`handarbeit/fabrik-test-gamma`, "alpha", this repo) imports a new function from the provider (`handarbeit/fabrik-test-delta`, "beta"). The provider's change must therefore land first. Because the work spans repositories, the Plan stage is expected to decompose it by spawning a child issue in the beta repo rather than doing everything in this issue.

The function name and string literal contain a per-run-unique token. The scenario's final assertion checks that this exact text lands in the beta repo, so both must be used exactly as written.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Provider repo exposes the new greeting (Priority: P1)

The beta repo (`handarbeit/fabrik-test-delta`) gains a new exported greeting function and a test for it. This is the prerequisite for everything else.

**Why this priority**: The consumer cannot compile against a function that does not exist, and the scenario's final assertion looks for this text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes and asserts the exact literal returned by the new function.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610190622361759()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261019-062236-1759`.
2. **Given** the beta repo's greeting tests, **When** the suite runs, **Then** a test covering the new function exists and passes.

---

### User Story 2 - Consumer repo uses the new greeting (Priority: P2)

This repo's program calls the new beta function and prints the result, in addition to its existing behaviour.

**Why this priority**: It depends on User Story 1 and proves the cross-repo dependency works end to end.

**Independent Test**: With the beta change available, build and run this repo's program. The output contains `e2e-cross-repo-spawn-20261019-062236-1759`, and the existing output is still present.

**Acceptance Scenarios**:

1. **Given** the beta repo provides `HelloE2E202610190622361759()`, **When** this repo's program runs, **Then** it prints the string returned by `greeting.HelloE2E202610190622361759()`.
2. **Given** the program's existing output, **When** the new call is added, **Then** the existing output is unchanged.
3. **Given** the dependency changes, **When** the module is tidied and built, **Then** the build succeeds.

---

### User Story 3 - Plan decomposes the work across repos in order (Priority: P1)

The pipeline's Plan stage splits the work by spawning a child issue in the beta repo. The beta change lands before this repo's change.

**Why this priority**: This is the behaviour the scenario exists to verify (#803, #1308).

**Independent Test**: After Plan runs, a child issue exists in `handarbeit/fabrik-test-delta` carrying the beta-side work, and this issue is blocked on it.

**Acceptance Scenarios**:

1. **Given** the work spans two repos, **When** Plan completes, **Then** it has emitted a child-spawn declaration targeting the beta repo.
2. **Given** the spawned beta child is not yet complete, **When** the pipeline continues, **Then** this repo's change does not land before the beta change.
3. **Given** the spawn target repo has not been initialised on demand, **When** the child is spawned, **Then** the spawn succeeds (#803).

---

### Edge Cases

- A previous run left a different token in either repo. The scenario's assertion must match only this run's token (#1308).
- The beta change lands but the alpha change is not yet made. The alpha build must not reference the function before it exists.
- `go mod tidy` is not needed if the dependency state is already correct. It should be run only if needed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function `HelloE2E202610190622361759() string` from its greeting package, at `pkg/greeting/greeting.go`.
- **FR-002**: `HelloE2E202610190622361759()` MUST return exactly the literal `e2e-cross-repo-spawn-20261019-062236-1759`.
- **FR-003**: The beta repo MUST include a test for `HelloE2E202610190622361759()` in `pkg/greeting/greeting_test.go` that asserts the returned literal.
- **FR-004**: This repo's `main.go` MUST call `greeting.HelloE2E202610190622361759()` and print the result, in addition to its existing behaviour.
- **FR-005**: The dependency metadata of this repo MUST be tidied (`go mod tidy`) if needed so that the program builds.
- **FR-006**: The function name and string literal MUST be used verbatim as given; no alternative name or token may be substituted.
- **FR-007**: The Plan stage MUST decompose the work by emitting a `FABRIK_SPAWN_CHILD_BEGIN` block targeting the beta repo.
- **FR-008**: The beta repo's change MUST land before this repo's change.

### Key Entities *(if applicable)*

- **Beta repo**: `handarbeit/fabrik-test-delta`, the provider of the greeting function.
- **Alpha repo**: `handarbeit/fabrik-test-gamma` (this repo), the consumer that prints the greeting.
- **Run token**: the per-run-unique text embedded in the function name and string literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline finishes, the beta repo's default branch contains `HelloE2E202610190622361759` and the literal `e2e-cross-repo-spawn-20261019-062236-1759`, which the scenario's final assertion checks.
- **SC-002**: The beta repo's greeting tests pass, including the new test.
- **SC-003**: This repo builds and its program output contains `e2e-cross-repo-spawn-20261019-062236-1759`.
- **SC-004**: A child issue was created in the beta repo by the Plan stage's spawn, and it reached completion before this repo's change landed.

## Assumptions

- The existing greeting package in the beta repo is the right home for the new function, as the issue states.
- This repo already depends on the beta repo's module, or can add it as part of this change. No new repository needs to be created manually.
- The `go mod tidy` step is conditional, as the issue says.
- Whether the pipeline picks up the beta repo's updated version through a dependency bump or an equivalent mechanism is a planning detail, not a product requirement.

## Out of Scope *(optional)*

- Any change to the greeting package beyond the one new function and its test.
- Any change to this repo's existing behaviour.
- Changes to Fabrik's engine or spawn logic. This issue only exercises it.
- Cleanup of fixtures from earlier runs.

## Source References *(optional)*

- #803: on-demand spawn-target repo init.
- #1308: per-run-unique fixture token.
