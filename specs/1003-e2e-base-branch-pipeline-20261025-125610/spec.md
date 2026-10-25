# Feature Specification: E2E base-branch pipeline marker

**Feature Branch**: `fabrik/issue-1003`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261025-125610 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261025-125610 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end run of the pipeline for an item carrying a `base:<branch>` label, so the work targets a non-default base branch. It is regression coverage for handarbeit/fabrik#1046. It validates two fixes: #1047 (issue-to-PR linkage for PRs targeting a non-default base) and #1050 (the data feed to the review gate). The code change is deliberately trivial, so that any failure points at the pipeline and not at the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline takes the item from Specify through to a merged change. The change adds one marker line to `e2e/markers/base-branch-pipeline.md`, on a PR that targets the `base:` branch.

**Why this priority**: This is the whole purpose of the issue. It exercises the non-default-base path end to end.

**Independent Test**: After the change merges, `e2e/markers/base-branch-pipeline.md` on `e2e-base-branch-20261025-125610` ends with the marker line, and no other file differs from the base.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is made, **Then** the file exists and its last line is `<!-- base-branch-pipeline-20261025-125610 -->`.
2. **Given** the file already exists, **When** the change is made, **Then** the marker line is appended at the very end and the existing content is unchanged.
3. **Given** the item carries `base:e2e-base-branch-20261025-125610`, **When** the PR is created, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The marker file already exists with no trailing newline: the marker must still sit on its own line.
- The marker line is already present (a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261025-125610 -->`. The file is created if it does not exist.
- **FR-002**: The change touches only that one file and adds only that one line, apart from existing content that is left as is.
- **FR-003**: The PR targets the branch `e2e-base-branch-20261025-125610` and is linked to this issue (#1047 contract).
- **FR-004**: The pipeline's review gate receives the PR's review data correctly for a non-default-base PR (#1050 contract).
- **FR-005**: No labels are added to or removed from the issue by the change itself.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch is exactly one added line in one file.
- **SC-002**: The PR is linked to the issue and targets the `base:` branch, with no manual linking.
- **SC-003**: The item moves through the pipeline without a pause caused by missing issue-PR linkage or missing review data.

## Assumptions

- The branch `e2e-base-branch-20261025-125610` exists on the remote and is the intended target (it is stated to be the repository's default base for this run).
- Labels `fabrik:cruise` and `base:e2e-base-branch-20261025-125610` are already set at filing and stay as they are.
- The work is a single task. Planning should not split it into child issues.

## Out of Scope *(optional)*

- Any change to engine code, tests or docs.
- Adding or removing labels.
- Multi-repo scenarios.
- Any marker content other than the one specified line.
