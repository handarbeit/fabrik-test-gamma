# Feature Specification: E2E smoke full-pipeline (20261116-151317)

**Feature Branch**: `fabrik/issue-2068`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261116-151317 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is a release-gate style smoke test of the Fabrik pipeline. It checks that a minimal issue can pass through every stage (Specify → Research → Plan → Implement → Review → Validate) and end at Done with a merged PR, with `fabrik:yolo` driving automatic advancement and merge. The change is deliberately trivial so that any failure points at the pipeline and not at the content of the change. Each smoke run adds one uniquely timestamped marker line to a shared marker file, which proves that this run's change, and not an earlier run's, reached the base branch.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator runs the smoke issue through the pipeline with `fabrik:yolo` set. At the end, the marker file on the default branch ends with this run's unique marker line, and the issue is Done.

**Why this priority**: It is the only story. It exercises the whole pipeline end to end.

**Independent Test**: After the issue reaches Done, read `e2e/markers/smoke-full-pipeline.md` on the default branch and confirm that its last line is `<!-- smoke-full-pipeline-20261116-151317 -->`. Confirm that the merged PR changes only that file by adding that one line.

**Acceptance Scenarios**:

1. **Given** the marker file already exists with earlier smoke lines, **When** the pipeline implements this issue, **Then** the file keeps all of its existing content and gains exactly one new final line, `<!-- smoke-full-pipeline-20261116-151317 -->`.
2. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances automatically through Research, Plan, Implement, Review and Validate, the linked PR is merged, and the issue ends in Done.
3. **Given** the work is a single-file, single-line change in one repository, **When** the Plan stage runs, **Then** it produces one implementation unit and spawns no child issues.

---

### Edge Cases

- The marker file does not exist at implementation time: it is created containing only the marker line.
- The existing file lacks a trailing newline: the marker line must still be on its own line and must not be joined to the previous line.
- The marker line is already present because of a retry or resume: it is not added a second time.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261116-151317 -->` after this issue's change is merged.
- **FR-002**: The change MUST append the line at the very end of the file and MUST NOT modify, reorder or remove any existing content.
- **FR-003**: If the file does not exist, the change MUST create it containing the marker line.
- **FR-004**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md` and add exactly one line.
- **FR-005**: The work MUST stay in this single repository, and the Plan stage MUST NOT decompose it or spawn child issues.
- **FR-006**: The linked PR MUST include `Closes #2068` so that the pipeline can link the PR to this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged into the default branch, with no human intervention beyond the initial `fabrik:yolo` label.
- **SC-002**: The merged PR's diff is exactly one added line in exactly one file.
- **SC-003**: The last line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals `<!-- smoke-full-pipeline-20261116-151317 -->`.
- **SC-004**: No child issues are created.

## Assumptions

- The default base branch is `main`.
- The marker file already exists, so the issue's "create it if missing" clause is the fallback and not the expected path.
- An HTML comment line in a Markdown file has no behavioral effect on any other part of the system.
- No documentation changes are needed, because the marker file is test scaffolding.

## Out of Scope

- Any change other than the single marker line.
- Cross-repository work or decomposition into child issues.
- Cleanup or deduplication of earlier smoke marker lines.
- Changes to pipeline behavior itself.
