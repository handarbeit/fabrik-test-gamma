# Feature Specification: e2e smoke full-pipeline (20261126-161753)

**Feature Branch**: `fabrik/issue-2730`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261126-161753 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. It checks that a single issue in a single repository moves through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR, without human intervention. The change is deliberately trivial so that any failure points to the pipeline, not the content. The run's timestamp (`20261126-161753`) is embedded in the change so this run's merged result can be told apart from other smoke runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on main via the full pipeline (Priority: P1)

An operator files this issue with the `fabrik:yolo` label. Fabrik drives it through every stage unattended. It produces a PR that appends one marker line to the smoke-test marker file, and it merges that PR. The issue then reaches Done.

**Why this priority**: This is the only story. It is the whole purpose of the smoke run.

**Independent Test**: After the run, check that `e2e/markers/smoke-full-pipeline.md` on the default branch ends with the exact marker line, that the PR is merged, and that the issue is in Done.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the default branch, **When** the pipeline completes, **Then** the file exists and contains the marker line `<!-- smoke-full-pipeline-20261126-161753 -->`.
2. **Given** the marker file already exists, **When** the pipeline completes, **Then** its prior content is unchanged and the marker line is the last line of the file.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage finishes, **Then** the issue advances automatically, the PR is merged, and the issue reaches Done.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker line must still be on its own line, at the end of the file.
- The marker file exists and already contains markers from earlier smoke runs: those are preserved, and only this run's line is appended.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- smoke-full-pipeline-20261126-161753 -->`, at the very end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created, and its only content is the marker line.
- **FR-003**: The change touches no file other than `e2e/markers/smoke-full-pipeline.md` (the spec file the engine commits for this issue is the engine's own projection and is excluded).
- **FR-004**: Work is confined to this single repository. The Plan stage does not decompose the work or spawn child issues.
- **FR-005**: The change is delivered through a PR that carries `Closes #2730` and is merged to the default branch (`main`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, `e2e/markers/smoke-full-pipeline.md` on `main` ends with `<!-- smoke-full-pipeline-20261126-161753 -->`, and the diff against its prior state is exactly one added line (or one new one-line file).
- **SC-002**: The linked PR is merged and the issue is in Done, with no human action after the issue was filed.
- **SC-003**: No child issues are spawned and no repository other than this one is modified.

## Assumptions

- The `fabrik:yolo` label stated in the issue is what lets the pipeline advance and merge unattended, so no approval gates are expected.
- "Appended at the very end of the file" means the marker occupies its own final line. A trailing newline after it is acceptable.
- The marker line is an HTML comment, so the file renders no visible change.

## Out of Scope

- Any change to application code, tests, or documentation.
- Cross-repository work, or decomposing the issue into child issues.
- Changes to the pipeline configuration or the stage definitions.
