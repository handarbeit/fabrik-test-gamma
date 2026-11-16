# Feature Specification: e2e base-branch pipeline marker (20261116-172243)

**Feature Branch**: `fabrik/issue-2146`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261116-172243 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261116-172243 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end regression check of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be forked from, rebased onto, and have its PR target that branch rather than the repository default. It gives regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The code change is deliberately trivial, so that any failure points at the pipeline's base-branch handling and not at the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An e2e run files this issue with the `base:e2e-base-branch-20261116-172243` label. The pipeline produces a one-line change to `e2e/markers/base-branch-pipeline.md` and delivers it through a PR that targets that base branch.

**Why this priority**: This is the whole purpose of the issue. It exercises the non-default-base path from Implement through landing.

**Independent Test**: After the pipeline finishes, inspect the PR and the base branch. The PR's base is `e2e-base-branch-20261116-172243`, its diff is exactly one added line, and that line is the final line of `e2e/markers/base-branch-pipeline.md`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists, **When** the change is applied, **Then** `<!-- base-branch-pipeline-20261116-172243 -->` is appended as the last line and no other line in the file changes.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `base:e2e-base-branch-20261116-172243`, **When** the PR is created, **Then** the PR targets that branch, not the repository's other branches.

---

### Edge Cases

- The file already exists with no trailing newline: the marker must still land on its own line at the end.
- The marker line is already present (for example on a re-run): it must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The exact line `<!-- base-branch-pipeline-20261116-172243 -->` is appended at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created before the line is appended.
- **FR-003**: The change touches exactly one file and adds exactly one line (plus file creation if needed).
- **FR-004**: The work is delivered by a PR targeting the base branch `e2e-base-branch-20261116-172243`, linked to this issue via `Closes #2146`.
- **FR-005**: The issue's labels (`fabrik:cruise`, `base:e2e-base-branch-20261116-172243`) are not added to or removed by the work itself.
- **FR-006**: The work is a single unit and is not decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261116-172243` contains exactly one changed file with exactly one added line, equal to the marker line.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261116-172243`, and the issue↔PR linkage and review-gate data feed work as for a default-base PR, with no manual intervention.
- **SC-003**: The pipeline advances this issue through all stages without any child issues being spawned.

## Assumptions

- The target base branch `e2e-base-branch-20261116-172243` already exists on the remote (it is also reported as the repository's default base for this run).
- The `fabrik:cruise` label governs auto-advance, so no human approval of this spec is needed.
- The `e2e/markers/` directory may need to be created together with the file.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Modifying labels, stage configuration, or engine behaviour.
- Decomposition into multiple issues or multi-repo work.
- Fixing #1046, #1047 or #1050 themselves; this issue only verifies them.
