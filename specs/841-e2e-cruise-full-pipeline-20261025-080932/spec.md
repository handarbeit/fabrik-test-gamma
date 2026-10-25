# Feature Specification: e2e cruise full pipeline (20261025-080932)

**Feature Branch**: `fabrik/issue-841`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261025-080932 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end smoke test of the `fabrik:cruise` label contract (#898). Cruise is meant to auto-advance an issue through every pipeline stage up to Validate-complete, but to stop there: the linked PR must not be auto-merged and the issue must not advance to Done on its own. The issue is expected to close only after a human merges the PR.

The change under test is deliberately trivial so that the pipeline behavior, not the change itself, is what gets verified. The change is one marker line appended to a shared marker file. That file already exists in the repository and carries one `cruise-pipeline-<timestamp>` marker per earlier run. This run's marker is unique to its timestamp, `20261025-080932`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

The pipeline produces a PR whose only change is one new line at the end of the cruise marker file.

**Why this priority**: This is the entire deliverable. The pipeline-contract checks in User Story 2 need a real PR to run against.

**Independent Test**: Inspect the PR diff. It should touch one file and add one line.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier cruise markers, **When** the Implement stage completes, **Then** the file's last line is `<!-- cruise-pipeline-20261025-080932 -->` and every earlier line is unchanged.
2. **Given** the PR is open, **When** its diff is reviewed, **Then** it shows exactly one file changed and one line added (the spec file the engine commits under `specs/` aside).

---

### User Story 2 - Cruise stops at Validate-complete (Priority: P1)

With `fabrik:cruise` on the issue, the pipeline advances unattended through every stage to Validate-complete, then waits for a human.

**Why this priority**: This is the contract the test exists to verify (#898).

**Independent Test**: Watch the issue from Specify through Validate with no human action, then check the PR and issue state.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise` and no `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the engine evaluates the item, **Then** the linked PR is still open and unmerged and the issue is still open and not in Done.
3. **Given** Validate-complete and an open PR, **When** a human merges the PR, **Then** the issue closes and reaches Done.

---

### Edge Cases

- The marker file does not exist: it is created, containing only the marker line.
- The marker file lacks a trailing newline: the marker still goes on its own final line, not appended to the previous line.
- The decomposition stage must not split this one-line change into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the exact line `<!-- cruise-pipeline-20261025-080932 -->` at the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: The change MUST NOT modify, reorder or remove any existing line in that file.
- **FR-003**: If the file does not exist, it MUST be created containing the marker line.
- **FR-004**: The change MUST touch no file other than the marker file (the engine's spec-file commit aside) and MUST add no more than that one line.
- **FR-005**: The Plan stage MUST NOT decompose the work or spawn child issues.
- **FR-006**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete without human intervention.
- **FR-007**: At Validate-complete, the engine MUST NOT auto-merge the PR and MUST NOT advance the issue to Done.
- **FR-008**: After a human merges the PR, the issue MUST close and reach Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR diff is exactly one added line, with the exact marker text, at the end of one file.
- **SC-002**: The issue reaches Validate-complete with zero human interactions between the start of Specify and Validate-complete.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is open and not in Done.
- **SC-004**: After the human merge, the issue is closed and in Done.
- **SC-005**: No child issues are created.

## Assumptions

- The marker file already exists in the repository, as the previous cruise runs imply. If it does not, FR-003 covers creating it.
- The timestamp `20261025-080932` is unique to this run and does not collide with any existing marker.
- The issue stays single-repo, as the Scope section stated.
- The marker is an HTML comment, so it does not change how the file renders.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Verifying `fabrik:yolo` behavior (auto-merge) or other autonomy modes.
- Multi-repo behavior.
- Changes to the cruise implementation itself (#898).

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract under test.
- `e2e/markers/cruise-full-pipeline.md`: the target file, which holds the markers from earlier cruise runs.
