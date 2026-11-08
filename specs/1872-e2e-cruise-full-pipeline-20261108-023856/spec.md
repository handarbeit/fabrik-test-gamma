# Feature Specification: e2e cruise full pipeline (20261108-023856)

**Feature Branch**: `fabrik/issue-1872`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261108-023856 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue exists to exercise the `fabrik:cruise` pipeline contract (#898) end to end. Under cruise, an issue should advance automatically through every stage up to and including Validate completion. Cruise must not auto-merge the linked PR or advance the issue to Done. After that, a human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial: one marker line appended to a marker file. The size of the change is not the point. It gives the pipeline a real, minimal diff to carry through every stage so the cruise behaviour can be observed. The unique timestamp in the marker line (`20261108-023856`) ties this change to this particular verification run.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

An operator running the cruise verification needs a small, unambiguous change on the issue branch so the pipeline has something real to carry through every stage.

**Why this priority**: Without this change there is no PR and nothing to verify cruise against.

**Independent Test**: Inspect the PR diff. It should contain exactly one added line, the marker comment, at the end of the marker file.

**Acceptance Scenarios**:

1. **Given** the marker file already exists, **When** the change is implemented, **Then** its last line is `<!-- cruise-pipeline-20261108-023856 -->` and its prior content is unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created and its content is that single marker line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** exactly one file is touched and exactly one line is added.

---

### User Story 2 - Cruise carries the issue to Validate-complete without merging (Priority: P1)

An operator needs to confirm that an issue labelled `fabrik:cruise` advances through all stages on its own, but stops short of merging.

**Why this priority**: This is the contract under verification and the purpose of the issue.

**Independent Test**: Watch the issue's stage progression and the PR state after Validate completes.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate has completed, **When** the pipeline settles, **Then** the linked PR is still open and unmerged, and the issue has not moved to Done.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

An operator needs to confirm that human-driven merging after Validate-complete leads to the correct terminal state.

**Why this priority**: It completes the contract but depends on stories 1 and 2.

**Independent Test**: Merge the PR by hand and observe the issue.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue ends up closed and in Done.

---

### Edge Cases

- The marker file already exists with or without a trailing newline: the new line must end up on its own line at the end of the file, with existing content untouched.
- The marker file or its parent directory does not exist: both are created.
- Plan must not split this work into child issues; a single change is the whole scope.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261108-023856 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: The change MUST consist of exactly one added line in exactly one file, with no other modifications.
- **FR-004**: The work MUST be delivered as a single issue and a single PR, with no decomposition into child issues.
- **FR-005**: With `fabrik:cruise` present, the issue MUST advance through all stages up to Validate-complete without human approval at each stage.
- **FR-006**: Cruise MUST NOT auto-merge the PR and MUST NOT advance the issue to Done at Validate completion.
- **FR-007**: After a human merges the PR, the issue MUST close correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged PR diff shows exactly one file changed and exactly one line added, equal to the specified marker comment.
- **SC-002**: The issue reaches Validate-complete with no manual stage approvals.
- **SC-003**: At Validate-complete the PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merge, the issue is closed.

## Assumptions

- The existing cruise contract (#898) is the reference definition of correct behaviour. This issue verifies it and does not redefine it.
- The marker line is an HTML comment, so it has no effect on rendered output or on any code.
- Appending is idempotent per run. The unique timestamp means the line is not already present.

## Out of Scope

- Any change beyond the single marker line.
- Changes to the cruise behaviour itself.
- Multi-repo scenarios.
- Decomposition into child issues.
