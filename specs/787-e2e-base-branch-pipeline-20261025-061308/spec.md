# Feature Specification: e2e base-branch pipeline marker (20261025-061308)

**Feature Branch**: `fabrik/issue-787`
**Created**: 2026-10-25
**Status**: Specified
**Input**: User description: "End-to-end verification of the base:<branch> (non-default base branch) pipeline contract — regression coverage for handarbeit/fabrik#1046, validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed) fixes. Trivial change: Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist): `<!-- base-branch-pipeline-20261025-061308 -->`. That is the entire change. One file, one line. Plan should NOT decompose. Scope: Single repo only, targeting the non-default base branch e2e-base-branch-20261025-061308 (already set via the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end probe of the `base:<branch>` pipeline contract, used as regression coverage for handarbeit/fabrik#1046. The earlier fixes it validates are #1047 (issue↔PR linkage for PRs targeting a non-default base branch) and #1050 (the data feed to the review gate for such PRs). The code change is deliberately trivial. Its value is that the item travels through every pipeline stage against the non-default base branch `e2e-base-branch-20261025-061308`, so linkage and review-gate behaviour can be observed. The `base:e2e-base-branch-20261025-061308` label was applied at filing and must not be added or removed.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended on the non-default base (Priority: P1)

An e2e operator wants a uniquely identifiable, harmless change on a PR that targets the non-default base branch. The issue-to-PR linkage and review-gate behaviour can then be checked on that PR.

**Why this priority**: This is the whole change, and the pipeline run it produces is the point of the issue.

**Independent Test**: Inspect `e2e/markers/base-branch-pipeline.md` on the PR branch and confirm it ends with the marker line, with no other file or line changed.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` already exists on the base branch, **When** the change is implemented, **Then** the file's last line is `<!-- base-branch-pipeline-20261025-061308 -->` and all previous content is unchanged.
2. **Given** the file does not exist on the base branch, **When** the change is implemented, **Then** the file is created and contains the marker line.
3. **Given** the PR is opened, **When** its base is inspected, **Then** it targets `e2e-base-branch-20261025-061308`.

---

### Edge Cases

- The file exists but does not end with a newline: the marker must still start on its own line and not be joined to the previous text.
- The marker line is already present (for example on a re-run): it must not be appended a second time.
- The `e2e/markers/` directory does not exist: it is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The file `e2e/markers/base-branch-pipeline.md` MUST end with the exact line `<!-- base-branch-pipeline-20261025-061308 -->`, appended after all existing content.
- **FR-002**: If the file does not exist, it MUST be created containing the marker line.
- **FR-003**: The change MUST consist of exactly one file and one added line; no other files are modified.
- **FR-004**: The marker line MUST appear exactly once in the file.
- **FR-005**: The work MUST target the base branch `e2e-base-branch-20261025-061308`, and the `base:` label MUST NOT be added or removed.
- **FR-006**: The work MUST be handled as a single unit, with no decomposition into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The merged diff against the base branch changes exactly one file, adding exactly one line.
- **SC-002**: A PR exists whose base is `e2e-base-branch-20261025-061308` and which is linked to this issue.
- **SC-003**: The issue progresses through the pipeline stages with no child issues created.

## Assumptions

- The default branch for this repository in this run is `e2e-base-branch-20261025-061308`, so the PR base and the `base:` label agree.
- Whether the file's existing content, if any, carries a trailing newline is not significant; the marker is placed on its own final line.
- A trailing newline after the marker line is acceptable.

## Out of Scope *(optional)*

- Any change to Fabrik engine code, configuration, or tests.
- Adding or removing labels on the issue.
- Changes to any file other than `e2e/markers/base-branch-pipeline.md`.
- Multi-repo scenarios.
- Decomposing the work into sub-issues.
