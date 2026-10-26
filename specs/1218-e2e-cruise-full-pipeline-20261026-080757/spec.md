# Feature Specification: e2e cruise full pipeline (20261026-080757)

**Feature Branch**: `fabrik/issue-1218`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261026-080757 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is a live end-to-end check of the `fabrik:cruise` label contract (#898). Under that contract, an issue carrying `fabrik:cruise` advances automatically through every pipeline stage up to Validate-complete, but the engine does not merge the linked PR or move the issue to Done. A human merges the PR, and the issue then closes. The work itself is deliberately trivial: a one-line marker appended to a single file, so that any failure points at the pipeline rather than at the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the cruise marker file (Priority: P1)

The repository's cruise marker file ends with a unique marker line identifying this e2e run.

**Why this priority**: This is the only code change in the issue. It gives the pipeline a real diff to carry through every stage.

**Independent Test**: Open `e2e/markers/cruise-full-pipeline.md` on the PR branch and check that its last line is exactly `<!-- cruise-pipeline-20261026-080757 -->` and that nothing else in the file has changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists, **When** the change is implemented, **Then** the file's final line is `<!-- cruise-pipeline-20261026-080757 -->` and all prior content is unchanged.
2. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and its final line is `<!-- cruise-pipeline-20261026-080757 -->`.

---

### User Story 2 - Cruise advances through all stages without merging (Priority: P1)

With `fabrik:cruise` on the issue, the pipeline moves it from Specify through Validate-complete with no human advancement, and the PR is left unmerged.

**Why this priority**: This is the contract under test. The marker change only supplies something for the pipeline to carry.

**Independent Test**: Watch the issue after Specify completes. Confirm it reaches Validate-complete on its own, that the linked PR is still open at that point, and that the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances it to the next stage without human action.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue is not moved to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue closes and reaches Done.

**Why this priority**: It completes the contract. It depends on Story 2 having reached Validate-complete.

**Independent Test**: Merge the PR by hand after Validate-complete and check that the issue ends up closed and in Done.

**Acceptance Scenarios**:

1. **Given** the issue is at Validate-complete with the PR open, **When** a human merges the PR, **Then** the issue is closed and moved to Done.

---

### Edge Cases

- The marker file does not exist yet: it is created, then the line is appended.
- The marker file exists without a trailing newline: the marker still ends up on its own final line.
- Plan must not split this issue into child issues, since the change is one file and one line.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/cruise-full-pipeline.md` MUST end with the single line `<!-- cruise-pipeline-20261026-080757 -->`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created before the line is appended.
- **FR-003**: The change MUST touch only that one file and add only that one line, plus any newline needed to place the marker on its own line.
- **FR-004**: Plan MUST NOT decompose this issue into child issues.
- **FR-005**: With `fabrik:cruise` present, the engine MUST advance the issue through all stages to Validate-complete without human intervention.
- **FR-006**: The engine MUST NOT merge the linked PR or move the issue to Done when Validate completes.
- **FR-007**: After a human merges the PR, the issue MUST end up closed and in Done.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file, `e2e/markers/cruise-full-pipeline.md`, with exactly one added marker line.
- **SC-002**: The issue reaches Validate-complete with no manual stage advancement.
- **SC-003**: The PR is still unmerged when the issue reaches Validate-complete.
- **SC-004**: Within one poll cycle of a human merging the PR, the issue is closed and in Done.

## Assumptions

- The marker text is taken exactly as given, including the timestamp `20261026-080757`.
- The pipeline's own `specs/<issue>-<slug>/spec.md` file is created by the engine and does not count against the "one file" rule in FR-003.
- "Single repo only" means no cross-repo work is needed.

## Out of Scope *(optional)*

- Any change other than the one marker line.
- Splitting the work into child issues.
- Modifying engine behavior; this issue only observes it.
- Merging the PR automatically; the merge is a human step.

## Source References *(optional)*

- #898: the `fabrik:cruise` pipeline contract this issue verifies.
