# Feature Specification: E2E smoke — full pipeline marker line

**Feature Branch**: `fabrik/issue-68`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261014-232701 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik SDLC pipeline. Its purpose is to confirm that a single-repo issue can travel through every stage (Specify → Research → Plan → Implement → Review → Validate) and finish in Done with its pull request merged. The code change is deliberately trivial so that any failure points at the pipeline rather than the change itself. The `fabrik:yolo` label requests automatic advancement between stages and automatic merge of the PR at the end of Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline. The only repository change is one new comment line at the end of `e2e/markers/smoke-full-pipeline.md`. The pipeline carries it through to a merged PR and the issue reaches Done.

**Why this priority**: This is the whole purpose of the issue; proving the pipeline completes end to end is the only deliverable.

**Independent Test**: After the issue reaches Done, inspect `e2e/markers/smoke-full-pipeline.md` on the default branch and confirm its last line is `<!-- smoke-full-pipeline-20261014-232701 -->` and the PR that introduced it is merged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** exactly one line, `<!-- smoke-full-pipeline-20261014-232701 -->`, is added at the very end of the file and no existing line is modified.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created containing that single line.
3. **Given** the PR is open with the one-line change, **When** Validate completes under `fabrik:yolo`, **Then** the PR is merged and the issue is moved to Done.

---

### Edge Cases

- The file already ends without a trailing newline: the new line must still begin on its own line, not be joined to the previous last line.
- The marker line is already present at the end of the file (for example, from a retried run): it must not be duplicated.
- Earlier smoke runs have left their own marker lines in the file; these must be left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- smoke-full-pipeline-20261014-232701 -->` as the final line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created, containing that line.
- **FR-003**: The change MUST touch only that one file and add only that one line; no existing content may be altered or removed.
- **FR-004**: The work MUST remain within this single repository; the Plan stage MUST NOT decompose the issue or spawn child issues.
- **FR-005**: The change MUST be delivered through a single PR that closes this issue and is merged automatically under `fabrik:yolo`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches the Done column with every stage having completed, without manual intervention beyond the initial labels.
- **SC-002**: The merged diff consists of exactly one file with exactly one added line, and zero removed lines.
- **SC-003**: On the default branch after merge, the marker line `<!-- smoke-full-pipeline-20261014-232701 -->` appears exactly once, as the last line of the file.

## Assumptions

- The file `e2e/markers/smoke-full-pipeline.md` already exists on `main` with earlier smoke-run marker lines; if it does not, creating it is explicitly permitted.
- The marker line is an HTML comment, so it does not affect how the markdown renders.
- The default base branch (`main`) is the PR target; no `base:<branch>` label applies.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration, or documentation.
- Cross-repo work or decomposition into child issues.
- Cleaning up or reorganising earlier smoke marker lines or files.
