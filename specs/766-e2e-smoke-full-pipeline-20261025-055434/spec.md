# Feature Specification: E2E smoke full-pipeline marker (20261025-055434)

**Feature Branch**: `fabrik/issue-766`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261025-055434 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline. It checks that a single-repo issue can move from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR. The code change is deliberately trivial so that any failure points at the pipeline rather than the change. Each smoke run appends one uniquely timestamped marker line to a shared marker file. The file already exists and holds markers from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run's marker line (Priority: P1)

An operator running the smoke test wants the issue to produce a merged PR whose only effect is appending this run's unique marker line to the end of the smoke marker file.

**Why this priority**: This is the whole change. Landing it through the pipeline is the point of the test.

**Independent Test**: After merge, the last line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261025-055434 -->`. All earlier lines are unchanged.

**Acceptance Scenarios**:

1. **Given** the marker file exists with prior marker lines, **When** the change is applied, **Then** exactly one new line, `<!-- smoke-full-pipeline-20261025-055434 -->`, is appended at the very end and no existing line is modified or removed.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created containing exactly that one line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the issue advances automatically through all stages, the PR is merged, and the issue reaches Done.

---

### Edge Cases

- The marker file lacks a trailing newline: the new line must still start on its own line, not be joined to the previous one.
- The marker line already exists in the file (for example on a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- smoke-full-pipeline-20261025-055434 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: The change MUST create the file with that single line if it does not exist.
- **FR-003**: The change MUST NOT alter, reorder or remove any existing content of the file.
- **FR-004**: The change MUST NOT touch any file other than the marker file (apart from the engine-managed spec file).
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-006**: The PR MUST include `Closes #766` so the pipeline can discover and merge it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged, with no human intervention beyond the `fabrik:yolo` label.
- **SC-002**: The merged diff to the marker file is exactly one added line, with zero deletions.
- **SC-003**: No child issues are spawned for this issue.

## Assumptions

- The marker file already exists on `main`; the create-if-missing path is a fallback only.
- The marker text is exactly as given in the request, including the timestamp `20261025-055434`.
- The marker file is plain text with HTML-comment lines; the appended line follows that format.

## Out of Scope *(optional)*

- Any change beyond the single appended line.
- Cross-repo work.
- Cleaning up or deduplicating earlier marker lines.
- Changes to the pipeline, stage configs or tests.
