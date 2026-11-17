# Feature Specification: e2e cruise full pipeline marker (20261117-103433)

**Feature Branch**: `fabrik/issue-2463`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261117-103433 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an item through every pipeline stage automatically, stop at Validate-complete without merging the linked PR, and leave the issue to close only after a human merges that PR. The code change is deliberately trivial, so that the pipeline behaviour, not the change, is what gets verified. The marker file `e2e/markers/cruise-full-pipeline.md` already exists in the repository and holds one marker line per earlier run; this run adds one more.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended by the pipeline (Priority: P1)

The pipeline produces a PR whose only change is one new HTML comment line at the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: The marker is the sole deliverable and the evidence that the run's changes reached a PR.

**Independent Test**: Inspect the PR diff and confirm it adds exactly one line, `<!-- cruise-pipeline-20261117-103433 -->`, at the end of the file and changes nothing else apart from the engine-managed spec file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier marker lines, **When** the Implement stage completes, **Then** the file's final line is `<!-- cruise-pipeline-20261117-103433 -->` and all earlier lines are unchanged.
2. **Given** the marker file does not exist, **When** the Implement stage completes, **Then** the file is created containing that one line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item advances automatically through every stage up to Validate-complete, and the PR is left unmerged.

**Why this priority**: This is the contract being verified.

**Independent Test**: Watch the issue move through each stage with no human intervention; at Validate-complete the linked PR is open and not merged and the issue is not advanced to Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances to the next stage without waiting for a human.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merge (Priority: P2)

After a human merges the PR, the issue closes and reaches Done.

**Why this priority**: It completes the verified lifecycle, but happens after the cruise-controlled part.

**Independent Test**: Merge the PR by hand and confirm the issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still sit on its own line, not be joined to the previous one.
- A marker line for this run id already exists in the file: it is not added a second time.
- Plan must not split this into child issues; the work is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the exact line `<!-- cruise-pipeline-20261117-103433 -->` as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The change MUST create `e2e/markers/cruise-full-pipeline.md` containing that line if the file does not exist.
- **FR-003**: Existing content of the marker file MUST NOT be modified.
- **FR-004**: The change MUST touch no files other than the marker file, apart from the spec file the engine commits itself.
- **FR-005**: The work MUST be delivered as a single issue with no child issues spawned.
- **FR-006**: The pipeline MUST NOT merge the PR as part of cruise; merging is left to a human.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in `e2e/markers/cruise-full-pipeline.md` and no other source changes.
- **SC-002**: The issue reaches Validate-complete with no human action between stages.
- **SC-003**: At Validate-complete the linked PR is open and unmerged.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file already exists in the repository, and each earlier run appended one similar line.
- The marker is an HTML comment, so it does not affect how the file renders.
- The timestamp in the marker is a unique run identifier and is used exactly as given.

## Out of Scope *(optional)*

- Any change to Fabrik engine behaviour or to the cruise logic itself.
- Any change to other marker files.
- Multi-repo behaviour.
- Merging the PR (done by a human after Validate).
