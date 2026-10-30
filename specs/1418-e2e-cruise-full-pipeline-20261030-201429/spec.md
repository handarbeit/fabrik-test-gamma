# Feature Specification: e2e cruise full pipeline marker (20261030-201429)

**Feature Branch**: `fabrik/issue-1418`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261030-201429 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` pipeline contract (#898). Under cruise, an item auto-advances through every stage up to Validate-complete, but the engine must not auto-merge the PR or move the issue to Done. After a human merges the PR, the issue must close correctly. The file change itself is deliberately trivial, so the pipeline's behaviour, not the content of the change, is what is being verified. Each run appends a uniquely timestamped marker line so that every run produces a real, reviewable diff.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker (Priority: P1)

A single timestamped HTML comment is appended to the end of the cruise full-pipeline marker file, so the pipeline has a real one-line change to carry through every stage.

**Why this priority**: This is the entire change. Without it there is no diff for the pipeline to carry.

**Independent Test**: Inspect the marker file on the issue branch. The last line is the marker for this run and every earlier line is unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is applied, **Then** the line `<!-- cruise-pipeline-20261030-201429 -->` is the final line of the file and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the line `<!-- cruise-pipeline-20261030-201429 -->`.

---

### User Story 2 - Cruise carries the item to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` on the issue, the item moves through every stage automatically and stops once Validate is complete, with the PR left unmerged.

**Why this priority**: This is the cruise contract the run exists to verify.

**Independent Test**: Observe the issue and PR after the pipeline settles. All stages are complete, the PR is open and unmerged, and the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the end of the pipeline, **Then** it neither merges the PR nor moves the issue to Done.

---

### User Story 3 - Issue closes after a human merge (Priority: P2)

Once a human merges the PR, the issue closes and reaches its terminal state correctly.

**Why this priority**: This verifies the hand-off from cruise's stopping point to normal completion.

**Independent Test**: Merge the PR manually and confirm the issue closes.

**Acceptance Scenarios**:

1. **Given** the PR is open and Validate is complete, **When** a human merges the PR, **Then** the issue is closed and the item reaches Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new marker must still occupy its own line.
- The marker file does not exist: it is created with only the new marker line.
- The marker line for this run already exists (a re-run): it is not appended a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261030-201429 -->`, as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created containing that line.
- **FR-003**: Existing content of the marker file is left unchanged.
- **FR-004**: The change touches only that one file.
- **FR-005**: The work is implemented as a single unit and is not decomposed into child issues.
- **FR-006**: The PR is not auto-merged by the engine under cruise, and the issue is not moved to Done by the engine at Validate completion.
- **FR-007**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line in exactly one file.
- **SC-002**: The item reaches Validate-complete with no human intervention between stages.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is open and not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file may already exist with earlier marker lines from prior runs; the new line is appended after them.
- A single repository is involved, and no cross-repo behaviour is exercised.
- The marker line is an HTML comment, so it has no visible effect on rendered content.

## Out of Scope *(optional)*

- Decomposing the work into child issues.
- Changes to any file other than the marker file.
- Changes to the cruise logic or the engine.
- Verifying yolo or other autonomy modes.
