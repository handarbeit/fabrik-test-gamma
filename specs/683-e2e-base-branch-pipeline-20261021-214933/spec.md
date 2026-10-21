# Feature Specification: e2e base-branch pipeline (20261021-214933)

**Feature Branch**: `fabrik/issue-683`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261021-214933 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261021-214933 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end test fixture for Fabrik's `base:<branch>` support, not a product feature. It exercises the full pipeline for an issue whose pull request targets a non-default base branch. It serves as regression coverage for handarbeit/fabrik#1046 and validates two fixes:

- #1047: the issue-to-PR linkage for PRs on a non-default base.
- #1050: the data feed into the review gate.

The code change is deliberately trivial so that any failure points at the pipeline contract, not at the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the base-branch e2e check files this issue with the `base:e2e-base-branch-20261021-214933` label. The pipeline produces a PR that adds the marker line and targets that branch.

**Why this priority**: This is the only deliverable. The change itself is what the pipeline is verified against.

**Independent Test**: Inspect the PR diff and the marker file on the PR branch.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file's last line is `<!-- base-branch-pipeline-20261021-214933 -->` and no other line is altered.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261021-214933` label, **When** the PR is created, **Then** it targets that branch and not the repository's other branches.

---

### Edge Cases

- The file exists but its last line has no trailing newline. The marker must still go on its own line at the end.
- The file already contains the exact marker line (for example, after a retry). The marker must not be duplicated.
- The target base branch is missing on the remote. This is existing engine behaviour, covered by the `base:` label rules and not changed here.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261021-214933 -->`, appended at the very end of the file.
- **FR-002**: If the file does not exist, it is created containing that line.
- **FR-003**: The change touches exactly one file and adds exactly one line, and existing content of the file is preserved unchanged.
- **FR-004**: The resulting PR targets the base branch `e2e-base-branch-20261021-214933` and carries `Closes #683`.
- **FR-005**: No labels are added to or removed from the issue by this work. The `base:<branch>` label is already set.
- **FR-006**: The work is a single unit and is not decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, an append-only file whose trailing HTML comment lines identify individual e2e runs.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line (the marker) in exactly one file.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261021-214933`.
- **SC-003**: The pipeline completes with the issue and PR correctly linked and the review gate receiving its data, so the base-branch regression (#1046) does not recur.

## Assumptions

- The marker line is HTML-comment syntax so it is invisible when the markdown is rendered.
- If the marker file already exists on the base branch, appending preserves all prior marker lines from earlier runs.
- The `fabrik:cruise` label governs advancement and is left as is.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, docs or stage configuration.
- Changing the issue's labels.
- Multi-repo scenarios and any base branch other than `e2e-base-branch-20261021-214933`.
- Decomposing the work into sub-issues.
