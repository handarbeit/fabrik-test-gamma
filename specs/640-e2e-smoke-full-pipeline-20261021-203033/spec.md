# Feature Specification: e2e smoke full-pipeline (20261021-203033)

**Feature Branch**: `fabrik/issue-640`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261021-203033 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. The goal is to confirm that Fabrik can carry an issue through every stage (Specify, Research, Plan, Implement, Review, Validate) to Done with a merged PR. The content change is deliberately trivial, so that any failure points at the pipeline rather than at the change itself. The issue carries `fabrik:yolo`, so stages auto-advance and the PR is auto-merged when Validate completes.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line lands on the base branch via the full pipeline (Priority: P1)

An operator files this issue and applies `fabrik:yolo`. Fabrik runs it through all stages unattended. When the pipeline finishes, the marker file on the base branch ends with the unique marker comment line, and the issue is Done.

**Why this priority**: This is the only story. It is the whole purpose of the smoke test.

**Independent Test**: After the pipeline completes, read `e2e/markers/smoke-full-pipeline.md` on `main` and confirm its last line is `<!-- smoke-full-pipeline-20261021-203033 -->`. Confirm the linked PR is merged and the issue is in Done.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist on the base branch, **When** the pipeline completes, **Then** the file exists and contains the marker line as its last line.
2. **Given** the marker file already exists, **When** the pipeline completes, **Then** the file's prior content is unchanged and the marker line is appended at the very end.
3. **Given** the issue has `fabrik:yolo`, **When** Validate completes, **Then** the linked PR is merged and the issue moves to Done without human intervention.

---

### Edge Cases

- The file exists but its last line has no trailing newline: the marker must still appear on its own line, not joined to the previous content.
- The marker line is already present (for example, a re-run after a partial success): it must not be duplicated.
- Plan must not decompose this work into child issues; there is nothing to split.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/smoke-full-pipeline.md` MUST end with the line `<!-- smoke-full-pipeline-20261021-203033 -->`.
- **FR-002**: If the file does not exist, it MUST be created (including any missing parent directory) with the marker line as its content.
- **FR-003**: If the file exists, its existing content MUST be preserved and the marker line appended once, at the very end.
- **FR-004**: The change MUST touch exactly one file and add exactly one line, apart from engine-managed artifacts such as the committed spec file.
- **FR-005**: The work MUST stay within this single repository, and Plan MUST NOT decompose it into child issues.
- **FR-006**: The change MUST be delivered through a PR that references this issue with `Closes #640` and is merged into the default branch `main`.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, an append-only log of smoke-run marker comments.
- **Marker line**: the HTML comment `<!-- smoke-full-pipeline-20261021-203033 -->`, unique to this run.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After the run, the last line of `e2e/markers/smoke-full-pipeline.md` on `main` is exactly the marker line, and it appears exactly once.
- **SC-002**: The PR diff for the change, apart from engine-managed spec files, is one file with one added line.
- **SC-003**: The issue reaches Done with its linked PR merged, with no human action after the `fabrik:yolo` label.
- **SC-004**: No child issues are created.

## Assumptions

- The marker file may or may not already exist; the issue explicitly allows creating it.
- The marker line is a single line of text terminated by a newline.
- `fabrik:yolo` is already applied, so no human approval is needed between stages.

## Out of Scope *(optional)*

- Any cross-repo work.
- Any change other than the single marker line.
- Decomposition into child or follow-up issues.
- Cleanup or deduplication of other smoke-run markers already in the file.
