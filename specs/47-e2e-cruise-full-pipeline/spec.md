# Feature Specification: E2E smoke full-pipeline marker (20261014-225751)

**Feature Branch**: `fabrik/issue-47`
**Created**: 2026-10-14
**Status**: Specified
**Input**: User description: "## Goal

End-to-end single-repo pipeline smoke. Verify Fabrik can take an issue from Specify all the way to Done with a merged PR.

## Trivial change

Append a single comment line to `e2e/markers/smoke-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- smoke-full-pipeline-20261014-225751 -->
```

That's the entire change. One file, one line.

## Scope

Single repo only — no cross-repo work. The Plan stage should NOT decompose.

## Labels

fabrik:yolo"

## Background

This issue is an end-to-end smoke test of the Fabrik pipeline in a single repository. Its purpose is to verify that an issue can move through every stage (Specify → Research → Plan → Implement → Review → Validate) to Done with a merged PR. The change itself is deliberately trivial, so that any failure points at the pipeline rather than at the work. The `fabrik:yolo` label is set, so stages auto-advance and the PR is auto-merged on Validate completion.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended and merged (Priority: P1)

An operator running the smoke test wants a uniquely identified marker line recorded in the repository's smoke-test marker file. The line is delivered through a PR that the pipeline merges on its own, which confirms the whole pipeline works.

**Why this priority**: This is the only deliverable. Without it the smoke test proves nothing.

**Independent Test**: After the PR merges, read `e2e/markers/smoke-full-pipeline.md` on the default branch. Its final line is exactly `<!-- smoke-full-pipeline-20261014-225751 -->`.

**Acceptance Scenarios**:

1. **Given** the marker file exists, **When** the change is applied, **Then** the file's last line is `<!-- smoke-full-pipeline-20261014-225751 -->` and its previous content is unchanged.
2. **Given** the marker file does not exist, **When** the change is applied, **Then** the file is created and contains the marker line.
3. **Given** the issue carries `fabrik:yolo`, **When** every stage completes, **Then** the PR is merged and the issue reaches Done without human intervention.

---

### Edge Cases

- The marker file exists but does not end with a newline: the marker still lands on its own line, not appended to the previous line.
- The marker line is already present (for example, on a re-run): it is not duplicated.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change MUST append the line `<!-- smoke-full-pipeline-20261014-225751 -->` as the last line of `e2e/markers/smoke-full-pipeline.md`.
- **FR-002**: If `e2e/markers/smoke-full-pipeline.md` does not exist, the change MUST create it, and it must contain the marker line.
- **FR-003**: The change MUST NOT modify any existing content in the marker file, and MUST NOT touch any other file.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST stay within this single repository, and the Plan stage MUST NOT decompose it into child issues.

### Key Entities

- **Marker file**: `e2e/markers/smoke-full-pipeline.md`, a record of smoke-test runs. Each run contributes one uniquely identified HTML comment line.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: After merge, the default branch's marker file ends with exactly `<!-- smoke-full-pipeline-20261014-225751 -->`.
- **SC-002**: The PR's diff touches exactly one file and adds exactly one line (plus the file's creation if it was absent).
- **SC-003**: The issue reaches Done with its PR merged and no manual action after the label was applied.

## Assumptions

- The marker is a single HTML comment line, so it does not change how the file renders.
- If the file is created, it holds only the marker line, since no further content is specified.
- Appending a newline-terminated line is the expected format, per the usual text-file convention.
- This repo's merge settings allow the PR to merge automatically under `fabrik:yolo`.

## Out of Scope *(optional)*

- Any cross-repo change.
- Any change beyond the single marker line, including edits to other files, docs, tests or code.
- Decomposing the work into child issues.
- Changes to pipeline or stage configuration.
