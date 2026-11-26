# Feature Specification: E2E smoke full-pipeline (20261126-152804)

**Feature Branch**: `fabrik/issue-2682`
**Created**: 2026-11-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261126-152804 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to confirm that an issue can move through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done, with its pull request merged and without human intervention. The `fabrik:yolo` label requests automatic advancement through the stages and an automatic merge when Validate completes.

The code change is deliberately minimal so that any failure points at the pipeline, not at the change. The change is a single marker comment line appended to a marker file. The timestamp in the marker makes this run's line unique and identifiable.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is added to the smoke marker file (Priority: P1)

An operator running the pipeline smoke wants the repository to contain a unique marker line for this run, so they can confirm the change travelled through the pipeline and was merged.

**Why this priority**: This is the only functional change. It is the observable proof that the pipeline completed.

**Independent Test**: After the change is merged, open `e2e/markers/smoke-full-pipeline.md` on the default branch and check that its last line is `<!-- smoke-full-pipeline-20261126-152804 -->`.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/smoke-full-pipeline.md` exists, **When** the change is applied, **Then** the file's final line is `<!-- smoke-full-pipeline-20261126-152804 -->` and all earlier content is unchanged.
2. **Given** `e2e/markers/smoke-full-pipeline.md` does not exist, **When** the change is applied, **Then** the file is created and its content is the marker line `<!-- smoke-full-pipeline-20261126-152804 -->`.

---

### User Story 2 - Issue completes the full pipeline to Done (Priority: P1)

An operator wants this single-repo issue to move from Specify to Done automatically, ending with a merged pull request.

**Why this priority**: This is the purpose of the smoke test. The marker line is only the payload.

**Independent Test**: Watch the issue after `fabrik:yolo` is applied: it should pass through every stage without manual advancement, and its linked pull request should end up merged with the issue in Done.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:yolo`, **When** each stage completes, **Then** the issue advances to the next stage without a human approving it.
2. **Given** Validate completes successfully, **When** the engine reaches the landing decision, **Then** the linked pull request is merged and the issue reaches Done.
3. **Given** the Plan stage runs, **When** it finishes, **Then** it produces a single-repo plan and spawns no child issues.

---

### Edge Cases

- If the marker file already exists and its last line has no trailing newline, the marker must still go on its own line, not be appended to the previous line.
- If the marker file already contains this exact marker line, it must not be duplicated.
- The `e2e/markers/` directory may not exist yet and should be created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The repository MUST contain the line `<!-- smoke-full-pipeline-20261126-152804 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` already exists, the change MUST only append the marker line and MUST leave all existing content unchanged.
- **FR-003**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it with the marker line as its content.
- **FR-004**: The change MUST touch exactly one file and add exactly one line of content.
- **FR-005**: The work MUST stay within this repository, and the Plan stage MUST NOT decompose the work into child issues.
- **FR-006**: The linked pull request MUST be merged and the issue MUST reach Done through the pipeline, with no manual stage advancement.

### Key Entities *(if applicable)*

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a text file that accumulates one marker comment line per smoke run.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261126-152804 -->`, which identifies this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the pull request merges, the final line of `e2e/markers/smoke-full-pipeline.md` on the default branch equals `<!-- smoke-full-pipeline-20261126-152804 -->`.
- **SC-002**: The merged pull request's diff contains one changed file with one added content line.
- **SC-003**: The issue reaches Done with a merged pull request and no child issues created.
- **SC-004**: The issue passes through all pipeline stages with no manual advancement.

## Assumptions

- The marker line must be its own line, ending with a trailing newline unless the repository's file convention says otherwise.
- The spec file that the engine writes during Specify is part of the pipeline's normal output. It does not count against the "one file, one line" limit, which applies to the implementation change.
- `fabrik:yolo` is already on the issue, so no stage waits for human approval.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change other than the single marker line.
- Changes to pipeline behaviour, stage configuration or documentation.
