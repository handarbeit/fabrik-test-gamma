# Feature Specification: E2E cruise full pipeline marker (20261017-175659)

**Feature Branch**: `fabrik/issue-277`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261017-175659 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end probe of the `fabrik:cruise` label contract (#898). Cruise should auto-advance the item through every pipeline stage up to Validate-complete, stop there without merging the linked PR, and let the issue close correctly once a human merges the PR. The code change is deliberately trivial: it exists only to give the pipeline a real, minimal PR to carry. The marker file already exists in the repository and holds one such comment line per earlier e2e run, so this run appends one more line carrying its own timestamp.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended (Priority: P1)

A maintainer running the cruise e2e check needs this run's unique marker recorded in the repository, as the single visible change that the pipeline carries.

**Why this priority**: Without the change there is no PR, and the cruise contract cannot be observed.

**Independent Test**: Inspect the resulting PR diff and the marker file. The diff adds exactly one line to `e2e/markers/cruise-full-pipeline.md`, and that line is the last line of the file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior run markers, **When** the change is implemented, **Then** the file ends with the line `<!-- cruise-pipeline-20261017-175659 -->` and all earlier lines are unchanged.
2. **Given** the marker file did not exist, **When** the change is implemented, **Then** the file is created containing that single line.

---

### User Story 2 - Cruise advances without merging (Priority: P1)

A maintainer wants confirmation that an item labelled `fabrik:cruise` auto-advances through all stages and stops at Validate-complete with the PR unmerged.

**Why this priority**: This is the contract under test (#898).

**Independent Test**: Observe the board and the PR after the item reaches Validate-complete. The item has passed through every stage automatically, and the PR is open and not merged.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage without human approval.
2. **Given** the item has completed Validate, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after human merge (Priority: P2)

After a human merges the PR, the issue closes correctly.

**Why this priority**: It completes the contract but depends on a manual step outside the engine's automation.

**Independent Test**: Merge the PR by hand, then check that the issue ends up closed and in Done.

**Acceptance Scenarios**:

1. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file's last line has no trailing newline: the new marker must still land on its own line, not be appended to the previous one.
- The marker file is absent: it is created with just the new line.
- Plan must not split this work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- cruise-pipeline-20261017-175659 -->`, to the end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: No existing line in the marker file may be altered or removed.
- **FR-004**: No file other than `e2e/markers/cruise-full-pipeline.md` may be changed.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: Under `fabrik:cruise`, the item MUST auto-advance through all stages to Validate-complete, and the linked PR MUST NOT be merged by the engine.
- **FR-007**: After a human merges the PR, the issue MUST close.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, an append-only log of run markers, one HTML comment per line.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line in exactly one file.
- **SC-002**: The item reaches Validate-complete with no human action between stages.
- **SC-003**: The PR is still unmerged when the item reaches Validate-complete.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker file is expected to already exist, with prior run markers; creating it is only a fallback.
- The run timestamp `20261017-175659` in the marker is fixed by this issue and must not be changed.
- A single repository is involved; no cross-repo behavior is exercised.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Cleaning up or deduplicating earlier markers.
- Changes to engine behavior or to cruise semantics.
- Decomposing the work into child issues.
