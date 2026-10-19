# Feature Specification: E2E smoke full-pipeline marker (20261019-154343)

**Feature Branch**: `fabrik/issue-507`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-154343 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can progress from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged pull request. The code change is deliberately minimal so that any failure points at the pipeline and not at the work. The issue carries the `fabrik:yolo` label, so each stage is expected to advance on its own and the pull request is expected to merge automatically once Validate completes.

The marker file already exists in the repository and holds one HTML comment line per earlier smoke run. This run adds its own line, identified by the timestamp `20261019-154343`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the smoke suite wants a unique, visible trace of this run in the repository. The trace is a single marker line appended to the end of the smoke marker file. Its presence on the default branch after the pipeline finishes shows that the run completed and merged.

**Why this priority**: This is the whole change. The smoke test only has value if the change lands through the pipeline.

**Independent Test**: After the pull request merges, read the marker file on the default branch. It ends with the exact line `<!-- smoke-full-pipeline-20261019-154343 -->`, and every earlier line is unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists and ends with earlier smoke marker lines, **When** the change is implemented, **Then** the file's last line is exactly `<!-- smoke-full-pipeline-20261019-154343 -->` and all previous lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing exactly the marker line.
3. **Given** the change is complete, **When** the pull request is inspected, **Then** it touches only the marker file and adds exactly one line.

---

### User Story 2 - Pipeline completes without human intervention (Priority: P2)

The operator wants the issue to reach Done with a merged pull request, with no manual steps after the `fabrik:yolo` label is applied.

**Why this priority**: This is the behaviour the smoke test exists to verify, but it depends on the change in User Story 1.

**Independent Test**: With the `fabrik:yolo` label applied, watch the issue move through every stage. It ends in Done and its linked pull request is merged.

**Acceptance Scenarios**:

1. **Given** the issue is labelled `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without human input.
2. **Given** Validate completes, **When** the pull request is ready, **Then** it is merged and the issue moves to Done.
3. **Given** the work is a single-file, single-repo change, **When** the Plan stage runs, **Then** it does not split the work into child issues.

---

### Edge Cases

- The marker file does not exist: it is created, with the marker line as its only content.
- The marker file's last line lacks a trailing newline: the marker goes on its own new line and does not join the previous line.
- The marker line is already present, for example on a re-run: it is not added a second time.
- Other smoke marker lines are present from earlier runs: they are left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The smoke marker file `e2e/markers/smoke-full-pipeline.md` ends with exactly one new line, `<!-- smoke-full-pipeline-20261019-154343 -->`, appended at the end of the file.
- **FR-002**: If the marker file does not exist, it is created with the marker line as its content.
- **FR-003**: Existing content of the marker file is preserved unchanged.
- **FR-004**: The change touches no file other than the marker file.
- **FR-005**: The work is a single-repository change, and the Plan stage does not decompose it into child issues.
- **FR-006**: The issue progresses from Specify to Done under `fabrik:yolo`, ending with the linked pull request merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the default branch's marker file ends with exactly `<!-- smoke-full-pipeline-20261019-154343 -->`, and a diff against the previous version shows a single added line.
- **SC-002**: The linked pull request is merged and the issue is in Done, with no human action taken after the issue was labelled `fabrik:yolo`.
- **SC-003**: No child issues are spawned from this issue.

## Assumptions

- The marker file already exists in the repository, with earlier smoke marker lines. The create-if-missing rule is a safeguard only.
- The marker is a Markdown/HTML comment, so it does not change how the file renders.
- Under `fabrik:yolo`, a clean pipeline run needs no clarification from the operator.

## Out of Scope *(optional)*

- Any cross-repository work.
- Any change to code, documentation, configuration or tests beyond the single marker line.
- Cleaning up or consolidating the earlier smoke marker lines.
- Changes to pipeline behaviour itself. This issue only exercises it.
