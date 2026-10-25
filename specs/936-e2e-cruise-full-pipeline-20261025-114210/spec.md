# Feature Specification: e2e cruise full pipeline (20261025-114210)

**Feature Branch**: `fabrik/issue-936`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261025-114210 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end verification of the `fabrik:cruise` pipeline contract (#898). The content of the change is deliberately trivial so that the only thing under test is the pipeline's behavior: with `fabrik:cruise` set, the item should auto-advance through every stage up to and including Validate-complete, stop there without merging the PR, and, once a human merges the PR, close the issue correctly. The run identifier `20261025-114210` makes this marker unique to this run.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261025-114210 -->`, added by this issue's pull request.

**Why this priority**: This is the whole change, and it gives the pipeline a real, reviewable diff to carry through every stage.

**Independent Test**: Inspect the PR diff and the file on the issue branch. The only change is one added line at the end of `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the new line is appended as the last line.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** exactly one file is changed and exactly one line is added.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

With `fabrik:cruise` on the issue, the engine moves the item through all stages automatically, and at Validate-complete it leaves the PR unmerged and the issue open.

**Why this priority**: This is the cruise contract the e2e run exists to verify.

**Independent Test**: Observe the item's stage progression and its final state after Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without a human approving it.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the linked PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P1)

After Validate-complete, a human merges the PR, and the issue then closes and reaches its terminal state correctly.

**Why this priority**: It verifies the second half of the cruise contract: a human-driven merge is correctly recognised.

**Independent Test**: Merge the PR by hand after Validate-complete and confirm the issue closes and the item reaches Done.

**Acceptance Scenarios**:

1. **Given** the item is at Validate-complete with the PR unmerged, **When** a human merges the PR, **Then** the issue is closed and the item reaches Done.

---

### Edge Cases

- The marker file may or may not exist beforehand. The change must handle both cases (create if missing, append if present).
- The file may or may not end with a trailing newline. The marker line must still end up on its own line at the very end.
- Plan must not split this issue into child issues, because the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261025-114210 -->` MUST be the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved unchanged and the marker line only appended.
- **FR-004**: The change MUST consist of one file and one added line, with no other files or edits.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: With `fabrik:cruise` set, the item MUST advance through all stages to Validate-complete without human approval at each step.
- **FR-007**: At Validate-complete, the PR MUST NOT be merged by the engine and the issue MUST NOT advance to Done.
- **FR-008**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `main` contains exactly one changed file with exactly one added line, equal to the marker line.
- **SC-002**: The item reaches Validate-complete with no manual stage advancement and the PR is still unmerged at that point.
- **SC-003**: After a human merge, the issue is closed and the item is in Done.

## Assumptions

- The marker line is appended verbatim, including the run identifier `20261025-114210`.
- The `e2e/markers/` directory may need to be created along with the file if it is absent.
- A trailing newline after the marker line is acceptable and conventional.
- The pipeline contract being verified is the one described in #898; this issue does not redefine it.

## Out of Scope *(optional)*

- Any change to engine behavior, stage configuration or documentation.
- Edits to any file other than `e2e/markers/cruise-full-pipeline.md`.
- Decomposing the work into child issues or multi-repo changes.
- Merging the PR automatically; the merge is done by a human as part of the verification.
