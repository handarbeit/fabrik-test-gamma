# Feature Specification: e2e smoke full-pipeline (20261126-044616)

**Feature Branch**: `fabrik/issue-2581`
**Created**: 2026-11-25
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261126-044616 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points at the pipeline, not at the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

An operator runs the issue through the pipeline with `fabrik:yolo`. The result is a merged PR whose only effect is one new marker line at the end of `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: It is the only story. It is the whole purpose of the smoke test.

**Independent Test**: After the issue reaches Done, inspect `e2e/markers/smoke-full-pipeline.md` on the default branch and the merged PR's diff.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** its last line is `<!-- smoke-full-pipeline-20261126-044616 -->` and the rest of the file is unchanged.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the pipeline advances automatically, the PR is merged, and the issue reaches Done.
4. **Given** the work is scoped to a single repository, **When** Plan runs, **Then** it does not decompose or spawn child issues.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still land on its own line.
- The file already contains this exact marker line: the marker must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The marker line `<!-- smoke-full-pipeline-20261126-044616 -->` is the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created first, with the marker line as its content.
- **FR-003**: The change touches exactly one file and adds exactly one line (plus any newline needed to place it on its own line).
- **FR-004**: All work happens in this single repository. No cross-repo work is done and no child issues are spawned.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged.
- **SC-002**: The merged PR's diff shows one file changed with one added marker line.
- **SC-003**: The pipeline passes through every stage with no human intervention beyond applying the `fabrik:yolo` label.

## Assumptions

- The marker line is added verbatim, including the HTML comment delimiters.
- The `fabrik:yolo` label is already on the issue, so stages auto-advance.
- A missing marker file is created with no other content.

## Out of Scope

- Any change other than the single marker line.
- Cross-repo work, decomposition or child issues.
- Changes to engine code, documentation or tests.
