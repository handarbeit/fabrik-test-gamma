# Feature Specification: E2E smoke full-pipeline marker (20261025-064153)

**Feature Branch**: `fabrik/issue-794`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261025-064153 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, with `fabrik:yolo` auto-advancing and auto-merging. The code change is deliberately trivial so that any failure points at the pipeline and not at the change.

The marker file `e2e/markers/smoke-full-pipeline.md` already exists. It holds one HTML comment line per earlier smoke run, and its last line is a prior run's marker (`<!-- smoke-full-pipeline-20261025-055434 -->`). The file ends with a trailing newline.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. The result is a merged PR that adds this run's unique marker line to the end of the marker file, and the issue ends in Done.

**Why this priority**: This is the only deliverable. It shows the pipeline works from end to end.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261025-064153 -->`, and the issue is in Done.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the change is implemented, **Then** exactly one new line `<!-- smoke-full-pipeline-20261025-064153 -->` is appended after the existing last line, and no existing line is modified or removed.
2. **Given** the PR is open with passing checks, **When** the issue reaches Validate under `fabrik:yolo`, **Then** the PR is merged and the issue moves to Done.

---

### Edge Cases

- The marker file is missing at implementation time: create it, then add the marker line.
- The existing file does not end with a newline: add a line break first so the new marker sits on its own line.
- The marker line is already present (a re-run): do not add a duplicate.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- smoke-full-pipeline-20261025-064153 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: The change MUST NOT modify or remove any existing content of that file.
- **FR-003**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created with the marker line as its content.
- **FR-004**: The change MUST touch exactly one file and add exactly one line, with no changes to any other file.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The PR MUST include `Closes #794` so Fabrik can link it to the issue.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a log of one HTML comment per smoke run.
- **Marker line**: `<!-- smoke-full-pipeline-20261025-064153 -->`, a unique ID for this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The diff of the merged PR is exactly one added line in exactly one file.
- **SC-002**: The issue reaches Done with its linked PR merged, without manual intervention beyond the `fabrik:yolo` label.
- **SC-003**: The marker line appears exactly once, as the last line of the file on the default branch.

## Assumptions

- The marker file already exists, so no creation is needed. FR-003 covers the case where it is missing.
- The trailing marker line is the only content that matters. No other check reads the file.
- `fabrik:yolo` handles auto-advance and auto-merge, so no human approval is needed between stages.

## Out of Scope *(optional)*

- Any cross-repo work or child-issue spawning.
- Any change to engine code, documentation, or other marker files.
- Cleanup or deduplication of earlier run markers in the file.
