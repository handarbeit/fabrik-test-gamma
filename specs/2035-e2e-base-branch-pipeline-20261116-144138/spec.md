# Feature Specification: e2e base-branch pipeline marker (20261116-144138)

**Feature Branch**: `fabrik/issue-2035`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261116-144138 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261116-144138 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261116-144138"

## Background

This issue is a live end-to-end test fixture. It exercises the pipeline for an issue carrying a `base:<branch>` label, so that the work forks from, and the PR targets, a non-default base branch. It provides regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed).

The code change is deliberately minimal. The value is in the pipeline traversal, not the content. The test run is identified by the timestamp `20261116-144138`.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An e2e operator files an issue labelled `base:e2e-base-branch-20261116-144138` and `fabrik:cruise`. The pipeline runs it through every stage. The result is a PR, targeting that base branch, that adds one marker line to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the only deliverable. It is the artifact the e2e run uses to confirm the base-branch pipeline worked end to end.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261116-144138`. It should show exactly one added line, `<!-- base-branch-pipeline-20261116-144138 -->`, at the end of `e2e/markers/base-branch-pipeline.md`, and no other changes.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** Implement completes, **Then** the file's final line is `<!-- base-branch-pipeline-20261116-144138 -->` and earlier content is unchanged.
2. **Given** the file does not exist on the base branch, **When** Implement completes, **Then** the file is created containing the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261116-144138`, **When** the PR is created, **Then** it targets that branch rather than the repository's other branches, and is linked to this issue.

---

### Edge Cases

- If the marker file already ends without a trailing newline, the marker is still placed on its own final line.
- If the marker line is already present (for example on a re-run), it is not duplicated.
- If the base branch is missing on the remote, the pipeline's existing fallback behaviour applies. That behaviour is not changed here.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261116-144138 -->`, at the end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST touch no file other than `e2e/markers/base-branch-pipeline.md` (the engine-written spec file for this issue aside).
- **FR-004**: The PR MUST target the base branch `e2e-base-branch-20261116-144138`, and the work MUST fork from that branch.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The issue's labels (`fabrik:cruise`, `base:e2e-base-branch-20261116-144138`) MUST NOT be added to, removed or altered by the work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261116-144138` contains exactly one added line in exactly one file.
- **SC-002**: The PR is linked to this issue and targets the non-default base branch.
- **SC-003**: The issue passes through all stages to completion under `fabrik:cruise` without manual intervention or a pause.

## Assumptions

- The base branch `e2e-base-branch-20261116-144138` exists on the remote, as the filing process sets it up.
- The `e2e/markers/` directory may need to be created along with the file if it is absent.
- The marker is a single HTML comment line so it is invisible when the markdown is rendered.

## Out of Scope

- Any change to engine, pipeline or test code.
- Any change to labels.
- Splitting the work into multiple issues or PRs.
- Marker files for other e2e scenarios.
