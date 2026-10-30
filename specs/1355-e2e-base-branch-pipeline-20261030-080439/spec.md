# Feature Specification: e2e base-branch pipeline (20261030-080439)

**Feature Branch**: `fabrik/issue-1355`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description:

> ## Goal
>
> End-to-end verification of the base:<branch> (non-default base branch)
> pipeline contract — regression coverage for handarbeit/fabrik#1046,
> validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
> fixes.
>
> ## Trivial change
>
> Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):
>
> ```
> <!-- base-branch-pipeline-20261030-080439 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261030-080439 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue is an end-to-end regression probe for the `base:<branch>` label contract. When an issue carries a `base:<branch>` label, the pipeline must fork its worktree from, target its PR at, and review and validate against that non-default base branch. The probe exists as regression coverage for handarbeit/fabrik#1046, and validates two fixes made for it:

- #1047: the issue↔PR linkage for PRs that target a non-default base branch.
- #1050: the data the review gate reads for such PRs.

The code change is intentionally trivial, so any failure in a run is attributable to the pipeline's handling of the base branch and not to the change itself. The item runs under `fabrik:cruise`, so it advances through the stages without human approval.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator files this issue with a `base:e2e-base-branch-20261030-080439` label. The pipeline works the issue, and the resulting change is a single appended marker line in `e2e/markers/base-branch-pipeline.md`. The PR targets the labelled base branch, not the repository's main branch.

**Why this priority**: This is the whole deliverable. It exercises the base-branch contract from worktree creation through PR creation and merge.

**Independent Test**: After the pipeline completes, inspect the PR diff against `e2e-base-branch-20261030-080439`. It must contain exactly one added line, the marker comment, at the end of `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is implemented, **Then** the file ends with the line `<!-- base-branch-pipeline-20261030-080439 -->` and no other content in the file is modified.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue has the label `base:e2e-base-branch-20261030-080439`, **When** the PR is created, **Then** its base branch is `e2e-base-branch-20261030-080439`.

---

### User Story 2 - Pipeline contract for non-default base holds (Priority: P1)

The pipeline's linkage and review-gate behaviour for a non-default-base PR is observable and correct (#1047, #1050).

**Why this priority**: This is the reason the probe exists. A trivial diff that merges but leaves broken linkage or review-gate data would defeat its purpose.

**Independent Test**: Observe the run from Implement through Validate. The issue is discovered as linked to its PR, and the review gate sees the PR's review data. Stage progression is not stalled by missing linkage or missing review data.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base branch with `Closes #1355` in its body, **When** the engine looks up the issue's linked PR, **Then** it resolves that PR despite GitHub creating no Development-panel link for non-default bases.
2. **Given** the PR has been marked ready, **When** the review gate evaluates it, **Then** it reads the PR's reviews and review requests and reaches a verdict instead of stalling on missing data.

---

### Edge Cases

- The marker file already exists on the base branch with content and a trailing newline. The new line is appended after the existing content and does not duplicate or alter prior lines.
- The marker file exists but lacks a trailing newline. The marker still ends up on its own line.
- A re-run after a partial failure must not append the marker twice. The file contains the marker line exactly once.
- The base branch is not the repository default. No step may silently fall back to the default branch.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- base-branch-pipeline-20261030-080439 -->` as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, the change MUST create it, with the marker line as its content.
- **FR-003**: The change MUST consist of exactly one file and one added line, with no other file or line modified.
- **FR-004**: The marker line MUST appear exactly once in the file, including after retries or re-runs.
- **FR-005**: The work MUST be based on, and the PR MUST target, the `e2e-base-branch-20261030-080439` branch.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: The issue's labels MUST NOT be added to or removed from by this work. The pre-set `fabrik:cruise` and `base:e2e-base-branch-20261030-080439` labels are left as filed.
- **FR-008**: The PR MUST carry `Closes #1355` so the engine can discover it as the issue's linked PR.

### Key Entities

- **Marker file** (`e2e/markers/base-branch-pipeline.md`): a text file that accumulates one HTML comment line per probe run.
- **Base branch** (`e2e-base-branch-20261030-080439`): the non-default branch that the work forks from and the PR targets.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR diff against `e2e-base-branch-20261030-080439` is one file and one added line, and that line is the marker.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261030-080439` for the whole life of the PR.
- **SC-003**: The issue moves through all pipeline stages to completion without a human intervening, and no stage stalls or pauses because the issue↔PR linkage or review-gate data was missing.
- **SC-004**: The marker line occurs exactly once in the final file.

## Assumptions

- The branch `e2e-base-branch-20261030-080439` already exists on the remote, because the label was applied at filing time.
- The marker line text is taken literally from the issue, including the timestamp suffix `20261030-080439`.
- The marker is appended even if an earlier run's marker, with a different timestamp, is already in the file.
- Without a human reviewer, the review gate is satisfied by whatever automated reviewers are configured, and this item does not depend on a particular reviewer.

## Out of Scope

- Any change to engine code, stage configs or documentation.
- Adding or removing labels, or changing the base branch.
- Multi-repo scenarios.
- Decomposing the work into child issues.
- Fixing #1046, #1047 or #1050. This item only verifies them.

## Source References

- handarbeit/fabrik#1046: the base-branch pipeline regression this probe covers.
- handarbeit/fabrik#1047: the issue↔PR linkage fix being validated.
- handarbeit/fabrik#1050: the review-gate data feed fix being validated.
