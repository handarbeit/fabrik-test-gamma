# Feature Specification: e2e smoke full-pipeline (20261025-075333)

**Feature Branch**: `fabrik/issue-830`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261025-075333 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, without human intervention (`fabrik:yolo`). The code change is deliberately trivial so that any failure points to the pipeline and not the change. Earlier smoke runs have each appended one timestamped marker line to the same file.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

A maintainer files the smoke issue. The pipeline appends one marker line to the smoke-test marker file, opens a PR, and merges it. The issue ends in Done.

**Why this priority**: This is the only purpose of the issue. It proves the pipeline works end to end.

**Independent Test**: After the pipeline finishes, `e2e/markers/smoke-full-pipeline.md` on `main` ends with `<!-- smoke-full-pipeline-20261025-075333 -->`, and the issue is in Done with its PR merged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier marker lines, **When** the pipeline implements this issue, **Then** exactly one new line `<!-- smoke-full-pipeline-20261025-075333 -->` is appended at the end and all existing lines are unchanged.
2. **Given** the marker file does not exist, **When** the pipeline implements this issue, **Then** the file is created containing that single line.
3. **Given** the PR passes its gates, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The file has no trailing newline: the marker must still land on its own line, not appended to the previous line.
- The marker line is already present: it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the exact line `<!-- smoke-full-pipeline-20261025-075333 -->` at the end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If the file does not exist, the change MUST create it with that line as its content.
- **FR-003**: The change MUST NOT modify, remove or reorder any existing line in the file.
- **FR-004**: The change MUST touch only that one file, apart from pipeline-generated artifacts such as the spec file the engine commits.
- **FR-005**: The work MUST stay in this repository only. The Plan stage MUST NOT decompose it into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, with no human intervention after the issue is created.
- **SC-002**: The merged diff to the marker file is exactly one added line, equal to the marker text.
- **SC-003**: No child issues are spawned.

## Assumptions

- The file path and marker text in the request are exact and do not need to be changed.
- The `fabrik:yolo` label drives automatic advancement and merging.
- A trailing-newline detail is handled by whoever implements the change and does not need to be decided here.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or documentation.
- Cross-repo work and issue decomposition.
- Changes to other marker files.
