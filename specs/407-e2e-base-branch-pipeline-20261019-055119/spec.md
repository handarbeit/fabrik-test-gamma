# Feature Specification: e2e base-branch pipeline marker (20261019-055119)

**Feature Branch**: `fabrik/issue-407`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line `<!-- base-branch-pipeline-20261019-055119 -->` to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist). That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261019-055119 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261019-055119"

## Background

This issue is a live end-to-end regression probe for the `base:<branch>` pipeline, which drives an issue through the pipeline against a non-default base branch. It is regression coverage for handarbeit/fabrik#1046, and it validates two fixes:

- #1047: issue↔PR linkage for PRs that target a non-default base.
- #1050: the data feed for the review gate on such PRs.

The code change is deliberately trivial so that the pipeline behaviour is the only thing under test. The issue carries `fabrik:cruise` and a `base:` label that were applied at filing. They are part of the test setup.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

A single marker line is appended to the marker file through the normal pipeline. The resulting PR targets the `base:` branch, not the repository default, and the issue and PR stay correctly linked and gated.

**Why this priority**: This is the entire purpose of the issue. It exercises the non-default-base contract end to end.

**Independent Test**: After the pipeline completes, inspect the merged PR's base branch and the contents of `e2e/markers/base-branch-pipeline.md` on that base.

**Acceptance Scenarios**:

1. **Given** the marker file exists on the base branch, **When** the change is applied, **Then** its last line is `<!-- base-branch-pipeline-20261019-055119 -->` and all prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line as its final line.
3. **Given** the issue has label `base:e2e-base-branch-20261019-055119`, **When** the Implement stage opens its PR, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker still goes on its own final line.
- The marker line is already present (for example on a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261019-055119 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created first, containing the marker line.
- **FR-003**: The change MUST touch only that one file and add only that one line (no other edits to the file or repo).
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261019-055119` as declared by the issue's `base:` label.
- **FR-005**: The issue's labels MUST NOT be added or removed as part of this work.
- **FR-006**: The work MUST NOT be decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a text file whose appended HTML comment lines record that a pipeline run reached the base branch.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The marker file on the base branch ends with the line `<!-- base-branch-pipeline-20261019-055119 -->` and the diff is exactly one added line (or one new file with that single line).
- **SC-002**: The PR for this issue targets `e2e-base-branch-20261019-055119` and is linked to issue #407.
- **SC-003**: The issue progresses through the pipeline with no manual intervention beyond the pre-applied labels.

## Assumptions

- The marker file may or may not already exist; both cases are handled by FR-001 and FR-002.
- A trailing newline after the marker line is acceptable and conventional.
- The branch `e2e-base-branch-20261019-055119` already exists on the remote.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configs or labels.
- Any other file or additional marker lines.
- Splitting this work into multiple issues.
