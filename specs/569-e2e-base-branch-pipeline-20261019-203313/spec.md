# Feature Specification: e2e base-branch pipeline (20261019-203313)

**Feature Branch**: `fabrik/issue-569`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261019-203313 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261019-203313 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be worked on, and have its pull request targeted at, that non-default base branch. It is regression coverage for handarbeit/fabrik#1046 and validates two fixes: #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The code change is deliberately trivial, a single marker line in one file, so that any failure points at the pipeline mechanics and not at the work itself. The marker line carries a run-unique timestamp (`20261019-203313`) so this run can be told apart from earlier ones.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e check files this issue with the `base:e2e-base-branch-20261019-203313` label. The pipeline appends one marker line to `e2e/markers/base-branch-pipeline.md` and delivers it through a pull request that targets that base branch, not the repository's usual default.

**Why this priority**: This is the core of the base-branch contract. If the change does not reach the labelled base branch, the check has failed.

**Independent Test**: After the run, inspect `e2e/markers/base-branch-pipeline.md` on the `e2e-base-branch-20261019-203313` branch. The last line must be exactly `<!-- base-branch-pipeline-20261019-203313 -->` and nothing else must have changed.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` exists, **When** the change is made, **Then** exactly one line, `<!-- base-branch-pipeline-20261019-203313 -->`, is appended at the very end and no existing line is altered.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created containing that single line.
3. **Given** the issue carries the `base:e2e-base-branch-20261019-203313` label, **When** the pull request is created, **Then** its base is `e2e-base-branch-20261019-203313`.

---

### User Story 2 - Issue and PR remain linked across the pipeline (Priority: P1)

The pull request created for this issue is discoverable from the issue (and the reverse) even though it targets a non-default base. This validates #1047.

**Why this priority**: Linkage is what lets the pipeline find the PR, post stage output and react to PR feedback. Without it the non-default-base flow cannot complete.

**Independent Test**: Check that the pipeline recognises the PR as belonging to issue #569 and can reach later stages and Done without manual intervention.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base that references this issue, **When** the pipeline looks up the issue's PR, **Then** it finds that PR.
2. **Given** the PR has merged to the non-default base, **When** the pipeline finishes, **Then** the issue ends up closed and in Done.

---

### User Story 3 - Review gate receives review data for the base-branch PR (Priority: P2)

Where a stage waits on reviews, the review gate gets the PR's review data and does not stall or skip because the PR's base is non-default. This validates #1050.

**Why this priority**: It is an important part of the contract, but it only matters once the PR exists and is linked (Stories 1 and 2).

**Independent Test**: Observe that the review and landing gates evaluate this PR's reviews and either clear or block on real review state, and never hang on missing data.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base, **When** a review-waiting stage evaluates the gate, **Then** the gate sees the PR's actual reviewers and review state.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line is still appended as its own final line, with no blank line between it and the previous content unless one was already there.
- The marker file already ends with this exact marker (a re-run): the line is not duplicated.
- The base branch does not exist on the remote: this is a harness precondition failure and is outside the scope of this change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261019-203313 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, the change MUST create it with that line as its only content.
- **FR-003**: The change MUST NOT modify, reorder or remove any other line in that file, and MUST NOT touch any other file other than the engine-managed spec file.
- **FR-004**: The work MUST be delivered on a branch and PR whose base is `e2e-base-branch-20261019-203313`, as set by the `base:` label.
- **FR-005**: The change MUST NOT add or remove any labels on the issue, including the `base:` and `fabrik:cruise` labels.
- **FR-006**: The work MUST remain a single task: no decomposition into child issues.
- **FR-007**: The PR MUST be linked to issue #569 so that the pipeline can discover it and close the issue on completion.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a plain text file used only as a trivially checkable payload for the e2e run.
- **Base branch**: `e2e-base-branch-20261019-203313`, the non-default branch the work forks from and targets.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On `e2e-base-branch-20261019-203313`, the marker file's last line equals `<!-- base-branch-pipeline-20261019-203313 -->`, and the diff against the previous tip of that branch is exactly one added line in that file (apart from the spec file).
- **SC-002**: The PR for this issue has `e2e-base-branch-20261019-203313` as its base, and no PR for it targets the repository default.
- **SC-003**: The pipeline carries the issue from Specify to Done with no manual intervention and no pause caused by missing PR linkage or review data.
- **SC-004**: The issue's label set at the end of the run is unchanged by the change itself, apart from the engine's own state labels.

## Assumptions

- The branch `e2e-base-branch-20261019-203313` already exists on the remote and is the target of the `base:` label, as stated in the issue.
- The marker line is written exactly as given, with no surrounding whitespace, and is followed by a trailing newline per normal text-file convention.
- Because the issue carries `fabrik:cruise`, stages advance automatically, but the PR is not auto-merged and the issue does not move to Done at Validate completion, per cruise semantics. The ultimate landing of the PR is outside this change's concern.
- The engine writes the spec file under `specs/569-*/spec.md` as part of this stage; that file is a projection of this issue body and is not part of the "one file, one line" payload.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration or skills.
- Any change to the content of the marker file other than the one appended line.
- Creating or managing the base branch, or changing labels.
- Decomposing the work into multiple issues.

## Source References *(optional)*

- handarbeit/fabrik#1046: base-branch pipeline contract regression
- handarbeit/fabrik#1047: issue↔PR linkage fix for non-default-base PRs
- handarbeit/fabrik#1050: review-gate data feed fix
