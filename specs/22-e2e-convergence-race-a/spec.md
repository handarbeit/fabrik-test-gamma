# Feature Specification: e2e base-branch pipeline (20261014-192638)

**Feature Branch**: `fabrik/issue-22`
**Created**: 2026-10-14
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
> `<!-- base-branch-pipeline-20261014-192638 -->`
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261014-192638 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue exists to exercise the pipeline for an item that carries a `base:<branch>` label, meaning its work is branched from and merged into a non-default base branch. It provides end-to-end regression coverage for handarbeit/fabrik#1046. It checks that two fixes hold for such items:

- **#1047**: the issue and its pull request stay linked.
- **#1050**: the review gate receives the correct review data.

The code change is deliberately trivial: one marker line in one file. It is a vehicle for exercising the pipeline, and the change itself has no product value. Any failure should come from the pipeline's handling of the non-default base, not from the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the base-branch e2e check files this issue with a `base:e2e-base-branch-20261014-192638` label. The pipeline carries it from Specify through Validate, and the resulting change lands on that base branch rather than the repository's default.

**Why this priority**: This is the only deliverable. It proves the end-to-end contract for non-default base branches.

**Independent Test**: After the pipeline finishes, inspect `e2e/markers/base-branch-pipeline.md` on branch `e2e-base-branch-20261014-192638`. The last line should be the marker comment, and the PR diff should contain only that one-line addition (plus the file itself if it was newly created).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the Implement stage completes, **Then** the file's final line is `<!-- base-branch-pipeline-20261014-192638 -->` and all earlier content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the Implement stage completes, **Then** the file exists and its sole content is `<!-- base-branch-pipeline-20261014-192638 -->`.
3. **Given** the issue carries the `base:e2e-base-branch-20261014-192638` label, **When** the PR is created, **Then** the PR targets that branch and not the repository default, if the two differ.

---

### User Story 2 - Issue and PR linkage and review-gate feed work on the non-default base (Priority: P2)

The operator confirms that the pipeline recognizes the PR as belonging to the issue and that the review gate sees its reviews, even though the PR targets a non-default base.

**Why this priority**: These are the #1047 and #1050 behaviors under test. They are secondary to the change landing, but they are the reason the issue exists.

**Independent Test**: Check that the pipeline discovers the PR as linked to issue #22. Check that the review gate evaluates that PR's review data, and that the item advances through the review and validate stages without a pause caused by missing linkage or review data.

**Acceptance Scenarios**:

1. **Given** a PR targeting the non-default base with `Closes #22` in its body, **When** the pipeline looks up the issue's linked PR, **Then** the PR is found.
2. **Given** that PR has reviews, **When** the review gate runs, **Then** it evaluates those reviews and does not treat the PR as having none.

---

### Edge Cases

- The marker file already exists with no trailing newline: the marker line must still end up on its own line, with earlier content preserved.
- The marker file already contains the exact marker line (for example after a retry): the line must not be duplicated.
- The `e2e/markers/` directory does not exist: it is created along with the file.
- The base branch is the same as the repository default: the change is identical, only the targeting check is moot.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261014-192638 -->`, as the final line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created, containing only that one line.
- **FR-003**: Existing content of the file MUST NOT be modified, removed or reordered.
- **FR-004**: The change MUST NOT touch any other file in the resulting PR (the spec file the engine commits is excluded).
- **FR-005**: The work MUST be delivered as a single unit and not decomposed into child issues.
- **FR-006**: The PR MUST target the base branch named by the issue's `base:` label, `e2e-base-branch-20261014-192638`.
- **FR-007**: The issue's labels MUST NOT be added to or removed by this work.
- **FR-008**: The PR body MUST include `Closes #22` so the pipeline links the PR to the issue.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a text file whose lines are HTML comments, one per e2e run.
- **Marker line**: `<!-- base-branch-pipeline-20261014-192638 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff against its base contains exactly one added line (the marker line) in exactly one file.
- **SC-002**: The marker line is the last line of the marker file on the base branch after landing.
- **SC-003**: The issue progresses through all pipeline stages without a pause attributable to missing issue–PR linkage or missing review data.
- **SC-004**: The PR's base branch is `e2e-base-branch-20261014-192638`.

## Assumptions

- The marker file may or may not already exist on the base branch. FR-001 and FR-002 cover both cases.
- The marker line is appended with a newline-terminated line (the file ends with a newline after it).
- The `base:` label already present on the issue is correct and stays as filed.
- "Plan should NOT decompose" means no child issues are spawned, and the whole change is a single implementation step.

## Out of Scope *(optional)*

- Any change to engine, pipeline or stage configuration.
- Any change to other marker files or other repositories.
- Adding, removing or changing labels on the issue.
- Fixing #1046, #1047 or #1050 themselves. This issue only verifies them.

## Source References *(optional)*

- handarbeit/fabrik#1046 — base-branch pipeline regression being covered
- handarbeit/fabrik#1047 — issue↔PR linkage fix
- handarbeit/fabrik#1050 — review-gate data feed fix
