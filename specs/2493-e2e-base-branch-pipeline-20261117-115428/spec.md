# Feature Specification: e2e base-branch pipeline (20261117-115428)

**Feature Branch**: `fabrik/issue-2493`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261117-115428 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261117-115428 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261117-115428"

## Background

This issue is a live end-to-end probe of the `base:<branch>` pipeline contract: an item whose work targets a non-default base branch must flow through every stage with correct issue↔PR linkage and review-gate data. It is regression coverage for handarbeit/fabrik#1046, and validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial, so that any failure observed comes from the pipeline and not from the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator files a one-line marker change with a `base:<branch>` label and expects the pipeline to carry it from Specify to Done against that base branch.

**Why this priority**: This is the only deliverable. It exercises the whole base-branch contract end to end.

**Independent Test**: After the pipeline finishes, `e2e/markers/base-branch-pipeline.md` on the base branch `e2e-base-branch-20261117-115428` ends with the marker line, and the change arrived through a PR that targeted that branch.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is applied, **Then** the file's last line is `<!-- base-branch-pipeline-20261117-115428 -->` and no other line is changed.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its final line.
3. **Given** the item carries `base:e2e-base-branch-20261117-115428`, **When** its PR is opened, **Then** the PR targets that branch and is linked to the issue.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must still go on its own final line, not be joined to the previous line.
- The marker line is already present (for example on a re-run): it must not be duplicated.
- The labels on the issue must not be added to or removed during the pipeline run.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261117-115428 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created first, so the marker line is its final line.
- **FR-003**: The change MUST touch only that one file and add only that one line.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261117-115428`, as set by the existing `base:` label.
- **FR-005**: The `fabrik:cruise` and `base:e2e-base-branch-20261117-115428` labels MUST NOT be added or removed by the work.
- **FR-006**: The work MUST be treated as a single unit, with no decomposition into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch consists of exactly one file and one added line (or one new file containing only that line).
- **SC-002**: The PR for the work targets `e2e-base-branch-20261117-115428` and is linked to issue #2493 so it is discovered by the pipeline.
- **SC-003**: The item reaches Done with no pause or manual intervention.

## Assumptions

- The marker file's current content, if any, is not otherwise relevant and is left unchanged.
- The target base branch already exists on the remote.
- The marker is written as a plain HTML comment, which has no effect on rendered Markdown.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Fixes to the pipeline itself (#1046, #1047, #1050), which this probe only verifies.
- Adding, removing or changing labels.
- Splitting the work into child issues.
