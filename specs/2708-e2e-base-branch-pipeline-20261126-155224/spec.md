# Feature Specification: e2e base-branch pipeline (20261126-155224)

**Feature Branch**: `fabrik/issue-2708`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261126-155224 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261126-155224 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261126-155224"

## Background

This issue is a live end-to-end probe of the `base:<branch>` pipeline contract. An item labelled `base:<branch>` must be forked from, rebased onto and have its PR target that non-default base branch. Earlier fixes (#1047 for issue↔PR linkage, #1050 for the review-gate data feed) addressed regressions first reported in #1046. This item exercises that path with the smallest possible change, so that any failure points at the pipeline rather than at the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline takes this issue through its stages and produces a PR that adds one marker line to `e2e/markers/base-branch-pipeline.md`, targeting the `e2e-base-branch-20261126-155224` base branch.

**Why this priority**: It is the only deliverable, and it proves the base-branch pipeline works end to end.

**Independent Test**: Inspect the merged result on `e2e-base-branch-20261126-155224`. `e2e/markers/base-branch-pipeline.md` must end with the marker line, and the diff must contain nothing else.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is implemented, **Then** the file exists and its last line is `<!-- base-branch-pipeline-20261126-155224 -->`.
2. **Given** the file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the issue carries the `base:e2e-base-branch-20261126-155224` label, **When** the PR is created, **Then** the PR targets that branch, not the repository's other branches, and is linked to this issue.

---

### Edge Cases

- The file may already end without a trailing newline; the marker must still land on its own final line.
- The marker line must appear exactly once; a re-run must not append a duplicate.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261126-155224 -->`, created first if it does not exist.
- **FR-002**: The change touches exactly one file and adds exactly one line (beyond creating the file, if needed).
- **FR-003**: The resulting PR targets the base branch `e2e-base-branch-20261126-155224` and is linked to this issue.
- **FR-004**: No labels are added to or removed from the issue as part of the work.
- **FR-005**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the final line of `e2e/markers/base-branch-pipeline.md` on the target base branch equals the marker string, and it occurs exactly once.
- **SC-002**: The PR's diff against its base contains one changed file.
- **SC-003**: The PR's base branch is `e2e-base-branch-20261126-155224` and the issue↔PR linkage and review-gate data feed work without manual intervention.

## Assumptions

- `e2e/markers/base-branch-pipeline.md` does not yet exist (other marker files do); it is created with the marker as its content.
- The `fabrik:cruise` label lets the item auto-advance through stages without further human approval.

## Out of Scope *(optional)*

- Any change beyond the one marker line.
- Decomposing the work into child issues.
- Changing labels on the issue.
- Changes to the Fabrik engine itself.
