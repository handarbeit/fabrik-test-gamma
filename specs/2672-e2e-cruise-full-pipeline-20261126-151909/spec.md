# Feature Specification: E2E cruise full-pipeline marker (20261126-151909)

**Feature Branch**: `fabrik/issue-2672`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261126-151909 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, an item should advance automatically through every pipeline stage up to Validate-complete, without Fabrik merging the linked PR. The issue should then close correctly once a human merges that PR. The code change is deliberately trivial: one marker line appended to a shared marker file. The change is only a vehicle for exercising the pipeline, so it must stay small and must not be decomposed into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the pipeline marker line (Priority: P1)

A maintainer running the cruise e2e check needs a unique, recognisable marker line recorded in the repository by the pipeline, so the run produces a real, reviewable PR.

**Why this priority**: The marker is the only deliverable. Without it the pipeline has no change to carry through the stages.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch and confirm the final line is the specified marker and nothing else changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is applied, **Then** its last line is `<!-- cruise-pipeline-20261126-151909 -->` and all earlier lines are unchanged.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise advances without merging (Priority: P1)

An operator wants to confirm that an item labelled `fabrik:cruise` moves through all stages to Validate-complete on its own, while the PR stays unmerged for a human to merge.

**Why this priority**: This is the contract under test (#898), and the reason the issue exists.

**Independent Test**: Observe the item reaching Validate-complete with the linked PR still open and not merged, then merge the PR by hand and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage with no human approval required.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR remains open and unmerged.
3. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file does not exist yet: it is created with the marker line as its content.
- The existing file does not end in a newline: the marker still lands on its own line at the end.
- Earlier cruise marker lines from previous runs must be preserved untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261126-151909 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST NOT modify or remove any existing line in the file, and MUST NOT touch any other file apart from the engine-maintained spec file.
- **FR-004**: The work MUST be delivered as a single, undecomposed unit; the Plan stage MUST NOT spawn child issues.
- **FR-005**: Under `fabrik:cruise`, the item MUST auto-advance through all stages to Validate-complete, and the linked PR MUST NOT be merged by Fabrik.
- **FR-006**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff for the implementation contains exactly one added line in one file, and that line is the specified marker.
- **SC-002**: The item reaches Validate-complete without any human stage approval.
- **SC-003**: The linked PR is still unmerged when Validate completes.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The file `e2e/markers/cruise-full-pipeline.md` already exists on `main` with earlier cruise marker lines, so the "create if missing" path is a fallback only.
- Only the spec file the engine commits is allowed alongside the one-line change.
- The merging human is the e2e harness or operator, not Fabrik.

## Out of Scope

- Any change beyond the single marker line.
- Decomposing the work into child issues.
- Merging the PR automatically (that is the `yolo` path, not `cruise`).
- Changes to the cruise logic itself.

## Source References

- #898 (the `fabrik:cruise` pipeline contract)
