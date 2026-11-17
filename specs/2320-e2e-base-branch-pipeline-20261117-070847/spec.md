# Feature Specification: e2e base-branch pipeline marker (20261117-070847)

**Feature Branch**: `fabrik/issue-2320`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261117-070847 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261117-070847 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end regression check for the `base:<branch>` pipeline contract, i.e. an issue whose work targets a non-default base branch. It covers handarbeit/fabrik#1046 and validates two fixes: #1047 (linkage between the issue and its PR) and #1050 (the data feed to the review gate). The code change is deliberately trivial. The point is to drive an item through the full pipeline against the base branch `e2e-base-branch-20261117-070847`, which was set by the `base:e2e-base-branch-20261117-070847` label at filing, and to observe that linkage and review gating behave correctly.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended on the non-default base (Priority: P1)

An operator running the e2e suite wants a minimal, deterministic change that travels the whole pipeline on a `base:<branch>` issue, so the base-branch linkage and review-gate behaviour can be checked end to end.

**Why this priority**: It is the only story. The change exists to exercise the pipeline.

**Independent Test**: After the change, the file `e2e/markers/base-branch-pipeline.md` on the PR branch ends with the line `<!-- base-branch-pipeline-20261117-070847 -->`, and the PR targets `e2e-base-branch-20261117-070847`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is made, **Then** exactly one new line `<!-- base-branch-pipeline-20261117-070847 -->` is added at the very end and no other content changes.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing that single line.
3. **Given** the issue carries the `base:e2e-base-branch-20261117-070847` label, **When** the PR is opened, **Then** it targets that branch and is linked to the issue.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line at the end.
- The marker line is already present from an earlier run: it is still appended, because this run's identifier is unique to this issue.
- The issue's labels (`fabrik:cruise`, `base:e2e-base-branch-20261117-070847`) must not be added, removed or altered by the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` ends with the single line `<!-- base-branch-pipeline-20261117-070847 -->` after the change.
- **FR-002**: If the file does not exist beforehand, it is created containing only that line.
- **FR-003**: The change touches exactly one file and adds exactly one line.
- **FR-004**: The work targets the non-default base branch `e2e-base-branch-20261117-070847`, and the PR is based on and targets that branch.
- **FR-005**: No labels are added to or removed from the issue by this work.
- **FR-006**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the resulting PR against `e2e-base-branch-20261117-070847` is exactly one file with one added line, being the marker comment.
- **SC-002**: The PR targets `e2e-base-branch-20261117-070847` and is linked to issue #2320.
- **SC-003**: The pipeline reaches completion for the item with no linkage or review-gate failure attributable to the non-default base.

## Assumptions

- The branch `e2e-base-branch-20261117-070847` already exists on the remote.
- The marker line is added exactly as written, with no surrounding blank line.
- The issue's existing labels stay as filed, including `fabrik:cruise`, which lets the item auto-advance.

## Out of Scope

- Any change beyond the single marker line.
- Any change to engine behaviour, labels or the base branch itself.
- Decomposition into child issues.
