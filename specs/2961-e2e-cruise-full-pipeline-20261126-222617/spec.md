# Feature Specification: e2e cruise full pipeline marker (20261126-222617)

**Feature Branch**: `fabrik/issue-2961`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Trivial change: append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261126-222617 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an item through every pipeline stage automatically, stop once Validate is complete without merging the linked PR, and leave the merge to a human. After that human merge, the issue should close correctly.

The code change is deliberately minimal. A single marker line is appended to a tracked file, so the work is a small, predictable diff that exercises the pipeline without any product-logic risk. The timestamp in the marker (`20261126-222617`) makes this run's change unique and identifiable.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

An operator running the e2e suite needs this run's unique marker line to appear at the end of `e2e/markers/cruise-full-pipeline.md` in the PR, so the run's change can be recognised and the pipeline's output checked.

**Why this priority**: It is the only code change. Without it there is no PR diff to carry through the pipeline.

**Independent Test**: Inspect the PR diff. It should contain exactly one changed file, `e2e/markers/cruise-full-pipeline.md`, with exactly one added line, `<!-- cruise-pipeline-20261126-222617 -->`, as the last line of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the marker line is the new last line.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator wants confirmation that an item labelled `fabrik:cruise` moves automatically through every stage up to and including Validate, and that the PR is left unmerged.

**Why this priority**: This is the stated purpose of the e2e run, namely verifying the cruise contract.

**Independent Test**: Watch the item progress from Specify to Validate-complete with no manual stage advancement. Confirm that the linked PR is still open and unmerged, and the issue is not moved to Done, when Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances it to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the linked PR is not merged and the issue is not advanced to Done.

---

### User Story 3 - Issue closes correctly after a human merge (Priority: P2)

After Validate-complete, a human merges the PR. The operator wants the issue to close and reach Done as expected.

**Why this priority**: It confirms the post-merge half of the cruise contract, but it depends on the earlier stories.

**Independent Test**: Merge the PR by hand. Confirm the issue is closed and the item reaches Done.

**Acceptance Scenarios**:

1. **Given** the item is Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue closes and the item reaches Done.

---

### Edge Cases

- The marker file does not exist: it is created, containing only the marker line.
- The file's existing last line has no trailing newline: the marker still appears on its own line.
- The change must not touch any file other than `e2e/markers/cruise-full-pipeline.md`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- cruise-pipeline-20261126-222617 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, the change MUST create it first, so that it contains the marker line.
- **FR-003**: The change MUST NOT modify any pre-existing content in the file or any other file.
- **FR-004**: The work MUST be delivered as one file and one added line, and Plan MUST NOT decompose it into child issues.
- **FR-005**: With `fabrik:cruise` present, the item MUST auto-advance through all stages to Validate-complete without a human approving each stage.
- **FR-006**: At Validate-complete, the engine MUST NOT merge the linked PR and MUST NOT move the issue to Done.
- **FR-007**: After a human merges the PR, the issue MUST close and the item MUST reach Done.

### Key Entities

- **Marker file** (`e2e/markers/cruise-full-pipeline.md`): The tracked file that receives the run's unique marker line.
- **Marker line**: The HTML comment `<!-- cruise-pipeline-20261126-222617 -->`, which identifies this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, and the added line is the marker line.
- **SC-002**: The item reaches Validate-complete with no manual stage advancement.
- **SC-003**: At Validate-complete, the linked PR is still open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed and in Done.
- **SC-005**: No child issues are spawned.

## Assumptions

- The `e2e/markers/` directory is acceptable for the new file if it does not yet exist.
- The marker line is a complete HTML comment on its own line, with no other content on that line.
- This issue covers a single repository only.

## Out of Scope *(optional)*

- Any change to engine, pipeline or cruise behaviour. This issue only verifies it.
- Any file other than `e2e/markers/cruise-full-pipeline.md`.
- Decomposing the work into child issues.
- Merging the PR automatically.
