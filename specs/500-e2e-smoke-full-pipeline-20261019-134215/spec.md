# Feature Specification: E2E smoke full-pipeline marker (20261019-134215)

**Feature Branch**: `fabrik/issue-500`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-134215 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline. It checks that an issue can move through Specify, Research, Plan, Implement, Review and Validate to Done, ending in a merged PR, without human intervention (`fabrik:yolo`). The code change is deliberately trivial so that any failure points at the pipeline and not at the change. The marker file `e2e/markers/smoke-full-pipeline.md` already exists and holds one marker comment line per earlier smoke run, the most recent being `<!-- smoke-full-pipeline-20261019-125645 -->`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker line (Priority: P1)

An operator running the pipeline smoke test needs this run to leave one unique, identifiable trace in the repository. That trace shows the issue went through every stage and the PR merged.

**Why this priority**: It is the only change in the issue. The marker line is the evidence that the smoke run completed.

**Independent Test**: After the PR merges, read the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch. It must be the marker line for this run.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with earlier markers, **When** the change is applied, **Then** the line `<!-- smoke-full-pipeline-20261019-134215 -->` is the last line of the file and all earlier lines are unchanged.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the PR for this issue passes CI and review, **When** the pipeline reaches the end of Validate under `fabrik:yolo`, **Then** the PR is merged and the issue ends in Done.

---

### Edge Cases

- The file's last line has no trailing newline: the marker must still go on its own line, not be joined to the previous line.
- The marker line is already present (for example on a re-run): it must not be added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The last line of `e2e/markers/smoke-full-pipeline.md` MUST be exactly `<!-- smoke-full-pipeline-20261019-134215 -->`, on its own line.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST NOT modify any existing line of that file, and MUST NOT touch any other file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log with one HTML comment line per smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the default branch's copy of the marker file ends with the line for run `20261019-134215`, and the diff against the previous version is exactly one added line.
- **SC-002**: The issue reaches Done with its linked PR merged, with no human action beyond applying `fabrik:yolo`.
- **SC-003**: No child issues are spawned and no other repository is touched.

## Assumptions

- The marker file already exists on `main`. The create-if-missing clause is a fallback only.
- The HTML comment form means the line renders as nothing in Markdown, so it has no visible effect.
- Normal CI and review gates apply and are not bypassed by this spec.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or documentation.
- Cross-repo work and issue decomposition.
- Cleaning up or reformatting earlier marker lines.
