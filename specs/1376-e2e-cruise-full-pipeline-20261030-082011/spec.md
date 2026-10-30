# Feature Specification: e2e cruise full pipeline (20261030-082011)

**Feature Branch**: `fabrik/issue-1376`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261030-082011 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end verification of the `fabrik:cruise` pipeline contract (#898). The item carries the `fabrik:cruise` label, so it should auto-advance through every stage up to Validate-complete without the PR being auto-merged. The issue should then close correctly once a human merges the PR. The code change is deliberately minimal, so that any behaviour observed comes from the pipeline and not from the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the repository (Priority: P1)

A single HTML comment marker line, unique to this test run, is appended to the end of the cruise pipeline marker file. This gives the pipeline a real, reviewable one-file change to carry through every stage.

**Why this priority**: This is the whole change. Without it the pipeline has nothing to carry to a PR.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its last line must be exactly `<!-- cruise-pipeline-20261030-082011 -->`, and the diff must touch only that file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the existing content is unchanged and the marker line is the final line of the file.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its final line.

---

### User Story 2 - Cruise carries the item through the pipeline without merging (Priority: P1)

An operator labels the issue `fabrik:cruise` and expects it to move through all stages to Validate-complete automatically. The PR must not be merged by the engine, and the issue must not move to Done until a human acts.

**Why this priority**: This is the contract under test (#898). The marker change exists only to exercise it.

**Independent Test**: Observe the item after Validate completes. The PR is open and unmerged, and the issue is still open.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances it to the next stage without waiting for human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** it does not merge the PR and does not advance the issue to Done.
3. **Given** a human merges the PR after Validate-complete, **When** the engine next processes the item, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still land on its own line at the end of the file.
- The marker file already contains an identical marker line: it is still appended once, so the file ends with exactly one new occurrence.
- Plan must not decompose the work into child issues, because the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261030-082011 -->`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: If the file exists, its prior content MUST be preserved unchanged and the marker line appended after it.
- **FR-004**: The change MUST consist of exactly one file and one added line, with no other files modified.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: With `fabrik:cruise` present, the pipeline MUST auto-advance through all stages to Validate-complete without human approval at intermediate stages.
- **FR-007**: At Validate-complete, the engine MUST NOT merge the PR and MUST NOT advance the issue to Done.
- **FR-008**: After a human merges the PR, the issue MUST close correctly.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a text file that accumulates one HTML comment line per cruise e2e run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, equal to `<!-- cruise-pipeline-20261030-082011 -->`.
- **SC-002**: The item reaches Validate-complete with no human action between stages.
- **SC-003**: At Validate-complete, the PR is open and unmerged and the issue is open.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The repository default branch is `main`, and the PR targets it.
- The marker line is an HTML comment so it is inert when the markdown file is rendered.
- Single-repo scope: only this repository is touched.

## Out of Scope *(optional)*

- Any change to the engine, stage configuration, or the cruise behaviour itself.
- Cross-repo behaviour.
- Decomposition into child issues.
- Any file other than `e2e/markers/cruise-full-pipeline.md`.
