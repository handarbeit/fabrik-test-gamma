# Feature Specification: E2E base-branch pipeline marker

**Feature Branch**: `fabrik/issue-157`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261017-144429 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261017-144429 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression probe for the `base:<branch>` pipeline contract, i.e. work that targets a non-default base branch. It guards against a recurrence of handarbeit/fabrik#1046 and exercises two fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review gate receiving the data it needs for such PRs). The code change is deliberately trivial so that any failure comes from the pipeline mechanics, not from the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator runs the issue through the full pipeline with the `base:e2e-base-branch-20261017-144429` label. The pipeline produces a PR that adds one marker line to the marker file and targets that base branch.

**Why this priority**: This is the whole change. Landing it on the correct base shows the base-branch contract holds end to end.

**Independent Test**: After the pipeline finishes, inspect the PR and the marker file on the base branch.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is made, **Then** the file exists and contains the marker line `<!-- base-branch-pipeline-20261017-144429 -->`.
2. **Given** the marker file already exists, **When** the change is made, **Then** the marker line is the last line of the file and all prior content is unchanged.
3. **Given** the issue carries the `base:e2e-base-branch-20261017-144429` label, **When** the PR is created, **Then** it targets that branch, not the repository's other branches.

---

### User Story 2 - Pipeline linkage and review gate work for a non-default base (Priority: P2)

The PR for this issue is linked back to the issue, and the review gate sees the PR data it needs, even though the PR's base is not the default branch.

**Why this priority**: These are the #1047 and #1050 regressions the probe exists to catch. They only show up while the pipeline is running.

**Independent Test**: Observe the pipeline progressing through Review and Validate for this issue without manual intervention beyond normal gates.

**Acceptance Scenarios**:

1. **Given** the PR targets the non-default base, **When** the pipeline looks for the issue's PR, **Then** it finds the PR and its comments and reviews.
2. **Given** the item has reached the review gate, **When** reviewers respond or none are expected, **Then** the gate decides on real PR data and does not stall or pause for missing data.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still be on its own line and must not be joined to the previous last line.
- The marker line is already present at the end of the file: it must not be added a second time.
- The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261017-144429`) must not be added or removed as part of this work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker file `e2e/markers/base-branch-pipeline.md` MUST end with exactly one occurrence of the line `<!-- base-branch-pipeline-20261017-144429 -->` after this change.
- **FR-002**: If the marker file does not exist, it MUST be created, containing the marker line.
- **FR-003**: If the marker file exists, its existing content MUST be preserved, and the marker line MUST be appended as a new final line.
- **FR-004**: The change MUST consist of that one file and one added line, with no other file or content changes in the PR.
- **FR-005**: The resulting PR MUST target the base branch `e2e-base-branch-20261017-144429`.
- **FR-006**: The work MUST be a single unit; it MUST NOT be decomposed into child issues.
- **FR-007**: The issue's labels MUST NOT be added or removed by this work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against `e2e-base-branch-20261017-144429` touches exactly one file and adds exactly one line.
- **SC-002**: The PR is linked to issue #157 and the pipeline reaches completion without a pause caused by missing PR linkage or review-gate data.
- **SC-003**: No child issues are spawned for this item.

## Assumptions

- The branch `e2e-base-branch-20261017-144429` exists on the remote and is the intended target. It is also reported as the repository's default base for this run.
- The marker line is terminated by a trailing newline, as is conventional for text files.
- The label setup done at filing (`fabrik:cruise`, `base:e2e-base-branch-20261017-144429`) is correct and final.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Changes to pipeline behavior or to fixes #1047/#1050; this issue only exercises them.
- Multi-repo scenarios.
- Decomposing the work into child issues.

## Source References *(optional)*

- handarbeit/fabrik#1046 (the regression this guards against), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
