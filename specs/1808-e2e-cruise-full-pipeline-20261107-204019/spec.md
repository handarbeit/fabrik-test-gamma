# Feature Specification: E2E cruise full pipeline marker (20261107-204019)

**Feature Branch**: `fabrik/issue-1808`
**Created**: 2026-11-07
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261107-204019 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This issue is a live end-to-end probe of the `fabrik:cruise` pipeline contract (#898). Under cruise, an item auto-advances through every stage up to Validate-complete, but the engine does not merge the PR or advance the item to Done. A human merges the PR, and the issue should then close correctly.

The code change is deliberately trivial: one marker line appended to one file. Its only purpose is to give the pipeline a real but minimal diff and PR to carry through every stage. The timestamp in the marker (`20261107-204019`) makes this run's change unique.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the e2e marker file (Priority: P1)

An operator running the e2e suite needs a trivial, deterministic change that produces a one-file PR. The change is a single HTML comment line at the end of `e2e/markers/cruise-full-pipeline.md`.

**Why this priority**: This is the whole deliverable. Without a PR carrying a real diff, the pipeline contract cannot be exercised.

**Independent Test**: Inspect the PR diff. Apart from the engine-committed spec artifact `specs/1808-e2e-cruise-full-pipeline-20261107-204019/spec.md`, it touches only `e2e/markers/cruise-full-pipeline.md` and adds only the line `<!-- cruise-pipeline-20261107-204019 -->` as the last line of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file exists and its last line is `<!-- cruise-pipeline-20261107-204019 -->`.
2. **Given** `e2e/markers/cruise-full-pipeline.md` already exists, **When** the change is implemented, **Then** its prior content is unchanged and the marker line is appended at the very end.

---

### User Story 2 - Cruise advances to Validate-complete without merging (Priority: P1)

An operator verifying the cruise contract needs the item to advance through every stage automatically and stop at Validate-complete with the PR still open.

**Why this priority**: This is the behavior the e2e run exists to verify.

**Independent Test**: After the pipeline runs, the item carries `stage:Validate:complete`, the linked PR is open and unmerged, and the issue is not in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the engine advances to the next stage without human approval.
2. **Given** Validate completes, **When** the engine evaluates the landing decision, **Then** it does not merge the PR and does not move the issue to Done.

---

### User Story 3 - Issue closes after a human merges the PR (Priority: P2)

After a human merges the PR, the issue should close and reach Done correctly.

**Why this priority**: It completes the contract, but it depends on the earlier stories.

**Independent Test**: Merge the PR manually and confirm the issue is closed and moved to Done.

**Acceptance Scenarios**:

1. **Given** the item is Validate-complete with an open PR, **When** a human merges the PR, **Then** the issue is closed and ends in Done.

---

### Edge Cases

- The marker file already exists with a trailing newline. The new line is appended after the last existing line and the file ends with exactly one marker line (a trailing newline is acceptable).
- The marker file already exists without a trailing newline. A newline is added so the marker sits on its own line rather than being joined to the previous line.
- The marker line is already present. It is not duplicated.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/cruise-full-pipeline.md` MUST contain the line `<!-- cruise-pipeline-20261107-204019 -->` as its last line.
- **FR-002**: If the file does not exist, it MUST be created, with the marker line as its content.
- **FR-003**: If the file exists, its existing content MUST be preserved unchanged and the marker line only appended.
- **FR-004**: The change MUST touch exactly one file and add exactly one line, apart from the engine-committed spec artifact `specs/1808-e2e-cruise-full-pipeline-20261107-204019/spec.md`, which is pipeline output and not part of the code change.
- **FR-005**: The work MUST NOT be decomposed into child issues.
- **FR-006**: The PR MUST NOT be merged by the engine under `fabrik:cruise`. A human merges it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Excluding the engine-committed spec artifact, the PR diff is exactly one file and one added line, matching the marker text.
- **SC-002**: The item reaches Validate-complete through automatic stage advancement with no human approval between stages.
- **SC-003**: At Validate-complete the linked PR is open and unmerged, and the issue is not in Done.
- **SC-004**: After a human merges the PR, the issue is closed.

## Assumptions

- The `e2e/markers/` directory may not exist yet. Creating it as part of creating the file is acceptable.
- The marker text is used exactly as given, including the timestamp.
- Single repo only: no cross-repo work or spawned children are involved.

## Out of Scope

- Any change beyond the single marker line.
- Decomposition into child issues.
- Merging the PR (done by a human after Validate-complete).
- Changes to the cruise logic itself.
