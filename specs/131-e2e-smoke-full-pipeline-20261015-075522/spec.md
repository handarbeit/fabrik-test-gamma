# Feature Specification: E2E smoke full-pipeline marker (20261015-075522)

**Feature Branch**: `fabrik/issue-131`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261015-075522 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The code change is deliberately trivial so that any failure points to the pipeline rather than to the change. Each smoke run appends a uniquely timestamped marker line to a shared marker file, and the file already holds lines from earlier runs.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the smoke marker line (Priority: P1)

An operator running the smoke suite needs this run's unique marker line to be present at the end of the shared marker file after the PR merges. That proves the pipeline carried the issue from spec to merged change.

**Why this priority**: It is the only deliverable. Without it the smoke run has no observable effect to verify.

**Independent Test**: After merge, read `e2e/markers/smoke-full-pipeline.md` on the default branch and check that its last line is exactly `<!-- smoke-full-pipeline-20261015-075522 -->`.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run lines, **When** the change is implemented, **Then** the file's final line is `<!-- smoke-full-pipeline-20261015-075522 -->` and all earlier content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and contains exactly that one line.
3. **Given** the change is merged via PR, **When** the issue reaches Done, **Then** the issue is closed and the PR is merged.

---

### Edge Cases

- The file already ends without a trailing newline: the new line must start on its own line, not be concatenated onto the previous one.
- The marker line is already present (for example on a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261015-075522 -->`.
- **FR-002**: If the file does not exist, it MUST be created before the line is added.
- **FR-003**: Existing content of the file MUST be preserved unchanged.
- **FR-004**: The change MUST touch only that one file and add only that one line.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals the marker line above.
- **SC-002**: The PR diff contains exactly one changed file and one added line.
- **SC-003**: The issue progresses from Specify through Validate to Done with its PR merged, with no human intervention beyond the `fabrik:yolo` label.

## Assumptions

- The marker file already exists, so no creation is expected. FR-002 covers the case where it does not.
- The marker line is an HTML comment, so it does not change how the file renders.
- The `fabrik:yolo` label is what lets the pipeline auto-advance and auto-merge.

## Out of Scope

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work and child-issue decomposition.
- Changes to the pipeline, the stage configs or the other marker files.
