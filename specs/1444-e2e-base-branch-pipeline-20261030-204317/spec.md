# Feature Specification: e2e base-branch pipeline marker (20261030-204317)

**Feature Branch**: `fabrik/issue-1444`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: 

> ## Goal
>
> End-to-end verification of the base:<branch> (non-default base branch)
> pipeline contract — regression coverage for handarbeit/fabrik#1046,
> validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
> fixes.
>
> ## Trivial change
>
> Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> ```
> <!-- base-branch-pipeline-20261030-204317 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261030-204317 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).
>
> ## Labels
>
> fabrik:cruise, base:e2e-base-branch-20261030-204317

## Background

This issue is a regression probe for the `base:<branch>` pipeline contract: an issue labelled with a non-default base branch must flow through the whole pipeline with its work branched from, and its PR targeted at, that base branch. It gives regression coverage for handarbeit/fabrik#1046 and validates two fixes: #1047 (the issue↔PR linkage for a PR on a non-default base) and #1050 (the data feed to the review gate).

The change itself is deliberately trivial, a single marker line in one file. That keeps the result about the pipeline's handling of the base branch and not about the content of the change. The marker is unique to this run (`20261030-204317`), so a reviewer can tell the run's change from earlier runs' markers.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the base-branch e2e check files this issue with the `base:e2e-base-branch-20261030-204317` label. The pipeline appends one marker line to `e2e/markers/base-branch-pipeline.md` and delivers it as a PR whose base is that non-default branch.

**Why this priority**: This is the whole change. Without it there is nothing for the base-branch pipeline checks to observe.

**Independent Test**: Inspect the resulting PR diff. It should show exactly one added line, the marker comment, at the end of `e2e/markers/base-branch-pipeline.md`, and the PR's base branch should be `e2e-base-branch-20261030-204317`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is made, **Then** the last line of the file is `<!-- base-branch-pipeline-20261030-204317 -->` and all earlier content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is made, **Then** the file is created and contains the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261030-204317` label, **When** the PR is created, **Then** the PR targets that branch, not the repository default.

---

### User Story 2 - Pipeline contract is observable end to end (Priority: P2)

The operator verifies that the issue↔PR linkage (#1047) and the review-gate data feed (#1050) work for a PR on a non-default base.

**Why this priority**: This is the reason the issue exists, but it is observed from the pipeline's behavior and needs no change beyond User Story 1.

**Independent Test**: After the PR exists, check that the issue and PR are linked as the pipeline expects for a non-default base, and that the review gate sees the PR's review data.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** the pipeline reaches the review gate, **Then** the gate receives the PR's review data.
2. **Given** a PR targeting the non-default base, **When** the pipeline discovers the PR for the issue, **Then** the issue and PR are associated.

---

### Edge Cases

- The marker file is missing on the base branch: create it, with the marker as its only content.
- The file exists but its last line has no trailing newline: the marker must still go on its own new line at the end.
- The marker line is already present, for example on a retried stage: do not add a second copy.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` ends with the line `<!-- base-branch-pipeline-20261030-204317 -->`.
- **FR-002**: If the file does not exist, it is created. If it exists, its prior content is preserved unchanged and the marker is appended at the very end.
- **FR-003**: The change touches exactly one file and adds exactly one line (plus any newline needed to start that line on its own line).
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The work branch is based on, and the PR targets, `e2e-base-branch-20261030-204317`.
- **FR-006**: The issue's labels are not changed: no labels are added or removed by this work.
- **FR-007**: The work is not decomposed into sub-issues or child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261030-204317` contains exactly one changed file and one added content line.
- **SC-002**: A search of `e2e/markers/base-branch-pipeline.md` for `base-branch-pipeline-20261030-204317` finds exactly one match, on the last line.
- **SC-003**: The PR's base branch is `e2e-base-branch-20261030-204317`, and the issue reaches its end state through the pipeline without manual intervention.

## Assumptions

- The branch `e2e-base-branch-20261030-204317` exists on the remote, since the issue was filed with the matching label.
- The marker file may or may not already exist. Either case is covered by FR-002.
- The file is Markdown. An HTML comment is invisible when rendered and so does not alter the rendered content.
- `fabrik:cruise` stays on the issue, so the pipeline advances automatically without waiting for spec approval, and the PR is not auto-merged by the pipeline.

## Out of Scope *(optional)*

- Any change to Fabrik's engine, stage configuration or tests.
- Any change to other files, other repositories or other branches.
- Adding, removing or editing labels.
- Splitting the work into multiple issues.
- Cleaning up earlier runs' markers or the e2e base branch.
