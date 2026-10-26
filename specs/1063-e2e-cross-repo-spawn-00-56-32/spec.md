# Feature Specification: E2E cross-repo spawn (beta-first decomposition)

**Feature Branch**: `fabrik/issue-1063`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202610260456321856() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261026-045632-1856\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202610260456321856() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This is an end-to-end regression scenario for cross-repo decomposition. A single piece of work needs changes in two repositories that depend on each other in one direction: the greeting library in `handarbeit/fabrik-test-delta` (the "beta" repo) must gain a new function before `handarbeit/fabrik-test-gamma` (this repo, "alpha") can import and use it. The pipeline is expected to recognize this ordering and decompose the work into a sub-issue in the beta repo that lands first, with the change in this repo following.

The scenario guards two earlier behaviors:

- **#803**: on-demand initialization of the repository that a spawned child targets.
- **#1308**: a per-run-unique fixture token, so a run's final assertion can tell its own work from artifacts of earlier runs. The function name and string literal below already contain this token and must be used exactly as written.

This repository's `main.go` already imports the beta repo's greeting package and prints the result of one helper call per line. The change here extends that existing pattern by one call.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the new greeting function (Priority: P1)

The greeting package in the beta repo exposes a new exported function that returns this run's unique token string, with a test proving it.

**Why this priority**: Everything else depends on it. This repo cannot reference a function that does not exist upstream, and the scenario's final assertion checks that this text landed in the beta repo.

**Independent Test**: In the beta repo alone, run its test suite and confirm the new function exists, returns the exact literal, and has a passing test.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202610260456321856()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261026-045632-1856`.
2. **Given** the beta repo's greeting tests, **When** the suite runs, **Then** a test covering the new function passes.

---

### User Story 2 - This repo prints the new greeting (Priority: P2)

After the beta change is available, this repo's entry point calls the new function and prints its result in addition to its existing output.

**Why this priority**: This is the consumer half of the cross-repo change. It can only be completed after User Story 1 has landed.

**Independent Test**: Run this repo's program and its tests after the dependency has been updated, and confirm the new token string appears in the output while all previous output is still present.

**Acceptance Scenarios**:

1. **Given** the beta repo change has landed and this repo's dependency resolves to a version containing it, **When** the program runs, **Then** its output includes `e2e-cross-repo-spawn-20261026-045632-1856` in addition to all prior output lines.
2. **Given** the change is complete, **When** the module is built and tested, **Then** the build and tests succeed with the dependency metadata consistent (tidy).

---

### User Story 3 - Pipeline decomposes the work across repos in order (Priority: P1)

The Plan stage recognizes the strict ordering between the two repos and spawns a child issue targeting the beta repo, so the beta change is delivered first and this repo's change is sequenced after it.

**Why this priority**: Verifying this decomposition is the actual purpose of the scenario.

**Independent Test**: Observe that a child issue is created in the beta repo from this issue, that it is blocking this issue, and that the beta function exists in the beta repo before this repo's change merges.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan identifies that the beta change must land first, **Then** it declares a spawned child targeting `handarbeit/fabrik-test-delta`.
2. **Given** the child issue exists, **When** the pipeline runs, **Then** the beta change lands before this repo's change is merged.
3. **Given** the full run completes, **When** the final assertion inspects the beta repo, **Then** it finds the exact function name and literal string from this issue.

---

### Edge Cases

- The beta repo may not yet be initialized for spawn targets; on-demand initialization (#803) must make the spawn succeed without manual setup.
- The name and literal must match exactly. Any substituted or regenerated token makes the run's final assertion fail.
- If this repo's dependency on the beta repo is not refreshed to include the new function, the build fails; the dependency must resolve to a version containing the new function.
- Existing output lines in this repo's program must not be removed or reordered in a way that breaks existing tests.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo's greeting package MUST export a function named exactly `HelloE2E202610260456321856` that takes no arguments and returns a string.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261026-045632-1856`.
- **FR-003**: The beta repo's greeting tests MUST include a test that verifies FR-002.
- **FR-004**: This repo's program MUST call `greeting.HelloE2E202610260456321856()` and print the returned value, in addition to its existing output.
- **FR-005**: This repo's module metadata MUST be consistent (tidy) after the change, with its dependency on the beta repo resolving to a version that contains the new function.
- **FR-006**: The Plan stage MUST declare the beta-repo work as a spawned child targeting `handarbeit/fabrik-test-delta`, so that it is delivered before this repo's change.
- **FR-007**: The function name and string literal MUST be used verbatim as given; no alternative names or tokens may be introduced.

### Key Entities

- **Beta repo (`handarbeit/fabrik-test-delta`)**: The upstream library repository holding the greeting package; the target of the spawned child issue.
- **Alpha repo (`handarbeit/fabrik-test-gamma`)**: This repository, the consumer of the greeting package.
- **Per-run-unique token**: The text embedded in the function name and literal, which ties this run's final assertion to this run's work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo's default branch contains the function `HelloE2E202610260456321856` returning `e2e-cross-repo-spawn-20261026-045632-1856`, with a passing test.
- **SC-002**: After the run, this repo's program output contains `e2e-cross-repo-spawn-20261026-045632-1856` and its build and tests pass.
- **SC-003**: A child issue targeting the beta repo was created from this issue by Plan, and the beta change merged before this repo's change.
- **SC-004**: The run completes without manual repository setup in the beta repo.

## Assumptions

- "Alpha" in the original request refers to this repo (`fabrik-test-gamma`) and "beta" to `fabrik-test-delta`; the names in the request are authoritative.
- The new output is an additional line printed after the existing greeting lines; exact position among them is not significant as long as existing output is preserved.
- The dependency update in this repo is done through the normal module mechanism once the beta change is available; the specific version reference is determined during planning and implementation.
- The `fabrik:yolo` label means the issue advances automatically through stages and merges when validated.

## Out of Scope

- Any change to the beta repo beyond the single function and its test.
- Changes to the pipeline or engine behavior being tested; this scenario only exercises it.
- Cleanup of earlier e2e scenario artifacts or fixtures in either repo.
