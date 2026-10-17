# Feature Specification: e2e base-branch pipeline (20261017-190519)

**Feature Branch**: `fabrik/issue-294`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261017-190519 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261017-190519 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This is an end-to-end test issue. It exercises the pipeline for an issue that carries a `base:<branch>` label, meaning its work targets a non-default base branch instead of the repository default. It is regression coverage for handarbeit/fabrik#1046. It validates two fixes:

- #1047: the issue↔PR linkage for PRs targeting a non-default base branch.
- #1050: the review-gate data feed for those PRs.

The code change is deliberately trivial, because the pipeline behavior is the subject under test, not the content. The issue carries the labels `fabrik:cruise` and `base:e2e-base-branch-20261017-190519`. These were applied at filing and must not be changed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An issue with a `base:<branch>` label moves through the pipeline. It produces a PR that adds one marker line to `e2e/markers/base-branch-pipeline.md`, targeting `e2e-base-branch-20261017-190519`.

**Why this priority**: This is the whole deliverable. It is the observable artifact that proves the non-default-base pipeline ran end to end.

**Independent Test**: After the pipeline completes, the file `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261017-190519` ends with the line `<!-- base-branch-pipeline-20261017-190519 -->`, and the diff against its previous content is exactly that one added line (or the whole file, if it was newly created).

**Acceptance Scenarios**:

1. **Given** the file already exists on the base branch, **When** the change is applied, **Then** the marker line is appended as the last line and no existing content is modified.
2. **Given** the file does not exist on the base branch, **When** the change is applied, **Then** the file is created containing the marker line.
3. **Given** the work is complete, **When** the PR is opened, **Then** the PR targets `e2e-base-branch-20261017-190519`, not the repository default branch.

---

### User Story 2 - Pipeline contract for non-default base is observable (Priority: P2)

The PR for this issue is linked to the issue and visible to the review gate, even though it targets a non-default base branch.

**Why this priority**: This is the regression the test exists to guard (#1047 and #1050). The marker line is only the vehicle.

**Independent Test**: Inspect the issue and its PR after the pipeline run. The engine discovers the PR as belonging to the issue, and review-gate evaluation sees the PR's reviews and review data.

**Acceptance Scenarios**:

1. **Given** the PR targets a non-default base, **When** the engine looks for the issue's linked PR, **Then** it finds the PR despite GitHub not creating a Development-panel link for non-default bases.
2. **Given** a Review stage with a review gate, **When** reviews arrive on the PR, **Then** the gate receives the PR's review data and evaluates it rather than waiting indefinitely on missing data.

---

### Edge Cases

- The marker file is missing at the base branch tip: it is created with the marker line.
- The marker file exists and does not end with a trailing newline: the marker still appears on its own final line.
- The marker line is already present: this should not happen on a fresh run. The timestamped marker is unique to this run.
- The base branch is missing on the remote: the engine falls back to the default and posts a comment. That fallback would be a failure of this e2e run, not an expected outcome.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change adds the exact line `<!-- base-branch-pipeline-20261017-190519 -->` as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created containing that line.
- **FR-003**: The change touches only that one file and adds only that one line, with no other file, line or edit.
- **FR-004**: The work is done on the issue's branch and delivered as a PR targeting `e2e-base-branch-20261017-190519`.
- **FR-005**: The existing labels `fabrik:cruise` and `base:e2e-base-branch-20261017-190519` are neither added nor removed by this work.
- **FR-006**: The work is not decomposed into child issues.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a plain Markdown file that only accumulates HTML-comment marker lines, one per e2e run.
- **Non-default base branch**: `e2e-base-branch-20261017-190519`, the branch this issue's PR targets, set by the `base:` label.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On `e2e-base-branch-20261017-190519`, the last line of `e2e/markers/base-branch-pipeline.md` is `<!-- base-branch-pipeline-20261017-190519 -->`.
- **SC-002**: The merged diff for this issue consists of exactly one file and one added line.
- **SC-003**: The issue's PR targets `e2e-base-branch-20261017-190519` and is associated with the issue so the pipeline advances it without manual linking.
- **SC-004**: The pipeline reaches completion through its stages without human intervention beyond the pre-applied `fabrik:cruise` label.

## Assumptions

- The marker file may or may not already exist. Either case is handled by FR-001 and FR-002.
- The repository's default base branch for this run is `e2e-base-branch-20261017-190519`, as stated in the pipeline context, so the `base:` label and the default coincide here. The PR target is still checked explicitly.
- Cruise mode means the PR is not auto-merged and the issue does not auto-advance to Done at Validate. Landing is out of this spec's control.
- The marker line is the exact text given in the request, with no added whitespace or annotations.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, tests or documentation.
- Adding or removing labels on the issue.
- Decomposing the work into sub-issues.
- Multi-repo scenarios.
- Fixing #1046, #1047 or #1050. They are only being verified.

## Source References *(optional)*

- handarbeit/fabrik#1046 (regression being covered), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix).
