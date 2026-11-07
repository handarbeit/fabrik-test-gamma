# Feature Specification: e2e base-branch pipeline marker (20261107-202814)

**Feature Branch**: `fabrik/issue-1794`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261107-202814 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261107-202814 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a deliberately trivial change used as a probe for the end-to-end pipeline when an issue carries a `base:<branch>` label. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (review-gate data feed). The content of the change is unimportant. What matters is that the issue moves through every stage and lands correctly on the non-default base branch `e2e-base-branch-20261107-202814`. The `fabrik:cruise` and `base:e2e-base-branch-20261107-202814` labels were applied at filing and must not be changed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added on the non-default base (Priority: P1)

An e2e run files this issue with a `base:` label and expects the pipeline to produce a PR against the named base branch that adds one marker line to a tracked file.

**Why this priority**: It is the only deliverable. It exercises the base-branch pipeline end to end.

**Independent Test**: After the change merges, inspect `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261107-202814`. Its last line is the marker comment.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is applied, **Then** the file ends with the line `<!-- base-branch-pipeline-20261107-202814 -->` and its prior content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261107-202814` label, **When** the PR is opened, **Then** the PR targets that branch, not the repository default, and is linked to this issue.

---

### Edge Cases

- The marker file already exists with no trailing newline: the marker is still appended on its own line, at the very end of the file.
- The marker line is already present (for example on a re-run): it is not duplicated.
- The marker file's parent directory `e2e/markers/` does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` MUST end with exactly one new line, `<!-- base-branch-pipeline-20261107-202814 -->`, as its last line. The file is created first if it does not exist.
- **FR-002**: The change MUST touch only that one file and add only that one line (plus the file's creation if needed).
- **FR-003**: The work MUST target the base branch `e2e-base-branch-20261107-202814`, and any resulting PR MUST be linked to this issue.
- **FR-004**: The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261107-202814`) MUST NOT be added, removed or altered by the work.
- **FR-005**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch consists of exactly one file and one added line (or one new file containing that line).
- **SC-002**: The PR's base branch is `e2e-base-branch-20261107-202814`.
- **SC-003**: The issue moves through the pipeline to completion with no child issues spawned and no manual label changes.

## Assumptions

- "At the very end of the file" means the last line, ending with a newline.
- If the file is created, it contains only the marker line.
- Running the pipeline again on an unchanged base is idempotent: the marker is not added twice.

## Out of Scope

- Any change to other files, code, or documentation.
- Changes to engine behavior; this issue only verifies existing behavior.
- Adding or removing labels, or changing the base branch.
- Decomposing the work into sub-issues.

## Source References

- handarbeit/fabrik#1046 (the base-branch pipeline regression), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
