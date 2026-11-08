# Feature Specification: E2E smoke full-pipeline marker (20261108-014638)

**Feature Branch**: `fabrik/issue-1827`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261108-014638 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is a smoke test of the whole Fabrik pipeline (Specify → Research → Plan → Implement → Review → Validate → Done) on a single repository. The change is deliberately trivial so that any failure points at the pipeline rather than at the work. Success means the issue travels the whole way and ends with a merged PR. The `fabrik:yolo` label is present, so each stage advances automatically and the PR is auto-merged when Validate completes.

The marker file `e2e/markers/smoke-full-pipeline.md` already exists and holds one HTML-comment line per earlier smoke run. Its latest entry is `smoke-full-pipeline-20261107-195615`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged (Priority: P1)

An operator running the release or smoke checks wants to see one issue go from Specify to Done unaided, ending in a merged PR that appends this run's marker line.

**Why this priority**: This is the only purpose of the issue. It exercises every pipeline stage and the auto-merge path.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` is `<!-- smoke-full-pipeline-20261108-014638 -->`. Every earlier line is unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier entries, **When** the pipeline implements this issue, **Then** exactly one new line, `<!-- smoke-full-pipeline-20261108-014638 -->`, is appended after the last existing line.
2. **Given** the marker file does not exist, **When** the pipeline implements this issue, **Then** the file is created and contains that single line.
3. **Given** the PR is open with passing checks, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The file's last line has no trailing newline: the new marker must start on its own line and must not be joined to the previous entry.
- The marker line is already present because of a retry or re-run: it must not be appended a second time.
- Another smoke run appends to the same file at the same time: the rebase or merge resolves it so that both markers are kept and no earlier line is lost.
- The Plan stage must not split this single-line change into child issues.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- smoke-full-pipeline-20261108-014638 -->`, at the end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created with that line as its only content.
- **FR-003**: The new line appears exactly once in the file.
- **FR-004**: No existing line in the file is modified, removed or reordered.
- **FR-005**: The only file changed by the implementation PR is `e2e/markers/smoke-full-pipeline.md`. The engine-committed spec artifact under `specs/` is excluded from this requirement.
- **FR-006**: The work stays within this repository; no child or cross-repo issues are created, and Plan does not decompose the work.
- **FR-007**: The linked PR merges, and the issue reaches Done, without manual intervention.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log of smoke runs with one HTML-comment line per run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of the marker file on `main` is `<!-- smoke-full-pipeline-20261108-014638 -->`, and that string occurs exactly once.
- **SC-002**: The diff against the previous file contents is a single added line, with no deletions.
- **SC-003**: The issue completes Specify, Research, Plan, Implement, Review and Validate and reaches Done with a merged PR, with no human input.
- **SC-004**: No child issues are spawned.

## Assumptions

- The `fabrik:yolo` label drives automatic advancement and merging, so no human approves the spec or the PR.
- The exact marker text is `<!-- smoke-full-pipeline-20261108-014638 -->`, with no surrounding whitespace or quotes.
- The engine-committed `specs/1827-*/spec.md` file is expected in the PR and does not break the "one file, one line" scope.

## Out of Scope *(optional)*

- Any change to engine, stage or workflow code.
- Cross-repo work or decomposition into sub-issues.
- Cleaning up or reordering earlier marker lines.
- Adding tests or documentation beyond the one marker line.
