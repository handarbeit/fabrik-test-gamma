# Feature Specification: E2E smoke full-pipeline marker (20261019-122342)

**Feature Branch**: `fabrik/issue-468`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261019-122342 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and finish in Done with a merged pull request. The code change is deliberately trivial so the pipeline itself is the thing under test. The issue carries the `fabrik:yolo` label, so stages auto-advance and the PR is auto-merged at the end of Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged via the full pipeline (Priority: P1)

An operator running the smoke test wants a unique, run-specific marker line added to the smoke-test marker file, delivered through the normal pipeline, so that a merged PR proves the pipeline works end to end.

**Why this priority**: This is the only deliverable. A merged PR containing exactly this change is the success signal for the whole smoke run.

**Independent Test**: After the PR merges, inspect the default branch: the marker file exists and its last line is the marker comment for this run.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's original content is unchanged and the line `<!-- smoke-full-pipeline-20261019-122342 -->` is appended at the very end.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the line `<!-- smoke-full-pipeline-20261019-122342 -->`.
3. **Given** the change is implemented, **When** the PR is opened, **Then** it modifies only that one file, adding only that one line, and links to this issue with `Closes #468`.
4. **Given** the `fabrik:yolo` label is present, **When** each stage completes, **Then** the issue advances automatically through Research, Plan, Implement, Review and Validate, the PR is merged, and the issue reaches Done.

---

### Edge Cases

- The file exists but its last line lacks a trailing newline: the marker must still land on its own line, not be joined to the previous content.
- The marker line is already present (for example on a retried run): it must not be added a second time.
- The file exists with other earlier markers: they are preserved untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/smoke-full-pipeline.md` MUST contain the line `<!-- smoke-full-pipeline-20261019-122342 -->` as its final line after the change.
- **FR-002**: If the file does not exist, it MUST be created (including any missing parent directory) with that line as its content.
- **FR-003**: If the file exists, existing content MUST be preserved unchanged and the marker line MUST be appended at the end, on its own line.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch no file other than `e2e/markers/smoke-full-pipeline.md` and add no line other than the marker line.
- **FR-006**: The work MUST stay within this single repository; the Plan stage MUST NOT decompose the issue or spawn child issues.
- **FR-007**: The PR MUST carry `Closes #468` so Fabrik can link it to this issue.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with its linked PR merged into the default branch, with no manual intervention beyond the initial labelling.
- **SC-002**: The merged PR's diff consists of exactly one file and exactly one added line (or one new file with that single line).
- **SC-003**: On the default branch after merge, the marker comment occurs exactly once and is the last line of `e2e/markers/smoke-full-pipeline.md`.

## Assumptions

- The default base branch is `main`, and the PR targets it.
- The `fabrik:yolo` label is already on the issue and is not changed by this work.
- The marker file contains only HTML comment lines meant for smoke tests, so appending a comment line has no functional effect.
- If a trailing newline is missing from the existing file, adding one before the marker is acceptable and is not counted as a second changed line.

## Out of Scope

- Any change to Fabrik engine code, stage configs, or documentation.
- Any cross-repo work, and any decomposition into child issues.
- Cleanup or reorganisation of other files in `e2e/markers/`.
