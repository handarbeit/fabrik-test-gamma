# Feature Specification: e2e smoke full-pipeline (20261117-063751)

**Feature Branch**: `fabrik/issue-2298`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261117-063751 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and end at Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline, not at the change. The `fabrik:yolo` label is set so the stages advance automatically and the PR is merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files an issue asking for one marker line to be appended to `e2e/markers/smoke-full-pipeline.md`. The pipeline carries it from Specify to Done with no human intervention beyond the initial labels.

**Why this priority**: This is the whole purpose of the issue. It proves the pipeline works end to end.

**Independent Test**: After the run, check the default branch for `e2e/markers/smoke-full-pipeline.md`. Its last line must be `<!-- smoke-full-pipeline-20261117-063751 -->`. Check that the issue is closed and sits in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's content before the change is unchanged and the new marker line is appended as the final line.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without a human approving it.
4. **Given** Validate completes successfully, **When** the PR's checks and reviews are satisfied, **Then** the PR is merged and the issue is moved to Done.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must go on its own line, not be joined to the previous content.
- The marker line is already present at the end of the file (for example, on a re-run): it must not be duplicated.
- The Plan stage must not split the work into child issues, because the work is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261117-063751 -->`.
- **FR-002**: If the file does not exist, it is created, and then the marker line is appended.
- **FR-003**: If the file exists, all existing content is preserved and the marker is appended on its own line at the very end.
- **FR-004**: The change touches only `e2e/markers/smoke-full-pipeline.md`, and the only change to it is that one added line.
- **FR-005**: The work stays in this one repository and the Plan stage does not decompose it into child issues.
- **FR-006**: The issue completes every pipeline stage, ends with its PR merged, and is moved to Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff contains exactly one file, with exactly one added line, and that line is the marker.
- **SC-002**: The issue reaches Done with its linked PR merged and no human intervention after the issue is filed.
- **SC-003**: No child issues are created.

## Assumptions

- The marker string includes the run timestamp `20261117-063751` exactly as given, so it identifies this run.
- The file's location and name are fixed as `e2e/markers/smoke-full-pipeline.md`.
- The `fabrik:yolo` label drives automatic advancement and merge, so no manual approval is expected.

## Out of Scope *(optional)*

- Any change other than the one marker line.
- Cross-repo work.
- Splitting the work into multiple issues.
- Changes to Fabrik's engine, stage configuration or documentation.
