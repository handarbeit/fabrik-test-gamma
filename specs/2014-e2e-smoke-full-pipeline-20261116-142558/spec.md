# Feature Specification: e2e smoke full-pipeline (20261116-142558)

**Feature Branch**: `fabrik/issue-2014`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261116-142558 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to show that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, with no human intervention beyond the `fabrik:yolo` label. The code change is deliberately trivial so that any failure points at the pipeline rather than at the work. The change is a single marker line appended to a marker file.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator running a smoke check files this issue with `fabrik:yolo`. Fabrik takes it through every stage unattended. The result is a merged PR whose only change is the marker line at the end of `e2e/markers/smoke-full-pipeline.md`, and the issue ends in Done.

**Why this priority**: This is the only story. It is the whole purpose of the issue.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is exactly `<!-- smoke-full-pipeline-20261116-142558 -->`, and the PR's diff touches only that file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
2. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and all existing content is unchanged.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged automatically and the issue reaches Done.
4. **Given** the issue is single-repo, **When** the Plan stage runs, **Then** it produces one plan with no child issues spawned.

---

### Edge Cases

- The file already exists and its last line has no trailing newline: the marker still goes on its own line, not appended to the previous line.
- The marker line is already present in the file (for example after a stage re-run): it is not added a second time.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261116-142558 -->`.
- **FR-002**: If the file does not exist, it is created (along with any missing parent directory) containing that line.
- **FR-003**: If the file exists, its existing content is preserved unchanged and the marker line is appended after it, on its own line.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change modifies no file other than `e2e/markers/smoke-full-pipeline.md`.
- **FR-006**: The work stays within this single repository, and the Plan stage does not decompose it or spawn child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with a merged PR, with no human action after the issue is filed and labeled `fabrik:yolo`.
- **SC-002**: The merged PR's diff contains exactly one changed file and, in the existing-file case, exactly one added line.
- **SC-003**: On the default branch after merge, the last line of `e2e/markers/smoke-full-pipeline.md` matches the marker string exactly.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The marker line is followed by a trailing newline, as is conventional for text files.
- The file's content is otherwise irrelevant to this test, so no other text needs to be added when the file is created.
- `fabrik:yolo` is already on the issue, so no manual approval is needed at any stage.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to engine code, stage configuration, documentation, or tests.
- Decomposing the work into child issues.
- Any content in the marker file beyond the single marker line.
