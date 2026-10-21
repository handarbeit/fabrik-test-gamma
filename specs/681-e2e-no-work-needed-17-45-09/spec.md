# Feature Specification: Verify README presence, test-bed description and license pointer (no-work-needed short-circuit)

**Feature Branch**: `fabrik/issue-681`
**Created**: 2026-10-21
**Status**: Specified
**Input**: User description: "## Goal

Verify that the FABRIK_NO_WORK_NEEDED short-circuit closes the issue without creating a PR.

## What this asks for

`README.md` in this repo already exists and contains a description of the test bed. Please verify that the README.md file is present, contains the substring \"`+`fabrik-test-alpha`+`\", and has the Apache 2.0 LICENSE pointer.

## Hint to Plan

This issue is purely a verification request. There is no code change to make. The Plan stage should recognise that and emit FABRIK_NO_WORK_NEEDED on a line of its own (instead of decomposing or proceeding to Implement).

This is an e2e regression test for #733 (the marker) and #742 (close-on-no-work)."

## Background

This is an end-to-end regression test for the engine's "no work needed" short-circuit (#733, the marker) and close-on-no-work behavior (#742). The scenario: an issue asks only for verification of existing state, so no code or documentation change is required. The pipeline should recognise this, mark the remaining stages complete, and move the issue to Done without creating a PR.

The requested verification covers a README that already exists in the repository: that it is present, that it describes the test bed and mentions the `fabrik-test-alpha` name, and that it points to the Apache 2.0 license. The issue author's hint is that the Plan stage should emit the no-work-needed marker rather than decomposing the issue or proceeding to Implement.

Findings from the current repository state (checked while specifying, not a design decision):
- `README.md` is present and describes the test bed (it is titled for this repository, a second bed paired with another test repository).
- `README.md` has a License section stating Apache 2.0, and a `LICENSE` file is present at the repository root.
- This repository was seeded from `fabrik-test-alpha` (per its git history), but the README text does not currently contain the literal string `fabrik-test-alpha`; it names this bed instead. The substring check is therefore an observation to record in the verification result, not a reason to change files: the issue states there is no code change to make, and the e2e test's purpose is the short-circuit, not README content.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Verification-only issue is closed without a PR (Priority: P1)

An operator files an issue that only asks to verify existing repository state. The pipeline determines that no change is needed and finishes the issue without opening a PR.

**Why this priority**: This is the entire purpose of the issue: regression coverage of the no-work-needed short-circuit and close-on-no-work.

**Independent Test**: Run this issue through the pipeline and confirm that it reaches Done and is closed, that remaining stages are marked skipped, and that no PR or implementation branch commits exist for it.

**Acceptance Scenarios**:

1. **Given** an issue that requests only verification, **When** the Plan stage determines no change is needed and signals no-work-needed together with stage completion, **Then** all remaining stages are marked complete as skipped and the issue moves to Done.
2. **Given** the same issue, **When** it reaches Done, **Then** the issue is closed and no PR was created for it.
3. **Given** the same issue, **When** the Plan stage runs, **Then** it does not decompose the issue into children or advance to Implement.

---

### User Story 2 - Report the README verification result (Priority: P2)

The verification the issue requests is performed and its outcome is reported in the stage output, so a reader can see what was checked.

**Why this priority**: The verification is the stated content of the request, but it is secondary to the short-circuit behavior under test.

**Independent Test**: Read the stage output and confirm it states, for each of the three checks, whether it held.

**Acceptance Scenarios**:

1. **Given** the repository's current state, **When** the verification is performed, **Then** the output states whether `README.md` is present.
2. **Given** the same state, **When** the verification is performed, **Then** the output states whether the README contains the substring `fabrik-test-alpha`.
3. **Given** the same state, **When** the verification is performed, **Then** the output states whether the README carries an Apache 2.0 license pointer.

---

### Edge Cases

- The README does not contain the literal `fabrik-test-alpha` substring (this repository is a different bed seeded from alpha). The result is reported as observed; it does not cause any file to be changed.
- The issue body's quoted substring is wrapped in stray template characters; the intended substring is the bare name `fabrik-test-alpha`.
- Under yolo autonomy the Specify stage advances without a human approval; this must not alter the expected end state (Done, closed, no PR).

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The pipeline MUST treat this issue as verification-only: no code, documentation or README change is made.
- **FR-002**: The Plan stage MUST signal no-work-needed (together with stage completion) instead of decomposing the issue or proceeding to Implement.
- **FR-003**: On no-work-needed, all remaining stages MUST be marked complete as skipped and the issue MUST move to Done and be closed.
- **FR-004**: No PR MUST be created for this issue.
- **FR-005**: The stage output MUST report whether `README.md` exists, whether it contains `fabrik-test-alpha`, and whether it has an Apache 2.0 license pointer.
- **FR-006**: A failed substring check MUST NOT cause the pipeline to make changes or to block; it is reported only.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue ends in Done and closed with zero PRs linked to it.
- **SC-002**: Every stage after the one that signals no-work-needed carries a "skipped" completion comment.
- **SC-003**: The verification result for all three README checks is stated in the pipeline output.

## Assumptions

- The intended substring is the bare name `fabrik-test-alpha`; the surrounding characters in the issue text are a templating artifact.
- The substring result does not affect the expected end state, since the issue states there is no change to make and the test targets the short-circuit.
- The autonomy label (`fabrik:yolo`) is part of the test setup and is expected to carry the issue through stages without manual approval.

## Out of Scope *(optional)*

- Changing `README.md` or `LICENSE`, including adding the alpha name to the README.
- Any engine change to the no-work-needed marker or close-on-no-work behavior (covered by #733 and #742).
- Creating a PR or any implementation branch commits.
