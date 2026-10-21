# Feature Specification: E2E cruise full pipeline marker (20261021-201758)

**Feature Branch**: `fabrik/issue-636`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261021-201758 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end test fixture for the `fabrik:cruise` pipeline contract (#898). Under that contract, an item carrying `fabrik:cruise` auto-advances through every stage up to Validate-complete. It does not auto-merge the PR or move the issue to Done. A human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial: one marker line appended to a marker file. Its only purpose is to give the pipeline a real, minimal diff and PR to carry through every stage. The product behavior under test is the cruise pipeline itself, not the file content.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended to the cruise marker file (Priority: P1)

An e2e operator needs a unique, traceable marker for this run in the shared marker file, so the run produces a minimal PR to carry through the pipeline.

**Why this priority**: The marker line is the entire deliverable. Without it there is no diff, no PR, and nothing for the pipeline stages to process.

**Independent Test**: Inspect `e2e/markers/cruise-full-pipeline.md` on the issue branch. Its last line is exactly `<!-- cruise-pipeline-20261021-201758 -->`, and it is the only line this change adds.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is made, **Then** the marker line is appended as the last line and no existing content is altered.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is made, **Then** the file is created containing the marker line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** it touches exactly one file and adds exactly one line (excluding any engine-generated spec file under `specs/`).

---

### User Story 2 - Cruise carries the item through all stages without merging (Priority: P1)

An operator verifying the cruise contract needs the item to auto-advance through every stage to Validate-complete and then stop, leaving the PR unmerged.

**Why this priority**: This is the actual contract under test (#898). The file change only enables it.

**Independent Test**: With `fabrik:cruise` on the issue, observe the item progress through Specify, Research, Plan, Implement, Review and Validate without human stage approvals. At the end the PR is open and unmerged, and the issue has not moved to Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the item advances to the next stage automatically.
2. **Given** Validate completes, **When** the engine evaluates landing, **Then** the PR is not merged and the issue is not advanced to Done.
3. **Given** Plan runs on this issue, **When** it produces its plan, **Then** the plan is a single unit of work and no child issues are spawned.

---

### User Story 3 - Issue closes correctly after a human merge (Priority: P2)

After a human merges the PR, the issue must close and reach its terminal state correctly.

**Why this priority**: It completes the cruise lifecycle, but only after the primary contract (no auto-merge) has been shown.

**Independent Test**: Merge the PR by hand after Validate-complete, then confirm the issue is closed and the item reaches Done.

**Acceptance Scenarios**:

1. **Given** Validate is complete and the PR is open, **When** a human merges the PR, **Then** the issue closes and the item reaches its terminal state.

---

### Edge Cases

- The marker file does not exist yet: it is created with the marker line as its content.
- The marker file exists but does not end with a newline: the marker still ends up on its own final line rather than being joined to the previous line.
- The marker file already contains markers from other e2e runs: they are left untouched, and only this run's unique marker is appended.
- The marker line is already present (for example on a re-run): it is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261021-201758 -->`.
- **FR-002**: If the file does not exist, it MUST be created so that the marker line is its content.
- **FR-003**: Existing content of the file MUST NOT be modified or removed.
- **FR-004**: The change MUST affect exactly one file and add exactly one line. The engine-generated spec file under `specs/` is not part of the change.
- **FR-005**: The marker line MUST appear exactly once in the file.
- **FR-006**: The Plan stage MUST NOT decompose this work or spawn child issues.
- **FR-007**: With `fabrik:cruise` present, the item MUST auto-advance through all stages to Validate-complete without the PR being auto-merged.
- **FR-008**: After a human merges the PR, the issue MUST close correctly.

### Key Entities *(if applicable)*

- **Cruise marker file** (`e2e/markers/cruise-full-pipeline.md`): The shared file in which each e2e cruise run records one unique HTML-comment marker.
- **Run marker**: The single line `<!-- cruise-pipeline-20261021-201758 -->`, unique to this run's timestamp.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR for this issue changes exactly one file with exactly one added line, and that line is the marker.
- **SC-002**: The item reaches Validate-complete with no human stage approvals between Specify and Validate.
- **SC-003**: At Validate-complete the PR is still open and unmerged, and the issue is still open and not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.
- **SC-005**: No child issues are created for this issue.

## Assumptions

- The directory `e2e/markers/` may or may not already exist. If it does not, creating it along with the file is acceptable.
- The marker is an HTML comment so that it is invisible when the markdown is rendered and cannot disturb any content the file already has.
- The exact marker text is fixed by the issue, and the timestamp in it is not to be altered.
- The engine-committed spec file under `specs/` is expected pipeline output and does not count against the "one file, one line" constraint.

## Out of Scope *(optional)*

- Any change to engine, stage or skill behavior, since this issue only verifies existing cruise behavior.
- Decomposing the work into sub-issues.
- Multi-repo scenarios.
- Auto-merging the PR (cruise explicitly does not do this).
- Cleaning up or consolidating other markers in the marker file.

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract this issue verifies.
