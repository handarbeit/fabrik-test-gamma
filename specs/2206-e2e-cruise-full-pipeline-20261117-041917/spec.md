# Feature Specification: e2e cruise full pipeline (20261117-041917)

**Feature Branch**: `fabrik/issue-2206`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261117-041917 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label's pipeline contract (#898). Under that contract, an item carrying `fabrik:cruise` advances automatically through every stage up to Validate-complete, but the engine does not merge the linked PR on its own. A human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial: one marker line added to a marker file. A trivial change keeps the focus on how the pipeline behaves, not on what the change contains. The marker line contains a unique timestamp, so this run's change can be told apart from other runs' changes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

The pipeline produces a PR that adds exactly one marker line, `<!-- cruise-pipeline-20261117-041917 -->`, at the very end of `e2e/markers/cruise-full-pipeline.md`. If the file does not exist, the PR creates it.

**Why this priority**: This is the only deliverable of the issue. The pipeline verification needs a real, minimal change to carry through the stages.

**Independent Test**: Inspect the PR diff. It touches only `e2e/markers/cruise-full-pipeline.md`, and the only added line is the marker comment, at the end of the file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is the last line of the file and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` applied, the item advances automatically through Specify, Research, Plan, Implement, Review and Validate with no human approval between stages. The PR is not merged when Validate completes.

**Why this priority**: Verifying this cruise behavior is the purpose of the issue.

**Independent Test**: Watch the item progress from Specify to Validate-complete without manual intervention. Confirm that the linked PR is still open and unmerged afterwards.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine moves the item to the next stage without waiting for human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR remains open and the issue stays open.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR. The issue then closes, and the item reaches its terminal state.

**Why this priority**: This confirms the second half of the cruise contract. It depends on the first two stories.

**Independent Test**: Merge the PR by hand and confirm that the issue is closed and the item reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate has completed and the PR is open, **When** a human merges the PR, **Then** the issue closes and the item reaches Done.

---

### Edge Cases

- The marker file does not exist: it is created with the marker line as its content.
- The marker file exists but does not end with a newline: the marker still goes on its own line at the end of the file.
- Plan must not decompose the work into child issues, because the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- cruise-pipeline-20261117-041917 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, the change MUST create it with the marker line.
- **FR-003**: The change MUST modify no file other than `e2e/markers/cruise-full-pipeline.md` and MUST add no line other than the marker line.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: The change MUST be confined to this repository.
- **FR-006**: The engine MUST NOT merge the linked PR when Validate completes under `fabrik:cruise`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line, the marker comment, in exactly one file.
- **SC-002**: The item reaches Validate-complete with no manual stage-to-stage intervention.
- **SC-003**: At Validate-complete, the PR is unmerged and the issue is open.
- **SC-004**: After a human merges the PR, the issue is closed.
- **SC-005**: No child issues are created.

## Assumptions

- The timestamp in the marker line is a run identifier and is used exactly as given.
- If the marker file exists, earlier content is left unchanged, and the marker line goes on its own line even if the previous last line had no trailing newline.
- A human performs the PR merge after the pipeline stops, outside this change's scope.

## Out of Scope

- Any change to engine behavior or to the cruise contract itself.
- Changes to any file other than the marker file.
- Multi-repo behavior.
- Splitting the work into child issues.
