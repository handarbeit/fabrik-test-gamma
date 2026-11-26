# Feature Specification: e2e base-branch pipeline marker (20261126-183919)

**Feature Branch**: `fabrik/issue-2810`
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
<!-- base-branch-pipeline-20261126-183919 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261126-183919 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is an end-to-end test fixture. Fabrik supports a `base:<branch>` label that makes an issue's worktree fork from, and its pull request target, a branch other than the repository default. Earlier defects in that path (tracked in handarbeit/fabrik#1046) broke the issue-to-PR linkage (fixed in #1047) and the data feed to the review gate (fixed in #1050). This run drives a deliberately tiny change through the whole pipeline on a non-default base branch, so a regression in either fix shows up as a pipeline failure rather than going unnoticed.

The change itself is intentionally trivial and carries no product value. What is under test is the pipeline's handling of the base-branch label, not the content of the file.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite files an issue labelled with a non-default `base:<branch>`. The pipeline adds one marker line to a tracked file, opens a pull request against that base branch, and carries the item through to completion.

**Why this priority**: It is the only story. Delivering the marker line through a PR targeting the non-default base is the whole purpose of the issue.

**Independent Test**: After the pipeline finishes, inspect the base branch `e2e-base-branch-20261126-183919`. `e2e/markers/base-branch-pipeline.md` should end with the marker line, and the diff of the merged PR should consist of exactly that one added line.

**Acceptance Scenarios**:

1. **Given** the issue carries the label `base:e2e-base-branch-20261126-183919`, **When** the Implement stage produces its pull request, **Then** the PR targets `e2e-base-branch-20261126-183919` and not the repository default.
2. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is applied, **Then** the line `<!-- base-branch-pipeline-20261126-183919 -->` is appended as the last line and the existing content is unchanged.
3. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is applied, **Then** the file is created and contains the marker line.
4. **Given** the PR carries `Closes #2810`, **When** the pipeline runs, **Then** the issue and the PR are linked, and the review gate receives the PR's review data (the #1047 and #1050 contracts).

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still be on its own line, not joined to the previous content.
- The marker line is already present at the end of the file (re-run): the file must not gain a duplicate line.
- The base branch differs from the repository default branch: the PR must never target the default branch.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261126-183919 -->`, at the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created and contain only that marker line.
- **FR-003**: The change MUST touch no file other than `e2e/markers/base-branch-pipeline.md`.
- **FR-004**: The pull request MUST target the branch `e2e-base-branch-20261126-183919`.
- **FR-005**: The issue's labels MUST NOT be added to or removed from as part of this work; the existing `base:e2e-base-branch-20261126-183919` label stays as filed.
- **FR-006**: The work MUST remain a single change in a single repository and MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch is exactly one added line in exactly one file.
- **SC-002**: The pull request's base branch equals `e2e-base-branch-20261126-183919`.
- **SC-003**: The issue is linked to its pull request and reaches its terminal state without manual intervention, showing that issue-to-PR linkage and review-gate data feed work on a non-default base.
- **SC-004**: No child issues are spawned.

## Assumptions

- The branch `e2e-base-branch-20261126-183919` already exists on the remote, since the harness created it before filing.
- The marker file lives under `e2e/markers/` and its directory may need to be created together with the file.
- The label is already correct as filed; the pipeline only needs to honour it.
- Because the issue carries `fabrik:cruise`, the stages advance without waiting for human approval. The spec has no open questions, so this is safe.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configs, or labels.
- Any change to other files, repositories, or branches.
- Any decomposition into sub-issues.

## Source References *(optional)*

- handarbeit/fabrik#1046 (base-branch pipeline regression), #1047 (issue/PR linkage fix), #1050 (review-gate data feed fix).
