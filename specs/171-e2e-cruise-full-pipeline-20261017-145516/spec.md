# Feature Specification: e2e cruise full pipeline marker (20261017-145516)

**Feature Branch**: `fabrik/issue-171`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

End-to-end verification of the fabrik:cruise pipeline contract (#898).

## Trivial change

Append a single HTML comment line to `e2e/markers/cruise-full-pipeline.md` at the very end of the file (create the file first if it doesn't already exist):

```
<!-- cruise-pipeline-20261017-145516 -->
```

That is the entire change. One file, one line. Plan should NOT decompose.

## Scope

Single repo only. This issue verifies that cruise auto-advances through all
stages to Validate-complete without merging the PR, and that the issue closes
correctly after a human merges the PR."

## Background

This issue is an end-to-end verification of the `fabrik:cruise` pipeline contract (#898). The code change is deliberately trivial so that the pipeline's behaviour, not the content of the change, is what is under test. The contract being verified is that an issue labelled `fabrik:cruise` auto-advances through every stage up to Validate-complete without the engine merging the PR, and that the issue closes correctly once a human merges the PR.

The change itself is a marker line appended to a tracking file, `e2e/markers/cruise-full-pipeline.md`, which is created first if it does not exist. The line is unique to this run, as the timestamp `20261017-145516` shows, so the run's effect can be identified unambiguously.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Marker line is recorded in the tracking file (Priority: P1)

The pipeline produces a change that appends one HTML comment line, `<!-- cruise-pipeline-20261017-145516 -->`, to the end of `e2e/markers/cruise-full-pipeline.md`, creating the file if it is absent.

**Why this priority**: This is the only deliverable of the issue. It gives the pipeline a real, reviewable diff to carry through every stage.

**Independent Test**: Inspect the PR diff. It touches only `e2e/markers/cruise-full-pipeline.md`, and the file's last line is the exact marker comment.

**Acceptance Scenarios**:

1. **Given** the marker file does not exist, **When** the change is implemented, **Then** the file exists and contains the marker line as its final line.
2. **Given** the marker file already exists with prior content, **When** the change is implemented, **Then** the prior content is unchanged and the marker line is appended after it as the final line.
3. **Given** the change is complete, **When** the PR diff is reviewed, **Then** exactly one file is changed and only one line is added.

---

### User Story 2 - Cruise carries the issue to Validate-complete without merging (Priority: P1)

With the `fabrik:cruise` label present, the issue advances automatically through every pipeline stage to Validate-complete, and the linked PR is left unmerged for a human.

**Why this priority**: This is the purpose of the verification, the contract from #898.

**Independent Test**: Observe the issue's progress after it enters the pipeline. It reaches Validate-complete with no manual stage approvals, and the PR is open and unmerged at that point.

**Acceptance Scenarios**:

1. **Given** the issue carries `fabrik:cruise`, **When** each stage completes, **Then** the issue advances to the next stage without human approval.
2. **Given** Validate completes, **When** the engine reaches its end-of-Validate decision, **Then** the PR is not merged and the issue does not move to Done.
3. **Given** Plan receives this issue, **When** it produces the plan, **Then** it does not decompose the work into child issues.

---

### User Story 3 - Issue closes correctly after a human merges the PR (Priority: P2)

After the pipeline stops at Validate-complete, a human merges the PR, and the issue then closes and reaches its final state.

**Why this priority**: It completes the contract by showing the post-merge side of the cruise flow works, but it depends on stories 1 and 2.

**Independent Test**: Merge the PR by hand. The issue closes and moves to Done.

**Acceptance Scenarios**:

1. **Given** the PR is open at Validate-complete, **When** a human merges it, **Then** the linked issue closes and reaches Done.

---

### Edge Cases

- The marker file already ends without a trailing newline. The marker must still land on its own line rather than being joined to the previous content.
- The marker file already contains this exact marker line (for example, a re-run). The change must not produce a duplicate line.
- The `e2e/markers/` directory does not exist yet. It is created along with the file.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The change appends exactly one line, `<!-- cruise-pipeline-20261017-145516 -->`, to the very end of `e2e/markers/cruise-full-pipeline.md`.
- **FR-002**: If `e2e/markers/cruise-full-pipeline.md` does not exist, it is created, containing the marker line.
- **FR-003**: Existing content in the marker file is left unchanged.
- **FR-004**: The change touches no file other than `e2e/markers/cruise-full-pipeline.md`.
- **FR-005**: The Plan stage does not decompose this issue into child issues.
- **FR-006**: With `fabrik:cruise` present, the issue advances through all stages to Validate-complete without manual approval, and the engine does not merge the linked PR.
- **FR-007**: After a human merges the PR, the issue closes correctly.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The PR diff contains exactly one changed file and exactly one added line, and that line is the specified marker comment.
- **SC-002**: The issue reaches Validate-complete with zero manual stage approvals.
- **SC-003**: At Validate-complete the linked PR is open and unmerged.
- **SC-004**: Within one poll cycle of a human merging the PR, the issue is closed.
- **SC-005**: No child issues are created.

## Assumptions

- The marker file may or may not already exist. The issue explicitly allows either case, and the behaviour is defined for both.
- The marker line is appended as its own line, terminated by a newline, so the file stays well-formed for later appends.
- This issue is a verification fixture, so no documentation or user-facing behaviour changes accompany it.

## Out of Scope

- Any change to engine, stage, or pipeline behaviour. This issue only verifies the existing cruise contract.
- Any file other than `e2e/markers/cruise-full-pipeline.md`.
- Decomposing the work into child issues.
- Merging the PR automatically. A human does that.
