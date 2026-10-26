# Feature Specification: E2E cruise full pipeline marker (20261026-050721)

**Feature Branch**: `fabrik/issue-1077`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-050721 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Cruise should carry an issue through every pipeline stage automatically, stop at Validate-complete without merging the PR, and close the issue correctly once a human merges the PR. The code change is deliberately trivial so that the pipeline behaviour, not the change itself, is what gets verified. Each run appends its own uniquely timestamped marker line to a shared marker file, so a successful run leaves a visible, attributable trace.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended (Priority: P1)

A maintainer running the cruise e2e check needs a PR that adds exactly this run's marker line to the end of the shared marker file, so the pipeline has a real, minimal change to carry through every stage.

**Why this priority**: Without the change there is no PR, and nothing else in the check can be verified.

**Independent Test**: Inspect the PR diff. It touches only `e2e/markers/cruise-full-pipeline.md` and adds only the line `<!-- cruise-pipeline-20261026-050721 -->` as the last line of the file.

**Acceptance Scenarios**:

1. **Given** the marker file exists with earlier run markers, **When** the change is made, **Then** the new marker line is the final line and all earlier lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is made, **Then** the file is created containing the marker line.

---

### User Story 2 - Cruise carries the issue to Validate-complete without merging (Priority: P1)

A maintainer applies `fabrik:cruise` and expects the issue to move through Specify, Research, Plan, Implement, Review and Validate on its own, ending at Validate-complete with the PR still open.

**Why this priority**: This is the contract under verification.

**Independent Test**: Observe the issue's stage labels and comments, and the PR state, after the pipeline settles.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the PR is open and unmerged and the issue has not advanced to Done.
3. **Given** the issue is a single one-line change, **When** Plan completes, **Then** no child issues are spawned.

---

### User Story 3 - Human merge closes the issue (Priority: P2)

After a human merges the PR, the issue should reach Done and be closed.

**Why this priority**: It verifies the tail of the contract, which depends on the first two stories.

**Independent Test**: Merge the PR by hand and check the issue state.

**Acceptance Scenarios**:

1. **Given** the issue is Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue moves to Done and is closed.

---

### Edge Cases

- The marker file already ends without a trailing newline: the new line must still land on its own line, not be joined to the previous one.
- Another run's marker is present in the file: it must be left untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- cruise-pipeline-20261026-050721 -->` as the last line of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST consist of that single file and that single added line, and MUST NOT modify existing lines.
- **FR-004**: The work MUST NOT be decomposed into child issues.
- **FR-005**: With `fabrik:cruise` present, the issue MUST advance through every stage to Validate-complete without the PR being merged by the engine.
- **FR-006**: After a human merges the PR, the issue MUST end up Done and closed.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file with exactly one added line, equal to the marker above.
- **SC-002**: The issue reaches Validate-complete with zero human approvals between stages.
- **SC-003**: At Validate-complete the PR is unmerged and zero child issues exist.
- **SC-004**: After the human merge, the issue is closed and in Done.

## Assumptions

- The marker file currently exists with earlier run markers, and the new line is simply appended.
- The timestamp in the marker (`20261026-050721`) is the run identifier and is used exactly as given.
- Merging the PR is performed by a human as part of the check, outside this issue's scope.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Decomposing the work or spawning child issues.
- Multi-repo behaviour.
- Engine-side merging of the PR.
