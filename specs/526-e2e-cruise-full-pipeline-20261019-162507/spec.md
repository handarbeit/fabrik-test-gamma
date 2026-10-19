# Feature Specification: e2e cruise full pipeline marker (20261019-162507)

**Feature Branch**: `fabrik/issue-526`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "End-to-end verification of the fabrik:cruise pipeline contract (#898). Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- cruise-pipeline-20261019-162507 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Single repo only. This issue verifies that cruise auto-advances through all stages to Validate-complete without merging the PR, and that the issue closes correctly after a human merges the PR."

## Background

This is an end-to-end smoke test of the `fabrik:cruise` label contract (#898). With `fabrik:cruise`, an issue should auto-advance through every pipeline stage up to Validate-complete, but the linked PR must not be auto-merged and the issue must not move to Done on its own. After a human merges the PR, the issue should close correctly.

The code change is deliberately trivial so that the pipeline's behavior is the thing under test, not the change. The marker file `e2e/markers/cruise-full-pipeline.md` already exists in the repository. It holds one HTML comment line per earlier run, and the most recent one is `<!-- cruise-pipeline-20261017-175659 -->`. This run appends its own line in the same format.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Append the run marker line (Priority: P1)

A maintainer running the cruise e2e check needs this run to leave a unique, identifiable trace in the repository. That trace is a PR that changes exactly one file by adding exactly one line.

**Why this priority**: This is the only change the issue asks for. The pipeline verification depends on a real, minimal PR existing.

**Independent Test**: Inspect the PR diff. It should touch only `e2e/markers/cruise-full-pipeline.md` and add only the line `<!-- cruise-pipeline-20261019-162507 -->` as the last line of the file.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/cruise-full-pipeline.md` exists with earlier marker lines, **When** the change is implemented, **Then** the file ends with `<!-- cruise-pipeline-20261019-162507 -->` and all prior lines are unchanged.
2. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file is created containing the single line `<!-- cruise-pipeline-20261019-162507 -->`.

---

### User Story 2 - Cruise pipeline contract holds (Priority: P1)

An operator wants confirmation that a `fabrik:cruise` issue advances through all stages without human intervention, stops at Validate-complete without merging the PR, and closes correctly once a human merges the PR.

**Why this priority**: This is the purpose of the issue.

**Independent Test**: Observe the issue's stage labels and board column as it moves from Specify to Validate. Confirm the PR stays open and unmerged at Validate-complete. Then merge the PR by hand and confirm the issue closes and reaches Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without a human approving it.
2. **Given** Validate has completed, **When** the engine evaluates the landing decision, **Then** the PR is not merged and the issue does not advance to Done.
3. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the issue closes and moves to Done.

---

### Edge Cases

- The marker file lacks a trailing newline. The new line must still sit on its own line, with no earlier line altered other than the necessary line terminator.
- Plan must not decompose the work or spawn child issues, because the change is a single line in a single file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/cruise-full-pipeline.md` MUST end with the line `<!-- cruise-pipeline-20261019-162507 -->`.
- **FR-002**: The change MUST add exactly one line and modify no other line in that file.
- **FR-003**: The change MUST touch no file other than `e2e/markers/cruise-full-pipeline.md`. Specify's own spec file under `specs/` is engine-managed and is not part of the implemented change.
- **FR-004**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it MUST be created containing only the marker line from FR-001.
- **FR-005**: The Plan stage MUST NOT decompose the work into child issues.
- **FR-006**: With `fabrik:cruise` present, the issue MUST auto-advance through all stages to Validate-complete.
- **FR-007**: The linked PR MUST NOT be merged automatically at Validate-complete, and the issue MUST NOT advance to Done until the PR is merged.
- **FR-008**: After a human merges the PR, the issue MUST close.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `main` is one file and one added line, matching FR-001.
- **SC-002**: The issue reaches Validate-complete with no manual stage advancement.
- **SC-003**: At Validate-complete the PR is still open and unmerged.
- **SC-004**: Within one poll cycle of a human merging the PR, the issue is closed.

## Assumptions

- The marker file already exists in the repo, with earlier cruise run lines in the same format. The create-if-missing clause stays as a fallback.
- Whether the file ends with a trailing newline is not specified, so the implementer handles it as long as FR-001 and FR-002 hold.
- The test is scoped to this single repository, as the issue states.

## Out of Scope *(optional)*

- Any change beyond the one marker line.
- Multi-repo behavior.
- Changes to engine cruise logic. This issue only verifies it.
- Merging the PR. A human does that.

## Source References *(optional)*

- Cruise pipeline contract: #898
