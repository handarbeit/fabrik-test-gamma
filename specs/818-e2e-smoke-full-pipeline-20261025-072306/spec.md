# Feature Specification: E2E smoke full-pipeline marker line (20261025-072306)

**Feature Branch**: `fabrik/issue-818`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261025-072306 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) and finish in Done with a merged PR. The code change is deliberately trivial: one marker line appended to one file. A failure here points at the pipeline, not at the change. The run is tagged `20261025-072306` so its marker line can be told apart from other smoke runs. The issue carries `fabrik:yolo`, so each stage advances automatically and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator runs the smoke issue through the pipeline. When it finishes, the repository's default branch holds the marker file with the run's marker line as its last line.

**Why this priority**: This is the only deliverable. It shows the pipeline produced and merged a real change.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. The last line must be `<!-- smoke-full-pipeline-20261025-072306 -->`, and the file's diff in the PR must be a one-line addition.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the file's original content is unchanged and the marker line is appended as the new last line.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the marker line is already the last line of the file, **When** the change is implemented, **Then** the line is not duplicated.

---

### User Story 2 - Issue reaches Done through the full pipeline (Priority: P1)

The issue passes through all pipeline stages without human intervention and ends in Done with its PR merged.

**Why this priority**: Exercising the pipeline is the point of the smoke test.

**Independent Test**: Inspect the issue after the run. It must be in the Done column and closed, with a merged PR that references it.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without waiting for a human.
2. **Given** Validate completes, **When** the PR passes its gates, **Then** the PR is merged and the issue moves to Done.
3. **Given** the Plan stage runs, **When** it produces its plan, **Then** the plan is a single unit of work and no child issues are spawned.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker must still land on its own line, not be joined to the previous line.
- The marker file exists with other marker lines from earlier runs: those lines are left untouched.
- The marker line is already present as the last line (for example on a retry): no second copy is added.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append exactly one line, `<!-- smoke-full-pipeline-20261025-072306 -->`, as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it so that it contains the marker line.
- **FR-003**: The change MUST NOT modify or remove any existing content of that file.
- **FR-004**: The change MUST NOT touch any file other than `e2e/markers/smoke-full-pipeline.md`.
- **FR-005**: The work MUST stay in this single repository, and the Plan stage MUST NOT decompose it or spawn child issues.
- **FR-006**: The PR for this work MUST reference this issue so that Fabrik can link them, and it MUST be merged by the pipeline.
- **FR-007**: The issue MUST end in Done with the PR merged.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the last line of `e2e/markers/smoke-full-pipeline.md` on the default branch is exactly `<!-- smoke-full-pipeline-20261025-072306 -->`.
- **SC-002**: The merged PR's diff contains one changed file and, for an existing file, one added line. For a newly created file the diff is that file with the marker line as its only addition.
- **SC-003**: The issue is in Done and closed, and exactly one merged PR is linked to it.
- **SC-004**: The issue passes through every pipeline stage with no human action beyond the initial `fabrik:yolo` label, and spawns zero child issues.

## Assumptions

- The marker file is a Markdown file, so an HTML comment line is valid content and does not affect rendering.
- If the file exists, appending to it is expected. If the existing content has no trailing newline, a newline is added before the marker so the marker sits on its own line.
- The marker string includes the run timestamp `20261025-072306` exactly as written in the request.
- The default base branch is `main`.

## Out of Scope

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repository work or decomposition into child issues.
- Changes to pipeline behaviour, stage configuration or tests.
- Cleaning up marker lines left by earlier smoke runs.
