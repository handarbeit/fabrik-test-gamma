# Feature Specification: E2E smoke full-pipeline marker line (20261015-062715)

**Feature Branch**: `fabrik/issue-77`
**Created**: 2026-10-15
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261015-062715 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with its pull request merged. The change itself is deliberately trivial, so that any failure points to the pipeline rather than to the work. The run is tagged `20261015-062715` so that this run's marker line is unique and identifiable. The issue carries the `fabrik:yolo` label, so stages are expected to advance automatically and the PR is expected to be merged automatically at the end of Validate.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged through the full pipeline (Priority: P1)

An operator running the smoke test wants a single, unique marker line recorded in the marker file, delivered through the normal pipeline, so that they can confirm the pipeline works from Specify to Done.

**Why this priority**: This is the entire purpose of the issue. The marker line is the observable artifact proving the pipeline completed.

**Independent Test**: After the pipeline finishes, inspect `e2e/markers/smoke-full-pipeline.md` on the default branch. The last line is the marker comment for this run, and the issue is Done with its PR merged.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is implemented, **Then** the exact line `<!-- smoke-full-pipeline-20261015-062715 -->` is the final line of the file and all pre-existing content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line `<!-- smoke-full-pipeline-20261015-062715 -->`.
3. **Given** the change is complete and `fabrik:yolo` is set, **When** the pipeline runs through Review and Validate, **Then** the linked PR is merged to `main` and the issue reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline: the marker must still appear on its own line at the end, not joined to the previous content.
- The marker line for this run is already present (for example after a re-run): the line must not be duplicated.
- Marker lines from earlier smoke runs already exist in the file: they must be left intact.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261015-062715 -->`.
- **FR-002**: If the file does not exist, it MUST be created containing that line.
- **FR-003**: Existing content of the file, if any, MUST be preserved unchanged, with the new line appended after it on its own line.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The change MUST touch only `e2e/markers/smoke-full-pipeline.md`; no other file is modified by the implementation.
- **FR-006**: The work MUST be confined to this single repository, and the Plan stage MUST NOT decompose it into child issues.
- **FR-007**: The pull request for this change MUST reference this issue with `Closes #77`.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only record of smoke-run marker comments.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261015-062715 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: On the default branch after the run, the last line of `e2e/markers/smoke-full-pipeline.md` is exactly `<!-- smoke-full-pipeline-20261015-062715 -->`, and it appears exactly once.
- **SC-002**: The diff of the merged PR consists of exactly one changed file with one added line, other than any engine-generated spec file.
- **SC-003**: The issue reaches Done with its PR merged, with no human intervention beyond applying the `fabrik:yolo` label.

## Assumptions

- The marker file may or may not already exist, and either case is handled as described above.
- Appending means adding a new final line, with a newline inserted first if the file lacks a trailing one.
- The engine-generated spec file under `specs/` is expected in the PR diff and is not counted as part of the implementation change.
- The content of the marker file is not otherwise validated by any tooling.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, stage configuration, or documentation.
- Cross-repo work or decomposition into child issues.
- Modifying or removing marker lines from previous smoke runs.
- Adding tests or tooling for the marker file.
