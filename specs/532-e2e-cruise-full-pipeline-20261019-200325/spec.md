# Feature Specification: e2e cruise full pipeline marker (20261019-200325)

**Feature Branch**: `fabrik/issue-532`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261019-200325 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end test fixture for the `fabrik:cruise` pipeline contract (#898). Cruise should auto-advance an item through every stage up to Validate-complete, without auto-merging the PR or moving the issue to Done. The issue should close correctly once a human merges the PR. The content change is deliberately trivial, a single marker line in a single file, so that any failure observed comes from the pipeline and not from the work itself. Because the change is one line in one file, Plan should not decompose it into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

A maintainer running the cruise e2e check needs a unique, harmless, verifiable change to flow through the pipeline. The change is one HTML comment line appended to the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: This is the whole deliverable. Without it the pipeline has nothing to carry.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue's branch. Its last line is `<!-- cruise-pipeline-20261019-200325 -->`, and the diff against `main` adds only that line, plus the file itself if it was newly created.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is applied, **Then** the marker line is appended as the last line and no existing content is altered.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise carries the item through the pipeline without merging (Priority: P1)

An operator wants confirmation that `fabrik:cruise` auto-advances the item through all stages to Validate-complete and then stops, leaving the PR unmerged.

**Why this priority**: This is the contract the issue exists to verify (#898).

**Independent Test**: Observe the item after Validate completes. Every stage up to Validate is complete, the linked PR is open and unmerged, and the issue is still open and not in Done.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage automatically.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the PR is not merged and the issue is not advanced to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue should close and reach its terminal state.

**Why this priority**: It completes the contract but only runs after the human step.

**Independent Test**: Merge the PR by hand and confirm the issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline. The appended line must still land on its own line.
- The marker file already contains the same marker line. It must not be duplicated.
- A leading blank line or a differently-cased marker must not be introduced. The text is exactly `<!-- cruise-pipeline-20261019-200325 -->`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261019-200325 -->`.
- **FR-002**: If the file does not already exist, it is created with that line as its content.
- **FR-003**: Existing content of the file is left unchanged. The only addition is the marker line.
- **FR-004**: The change touches only that one file in this one repository.
- **FR-005**: The work is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `main` consists of exactly one added marker line in `e2e/markers/cruise-full-pipeline.md`, plus the file's creation if it was missing. The only other files are the engine-generated spec file and any other engine-generated artifacts.
- **SC-002**: The item reaches Validate-complete through automatic advancement with no manual stage moves.
- **SC-003**: At Validate-complete the PR is unmerged and the issue is open.
- **SC-004**: After a human merges the PR, the issue closes.

## Assumptions

- The marker line is appended verbatim, with no surrounding whitespace changes.
- The directory `e2e/markers/` may or may not exist. If it does not, creating the file creates the directory.
- Spec-file generation by the engine is expected pipeline behavior and does not count as scope creep.

## Out of Scope *(optional)*

- Any change to the engine, stages or tests.
- Changes to other marker files.
- Cross-repo work.
- Auto-merging the PR. Cruise deliberately does not do this.
