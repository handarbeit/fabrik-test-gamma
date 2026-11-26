# Feature Specification: E2E smoke full-pipeline marker (20261126-211854)

**Feature Branch**: `fabrik/issue-2903`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261126-211854 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline. Its purpose is to verify that a single-repo issue can travel from Specify through Research, Plan, Implement, Review and Validate to Done with a merged PR. The change itself is deliberately trivial so that any failure is attributable to the pipeline rather than the content of the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged on Validate completion.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the smoke suite wants a unique, timestamped marker line appended to the smoke marker file, delivered through a merged PR, to prove the full pipeline works end to end.

**Why this priority**: It is the only deliverable; the pipeline run is the test.

**Independent Test**: After the pipeline completes, the marker file on the default branch ends with the marker line, and the issue is in Done with its PR merged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is applied, **Then** the file's final line is `<!-- smoke-full-pipeline-20261126-211854 -->` and all prior lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done without human intervention.

---

### Edge Cases

- The marker file already exists (it does on the default branch today): the line is appended, not the file replaced.
- The existing file may lack a trailing newline: the new line must still appear on its own line.
- The marker line is already present (e.g. a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261126-211854 -->` after the change.
- **FR-002**: The change MUST be an append only; existing content of the file MUST remain unchanged.
- **FR-003**: If the file does not exist, it MUST be created containing the marker line.
- **FR-004**: The change MUST touch exactly one file and add exactly one line.
- **FR-005**: The work MUST be confined to this repository; the Plan stage MUST NOT decompose it into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff consists of exactly one added line in exactly one file (excluding the Specify stage's spec projection).
- **SC-002**: The issue reaches Done with its linked PR merged, with no human intervention beyond the initial labelling.
- **SC-003**: No child issues are spawned.

## Assumptions

- The marker file already exists on the default branch, so the change is an append.
- The `specs/` spec projection committed by the Specify stage is expected pipeline behaviour and is not counted as part of the "one file" change.

## Out of Scope *(optional)*

- Cross-repo work.
- Any change to engine behaviour, docs, or tests.
- Any change to other marker files.
