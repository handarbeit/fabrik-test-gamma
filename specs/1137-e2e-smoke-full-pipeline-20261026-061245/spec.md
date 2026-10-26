# Feature Specification: E2E smoke full-pipeline marker (20261026-061245)

**Feature Branch**: `fabrik/issue-1137`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261026-061245 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to verify that an issue can progress through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR, under `fabrik:yolo` auto-advance. The code change is deliberately trivial so that any failure points to the pipeline rather than the change. Each smoke run appends one uniquely timestamped marker line to a shared marker file; this run's marker is `smoke-full-pipeline-20261026-061245`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker and land it through the full pipeline (Priority: P1)

An operator running the smoke test wants this run's marker line recorded at the end of the shared marker file and delivered as a merged PR, proving the pipeline works end to end.

**Why this priority**: It is the only deliverable; the marker landing on the default branch is the evidence the smoke run succeeded.

**Independent Test**: After the issue reaches Done, the marker file on `main` ends with the run's marker line, the PR is merged, and the issue is closed.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the change is implemented, **Then** the file's last line is `<!-- smoke-full-pipeline-20261026-061245 -->` and all earlier lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing the run's marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically through to Done with the PR merged.

---

### Edge Cases

- The marker file lacks a trailing newline: the new marker must still occupy its own line.
- The marker file is missing: it is created rather than the stage failing.
- Another smoke run appended concurrently: the change is rebased and both markers are retained, with this run's marker last.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261026-061245 -->`, appended as the last line of the file.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST add exactly one line to exactly one file and MUST NOT modify any existing content.
- **FR-004**: The change MUST be confined to this repository, and the Plan stage MUST NOT decompose the work into child issues.
- **FR-005**: The change MUST be delivered as a PR that references this issue with a closing keyword for this issue and is merged to the default branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, with no manual intervention beyond the initial `fabrik:yolo` label.
- **SC-002**: The merged diff contains exactly one added line in one file, matching the specified marker text.
- **SC-003**: No child issues are created.

## Assumptions

- The marker file currently exists and ends with prior smoke-run markers, one HTML comment per line.
- The default base branch is `main`.
- Marker lines are independent comments, so ordering relative to other runs' lines does not affect correctness.

## Out of Scope *(optional)*

- Any cross-repo work.
- Changes to engine behavior, stage configs, or tests.
- Cleaning up or deduplicating earlier markers in the file.
