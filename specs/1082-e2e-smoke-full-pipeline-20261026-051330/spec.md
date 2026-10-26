# Feature Specification: E2E smoke full-pipeline marker (20261026-051330)

**Feature Branch**: `fabrik/issue-1082`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261026-051330 -->`. That's the entire change. One file, one line. Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can go from Specify through Research, Plan, Implement, Review and Validate to Done with a merged PR. The code change is deliberately trivial so that any failure points to the pipeline and not to the change. The marker file records each smoke run as one HTML comment line, so each run leaves a unique, harmless trace in the repository.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Record the smoke run marker (Priority: P1)

An operator running the pipeline smoke needs a single, unique marker line added to the smoke marker file. The line is the only change the pipeline produces, and it is delivered by a merged PR.

**Why this priority**: It is the only deliverable. The marker line is the observable proof that the pipeline completed.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261026-051330 -->`, and the PR diff contains only that one added line.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the change is implemented, **Then** the file ends with the new marker line and all earlier lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing exactly the new marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done without manual intervention.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still begin on its own line and must not be joined to the previous line.
- The marker line already exists in the file because of a re-run: it must not be duplicated.
- The Plan stage must not split this work into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the exact line `<!-- smoke-full-pipeline-20261026-051330 -->`.
- **FR-002**: If the file does not exist, it MUST be created with the marker line as its content.
- **FR-003**: The change MUST be an append only. No existing line in the file is modified or removed.
- **FR-004**: The change MUST touch exactly one file and add exactly one line.
- **FR-005**: The work MUST stay within this single repository and MUST NOT be decomposed into child issues.

### Key Entities

- **Smoke marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log of HTML comment lines, one per smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff consists of exactly one added line in one file.
- **SC-002**: The issue reaches Done with its PR merged and no manual intervention beyond the `fabrik:yolo` label.
- **SC-003**: No child issues are spawned for this work.

## Assumptions

- The marker file already exists in the repository. The create-if-missing rule is kept as a fallback.
- The marker line is unique per run because of its timestamp, so it will not collide with existing lines.

## Out of Scope *(optional)*

- Any change to engine code, stage configuration, or documentation.
- Cross-repo work.
- Decomposing the work into sub-issues.
