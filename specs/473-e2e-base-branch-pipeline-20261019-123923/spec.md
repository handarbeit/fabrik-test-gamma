# Feature Specification: e2e base-branch pipeline (20261019-123923)

**Feature Branch**: `fabrik/issue-473`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "

## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261019-123923 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261019-123923 (already set via
the base:<branch> label applied at filing — do not add or remove labels).
"

## Background

This issue exists to exercise the `base:<branch>` (non-default base branch) pipeline end to end. It is regression coverage for handarbeit/fabrik#1046, and it validates two fixes made in response to it: #1047 (the issue↔PR linkage for a PR that targets a non-default base) and #1050 (the data feed to the review gate for such PRs).

The change itself is deliberately trivial, a single marker line in a single file. The value is in the pipeline path the issue takes, not in the content of the change. The issue already carries a `base:e2e-base-branch-20261019-123923` label and the `fabrik:cruise` label; both are inputs to the test and must stay as they are.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite files this issue against the non-default base branch `e2e-base-branch-20261019-123923`. The pipeline processes it through every stage. The result is one new line appended to `e2e/markers/base-branch-pipeline.md`, delivered by a PR that targets the non-default base branch.

**Why this priority**: This is the whole change. It is the observable artifact that shows the pipeline ran correctly against a non-default base.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR head and confirm that its last line is exactly `<!-- base-branch-pipeline-20261019-123923 -->` and that no other content changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists, **When** the change is made, **Then** the file's content is unchanged except for one new final line, `<!-- base-branch-pipeline-20261019-123923 -->`.
2. **Given** `e2e/markers/base-branch-pipeline.md` does not exist, **When** the change is made, **Then** the file is created and contains the marker line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** it touches only `e2e/markers/base-branch-pipeline.md` and adds exactly one line (plus the spec file the engine commits for this issue, which is not part of the requested change).

---

### User Story 2 - Pipeline honours the non-default base contract (Priority: P1)

The same run also verifies the `base:<branch>` contract. The PR targets the labelled base branch rather than the repository default. The issue↔PR linkage (#1047) and the review-gate data feed (#1050) work for that PR.

**Why this priority**: This is the purpose of the test. A correct marker line on the wrong base, or with broken linkage, would be a failure.

**Independent Test**: Check that the PR's base branch is `e2e-base-branch-20261019-123923`. Check that the engine discovers the PR as linked to this issue. Check that the review gate sees the PR's review data.

**Acceptance Scenarios**:

1. **Given** the issue carries the label `base:e2e-base-branch-20261019-123923`, **When** the pipeline opens a PR, **Then** the PR's base branch is `e2e-base-branch-20261019-123923`.
2. **Given** that PR exists, **When** the engine looks for the PR linked to this issue, **Then** it finds it, even though GitHub itself creates no Development-panel link for a non-default-base PR.
3. **Given** the PR is open, **When** the review gate evaluates it, **Then** it receives the PR's review data and does not stall or misreport.

---

### Edge Cases

- The marker file does not exist on the base branch: it is created containing the marker line.
- The marker file exists but its final line has no trailing newline: the marker still ends up on its own line, so that the file's last line is exactly the marker.
- The marker line is already present (for example, a re-run): no duplicate line is added.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261019-123923 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, with the marker line as its content.
- **FR-003**: The change MUST NOT modify any other file or any other existing line of the marker file.
- **FR-004**: The work MUST be delivered through a PR whose base branch is `e2e-base-branch-20261019-123923`, as set by the issue's `base:` label.
- **FR-005**: The pipeline MUST NOT add or remove the `base:<branch>` label on the issue.
- **FR-006**: The work MUST NOT be decomposed into child issues; it stays a single issue with a single PR.
- **FR-007**: The work MUST be confined to this single repository.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the last line of `e2e/markers/base-branch-pipeline.md` on the PR head is exactly `<!-- base-branch-pipeline-20261019-123923 -->`.
- **SC-002**: The PR diff against its base contains exactly one added line in exactly one file, apart from the engine-committed spec file.
- **SC-003**: The PR's base branch is `e2e-base-branch-20261019-123923`, and the PR is associated with this issue.
- **SC-004**: The issue reaches the end of the pipeline without operator intervention, other than the label-driven automation already configured, and without a stall at the review gate.

## Assumptions

- The marker file may or may not already exist on the base branch; both cases are handled (FR-002).
- "Append at the very end" means the marker line is the last line of the file.
- The `base:` and `fabrik:cruise` labels are already applied correctly and are test inputs, so they are not touched.
- The spec file the engine commits for this issue is part of normal pipeline behaviour and is not counted as part of the requested change.

## Out of Scope *(optional)*

- Any change other than the single marker line in the single marker file.
- Decomposing the work into sub-issues.
- Changes to the engine, the stage configuration, or the labels.
- Multi-repo scenarios.

## Source References *(optional)*

- handarbeit/fabrik#1046 — the regression this test covers
- handarbeit/fabrik#1047 — issue↔PR linkage fix for non-default base
- handarbeit/fabrik#1050 — review-gate data feed fix
