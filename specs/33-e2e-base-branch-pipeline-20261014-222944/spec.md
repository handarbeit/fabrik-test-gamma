# Feature Specification: e2e base-branch pipeline marker (20261014-222944)

**Feature Branch**: `fabrik/issue-33`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261014-222944 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261014-222944 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end exercise of the `base:<branch>` pipeline contract: an issue whose work targets a non-default base branch must flow through the pipeline stages with its issue↔PR linkage and review-gate data intact. It is regression coverage for handarbeit/fabrik#1046, validating the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial so that the pipeline behavior, not the change, is what is being verified.

The target base branch is `e2e-base-branch-20261014-222944`, already set through the `base:<branch>` label at filing. The issue also carries `fabrik:cruise`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator files an issue labelled `base:e2e-base-branch-20261014-222944`. The pipeline appends one marker line to `e2e/markers/base-branch-pipeline.md`, opens a PR that targets that base branch, and carries the issue through to completion.

**Why this priority**: This is the entire change and the artifact the e2e run asserts on.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR head: its last line is exactly `<!-- base-branch-pipeline-20261014-222944 -->` and the PR diff touches only that file.

**Acceptance Scenarios**:

1. **Given** the file does not exist, **When** the change is made, **Then** the file is created and contains the marker line as its last line.
2. **Given** the file already exists, **When** the change is made, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the issue has the `base:e2e-base-branch-20261014-222944` label, **When** the PR is opened, **Then** the PR targets that base branch rather than the repository's other branches.

---

### User Story 2 - Issue↔PR linkage and review-gate data work for a non-default base (Priority: P2)

The pipeline discovers the PR for the issue and feeds review-gate data correctly even though GitHub creates no Development-panel link for a PR on a non-default base.

**Why this priority**: This is the regression being guarded (#1046, #1047, #1050), but it is observed from pipeline behavior rather than produced by the change.

**Independent Test**: The issue progresses through its stages with the PR found and the review gate evaluated, without a broken-linkage pause.

**Acceptance Scenarios**:

1. **Given** a PR on the non-default base that references the issue, **When** the pipeline reaches PR-dependent stages, **Then** the PR is recognized as the issue's linked PR.
2. **Given** the PR is under review, **When** the review gate is evaluated, **Then** it uses the PR's review data and does not pause for missing linkage.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the line must still end up on its own final line, with nothing after it.
- The marker line must not be duplicated if the stage is re-run.
- Labels must not be added or removed by the work itself.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` MUST end with the single line `<!-- base-branch-pipeline-20261014-222944 -->`, created if the file does not exist.
- **FR-002**: Any pre-existing content of that file MUST be left unchanged; the marker is appended at the very end.
- **FR-003**: The change MUST consist of exactly one file and one added line (excluding the engine-written spec file).
- **FR-004**: The PR MUST target the base branch `e2e-base-branch-20261014-222944` and include `Closes #33` in its body.
- **FR-005**: The work MUST NOT add or remove any issue labels.
- **FR-006**: The work MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The final line of `e2e/markers/base-branch-pipeline.md` on the merged result equals the marker string exactly, appearing once.
- **SC-002**: The PR diff against `e2e-base-branch-20261014-222944` contains one changed file with one added line (plus the engine-committed spec file).
- **SC-003**: The issue completes the pipeline with its PR linked and no linkage- or review-data-related pause.

## Assumptions

- The marker file does not currently exist on the base branch, so it will be created; if it does exist, appending applies.
- The `fabrik:cruise` label lets the pipeline advance through stages automatically without auto-merging.
- The base branch `e2e-base-branch-20261014-222944` exists on the remote.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Modifying labels, stage configuration, or engine behavior.
- Splitting the work into child issues.
- Changes to other repositories.
