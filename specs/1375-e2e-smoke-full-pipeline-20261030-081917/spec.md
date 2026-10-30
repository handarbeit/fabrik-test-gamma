# Feature Specification: e2e smoke full-pipeline (20261030-081917)

**Feature Branch**: `fabrik/issue-1375`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261030-081917 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can progress through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR. The code change is intentionally trivial so that any failure points at the pipeline, not at the work. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator runs the smoke issue through the pipeline. The pipeline produces a PR that appends one marker comment line to `e2e/markers/smoke-full-pipeline.md`, and that PR is merged to the default branch (`main`).

**Why this priority**: This is the only deliverable; a merged PR containing exactly this line proves the full pipeline works.

**Independent Test**: After merge, inspect `e2e/markers/smoke-full-pipeline.md` on `main` and confirm its last line is `<!-- smoke-full-pipeline-20261030-081917 -->` and that it is the only line added.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with prior marker lines, **When** the pipeline completes, **Then** the file's final line is `<!-- smoke-full-pipeline-20261030-081917 -->` and all prior lines are unchanged.
2. **Given** the file does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the issue has `fabrik:yolo`, **When** Validate completes, **Then** the linked PR is merged and the issue reaches Done.

---

### Edge Cases

- The file lacks a trailing newline: the marker must still land on its own final line, not appended to the previous line.
- The marker line is already present (e.g. a re-run): it must not be duplicated.
- Other issues' smoke runs may have appended lines concurrently: those lines must be preserved, and the new line stays at the very end after any rebase.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- smoke-full-pipeline-20261030-081917 -->`, at the very end of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it is created containing that line.
- **FR-003**: No other file is modified and no existing line in the marker file is altered or removed.
- **FR-004**: The work is confined to this single repository; the Plan stage does not decompose it into child issues.
- **FR-005**: The change is delivered via a PR that closes this issue and is merged to the default branch `main`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff consists of exactly one added line in exactly one file.
- **SC-002**: The issue reaches Done with its PR merged, without human intervention beyond applying `fabrik:yolo`.
- **SC-003**: No child issues are spawned for this issue.

## Assumptions

- The marker file already exists on `main` (it holds earlier smoke markers); creation is only a fallback.
- The marker is an HTML comment so it is invisible when the markdown is rendered.
- The timestamp in the marker (`20261030-081917`) is used exactly as given.

## Out of Scope *(optional)*

- Any cross-repo work or issue decomposition.
- Any change to engine code, docs, stage configs, or other marker files.
- Cleaning up or deduplicating earlier smoke markers.
