# Feature Specification: E2E smoke full-pipeline marker line (20261117-043757)

**Feature Branch**: `fabrik/issue-2227`
**Created**: 2026-11-16
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261117-043757 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can travel from Specify through Research, Plan, Implement, Review and Validate to Done, ending in a merged PR, with no human intervention (the `fabrik:yolo` label requests automatic advancement and merge). The code change is deliberately trivial so that any failure points at the pipeline, not at the work. Each run uses a unique timestamped marker line (here `20261117-043757`) so that the run's change is distinguishable from earlier runs and the merged result can be checked.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator files this issue with `fabrik:yolo`. Fabrik carries it through every pipeline stage and merges a PR whose only effect is that the smoke-test marker file ends with the run's marker line.

**Why this priority**: This is the whole purpose of the issue; it is the only story.

**Independent Test**: After the issue reaches Done, read `e2e/markers/smoke-full-pipeline.md` on the default branch (`main`) and confirm that its last line is `<!-- smoke-full-pipeline-20261117-043757 -->`. Confirm the PR is merged and the issue is closed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` already exists, **When** the change is implemented, **Then** the exact line `<!-- smoke-full-pipeline-20261117-043757 -->` is appended at the end of the file and all existing content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is implemented, **Then** the file is created and contains the line `<!-- smoke-full-pipeline-20261117-043757 -->`.
3. **Given** the issue carries `fabrik:yolo`, **When** Validate completes, **Then** the PR is merged automatically and the issue moves to Done without manual action.
4. **Given** the work is a single one-line change in one repo, **When** the Plan stage runs, **Then** it does not decompose the work into child issues.

---

### Edge Cases

- The marker file exists but its last line has no trailing newline: the marker line must still appear on its own line, not joined to the previous content.
- The marker file already contains the same marker line (for example after a retry): the line must not be duplicated.
- The marker file exists with lines from earlier smoke runs: those lines are preserved, and the new line goes after them.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261117-043757 -->` after the change is merged.
- **FR-002**: If the file does not exist, the change MUST create it; if it exists, the change MUST only append, leaving all prior content unchanged.
- **FR-003**: The marker line MUST appear exactly once in the file and on its own line.
- **FR-004**: The change MUST consist of exactly one file and one added line, and touch no other repository.
- **FR-005**: The work MUST NOT be decomposed into child issues; it is handled as a single issue through all stages.
- **FR-006**: The issue MUST reach Done with its PR merged into the default branch (`main`).

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a Markdown file accumulating one HTML-comment marker line per smoke run.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261117-043757 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pipeline finishes, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` equals `<!-- smoke-full-pipeline-20261117-043757 -->`, and that string occurs exactly once in the file.
- **SC-002**: The merged PR's diff contains exactly one added line outside of the Specify stage's spec file, plus the file creation if the file did not exist beforehand.
- **SC-003**: The issue is closed and in Done with its PR merged, with no manual intervention after the issue was filed.
- **SC-004**: No child issues are created for this work.

## Assumptions

- The default base branch is `main`, and the PR targets it.
- Appending to an existing file means adding a new final line, preserving a single newline between it and the previous content.
- The `fabrik:yolo` label provides the automatic advancement and merge; no manual approval of the spec is expected.
- The spec file the engine commits under `specs/` for this issue is part of the normal Specify output and is not counted as part of the one-line change.

## Out of Scope *(optional)*

- Any change to files other than `e2e/markers/smoke-full-pipeline.md`.
- Cross-repo work or decomposition into child issues.
- Changes to Fabrik's engine, stage configuration or documentation.
- Cleaning up or rewriting marker lines from earlier smoke runs.
