# Feature Specification: e2e cruise full pipeline (20261031-022220)

**Feature Branch**: `fabrik/issue-1503`
**Created**: 2026-10-30
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261031-022220 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` contract (#898). Cruise should auto-advance an issue through every pipeline stage up to Validate-complete, but it must not auto-merge the linked PR or advance the issue to Done. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial so the pipeline's behaviour, not the change, is what gets verified. The change is one marker line appended to one file. It is small enough that the Plan stage has no reason to split it into child issues.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

A maintainer wants a unique, identifiable marker for this run committed to the repository. The marker proves the pipeline produced a real change on the issue branch and that the change can be merged.

**Why this priority**: It is the only code change in the issue. Without it there is no PR for cruise to hold at Validate-complete.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its last line is exactly `<!-- cruise-pipeline-20261031-022220 -->`, and the diff against the base branch adds only that line (plus the file itself if it did not exist).

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- cruise-pipeline-20261031-022220 -->` and all earlier content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the file already ends with a different marker, **When** the change is implemented, **Then** the new marker is appended after it and the existing lines are preserved.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

An operator applies `fabrik:cruise` to the issue and expects it to move through Specify, Research, Plan, Implement, Review and Validate without further human action, stopping at Validate-complete with the PR open.

**Why this priority**: This is the contract under verification (#898).

**Independent Test**: Watch the issue move from Specify to Validate with each stage's completion recorded. At Validate-complete, check that the linked PR is open and unmerged and that the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without a human approving it.
2. **Given** Validate has completed, **When** the engine evaluates the next step, **Then** the PR is not merged and the issue is not moved to Done.
3. **Given** the Plan stage runs, **When** it finishes, **Then** it produces a single-task plan and spawns no child issues.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR. The issue should then reach its terminal state correctly.

**Why this priority**: It completes the contract but depends on the first two stories.

**Independent Test**: Merge the PR by hand and confirm the issue is closed and reaches Done.

**Acceptance Scenarios**:

1. **Given** the issue is at Validate-complete with its PR open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file may already end without a trailing newline. The appended marker must still land on its own line.
- The marker file may already exist with earlier markers from other runs. Those lines must not be altered.
- The `e2e/markers/` directory may not exist. It is created along with the file.
- Both `fabrik:cruise` and `fabrik:yolo` may be present. Cruise takes precedence, so no auto-merge happens.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The final line of `e2e/markers/cruise-full-pipeline.md` is exactly `<!-- cruise-pipeline-20261031-022220 -->`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created containing the marker line.
- **FR-003**: Existing content of `e2e/markers/cruise-full-pipeline.md` is left unchanged. The marker is only appended.
- **FR-004**: The change touches no file other than `e2e/markers/cruise-full-pipeline.md`.
- **FR-005**: The Plan stage does not decompose the work and spawns no child issues.
- **FR-006**: With `fabrik:cruise` applied, the issue advances through all pipeline stages to Validate-complete without a human approving each stage.
- **FR-007**: At Validate-complete, the linked PR remains open and unmerged, and the issue is not moved to Done.
- **FR-008**: After a human merges the PR, the issue is closed correctly.

### Key Entities

- **Marker file**: `e2e/markers/cruise-full-pipeline.md`, a Markdown file whose lines are HTML comments identifying e2e runs.
- **Marker line**: The unique comment `<!-- cruise-pipeline-20261031-022220 -->` identifying this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one added marker line in exactly one file, with no other changed lines.
- **SC-002**: The issue reaches Validate-complete with no human advancement action between Specify and Validate.
- **SC-003**: At Validate-complete, 0 merged PRs are linked to the issue and the issue is not in Done.
- **SC-004**: After the human merge, the issue is closed with no manual label or board correction.
- **SC-005**: Zero child issues are spawned from this issue.

## Assumptions

- The marker line is written exactly as given, with no added or changed characters.
- The "very end of the file" means the last line. A trailing newline after the marker is acceptable.
- The issue number is 1503 and the repository's default base branch is `main`.
- The marker file may or may not already exist. The spec covers both cases.
- `fabrik:cruise` is already applied to the issue, as the Labels section states.

## Out of Scope *(optional)*

- Any change other than the single marker line.
- Decomposing the work into child issues.
- Auto-merging the PR. Cruise deliberately does not do this.
- Testing multi-repo behaviour. This issue is single-repo only.
- Changes to engine behaviour or to the cruise contract itself (#898).

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract this issue verifies.
