# Feature Specification: e2e smoke full-pipeline (20261017-150134)

**Feature Branch**: `fabrik/issue-175`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261017-150134 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to verify that Fabrik can carry an issue through every stage, from Specify to Done, and end with a merged PR. The code change is deliberately trivial, so that any failure points at the pipeline and not at the work. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue and applies `fabrik:yolo`. Without further intervention, the pipeline specifies, researches, plans, implements, reviews and validates the change. It merges a PR that appends one marker comment line to `e2e/markers/smoke-full-pipeline.md`, and the issue ends in Done.

**Why this priority**: This is the only story. It is the whole point of the smoke test.

**Independent Test**: After the run, inspect the default branch. `e2e/markers/smoke-full-pipeline.md` exists and its last line is exactly `<!-- smoke-full-pipeline-20261017-150134 -->`. The PR that made the change is merged, and the issue is in Done and closed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the marker line is added at the very end of the file and existing content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue has the `fabrik:yolo` label, **When** each stage completes, **Then** the issue advances to the next stage without human input, the PR is merged at the end of Validate, and the issue reaches Done.
4. **Given** the issue is single-repo and single-file, **When** the Plan stage runs, **Then** it does not decompose the work into child issues.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must still land on its own line at the end of the file.
- The marker line is already present at the end of the file (for example on a retry): it must not be added a second time.
- The file exists with other markers from earlier smoke runs: those lines are left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- smoke-full-pipeline-20261017-150134 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it, containing the marker line.
- **FR-003**: The change MUST NOT modify any existing content of that file, and MUST NOT touch any other file in the implementation diff.
- **FR-004**: The marker line MUST appear exactly once in the file after the change.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The issue MUST reach Done via a merged PR, with the stages auto-advancing under `fabrik:yolo`.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the default branch after the run, the last line of `e2e/markers/smoke-full-pipeline.md` is exactly `<!-- smoke-full-pipeline-20261017-150134 -->`, and that line occurs once.
- **SC-002**: The implementation diff changes exactly one file and adds at most one line (plus the file creation if it did not exist).
- **SC-003**: The issue ends in Done and closed, and its linked PR is merged, with no manual stage advancement or merge.
- **SC-004**: No child issues are spawned for this issue.

## Assumptions

- The marker line is added as the final line of the file with a trailing newline, following normal text-file convention.
- The default base branch `main` is the target of the PR.
- Spec persistence (`specs/<issue>-<slug>/spec.md`) is produced by the engine and is part of the pipeline's normal output, not part of the "one file, one line" implementation change.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change other than the single marker line (no refactors, docs, or tests).
- Decomposing the work into child issues.

## Source References *(optional)*

- Original issue body (preserved verbatim in **Input**).
