# Feature Specification: E2E cruise full pipeline verification (20261116-144435)

**Feature Branch**: `fabrik/issue-2038`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261116-144435 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

The `fabrik:cruise` label is meant to auto-advance an issue through every pipeline stage up to Validate-complete. It must stop there: it does not merge the PR or move the issue to Done. A human then merges the PR, and the issue should close correctly. This issue exercises that contract end to end with the smallest possible change, so any failure points at the pipeline and not at the content of the work. The change is one marker line appended to one file. It is intentionally not decomposable, so Plan must not spawn child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

An operator running the cruise end-to-end check needs a single, uniquely identifiable marker line added to the cruise marker file. The line shows that this specific run's change went through the pipeline.

**Why this priority**: It is the only deliverable of the issue, and the rest of the verification depends on it.

**Independent Test**: Inspect the final content of the marker file on the PR branch. The last line must be exactly the marker, and it must be the only line this change adds.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is applied, **Then** the file's existing content is unchanged and the marker line `<!-- cruise-pipeline-20261116-144435 -->` is its final line.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

An operator wants to confirm that cruise carries the issue through every stage automatically and stops short of merging.

**Why this priority**: This is the pipeline contract (#898) the issue exists to verify.

**Independent Test**: Observe the issue and its PR after the pipeline settles.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate completes, **When** the pipeline settles, **Then** the PR is still open and unmerged, and the issue has not been moved to Done.
3. **Given** both the PR and the issue remain open after Validate-complete, **When** a human merges the PR, **Then** the issue closes correctly.

---

### Edge Cases

- If the marker file exists but its last line has no trailing newline, the marker must still end up on its own line.
- If the marker is already present (for example after a re-run), it must not be duplicated.
- Plan must not split this one-line change into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker file `e2e/markers/cruise-full-pipeline.md` ends with the line `<!-- cruise-pipeline-20261116-144435 -->`.
- **FR-002**: If the marker file does not exist, it is created, and it contains the marker line.
- **FR-003**: The change touches exactly one file and adds exactly one line, apart from the creation case in FR-002 and the spec projection the engine adds under `specs/`.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The Plan stage does not decompose the work or spawn child issues.
- **FR-006**: With `fabrik:cruise` present, the issue advances through all stages to Validate-complete without human approval at each stage.
- **FR-007**: The PR is not auto-merged and the issue is not advanced to Done when Validate completes.
- **FR-008**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The final PR diff contains exactly one added marker line in the marker file, apart from the engine's spec projection.
- **SC-002**: The issue reaches Validate-complete with no human stage approvals, and its PR is still open and unmerged at that point.
- **SC-003**: After a human merges the PR, the issue is closed with no manual label or board intervention.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The marker file path is relative to the repository root of the single target repository.
- The marker file may or may not already exist; either case is handled as described above.
- This issue exists only for pipeline verification, and the marker line has no functional effect on the repository.

## Out of Scope *(optional)*

- Any change beyond the one marker line.
- Changes to the cruise feature or the pipeline itself.
- Multi-repo behavior.
- Decomposing the work or spawning child issues.
