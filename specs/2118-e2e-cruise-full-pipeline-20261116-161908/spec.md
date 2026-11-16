# Feature Specification: E2E cruise full pipeline marker (20261116-161908)

**Feature Branch**: `fabrik/issue-2118`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261116-161908 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, an item auto-advances through every stage up to Validate-complete, but the PR is not auto-merged and the issue does not move to Done until a human merges the PR. The change itself is deliberately trivial (one appended marker line) so that the pipeline's behaviour, not the change, is what is being verified. A unique timestamped marker per run makes each run's change distinguishable and guarantees a real, non-empty diff.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the e2e suite needs a PR containing exactly one new marker line so the cruise pipeline has real work to carry through all stages.

**Why this priority**: Without the change there is no diff, so no PR to carry through the pipeline.

**Independent Test**: Inspect the PR diff; it adds exactly one line to `e2e/markers/cruise-full-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the last line of the file is `<!-- cruise-pipeline-20261116-161908 -->` and all previously existing lines are unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

An operator needs to confirm that cruise auto-advances through all stages without merging.

**Why this priority**: This is the purpose of the e2e run.

**Independent Test**: Observe the item reach Validate-complete with the PR open and unmerged; merge the PR by hand and observe the issue close.

**Acceptance Scenarios**:

1. **Given** the item carries `fabrik:cruise`, **When** each stage completes, **Then** it advances automatically through to Validate-complete.
2. **Given** Validate is complete, **When** no human has merged the PR, **Then** the PR remains open and the issue is not moved to Done.
3. **Given** a human merges the PR, **When** the engine next settles, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- If the marker file already exists, earlier marker lines from other runs are preserved untouched.
- If the existing file lacks a trailing newline, the new marker still appears on its own final line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- cruise-pipeline-20261116-161908 -->` is appended as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If the file does not exist it is created, containing that line.
- **FR-003**: The change touches only that one file and adds only that one line; no existing line is modified or removed.
- **FR-004**: The work is a single unit and is not decomposed into child issues.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, an append-only log of per-run HTML comment markers.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line in exactly one file.
- **SC-002**: The item reaches Validate-complete via cruise with no human stage approvals, and the PR is unmerged at that point.
- **SC-003**: After a human merges the PR, the issue closes.

## Assumptions

- The marker timestamp `20261116-161908` is used exactly as given.
- The default base branch `main` is the PR target.

## Out of Scope *(optional)*

- Any change to engine behaviour or other marker files.
- Merging the PR automatically (cruise must not do so).
- Decomposing the work into child issues.
