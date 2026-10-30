# Feature Specification: E2E smoke test of the full single-repo pipeline

**Feature Branch**: `fabrik/issue-1424`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261030-202609 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue exists to exercise the pipeline, not to deliver product behavior. It is a smoke test: a deliberately trivial change that lets an operator confirm Fabrik can carry one issue through every stage (Specify → Research → Plan → Implement → Review → Validate) to a merged PR and a Done issue, in a single repository, under `fabrik:yolo` auto-advance. The change is kept to one line in one marker file so that any failure is attributable to the pipeline rather than to the work.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the default branch via the full pipeline (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik runs every stage without manual intervention, opens a PR containing only the marker line, merges it, and moves the issue to Done.

**Why this priority**: It is the sole purpose of the issue; a passing run proves the pipeline works end to end.

**Independent Test**: After the run, check that `e2e/markers/smoke-full-pipeline.md` on the default branch ends with the marker line, that the PR touching it is merged, and that the issue is in Done.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the pipeline completes, **Then** its last line is `<!-- smoke-full-pipeline-20261030-202609 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the pipeline completes, **Then** the file exists and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage and the PR is merged without human action.

---

### Edge Cases

- The marker file exists without a trailing newline: the marker must still occupy its own final line rather than being joined to the previous line.
- The marker line is already present (for example on a retried run): it must not be appended a second time.
- Plan is tempted to split the work: the change is a single file and one line, so no child issues are spawned.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with exactly the line `<!-- smoke-full-pipeline-20261030-202609 -->`.
- **FR-002**: If the file does not exist, it MUST be created, containing the marker line.
- **FR-003**: If the file exists, existing content MUST be preserved and the marker line appended at the very end.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST consist of exactly one file and one added line, with no other files modified.
- **FR-006**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-007**: The change MUST reach the default branch through a merged PR that closes this issue, and the issue MUST end in Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue progresses from Specify to Done with no manual intervention beyond the initial `fabrik:yolo` label.
- **SC-002**: The merged PR's diff is exactly one added line in one file.
- **SC-003**: The default branch's `e2e/markers/smoke-full-pipeline.md` ends with the marker line, which appears once.
- **SC-004**: No child issues are created.

## Assumptions

- `fabrik:yolo` is applied, so every stage auto-advances and the PR is auto-merged after Validate.
- The marker directory `e2e/markers/` already exists in the repository; if not, creating it is implied by creating the file.
- The marker text is used literally, including the timestamp suffix `20261030-202609`.

## Out of Scope

- Cross-repo work or multi-repo pipelines.
- Any change other than the single marker line.
- Modifying pipeline or engine behavior; this issue only observes it.
