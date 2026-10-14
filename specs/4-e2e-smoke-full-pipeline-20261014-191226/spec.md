# Feature Specification: e2e smoke full-pipeline marker (20261014-191226)

**Feature Branch**: `fabrik/issue-4`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261014-191226 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. It checks that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial, so a failure points at the pipeline and not at the change. The issue carries `fabrik:yolo`, so stages auto-advance and the linked PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants the issue to produce a merged PR whose only effect is one new marker line at the end of `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: This is the whole change, and its merged result is the evidence that the pipeline works end to end.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261014-191226 -->`, and the PR diff touches only that file and adds only that line.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the marker line is its final line.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the PR has been reviewed and validated, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The file's last line lacks a trailing newline: the marker must still land on its own line, not appended onto the previous line.
- The marker line is already present (for example on a re-run): it must not be duplicated.
- The Plan stage must not decompose the work or spawn child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final line of `e2e/markers/smoke-full-pipeline.md` MUST be exactly `<!-- smoke-full-pipeline-20261014-191226 -->`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing the marker line.
- **FR-003**: Existing content of the file MUST NOT be modified, reordered or removed.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md`, in a single repository.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: The issue MUST pass through all pipeline stages to Done with the change delivered through a merged PR.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added line and no deletions or other file changes.
- **SC-002**: The linked PR is merged into the default branch without manual intervention beyond the `fabrik:yolo` label.
- **SC-003**: The issue ends in the Done column and is closed.
- **SC-004**: No child issues are created.

## Assumptions

- The marker file lives under `e2e/markers/` at the repository root, as the issue states.
- The file is a text/Markdown file where an HTML comment is harmless content.
- `fabrik:yolo` is already applied, so no human approval is needed between stages.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to the marker files for other e2e scenarios.
- Any code, test or documentation change other than the single marker line.
- Decomposing the work into sub-issues.
