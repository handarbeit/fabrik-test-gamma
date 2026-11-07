# Feature Specification: E2E smoke full-pipeline marker (20261107-195615)

**Feature Branch**: `fabrik/issue-1761`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: 

> ## Goal
>
> End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.
>
> ## Trivial change
>
> Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> ```
> <!-- smoke-full-pipeline-20261107-195615 -->
> ```
>
> That's the entire change. One file, one line.
>
> ## Scope
>
> Single repo only — no cross-repo work. The Plan stage should NOT decompose.
>
> ## Labels
>
> fabrik:yolo

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) and end at Done with a merged PR. The code change is deliberately minimal, so any failure points at the pipeline and not at the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator runs the smoke test. The pipeline appends a unique marker comment line to the smoke-test marker file, and the change reaches the default branch through a merged PR with the issue ending in Done.

**Why this priority**: This is the whole purpose of the issue. Without it, the smoke test proves nothing.

**Independent Test**: After the pipeline finishes, check that the default branch's `e2e/markers/smoke-full-pipeline.md` ends with the marker line `<!-- smoke-full-pipeline-20261107-195615 -->`, that the PR is merged, and that the issue is at Done.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line as its last line.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and the existing content is unchanged.
3. **Given** the change is complete and validated, **When** the issue has `fabrik:yolo`, **Then** the linked PR is merged and the issue reaches Done without manual intervention.

---

### Edge Cases

- The marker file exists and does not end with a newline: the marker line must still land on its own line, at the very end of the file.
- The marker line is already present (for example on a re-run): the file must not end up with a duplicate of the same line.
- Other marker files in the same directory are left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the exact line `<!-- smoke-full-pipeline-20261107-195615 -->`.
- **FR-002**: If the file does not exist, it MUST be created, with the marker line as its content.
- **FR-003**: If the file exists, its prior content MUST be preserved and the marker line appended after it.
- **FR-004**: The change MUST touch only this one file, adding only this one line (plus the file's creation if it was absent).
- **FR-005**: The work MUST be done within this single repository; the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The change MUST be delivered through a PR that is merged, with the issue ending at Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline completes, the marker file on the default branch ends with the marker line, and the diff of the merged PR is exactly one added line (or one new file holding that line).
- **SC-002**: The issue progresses from Specify to Done with no human intervention beyond applying the `fabrik:yolo` label.
- **SC-003**: No child issues are spawned.

## Assumptions

- The `e2e/markers/` directory may not exist yet. Creating it is implied by creating the file.
- The marker line is plain text with no surrounding blank lines or extra formatting beyond a single trailing newline.
- The `fabrik:yolo` label already on the issue is sufficient for auto-advance and auto-merge, so no human approval of the spec is needed.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration, or documentation.
- Cross-repo work or decomposition into child issues.
- Any file other than `e2e/markers/smoke-full-pipeline.md`.
