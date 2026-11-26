# Feature Specification: e2e cruise full pipeline marker (20261126-053948)

**Feature Branch**: `fabrik/issue-2632`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261126-053948 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an automated end-to-end check of the `fabrik:cruise` label contract (#898). Under cruise, the pipeline auto-advances through every stage up to Validate-complete but does not merge the PR or move the issue to Done. The issue should close only after a human merges the PR. To exercise that contract, the issue makes one deliberately trivial change: it appends a uniquely timestamped HTML comment marker to a shared marker file. The change is small on purpose, so the pipeline mechanics are the only thing under test and Plan has nothing to decompose.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended (Priority: P1)

An operator running the e2e suite needs a unique, traceable marker for this run committed to the repository, so the run's PR has a real, minimal diff to carry through the pipeline.

**Why this priority**: It is the only deliverable. The cruise pipeline cannot be verified without a PR that has a change.

**Independent Test**: After the change, the last line of `e2e/markers/cruise-full-pipeline.md` is exactly `<!-- cruise-pipeline-20261126-053948 -->`, and the PR diff touches only that file and adds only that line.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the change is made, **Then** the new marker line is the last line of the file and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

An operator wants to confirm that, with `fabrik:cruise` set, the item advances automatically through all stages to Validate-complete without merging, and that the issue closes once a human merges the PR.

**Why this priority**: This is the purpose of the issue; the file change exists only to drive it.

**Independent Test**: Observe the issue's stage progression and PR state across the run, then merge the PR by hand and observe the issue closing.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the next stage starts without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR remains open and unmerged and the issue is not moved to Done.
3. **Given** a human merges the PR, **When** the merge lands, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the new marker must still begin on its own line, so it does not join the previous marker.
- The marker file is absent: it is created, and the marker line is its content.
- Plan must not decompose the work into child issues; it is a single-file, single-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261126-053948 -->`, as the final line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created and contains the marker line.
- **FR-003**: Existing content of the marker file is not modified, reordered or removed.
- **FR-004**: The change touches only `e2e/markers/cruise-full-pipeline.md`; no other file is added, modified or removed.
- **FR-005**: The work is delivered as a single unit, with no child issues spawned.
- **FR-006**: With `fabrik:cruise` set, the pipeline advances through all stages to Validate-complete without human approval at each stage.
- **FR-007**: The pipeline does not merge the PR and does not move the issue to Done on its own at Validate-complete.
- **FR-008**: After a human merges the PR, the issue closes.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a shared file accumulating one timestamped HTML comment per cruise e2e run.
- **Marker line**: `<!-- cruise-pipeline-20261126-053948 -->`, the unique identifier of this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one added line in exactly one file.
- **SC-002**: The last line of the marker file after the change equals the specified marker string exactly.
- **SC-003**: The issue reaches Validate-complete with the PR still open and unmerged, with no manual stage-advance actions.
- **SC-004**: After the PR is merged by a human, the issue is closed.

## Assumptions

- The marker file normally already exists with earlier markers; the create-if-missing path is a fallback.
- The marker string contains the timestamp `20261126-053948` and is used verbatim.
- Cruise precedence and its no-merge behaviour are as documented for `fabrik:cruise`.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Changes to engine behaviour or to the cruise label contract itself.
- Cross-repo scenarios.
- Decomposition into child issues.

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract being verified.
