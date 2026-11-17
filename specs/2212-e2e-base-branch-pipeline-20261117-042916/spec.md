# Feature Specification: e2e base-branch pipeline (20261117-042916)

**Feature Branch**: `fabrik/issue-2212`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261117-042916 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261117-042916 (already set via
the base:<branch> label applied at filing — do not add or remove labels).

## Labels

fabrik:cruise, base:e2e-base-branch-20261117-042916"

## Background

This issue is a live end-to-end run that exercises the pipeline for an issue carrying a `base:<branch>` label, i.e. work that targets a non-default base branch. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes made in #1047 (linking the issue to its pull request) and #1050 (feeding the review gate its data). The code change is deliberately minimal; the value is in the pipeline passing through every stage correctly on the non-default base, not in the change itself.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base (Priority: P1)

An operator running the base-branch e2e wants a single, recognizable marker line added to a marker file so the pipeline has a real, minimal change to carry through every stage against the non-default base branch.

**Why this priority**: It is the only change in the issue; without it there is nothing for the pipeline to carry.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR branch and confirm its last line is the marker comment, with no other file changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists, **When** the change is made, **Then** the file's final line is `<!-- base-branch-pipeline-20261117-042916 -->` and all prior content is unchanged.
2. **Given** the file does not exist, **When** the change is made, **Then** the file is created and contains the marker comment line as its last line.

---

### User Story 2 - Pipeline honours the non-default base (Priority: P2)

The operator wants the item to move through the pipeline with its PR targeting `e2e-base-branch-20261117-042916`, correctly linked to the issue, and with the review gate receiving its data.

**Why this priority**: This is what the e2e actually verifies (#1047, #1050), but it depends on the change in Story 1 existing.

**Independent Test**: Observe the resulting PR: its base is the non-default branch, it is linked to this issue, and the review gate evaluates it.

**Acceptance Scenarios**:

1. **Given** the issue carries `base:e2e-base-branch-20261117-042916`, **When** the pipeline creates a PR, **Then** the PR targets that branch rather than the repository's other branches.
2. **Given** the PR exists, **When** the pipeline evaluates it, **Then** the issue and PR are recognised as linked and the review gate is fed the PR's review data.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still appear on its own final line.
- The marker file already contains the marker line (e.g. a re-run): the line must not be duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The line `<!-- base-branch-pipeline-20261117-042916 -->` is appended as the last line of `e2e/markers/base-branch-pipeline.md`.
- **FR-002**: If `e2e/markers/base-branch-pipeline.md` does not exist, it is created containing that line.
- **FR-003**: The change touches exactly one file and adds exactly one line (apart from the spec file the engine commits).
- **FR-004**: The work is delivered as a single unit and is not decomposed into child issues.
- **FR-005**: The PR targets the base branch `e2e-base-branch-20261117-042916`.
- **FR-006**: No labels are added to or removed from the issue as part of this work.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch changes `e2e/markers/base-branch-pipeline.md` only, by exactly one added line.
- **SC-002**: The PR's base branch is `e2e-base-branch-20261117-042916` and the PR is linked to this issue.
- **SC-003**: The item completes every pipeline stage without manual intervention beyond the existing `fabrik:cruise` automation.

## Assumptions

- The default base branch for this repository's pipeline run is `e2e-base-branch-20261117-042916`, matching the `base:` label.
- The marker file may or may not exist yet; creating it is acceptable.
- The marker comment is harmless content and is not interpreted by any tool.

## Out of Scope *(optional)*

- Any change beyond the single marker line.
- Splitting the work into child issues.
- Changing labels or the base-branch configuration.
- Multi-repo scenarios.
