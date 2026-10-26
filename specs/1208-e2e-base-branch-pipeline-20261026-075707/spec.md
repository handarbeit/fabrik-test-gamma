# Feature Specification: e2e base-branch pipeline marker (20261026-075707)

**Feature Branch**: `fabrik/issue-1208`
**Created**: 2026-10-26
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the base:<branch> (non-default base branch)
pipeline contract — regression coverage for handarbeit/fabrik#1046,
validating the #1047 (issue<->PR linkage) and #1050 (review-gate data feed)
fixes.

## Trivial change

Append a single HTML comment line to `e2e/markers/base-branch-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- base-branch-pipeline-20261026-075707 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only, targeting the non-default base branch e2e-base-branch-20261026-075707 (already set via
the base:<branch> label applied at filing — do not add or remove labels)."

## Background

This issue is a live end-to-end check of the `base:<branch>` pipeline contract: an issue carrying a `base:<branch>` label must be worked on a branch forked from that base, produce a PR targeting it, and be carried through the pipeline. It is regression coverage for handarbeit/fabrik#1046 and validates the fixes in #1047 (issue↔PR linkage for a PR on a non-default base) and #1050 (review-gate data feed). The change itself is deliberately trivial so that any failure is attributable to the base-branch pipeline machinery and not to the content of the change.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is appended on the non-default base (Priority: P1)

The pipeline appends one HTML comment line to the base-branch marker file, producing a PR that targets the non-default base branch `e2e-base-branch-20261026-075707`.

**Why this priority**: This is the whole change; it is the artifact the e2e run exercises.

**Independent Test**: Inspect the PR diff and the marker file on the PR branch.

**Acceptance Scenarios**:

1. **Given** `e2e/markers/base-branch-pipeline.md` does not exist on the base branch, **When** the change is made, **Then** the file exists and its last line is `<!-- base-branch-pipeline-20261026-075707 -->`.
2. **Given** the file already exists, **When** the change is made, **Then** the line is appended at the very end and existing content is unchanged.
3. **Given** the issue carries the `base:e2e-base-branch-20261026-075707` label, **When** the PR is created, **Then** the PR targets that branch and is linked to this issue.

---

### Edge Cases

- The file exists without a trailing newline: the marker must still land on its own line at the end of the file.
- The file already contains the marker line (e.g. a re-run): it must not be duplicated.
- Labels on the issue must not be added or removed as part of the work.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: `e2e/markers/base-branch-pipeline.md` MUST end with the exact line `<!-- base-branch-pipeline-20261026-075707 -->`, appended at the very end of the file; the file is created first if it does not exist.
- **FR-002**: The change MUST consist of exactly one file and one added line, with no other files or content modified.
- **FR-003**: The work MUST be based on, and its PR MUST target, the branch `e2e-base-branch-20261026-075707`, as set by the existing `base:` label.
- **FR-004**: The work MUST NOT add or remove any issue labels, and MUST NOT be decomposed into child issues.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff against `e2e-base-branch-20261026-075707` contains exactly one changed file with exactly one added line, the marker above.
- **SC-002**: The PR is linked to this issue and targets the non-default base branch, with the issue progressing through the pipeline stages without manual intervention.

## Assumptions

- The marker file may not exist yet; creating it with just the marker line is acceptable.
- The marker line is the exact text given, with no surrounding whitespace or extra blank lines beyond a terminating newline.

## Out of Scope

- Any change to engine code, stage configuration or documentation.
- Adding, removing or changing labels.
- Multi-repo scenarios or decomposition into child issues.
