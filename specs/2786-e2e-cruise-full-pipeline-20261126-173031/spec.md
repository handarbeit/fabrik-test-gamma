# Feature Specification: e2e cruise full pipeline (20261126-173031)

**Feature Branch**: `fabrik/issue-2786`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261126-173031 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` label end to end (#898). Cruise should auto-advance an item through every pipeline stage up to Validate-complete, stop without merging the linked PR, and leave the merge to a human. The issue should then close correctly once the human merges. The code change is deliberately trivial so that the pipeline behaviour, not the change, is what is being verified. The change must stay a single, non-decomposable unit of work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

A single marker line is added to the end of the cruise-pipeline marker file, giving the pipeline a real, reviewable change to carry through every stage.

**Why this priority**: Without a change there is no PR to carry through the stages, so nothing else can be verified.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the PR branch and confirm its last line is the marker comment and that no other content was altered.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created and its final line is `<!-- cruise-pipeline-20261126-173031 -->`.
2. **Given** the marker file already exists, **When** the change is made, **Then** the existing content is unchanged and the marker line is appended as the new final line.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item moves through all stages automatically and halts at Validate-complete, with its PR still open and unmerged.

**Why this priority**: This is the contract being verified (#898).

**Independent Test**: After the pipeline settles, observe that the issue has Validate-complete, the linked PR is open, and the PR has not been merged by Fabrik.

**Acceptance Scenarios**:

1. **Given** an issue labelled `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate completes, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

Once a human merges the PR, the issue closes correctly.

**Why this priority**: It completes the contract, but depends on the stories above.

**Independent Test**: Merge the PR by hand and confirm the issue ends closed and in Done.

**Acceptance Scenarios**:

1. **Given** Validate-complete and an open PR, **When** a human merges the PR, **Then** the issue is closed and reaches Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the marker must still land on its own final line.
- Plan must not decompose the work into child issues, since the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261126-173031 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST consist of exactly one file and one added line, with no other content altered.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: Under `fabrik:cruise`, the item MUST auto-advance through all stages to Validate-complete without human approval.
- **FR-006**: Fabrik MUST NOT merge the PR, and MUST NOT advance the issue to Done, at Validate-complete.
- **FR-007**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff touches exactly one file and adds exactly one line, the marker comment.
- **SC-002**: The item reaches Validate-complete with no human approval between stages.
- **SC-003**: At Validate-complete the PR is open and unmerged.
- **SC-004**: After the human merge, the issue is closed.

## Assumptions

- The single repository named by the issue is the only one involved.
- The marker string is the exact timestamped value given in the request.
- A human performs the PR merge, as the issue states.

## Out of Scope

- Any change beyond the single marker line.
- Multi-repo behaviour.
- Decomposing the work into child issues.
- Auto-merging the PR.
