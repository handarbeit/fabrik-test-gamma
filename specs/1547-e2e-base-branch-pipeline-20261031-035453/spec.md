# Feature Specification: e2e base-branch pipeline marker (20261031-035453)

**Feature Branch**: `fabrik/issue-1547`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261031-035453 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261031-035453 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of Fabrik's `base:<branch>` contract, which lets an issue's work fork from, and target its pull request at, a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046, and it validates the fixes in #1047 (the issue-to-PR linkage for a PR on a non-default base) and #1050 (the data the review gate reads for such a PR). The code change is deliberately trivial, so that any failure comes from the pipeline and not from the change. The issue carries the labels `fabrik:cruise` and `base:e2e-base-branch-20261031-035453`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

The pipeline takes this issue through every stage and produces a single-line change to the marker file `e2e/markers/base-branch-pipeline.md`. That change is carried on a pull request that targets `e2e-base-branch-20261031-035453` and not the repository's usual base.

**Why this priority**: This is the whole purpose of the issue. It exercises the non-default-base path from start to finish with the smallest possible change.

**Independent Test**: After the pipeline finishes, compare the marker file on the target base branch with its state before the change. The only difference is one appended line, `<!-- base-branch-pipeline-20261031-035453 -->`, at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the target base branch, **When** the change is made, **Then** the file is created and its content is exactly the one line `<!-- base-branch-pipeline-20261031-035453 -->`.
2. **Given** the marker file already exists, **When** the change is made, **Then** the existing content is unchanged and the one line is appended at the very end of the file.
3. **Given** the issue carries the `base:e2e-base-branch-20261031-035453` label, **When** the pull request is opened, **Then** its base is `e2e-base-branch-20261031-035453`.

---

### User Story 2 - Pipeline contract for a non-default base holds (Priority: P1)

The behaviours covered by #1047 and #1050 are observable while the issue moves through the pipeline.

**Why this priority**: This is the regression coverage the issue exists to provide. A correct file change on a pull request the pipeline cannot link or gate would still be a failure.

**Independent Test**: Inspect the issue and the pull request after the run: the pull request is discovered as the issue's linked PR, and the review gate sees the pull request's review data.

**Acceptance Scenarios**:

1. **Given** the pull request targets a non-default base, **When** the pipeline looks for the issue's pull request, **Then** it finds it and associates it with this issue (#1047).
2. **Given** the pull request has been opened, **When** the review gate evaluates it, **Then** it reads the pull request's review data and reaches a decision (#1050).
3. **Given** the item runs under `fabrik:cruise`, **When** each stage completes, **Then** the item advances without waiting for human approval, and no `Closes`-style keyword links to any issue other than this one.

---

### Edge Cases

- The marker file is absent on the base branch: it is created containing only the marker line.
- The marker file exists but its last line has no trailing newline: the appended marker must still be on its own line.
- The marker line is already present at the end of the file (for example on a re-run): the file must not end up with the line twice.
- The target base branch is the repository's default for this run: the change and its pull request must still target `e2e-base-branch-20261031-035453`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261031-035453 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it before appending the line.
- **FR-003**: The change MUST touch exactly one file and add exactly one line, with no other edits.
- **FR-004**: The work MUST be based on, and its pull request MUST target, the branch `e2e-base-branch-20261031-035453`.
- **FR-005**: The pull request MUST be linked to this issue (`Closes #1547`) and MUST NOT use a closing keyword for any other issue or report.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: The work MUST NOT add or remove any labels on the issue beyond those the pipeline manages itself.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the pull request against `e2e-base-branch-20261031-035453` is one file with one added line, the marker line.
- **SC-002**: The pull request's base branch is `e2e-base-branch-20261031-035453`.
- **SC-003**: The issue reaches the end of the pipeline without a human having to relink the pull request or intervene in the review gate.
- **SC-004**: No child issues are created from this issue.

## Assumptions

- The branch `e2e-base-branch-20261031-035453` exists on the remote and is the intended base; the `base:` label is already applied and stays as is.
- The `e2e/markers/` directory already exists on the base branch (other marker files live there); if it does not, creating it is part of creating the file.
- The marker line is appended as its own line, ending with a newline, which is the file's convention.
- Re-running the change must be idempotent, so the marker line appears once.

## Out of Scope *(optional)*

- Any change beyond the single marker line, including to other marker files, code, or documentation.
- Changes to the pipeline or engine behaviour for non-default bases; this issue only verifies it.
- Splitting the work into several issues or child issues.
- Adding or removing labels.
