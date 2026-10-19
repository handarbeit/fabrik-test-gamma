# Feature Specification: E2E smoke full-pipeline marker line (20261019-070912)

**Feature Branch**: `fabrik/issue-441`
**Created**: 2026-10-19
**Status**: Specified
**Input**: User description: "e2e smoke full-pipeline (20261019-070912). Goal: End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR. Trivial change: Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- smoke-full-pipeline-20261019-070912 -->`. That's the entire change. One file, one line. Scope: Single repo only — no cross-repo work. The Plan stage should NOT decompose. Labels: fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline on a single repository. Its purpose is to show that Fabrik can take an issue through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done, ending with a merged pull request. The code change is deliberately trivial so that any failure points at the pipeline and not at the work. The issue carries `fabrik:yolo`, so stages advance automatically and the PR is merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line appended and merged through the full pipeline (Priority: P1)

An operator files this issue. Fabrik carries it through all stages without manual intervention. The result is a merged PR that adds one marker comment line to the end of `e2e/markers/smoke-full-pipeline.md`.

**Why this priority**: This is the only purpose of the issue. It exercises the whole pipeline with the smallest possible change.

**Independent Test**: After the issue reaches Done, check the default branch. `e2e/markers/smoke-full-pipeline.md` must exist and its last line must be `<!-- smoke-full-pipeline-20261019-070912 -->`. The PR that made the change must be merged and the issue closed.

**Acceptance Scenarios**:

1. **Given** the file `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's previous content is unchanged and the marker line is appended as the new last line.
2. **Given** the file does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the PR is merged and the issue reaches Done.
4. **Given** the Plan stage runs, **When** it produces its plan, **Then** the plan is a single unit of work and no child issues are spawned.

---

### Edge Cases

- The file exists but its last line has no trailing newline. The marker must still be on its own line, not joined to the previous content.
- The marker line is already present at the end of the file, for example from a re-run. It must not be added a second time.
- The file exists with other marker lines from earlier smoke runs. These stay untouched.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` ends with the exact line `<!-- smoke-full-pipeline-20261019-070912 -->`.
- **FR-002**: If the file does not exist, it is created, and the marker line is its content.
- **FR-003**: If the file exists, its existing content is preserved and the marker line is appended after it on its own line.
- **FR-004**: The marker line appears exactly once in the file.
- **FR-005**: The change touches only `e2e/markers/smoke-full-pipeline.md` (plus the engine-managed spec file).
- **FR-006**: The work is done in a single repository, and the Plan stage does not decompose it into child issues.
- **FR-007**: The change is delivered through a PR that links to this issue and is merged to the default branch (`main`).

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done with no manual intervention after it is filed.
- **SC-002**: The linked PR is merged into `main`.
- **SC-003**: On `main`, the last line of `e2e/markers/smoke-full-pipeline.md` is `<!-- smoke-full-pipeline-20261019-070912 -->`, and that string appears exactly once.
- **SC-004**: The merged PR's diff contains the marker line and, apart from the engine-managed `specs/` file, no other file changes.

## Assumptions

- A trailing newline after the marker line is fine and is not treated as an extra line.
- The directory `e2e/markers/` may need to be created if it does not exist.
- Because the issue carries `fabrik:yolo`, no human approval is needed between stages.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change to engine code, stage configuration, tests, or documentation.
- Decomposing the work into child issues.
- Changing other marker files or earlier smoke-run markers.
