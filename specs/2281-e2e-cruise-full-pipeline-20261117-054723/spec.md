# Feature Specification: e2e cruise full pipeline (20261117-054723)

**Feature Branch**: `fabrik/issue-2281`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261117-054723 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end verification of the `fabrik:cruise` pipeline contract (#898). Cruise should carry an issue through every pipeline stage automatically, stopping at Validate-complete without merging the linked PR. After a human merges the PR, the issue should close correctly. The code change itself is deliberately trivial so that any failure observed is a pipeline failure and not a failure of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended to the cruise marker file (Priority: P1)

The marker file `e2e/markers/cruise-full-pipeline.md` ends with the run-specific HTML comment line `<!-- cruise-pipeline-20261117-054723 -->`.

**Why this priority**: This is the whole change. It gives the pipeline a real, reviewable diff to carry through every stage.

**Independent Test**: Inspect the final line of `e2e/markers/cruise-full-pipeline.md` on the issue's branch. It must be exactly the marker line.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is made, **Then** the marker line is added as the new last line and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created and contains the marker line.

---

### User Story 2 - Cruise carries the issue through all stages without merging (Priority: P1)

With `fabrik:cruise` set, the issue advances through every stage to Validate-complete, and the linked PR is left unmerged for a human.

**Why this priority**: This is the contract being verified (#898).

**Independent Test**: Observe the issue after Validate completes. It must carry the Validate-complete state, the PR must be open and unmerged, and the issue must not be in Done.

**Acceptance Scenarios**:

1. **Given** an issue with `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is not merged and the issue is not moved to Done.
3. **Given** a human then merges the PR, **When** the merge lands, **Then** the issue closes correctly.

---

### Edge Cases

- The marker file does not exist yet: it is created with just the marker line.
- The marker file lacks a trailing newline: the marker still ends up on its own line at the very end.
- Plan must not decompose the work or spawn child issues, since the change is one line in one file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The last line of `e2e/markers/cruise-full-pipeline.md` MUST be exactly `<!-- cruise-pipeline-20261117-054723 -->`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: Existing content of the marker file MUST be left unchanged.
- **FR-004**: The change MUST touch only that one file, adding only that one line.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: With `fabrik:cruise` present, the issue MUST advance through all stages to Validate-complete without human approval.
- **FR-007**: The pipeline MUST NOT merge the linked PR, and MUST NOT move the issue to Done, at Validate completion.
- **FR-008**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff consists of exactly one file with exactly one added line, the marker line (plus the generated spec file for this issue).
- **SC-002**: The issue reaches Validate-complete with zero manual stage approvals.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The marker text is taken literally from the issue, including the timestamp `20261117-054723`.
- The marker file's location (`e2e/markers/`) is relative to the repository root.
- Only this single repository is involved; no cross-repo behavior is exercised.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Cross-repo behavior.
- Child issue spawning or decomposition.
- Changes to the pipeline or cruise behavior itself.
