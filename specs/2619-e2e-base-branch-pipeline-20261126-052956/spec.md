# Feature Specification: e2e base-branch pipeline (20261126-052956)

**Feature Branch**: `fabrik/issue-2619`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261126-052956 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261126-052956 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` pipeline contract, that is, how an issue is processed when its work targets a non-default base branch. It is regression coverage for handarbeit/fabrik#1046. It validates two earlier fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The code change is deliberately trivial, since the point is to push a minimal change through the whole pipeline against the non-default base `e2e-base-branch-20261126-052956`. The issue carries the labels `fabrik:cruise` and `base:e2e-base-branch-20261126-052956`, applied at filing.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

A single marker line is appended to `e2e/markers/base-branch-pipeline.md`. The change travels through the pipeline as a PR targeting `e2e-base-branch-20261126-052956` and not the repository's other branches.

**Why this priority**: It is the only deliverable, and it exercises the base-branch pipeline contract end to end.

**Independent Test**: After the pipeline completes, the marker file on the target base branch ends with the line `<!-- base-branch-pipeline-20261126-052956 -->`, and the diff against the base touches only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- base-branch-pipeline-20261126-052956 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.
3. **Given** the issue has label `base:e2e-base-branch-20261126-052956`, **When** the PR is created, **Then** the PR targets that branch and its body contains `Closes #2619`.

---

### Edge Cases

- The file already ends without a trailing newline: the marker must still occupy its own final line.
- The marker line already exists in the file: it must not be duplicated.
- The `base:<branch>` label is not touched at any point: do not add, remove or change labels.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261126-052956 -->` is appended as the last line of `e2e/markers/base-branch-pipeline.md`, creating the file first if it does not exist.
- **FR-002**: The change consists of exactly one file and one added line (plus the file's creation if it was absent), with no other modifications.
- **FR-003**: The resulting PR targets the base branch `e2e-base-branch-20261126-052956`.
- **FR-004**: The PR body includes `Closes #2619` so the issue and PR are linked.
- **FR-005**: The `base:` label and every other label on the issue are left untouched by the change.
- **FR-006**: The work is not decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a plain text file holding one HTML comment line per e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker file on the base branch `e2e-base-branch-20261126-052956` ends with exactly one occurrence of `<!-- base-branch-pipeline-20261126-052956 -->`.
- **SC-002**: The PR diff against that base branch contains one file and one added line, or one new file containing only the marker line.
- **SC-003**: The issue is linked to its PR and moves through the pipeline stages without the base-branch contract (#1047 linkage, #1050 review-gate feed) causing a pause or stall.

## Assumptions

- The repository's default branch is `main`. The target branch `e2e-base-branch-20261126-052956` is a non-default base, selected by the `base:` label, which is what exercises the `base:<branch>` contract.
- The marker is a plain HTML comment, so it needs no format or build handling.
- No documentation changes are needed.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Splitting the work into sub-issues.
- Adding or removing labels.
- Multi-repo behaviour.
- Changes to engine code or to the e2e harness.
