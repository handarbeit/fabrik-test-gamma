# Feature Specification: e2e base-branch pipeline marker (20261108-034143)

**Feature Branch**: `fabrik/issue-1899`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261108-034143 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: single repo only, targeting the non-default base branch e2e-base-branch-20261108-034143 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261108-034143"

## Background

This issue is an automated end-to-end probe of the pipeline's handling of issues that carry a `base:<branch>` label, meaning the work targets a branch other than the repository default. It is regression coverage for handarbeit/fabrik#1046. The probe checks two earlier fixes:

- #1047: the issue and its pull request stay linked when the PR targets a non-default base.
- #1050: the review gate receives the data it needs when the PR targets a non-default base.

The code change is deliberately trivial so the result depends only on how the pipeline behaves, not on the content of the change. The issue was filed with `fabrik:cruise` and a `base:` label. Those labels are part of the test setup and must not be altered.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite files this issue and lets it run through the pipeline unattended. The result is a pull request that adds exactly one marker line to the marker file and targets the non-default base branch.

**Why this priority**: This is the whole deliverable. The rest of the pipeline's behavior is observed through this change.

**Independent Test**: After the pipeline finishes, inspect the PR diff. It contains one added line, the marker, at the end of `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the change is made, **Then** the file is created and its last line is `<!-- base-branch-pipeline-20261108-034143 -->`.
2. **Given** the marker file already exists, **When** the change is made, **Then** the marker line is appended after all existing content and the existing content is unchanged.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** exactly one file is changed and exactly one line is added.

---

### User Story 2 - Pipeline honors the non-default base (Priority: P1)

The pipeline treats `e2e-base-branch-20261108-034143` as the base for this issue throughout, which is what the #1047 and #1050 regression checks verify.

**Why this priority**: This is the actual purpose of the e2e run. A correct file change on the wrong base would still fail the test.

**Independent Test**: Check the PR's base branch, the issue-to-PR linkage, and that the review gate sees the PR's review data.

**Acceptance Scenarios**:

1. **Given** the issue carries the `base:e2e-base-branch-20261108-034143` label, **When** the PR is created, **Then** its base branch is `e2e-base-branch-20261108-034143`.
2. **Given** the PR targets a non-default base, **When** the pipeline looks up the issue's linked PR, **Then** it finds that PR (#1047 behavior).
3. **Given** the PR is open against the non-default base, **When** the review gate evaluates it, **Then** it receives the PR's review data (#1050 behavior).

---

### Edge Cases

- The marker file does not exist yet: it is created, containing only the marker line.
- The existing file does not end in a newline: the marker still goes on its own line at the end.
- The marker line is already present, for example after a retried run: it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the single line `<!-- base-branch-pipeline-20261108-034143 -->` at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it, and the marker line MUST be its last line.
- **FR-003**: The change MUST NOT modify any other file or any existing content of the marker file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST target the base branch `e2e-base-branch-20261108-034143`.
- **FR-006**: The work MUST NOT add or remove any labels on the issue, including `fabrik:cruise` and `base:e2e-base-branch-20261108-034143`.
- **FR-007**: The work MUST be treated as a single unit. It MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The resulting PR's diff changes exactly one file and adds exactly one line, and that line is the marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261108-034143`.
- **SC-003**: The issue is linked to its PR, and the pipeline reaches the review gate with the PR's review data available, without manual intervention.
- **SC-004**: The issue's labels at the end of the run are the same as at the start, apart from the engine's own state labels.

## Assumptions

- The marker file is a Markdown file under `e2e/markers/` in the target repository, and an HTML comment line is invisible when rendered.
- The base branch `e2e-base-branch-20261108-034143` already exists on the remote, since the filing process set it up.
- "Entire change" means no tests, docs, or other edits are wanted beyond the marker line.
- Engine-managed state labels such as `stage:*` and `fabrik:locked:*` are expected to change during the run. The "do not add or remove labels" rule refers to the user-facing configuration labels.

## Out of Scope *(optional)*

- Any change other than the one marker line.
- Decomposing the work into sub-issues.
- Changing pipeline or engine behavior. This issue only observes it.
- Adding or removing labels by hand.

## Source References *(optional)*

- handarbeit/fabrik#1046: regression being covered
- handarbeit/fabrik#1047: issue↔PR linkage fix under test
- handarbeit/fabrik#1050: review-gate data feed fix under test
