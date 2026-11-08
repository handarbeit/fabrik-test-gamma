# Feature Specification: e2e cross-repo spawn

**Feature Branch**: `fabrik/issue-1846`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

Verify cross-repo decomposition works end-to-end. Plan should spawn a sub-issue in handarbeit/fabrik-test-delta because the work requires changes in both repos in strict order.

## What needs to change

### In handarbeit/fabrik-test-delta (the beta repo)
- Add an exported function HelloE2E202611080205286076() string to pkg/greeting/greeting.go returning the literal string \"e2e-cross-repo-spawn-20261108-020528-6076\".
- Add a test for it in pkg/greeting/greeting_test.go.

### In handarbeit/fabrik-test-gamma (this repo)
- Update main.go to call greeting.HelloE2E202611080205286076() (in addition to whatever it already does) and print the result.
- Run go mod tidy if needed.

## Constraint

Beta must land first because alpha cannot import a function that does not yet exist. Plan should emit a FABRIK_SPAWN_CHILD_BEGIN block targeting the beta repo to decompose. Use the exact function name and string literal given above verbatim — do not invent your own; a per-run-unique token is already baked into them, and this scenario's final assertion depends on that exact text landing in the beta repo.

This is an e2e regression test for #803 (on-demand spawn-target repo init) and #1308 (per-run-unique fixture token)."

## Background

This issue is an end-to-end regression scenario for Fabrik's cross-repo decomposition. It covers two earlier fixes:

- **#803**: on-demand initialisation of a spawn-target repository.
- **#1308**: a per-run-unique fixture token, so that each run's final assertion can tell its own output from leftovers of earlier runs.

The work is deliberately split across two repositories with a hard ordering dependency. The shared greeting function lives in the "beta" repo (`handarbeit/fabrik-test-delta`). The "alpha" repo (this one, `handarbeit/fabrik-test-gamma`) consumes it from its entry point, so the function must exist in beta before alpha can use it. A correct pipeline run must therefore decompose the work. The beta change becomes a child issue in the other repo, and this (parent) issue's own change waits until that child has landed.

The function name and string literal embed a per-run-unique token (`20261108-020528-6076`). The scenario's final assertion checks for that exact text in the beta repo, so both must be used character for character.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Beta repo gains the greeting function (Priority: P1)

The beta repo exposes a new greeting function whose return value is the run-unique token string, covered by a test. This is the prerequisite slice. The final assertion for the whole scenario depends on it.

**Why this priority**: Alpha cannot compile against a function that does not exist. The scenario's final assertion also inspects this text in the beta repo.

**Independent Test**: In the beta repo, run the greeting package's tests. The new test passes, and calling the function returns exactly `e2e-cross-repo-spawn-20261108-020528-6076`.

**Acceptance Scenarios**:

1. **Given** the beta repo's greeting package, **When** `HelloE2E202611080205286076()` is called, **Then** it returns exactly `e2e-cross-repo-spawn-20261108-020528-6076`.
2. **Given** the beta repo's greeting tests, **When** the test suite runs, **Then** a test covering the new function exists and passes.
3. **Given** the work was decomposed, **When** the beta change is made, **Then** it is made by a child issue in the beta repo and not by this issue.

---

### User Story 2 - Alpha repo consumes and prints the beta function (Priority: P1)

After beta has landed, the alpha repo's entry point calls the new beta function in addition to its existing behaviour and prints the result.

**Why this priority**: This is the user-visible end of the cross-repo change. It can only be done after User Story 1.

**Independent Test**: Build and run the alpha program with the beta change available. Its output includes `e2e-cross-repo-spawn-20261108-020528-6076`, and its existing output is unchanged.

**Acceptance Scenarios**:

1. **Given** the beta function is available to alpha, **When** alpha's entry point runs, **Then** the printed output contains `e2e-cross-repo-spawn-20261108-020528-6076`.
2. **Given** alpha's existing output before this change, **When** the entry point runs after the change, **Then** the pre-existing output is still produced.
3. **Given** the dependency set changed, **When** the module is built, **Then** it builds cleanly with dependency metadata consistent (tidied if needed).

---

### User Story 3 - Pipeline decomposes across repos in the right order (Priority: P1)

The pipeline itself recognises that the work spans two repos with a strict order. It spawns a child issue targeting the beta repo, and it holds the parent's implementation until the child is done.

**Why this priority**: This is what the scenario exists to verify (#803, #1308).

**Independent Test**: Run the issue through the pipeline. A child issue appears in the beta repo, linked as a blocker of this issue, and this issue's implementation does not proceed until the child is complete.

**Acceptance Scenarios**:

1. **Given** this issue reaches Plan, **When** Plan decomposes the work, **Then** it declares a spawn targeting `handarbeit/fabrik-test-delta`.
2. **Given** the beta repo has not been prepared for spawning, **When** the child is created, **Then** the spawn-target repo is initialised on demand (#803).
3. **Given** the child issue has not landed, **When** the parent would otherwise proceed, **Then** the parent remains blocked on the child.
4. **Given** the child has landed, **When** the parent proceeds, **Then** its change lands afterwards.

---

### Edge Cases

- The function name or string literal is altered, shortened or re-tokenised: the final assertion fails. Both must appear verbatim.
- A leftover token from an earlier run exists in the beta repo: the unique token must still distinguish this run's output (#1308).
- The beta repo is not yet set up for spawn targets: it must be initialised on demand (#803), not fail the spawn.
- Alpha is attempted before beta lands: this must not happen, because the parent is blocked on the child.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The beta repo (`handarbeit/fabrik-test-delta`) MUST export a function named exactly `HelloE2E202611080205286076`, taking no arguments and returning a string, in its greeting package.
- **FR-002**: That function MUST return exactly the literal `e2e-cross-repo-spawn-20261108-020528-6076`.
- **FR-003**: The beta repo MUST include a test for that function in its greeting package, and that test MUST pass.
- **FR-004**: The alpha repo (`handarbeit/fabrik-test-gamma`, this repo) entry point MUST call `greeting.HelloE2E202611080205286076()` from the beta repo's greeting package and print the returned value, in addition to its existing behaviour.
- **FR-005**: The alpha repo's dependency metadata MUST be tidied if the change requires it, so that the module builds cleanly.
- **FR-006**: The beta-repo change MUST be delivered through a child issue spawned in the beta repo, and the alpha-repo change MUST NOT land before the beta-repo change.
- **FR-007**: The function name and string literal MUST be used verbatim as given, with no substitution of a different name or token.

### Key Entities

- **Parent issue (alpha)**: This issue, in `handarbeit/fabrik-test-gamma`. It owns the entry-point change.
- **Child issue (beta)**: The spawned issue in `handarbeit/fabrik-test-delta`. It owns the greeting function and its test.
- **Run-unique token**: `20261108-020528-6076`, embedded in the function name and the returned literal.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the beta repo contains the exported function `HelloE2E202611080205286076` returning exactly `e2e-cross-repo-spawn-20261108-020528-6076`, and its test passes.
- **SC-002**: After the run, the alpha program's output includes `e2e-cross-repo-spawn-20261108-020528-6076`, and its previous output is preserved.
- **SC-003**: The beta-repo change is merged before the alpha-repo change, and exactly one child issue was spawned in the beta repo for it.
- **SC-004**: The scenario's final assertion, which searches the beta repo for the run-unique text, passes.

## Assumptions

- The beta repo already has a greeting package with a source file and a test file, as the request describes.
- "In addition to whatever it already does" means the existing output and behaviour of the alpha entry point are retained.
- Where the printed value appears in alpha's output (its own line, for example) is not significant, as long as the exact literal is printed.
- The alpha module is already able to import from the beta repo's module. Only dependency tidying is needed, as the request says.
- The `fabrik:yolo` label lets the pipeline advance through stages without waiting for manual approval.

## Out of Scope

- Any greeting-package changes in the beta repo beyond the one function and its test.
- Refactoring or changing alpha's existing behaviour.
- Changes to the Fabrik engine or its spawn mechanism. This scenario only exercises them.
- Using a different function name or token, even if it looks equivalent.

## Source References

- #803: on-demand spawn-target repo init
- #1308: per-run-unique fixture token
