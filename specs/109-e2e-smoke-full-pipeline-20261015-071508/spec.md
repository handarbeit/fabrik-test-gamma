# Feature Specification: e2e smoke full-pipeline (20261015-071508)

**Feature Branch**: `fabrik/issue-109`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261015-071508 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline, not at the change. Each smoke run appends one uniquely identified comment line to a shared marker file, so the file already holds lines from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants a single, unique marker line for this run added to the end of the smoke marker file and merged to the default branch, with no other change.

**Why this priority**: It is the only deliverable. Its merge is the evidence that the pipeline completed.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` is `<!-- smoke-full-pipeline-20261015-071508 -->`, and, apart from the engine-committed spec file, the merged diff touches only that file and adds only that line.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior run lines, **When** the change is implemented, **Then** the file ends with `<!-- smoke-full-pipeline-20261015-071508 -->` as a new line and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing that single line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the item advances automatically through all stages, the linked PR is merged and the issue reaches Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still start on its own line, not be joined to the previous one.
- The marker file already contains this exact line (for example after a retried run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261015-071508 -->` after the change.
- **FR-002**: The line MUST be appended at the very end of the file, with all existing content left unchanged.
- **FR-003**: If the file does not exist, it MUST be created containing that line.
- **FR-004**: Apart from this spec file (`specs/109-e2e-smoke-full-pipeline-20261015-071508/spec.md`), which the engine commits to the issue branch when Specify completes, the change MUST be limited to that one marker file and that one added line.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The linked PR MUST include `Closes #109` so Fabrik can link it to this issue.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log with one HTML comment line per smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged into `main`, driven only by the `fabrik:yolo` auto-advance and with no manual intervention.
- **SC-002**: Apart from the engine-committed spec file, the merged PR's diff is exactly one added line in exactly one file (`e2e/markers/smoke-full-pipeline.md`).
- **SC-003**: No child issues are spawned and no other repository is touched.

## Assumptions

- The marker file currently exists, so it is appended to rather than created.
- Each run's marker line is unique, so a duplicate line only appears if a retry re-applies the change.
- `fabrik:yolo` is the only automation label needed to reach merge and Done.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to the engine, stage configs or documentation.
- Any other change to the marker file, such as cleaning up prior run lines.
- Splitting the work into child issues.
