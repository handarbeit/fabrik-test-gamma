# Feature Specification: E2E smoke full-pipeline (20261019-060119)

**Feature Branch**: `fabrik/issue-410`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261019-060119 -->`. That's the entire change. One file, one line. Scope: single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to confirm that a single-repo issue can travel through every stage (Specify → Research → Plan → Implement → Review → Validate) and end in Done with a merged PR. The change itself is deliberately trivial so that any failure points at the pipeline, not the work. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator runs the smoke test and expects the marker line for this run to appear at the end of the marker file on the default branch, delivered through a merged PR.

**Why this priority**: This is the entire change and the observable proof the pipeline worked end to end.

**Independent Test**: After the issue reaches Done, read `e2e/markers/smoke-full-pipeline.md` on `main` and confirm its last line is `<!-- smoke-full-pipeline-20261019-060119 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with earlier run markers, **When** the pipeline completes, **Then** the file's final line is `<!-- smoke-full-pipeline-20261019-060119 -->` and all earlier lines are unchanged.
2. **Given** the file does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the pipeline runs, **When** it finishes, **Then** the PR linked to this issue is merged and the issue is in Done.

---

### Edge Cases

- The file already exists with or without a trailing newline: the new marker must land on its own line, not appended to the previous line.
- The marker for this run is already present (re-run): it should not be duplicated.
- Plan must not split the work into child issues; this is a single-file, single-line change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261019-060119 -->`, created if absent.
- **FR-002**: The change adds only that one line to that one file; existing content is untouched.
- **FR-003**: The change is delivered via a PR that references this issue and is merged.
- **FR-004**: The work is not decomposed into child issues and touches no repository other than this one.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, without manual intervention beyond the `fabrik:yolo` label.
- **SC-002**: The merged diff is exactly one added line in one file.
- **SC-003**: No child issues are spawned.

## Assumptions

- The marker file currently exists with earlier run markers, but the spec does not depend on that; if it is absent it is created.
- The `fabrik:yolo` label provides auto-advance and auto-merge, so no human approval of this spec is needed.

## Out of Scope

- Any change beyond the single marker line.
- Cross-repo work.
- Changes to pipeline behavior or stage configuration.
