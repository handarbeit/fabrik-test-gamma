# Feature Specification: e2e smoke full-pipeline (20261017-161327)

**Feature Branch**: `fabrik/issue-215`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261017-161327 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to check that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points to the pipeline and not to the work: one comment line is appended to a marker file. The issue carries `fabrik:yolo`, so stages advance automatically and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the default branch via the full pipeline (Priority: P1)

An operator running the e2e smoke suite files this issue. Fabrik takes it through all pipeline stages without human intervention. The result is a merged PR whose only effect is one new comment line at the end of `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: This is the whole purpose of the issue. It checks that the pipeline works from Specify to Done.

**Independent Test**: After the pipeline finishes, read `e2e/markers/smoke-full-pipeline.md` on the default branch (`main`). Its last line must be `<!-- smoke-full-pipeline-20261017-161327 -->`. The issue must be closed and in Done, and the PR merged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the existing content is unchanged and `<!-- smoke-full-pipeline-20261017-161327 -->` is the final line of the file.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains `<!-- smoke-full-pipeline-20261017-161327 -->` as its final line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the issue advances to Done through the pipeline and its linked PR is merged.
4. **Given** the issue is single-repo and one-line in scope, **When** the Plan stage runs, **Then** the plan does not decompose the work into child issues.

---

### Edge Cases

- **Marker file lacks a trailing newline**: the new line must still start on its own line, not be joined to the previous content.
- **Marker file is empty**: the file then contains only the marker line.
- **Marker line already present** (for example after a retry): it must not be added twice.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The default branch must end up with `e2e/markers/smoke-full-pipeline.md` containing `<!-- smoke-full-pipeline-20261017-161327 -->` as its last line.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, it must be created. If it exists, its prior content must be preserved and the marker line appended at the very end.
- **FR-003**: The marker line must appear exactly once in the file and be written exactly as `<!-- smoke-full-pipeline-20261017-161327 -->`.
- **FR-004**: The change must touch only `e2e/markers/smoke-full-pipeline.md`.
- **FR-005**: The work must stay within this one repository, and the Plan stage must not decompose it into child issues.
- **FR-006**: The issue must pass through the pipeline to Done with its PR merged, without human intervention beyond the `fabrik:yolo` label already applied.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a file whose only role is to carry smoke-run marker comments.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261017-161327 -->`, unique to this smoke run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the marker file on the default branch ends with the marker line, and that line appears exactly once.
- **SC-002**: The diff of the merged PR contains one changed file and, for an existing file, one added line.
- **SC-003**: The issue reaches Done and is closed, and its linked PR is merged, with every pipeline stage completed.
- **SC-004**: No child issues are created from this issue.

## Assumptions

- The marker file's current content, if any, is irrelevant to this smoke test and only needs to be preserved.
- If the marker file is created, it needs no content other than the marker line.
- Appending the marker line to an existing file means adding it as a new line after any existing content.
- `fabrik:yolo` is already applied, so no human approval is expected at any stage.

## Out of Scope

- Any change other than appending the one marker line to the one marker file.
- Cross-repo work or decomposition into child issues.
- Changes to engine behaviour, stage configuration or documentation.
