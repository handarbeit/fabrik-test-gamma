# Feature Specification: E2E smoke full-pipeline marker (20261031-023827)

**Feature Branch**: `fabrik/issue-1507`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261031-023827 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the single-repo Fabrik pipeline. It checks that an issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR, with the smallest possible code change. The change itself is deliberately trivial so that any failure points at the pipeline, not the work. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged once Validate completes.

The marker file `e2e/markers/smoke-full-pipeline.md` already exists. Earlier smoke runs each appended one timestamped HTML comment line to it, and the last line is currently `<!-- smoke-full-pipeline-20261030-202609 -->`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the pipeline smoke wants a one-line change that is unique to this run, so a merged PR proves the whole pipeline completed.

**Why this priority**: It is the only change in the issue; the smoke test has no other purpose.

**Independent Test**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` is `<!-- smoke-full-pipeline-20261031-023827 -->`, and the PR diff touches only that file and adds only that line.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists with prior marker lines, **When** the change is implemented, **Then** the exact line `<!-- smoke-full-pipeline-20261031-023827 -->` is appended after the existing last line and no existing line is modified or removed.
2. **Given** the PR carries `Closes #1507` and the issue has `fabrik:yolo`, **When** Validate completes, **Then** the PR is merged and the issue reaches Done.

---

### Edge Cases

- The existing file does not end with a trailing newline: the new marker must still start on its own line, not be joined to the previous line.
- The file is absent on the base branch: it is created containing only the marker line.
- The marker line is already present (for example on a re-run): it is not added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- smoke-full-pipeline-20261031-023827 -->` MUST be appended as the final line of `e2e/markers/smoke-full-pipeline.md`, exactly as written, with no leading or trailing characters on that line.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it MUST be created containing that single line.
- **FR-003**: Existing lines in the file MUST be left unchanged.
- **FR-004**: The change MUST consist of that one file and that one added line, and no other file is modified by the implementation.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue passes through every pipeline stage to Done without human intervention beyond the initial `fabrik:yolo` label.
- **SC-002**: A PR linked to this issue is merged, and its diff is exactly one added line in one file.
- **SC-003**: On the base branch after the merge, `e2e/markers/smoke-full-pipeline.md` ends with `<!-- smoke-full-pipeline-20261031-023827 -->` and contains that line exactly once.

## Assumptions

- The marker file already exists on `main`, as it does today; the "create if missing" instruction is a fallback only.
- The HTML comment line is inert, so it does not affect any behavior or test.
- Under `fabrik:yolo` no human approval is required between stages.

## Out of Scope

- Any cross-repo work or child-issue spawning.
- Changes to any file other than `e2e/markers/smoke-full-pipeline.md`.
- Edits to, or cleanup of, earlier marker lines in the file.
- Changes to Fabrik engine code, stage configuration or documentation.
