# Feature Specification: E2E base-branch pipeline marker (20261026-054340)

**Feature Branch**: `fabrik/issue-1116`
**Created**: 2026-10-26
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
> <!-- base-branch-pipeline-20261026-054340 -->
> ```
>
> That is the entire change. One file, one line. Plan should NOT decompose.
>
> ## Scope
>
> Single repo only, targeting the non-default base branch e2e-base-branch-20261026-054340 (already set via
> the base:<branch> label applied at filing — do not add or remove labels).

## Background

This issue is an end-to-end regression check of the `base:<branch>` contract, where an issue's work targets a non-default base branch. It guards against a repeat of handarbeit/fabrik#1046 and exercises the fixes for #1047 (issue↔PR linkage for non-default-base PRs) and #1050 (the review-gate data feed). The code change is deliberately trivial so that the pipeline mechanics are what is under test: branch creation from the base, PR targeting, linkage, review gate, and landing. The issue carries the `base:e2e-base-branch-20261026-054340` and `fabrik:cruise` labels, set at filing.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

A single marker line is appended to the marker file, and the work flows through the pipeline against the labelled non-default base branch rather than the repository's default branch.

**Why this priority**: This is the only deliverable. It is the observable artifact the e2e run verifies.

**Independent Test**: After the pipeline finishes, `e2e/markers/base-branch-pipeline.md` on the base branch `e2e-base-branch-20261026-054340` ends with the exact marker line, and the PR targeted that branch.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` exists on the base branch, **When** the change is applied, **Then** the file's last line is `<!-- base-branch-pipeline-20261026-054340 -->` and no other line is altered.
2. **Given** the file does not exist, **When** the change is applied, **Then** the file is created containing the marker line.
3. **Given** the issue carries the `base:e2e-base-branch-20261026-054340` label, **When** the PR is opened, **Then** the PR's base is `e2e-base-branch-20261026-054340` and the PR is linked to this issue.

---

### Edge Cases

- The file exists but lacks a trailing newline: the marker must still occupy its own final line.
- The marker line is already present (re-run): it must not be duplicated.
- The file is created fresh: it contains only the marker line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- base-branch-pipeline-20261026-054340 -->`, at the very end of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it MUST be created first, containing the marker line.
- **FR-003**: The change MUST touch only that one file and add only that one line.
- **FR-004**: The work MUST target the base branch `e2e-base-branch-20261026-054340`, as set by the existing `base:` label; labels MUST NOT be added or removed as part of the change.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The change MUST be scoped to this single repository only.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR's diff against its base contains exactly one changed file with exactly one added line, the marker line.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261026-054340`, and the PR is recognised as linked to issue #1116.
- **SC-003**: The item progresses through the pipeline stages and the review gate to completion, with no manual intervention beyond what `fabrik:cruise` implies.
- **SC-004**: The marker line is present exactly once at the end of the file on the base branch after landing.

## Assumptions

- The base branch `e2e-base-branch-20261026-054340` already exists on the remote, as stated in the issue.
- The marker line is a fixed literal that includes the timestamp suffix; no other timestamp is substituted.
- A trailing newline after the marker line is acceptable.
- Existing content of the marker file, if any, is preserved unchanged.

## Out of Scope *(optional)*

- Any change to engine, stage configuration, labels, or documentation.
- Child issues, decomposition, or multi-repo work.
- Adding or removing the `base:` or `fabrik:cruise` labels.
- Changes to any file other than `e2e/markers/base-branch-pipeline.md`.

## Source References *(optional)*

- handarbeit/fabrik#1046 (regression being guarded), #1047 (issue↔PR linkage fix), #1050 (review-gate data feed fix)
