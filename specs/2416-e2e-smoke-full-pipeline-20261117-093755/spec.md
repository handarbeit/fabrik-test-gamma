# Feature Specification: e2e smoke full-pipeline (20261117-093755)

**Feature Branch**: `fabrik/issue-2416`
**Created**: 2026-11-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261117-093755 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This is an end-to-end smoke test of the Fabrik pipeline in a single repository. The goal is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The change itself is deliberately trivial so that any failure points to the pipeline and not to the content of the change. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is merged automatically once Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik carries it through all stages unattended. The result is a merged PR that adds one marker comment line at the end of `e2e/markers/smoke-full-pipeline.md`, and the issue ends in Done.

**Why this priority**: It is the only story. It exercises the whole pipeline end to end.

**Independent Test**: After the run, check `main` for `e2e/markers/smoke-full-pipeline.md` with the marker line as its last line. Check that the issue is in Done and the PR is merged.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the line `<!-- smoke-full-pipeline-20261117-093755 -->` is appended at the very end and no existing content is modified.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue has `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done with no human intervention.

---

### Edge Cases

- If the file exists and its last line has no trailing newline, the marker must still land on its own line, not appended to the previous line's text.
- If the marker line is already present (for example on a re-run), it must not be duplicated.
- Plan must not decompose the work into child issues, because the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the exact line `<!-- smoke-full-pipeline-20261117-093755 -->`.
- **FR-002**: If the file does not exist, it MUST be created, containing the marker line.
- **FR-003**: Existing content of the file, if any, MUST remain unchanged; the marker is only appended.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: The work MUST be confined to this repository, and Plan MUST NOT decompose the issue into child issues.
- **FR-006**: The marker line MUST appear exactly once in the file.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with a merged PR and no manual action beyond the initial `fabrik:yolo` label.
- **SC-002**: The merged diff contains exactly one changed file and exactly one added line (plus the file creation if it did not exist).
- **SC-003**: The last line of `e2e/markers/smoke-full-pipeline.md` on `main` equals the marker string.
- **SC-004**: No child issues are spawned.

## Assumptions

- The marker directory `e2e/markers/` may or may not exist; creating it is part of creating the file.
- The file is a Markdown file, so an HTML comment line is a valid, invisible addition.
- The timestamp in the marker (`20261117-093755`) is a fixed value from the issue and is used as written.
- The `fabrik:yolo` label governs auto-advance and merge, so no human spec approval is expected.

## Out of Scope

- Cross-repo work.
- Any change to engine code, docs, tests or other files.
- Decomposition into child issues.
- Changes to the content of the marker file beyond the appended line.
