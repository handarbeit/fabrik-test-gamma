# Feature Specification: e2e base-branch pipeline marker (20261015-064341)

**Feature Branch**: `fabrik/issue-96`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261015-064341 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261015-064341 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end test fixture, not a product feature. It exercises the pipeline for an issue carrying a `base:<branch>` label, where the work targets a non-default base branch. It provides regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage) and #1050 (review-gate data feed). The code change is deliberately trivial so that the only variable under test is the pipeline's handling of the non-default base branch. The marker file does not currently exist on this branch, so the change creates it.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the non-default base branch (Priority: P1)

An e2e harness files this issue with `base:e2e-base-branch-20261015-064341` and `fabrik:cruise`. The pipeline carries it through every stage, producing a PR that targets the non-default base branch and adds the marker line.

**Why this priority**: This is the entire purpose of the issue. The marker line is the observable artifact that proves the pipeline ran end to end against the non-default base.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR head. The last line must be exactly `<!-- base-branch-pipeline-20261015-064341 -->`, and the PR diff must touch only that file.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file exists and contains the marker line as its final line.
2. **Given** the marker file already exists, **When** the change is implemented, **Then** the marker line is appended at the very end and existing content is unchanged.
3. **Given** the issue carries `base:e2e-base-branch-20261015-064341`, **When** the PR is created, **Then** the PR targets that branch, not the repository's other branches.

---

### User Story 2 - Issue↔PR linkage and review-gate data on a non-default base (Priority: P2)

The pipeline must discover the PR for this issue and feed review state to the review gate even though GitHub creates no automatic Development-panel link for a PR on a non-default base branch.

**Why this priority**: These are the contracts #1047 and #1050 fixed, and the reason this regression test exists. They are secondary to the change itself because they are verified by the pipeline's behaviour rather than the diff.

**Independent Test**: Observe that the pipeline finds the PR and advances through Review and Validate. The review gate must see review data for it.

**Acceptance Scenarios**:

1. **Given** the PR body contains `Closes #96`, **When** the pipeline polls, **Then** the PR is discovered as this issue's linked PR.
2. **Given** reviews or threads exist on the PR, **When** the review gate evaluates, **Then** it sees that data rather than treating the PR as unlinked.

---

### Edge Cases

- If the marker file exists with no trailing newline, the marker must still end up on its own line, so the appended line must not merge into the previous line.
- If the marker line is already present as the last line, no duplicate is added.
- The labels on the issue (`fabrik:cruise`, `base:e2e-base-branch-20261015-064341`) must not be added or removed.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` MUST end with the line `<!-- base-branch-pipeline-20261015-064341 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: If the file exists, existing content MUST be preserved and the line appended after it, on its own line.
- **FR-004**: The change MUST consist of exactly one file and one added line of content, with no other files modified.
- **FR-005**: The work MUST target the non-default base branch `e2e-base-branch-20261015-064341`, and the PR MUST include `Closes #96`.
- **FR-006**: The work MUST NOT be decomposed into child issues.
- **FR-007**: Labels on the issue MUST NOT be added or removed as part of this work.

### Key Entities

- **Marker file**: `e2e/markers/base-branch-pipeline.md`, a text file whose trailing HTML comment lines record e2e runs. HTML comments are invisible in rendered Markdown.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR's diff contains exactly one changed file and one added content line, the marker line.
- **SC-002**: The PR targets `e2e-base-branch-20261015-064341`, and the pipeline links it to issue #96 without manual intervention.
- **SC-003**: The issue progresses through all pipeline stages to completion with no human input and no pause.

## Assumptions

- The marker file does not exist yet on the base branch (none was found in `e2e/markers/`), so it will be created.
- The timestamp in the marker line is a unique run identifier, so it is used verbatim.
- Because the file is Markdown, a created file may contain only the marker line, with no heading.

## Out of Scope

- Any change to engine code, stage configuration or docs.
- Changes to other marker files in `e2e/markers/`.
- Decomposing the work or spawning child issues.
- Modifying issue labels.

## Source References

- handarbeit/fabrik#1046 (regression being covered)
- handarbeit/fabrik#1047 (issue↔PR linkage fix)
- handarbeit/fabrik#1050 (review-gate data feed fix)
