# Feature Specification: E2E smoke full-pipeline marker (20261019-125645)

**Feature Branch**: `fabrik/issue-478`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-125645 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. It checks that an issue can move through Specify, Research, Plan, Implement, Review and Validate, and end with a merged PR and the issue in Done. The code change is deliberately tiny so that any failure points at the pipeline, not the change. The marker file `e2e/markers/smoke-full-pipeline.md` already exists and holds one HTML-comment line per earlier smoke run. This run adds its own line to that history. The issue carries `fabrik:yolo`, so stages auto-advance and the PR auto-merges after Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append this run's marker line (Priority: P1)

An operator running the smoke test needs the marker file to end with a line unique to this run. That line is the evidence that the change travelled through the whole pipeline and was merged.

**Why this priority**: It is the only change in the issue, and the pass/fail signal for the smoke test.

**Independent Test**: After merge, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its last line is `<!-- smoke-full-pipeline-20261019-125645 -->`, and every earlier line is unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run lines, **When** the change is applied, **Then** the file ends with exactly one new line, `<!-- smoke-full-pipeline-20261019-125645 -->`, and the earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing that single line.
3. **Given** the PR for the change is open and linked to the issue, **When** the pipeline completes, **Then** the PR is merged and the issue is in Done.

---

### Edge Cases

- The file already ends without a trailing newline: the new line must start on its own line and must not be joined to the previous one.
- The marker line is already present (for example after a retry): it must not be added a second time.
- Other marker files in `e2e/markers/` must not be touched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261019-125645 -->`, appended at the very end of the file.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST add exactly one line to exactly one file, and MUST NOT modify any existing line or any other file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST stay within this single repository. The Plan stage MUST NOT decompose it or spawn child issues.
- **FR-006**: The PR MUST reference this issue with `Closes #478` so that Fabrik can discover it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the PR merges, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is `<!-- smoke-full-pipeline-20261019-125645 -->`.
- **SC-002**: The merged diff touches exactly one file and adds exactly one line.
- **SC-003**: The issue reaches Done with its PR merged, without manual intervention beyond applying the `fabrik:yolo` label.

## Assumptions

- The marker file exists on the default branch. The issue says to create it if missing, and that case is covered by FR-002.
- The content is a plain HTML comment and has no functional effect.
- Existing content is not reordered, deduplicated or cleaned up.

## Out of Scope

- Cross-repo work.
- Decomposition into child issues.
- Changes to any other marker file or to engine code, docs or tests.
- Cleaning up old marker lines.
- Any behavior change in the product.
- Pipeline fixes found during the run; these belong in separate issues.
