# Feature Specification: e2e base-branch pipeline marker (20261126-232752)

**Feature Branch**: `fabrik/issue-2981`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261126-232752 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261126-232752 (already set via the base:<branch> label applied at filing — do not add or remove labels). Labels: fabrik:cruise, base:e2e-base-branch-20261126-232752"

## Background

This issue is a live end-to-end test fixture. Its purpose is to push a deliberately minimal change through the whole Fabrik pipeline on an item carrying a `base:<branch>` label, so that the non-default base-branch contract is exercised for real. It is regression coverage for #1046 and validates two fixes: #1047 (issue-to-PR linkage when the PR targets a non-default base) and #1050 (the data feed to the review gate for such PRs).

The change itself is intentionally trivial so that any failure points at the pipeline and not at the content. The marker line embeds a run-unique timestamp (`20261126-232752`) so this run's change can be told apart from other runs'.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An operator running the e2e suite files this issue with a `base:e2e-base-branch-20261126-232752` label. The pipeline processes it and produces a PR, targeting that base branch, that appends the run-specific marker line to `e2e/markers/base-branch-pipeline.md`.

**Why this priority**: This is the whole change. Without it there is nothing for the pipeline to carry through the non-default-base path.

**Independent Test**: Inspect the PR diff against `e2e-base-branch-20261126-232752`. It touches exactly one file, `e2e/markers/base-branch-pipeline.md`, and adds exactly one line, which is the final line of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file keeps its prior content and ends with `<!-- base-branch-pipeline-20261126-232752 -->` as its last line.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and its content is that single marker line.
3. **Given** the item carries the `base:e2e-base-branch-20261126-232752` label, **When** the PR is created, **Then** the PR targets `e2e-base-branch-20261126-232752` and not any other branch.

---

### User Story 2 - Pipeline contract for non-default base holds (Priority: P1)

The same run verifies the behaviours #1047 and #1050 fixed, using this item as the subject.

**Why this priority**: This is the reason the issue exists. The one-line edit is only the vehicle for it.

**Independent Test**: Observe the run from the PR to the review gate, and check the issue-to-PR linkage and the review-gate inputs.

**Acceptance Scenarios**:

1. **Given** a PR on a non-default base that contains `Closes #2981`, **When** the pipeline discovers the PR for the issue, **Then** the issue and PR are linked so that stage comments, review feedback and the review gate see the PR.
2. **Given** the PR is open on the non-default base, **When** the review gate evaluates it, **Then** it receives that PR's review data and does not stall or skip for lack of linkage.

---

### Edge Cases

- The marker file already exists with no trailing newline: the marker must still land on its own final line and must not be joined to the previous content.
- The marker file already contains the same marker line from an earlier attempt: it must not be added twice (one line, one occurrence).
- The base branch is missing on the remote: the engine's existing fallback applies. That behaviour is under test elsewhere and is not part of this change.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- base-branch-pipeline-20261126-232752 -->`, as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created first, and the marker line is then its only content.
- **FR-003**: No file other than `e2e/markers/base-branch-pipeline.md` is added, modified or removed.
- **FR-004**: Existing content of `e2e/markers/base-branch-pipeline.md` is preserved unchanged ahead of the new line.
- **FR-005**: The PR targets `e2e-base-branch-20261126-232752`, the branch named by the item's `base:` label.
- **FR-006**: The work is carried out as a single unit, with no child issues spawned and no decomposition of the plan.
- **FR-007**: No labels are added to or removed from the issue as part of this work. The `base:` label and `fabrik:cruise` remain as filed, apart from engine-managed state labels.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a Markdown file whose only purpose is to carry per-run HTML-comment markers.
- **Marker line**: `<!-- base-branch-pipeline-20261126-232752 -->`, an HTML comment with a run-unique timestamp suffix.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains one changed file and one added line, and the added line is exactly the marker line.
- **SC-002**: The PR base is `e2e-base-branch-20261126-232752`.
- **SC-003**: The issue reaches Done through the pipeline with no manual intervention beyond what `fabrik:cruise` permits, and with no stall at the review gate caused by missing issue-to-PR linkage.
- **SC-004**: The rendered Markdown of the marker file shows no visible change from the added line, because it is an HTML comment.

## Assumptions

- The base branch `e2e-base-branch-20261126-232752` exists on the remote, as the issue states it is already set.
- The marker file may or may not exist on the base branch, and both cases are handled by FR-001 and FR-002.
- The timestamp in the marker is taken as given and is not to be regenerated or altered.
- `fabrik:cruise` means the stages auto-advance without a human spec approval. It does not change what is built.

## Out of Scope

- Any change besides the single marker line.
- Adding or removing labels, or changing the `base:` label.
- Changes to engine behaviour for non-default bases. This issue only verifies the existing behaviour.
- Multi-repo scenarios.
- Decomposing the work into child issues.

## Source References

- handarbeit/fabrik#1046 (regression being covered)
- handarbeit/fabrik#1047 (issue-to-PR linkage fix being validated)
- handarbeit/fabrik#1050 (review-gate data feed fix being validated)
