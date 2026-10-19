# Feature Specification: E2E smoke full-pipeline marker line (20261019-063106)

**Feature Branch**: `fabrik/issue-420`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261019-063106 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single issue in a single repository can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and end in Done with a merged PR. The content change is deliberately trivial so that any failure points at the pipeline rather than at the work. The issue carries `fabrik:yolo`, so stages advance automatically and the PR is merged automatically when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

An operator running the smoke test wants a unique, run-specific marker line appended to the end of the smoke marker file, delivered through the normal pipeline. They can then check that the pipeline completed end to end.

**Why this priority**: This is the whole purpose of the issue. There is no other story.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is exactly `<!-- smoke-full-pipeline-20261019-063106 -->`, and the issue is closed on the board's Done column.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's existing content is unchanged and the marker line is appended as the new last line.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances through Research, Plan, Implement, Review and Validate without waiting for human approval, and the linked PR is merged and the issue reaches Done.

---

### Edge Cases

- The file exists but does not end with a trailing newline: the marker must still land on its own line, not be glued onto the previous line.
- The marker line is already present at the end of the file (for example on a re-run): the line must not be duplicated.
- The Plan stage must not split this into child issues; it is a single change in a single repo.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the line `<!-- smoke-full-pipeline-20261019-063106 -->` after the change.
- **FR-002**: If the file does not exist, it is created, and the marker line is its content.
- **FR-003**: Any existing content of the file is preserved unchanged; the marker is appended after it, on its own line.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change touches only this one file and adds only this one line (plus a file-creation if needed). No other files are modified by the implementation.
- **FR-006**: The work stays within a single repository, and the Plan stage does not decompose it into child issues.
- **FR-007**: The PR for the change links to this issue so the pipeline can discover it, and is merged by the pipeline.

### Key Entities

- **Smoke marker file**: `e2e/markers/smoke-full-pipeline.md`, which holds one marker comment line per smoke run.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261019-063106 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the default branch after merge, the last line of `e2e/markers/smoke-full-pipeline.md` is exactly `<!-- smoke-full-pipeline-20261019-063106 -->`, and it occurs once.
- **SC-002**: The merged PR's diff contains exactly one added content line in exactly one file (spec documentation committed by the Specify stage aside).
- **SC-003**: The issue reaches the Done column and is closed with a merged PR, with no human intervention beyond applying the `fabrik:yolo` label.
- **SC-004**: No child issues are created for this issue.

## Assumptions

- The marker is a plain Markdown-compatible HTML comment, so it does not alter how the file renders.
- If the file already exists, the new line is added at the very end, after all existing content, even if earlier runs left other marker lines.
- The `e2e/markers/` directory may need to be created together with the file if it is missing.
- The spec file the Specify stage commits under `specs/` is pipeline output and is not counted as part of the change to the marker file.

## Out of Scope *(optional)*

- Any change beyond the single marker line, including edits to other marker files or documentation.
- Cross-repo work or decomposition into child issues.
- Changes to Fabrik's engine, stage configuration or tests.
- Cleanup of old marker lines left by earlier smoke runs.
