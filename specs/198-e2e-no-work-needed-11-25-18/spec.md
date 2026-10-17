# Feature Specification: e2e no-work-needed verification

**Feature Branch**: `fabrik/issue-198`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal

Verify that the FABRIK_NO_WORK_NEEDED short-circuit closes the issue without creating a PR.

## What this asks for

`README.md` in this repo already exists and contains a description of the test bed. Please verify that the README.md file is present, contains the substring "`+`fabrik-test-alpha`+`", and has the Apache 2.0 LICENSE pointer.

## Hint to Plan

This issue is purely a verification request. There is no code change to make. The Plan stage should recognise that and emit FABRIK_NO_WORK_NEEDED on a line of its own (instead of decomposing or proceeding to Implement).

This is an e2e regression test for #733 (the marker) and #742 (close-on-no-work)."

## Background

This issue is an end-to-end regression test for the engine's "no work needed" short-circuit (#733, the `FABRIK_NO_WORK_NEEDED` marker) and its close-on-no-work behaviour (#742). The requested work is a read-only verification of facts about `README.md`; the artifact already exists, so no code or documentation change is wanted. The purpose is to confirm that a stage which finds nothing to do can end the pipeline early: all remaining stages are marked complete as skipped, and the issue moves to Done and is closed without a PR being created.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline short-circuits on a verification-only issue (Priority: P1)

An operator files an issue that only asks for a verification of existing state. A stage confirms that nothing needs changing and signals "no work needed". The issue ends in Done and is closed, with no branch work turned into a PR.

**Why this priority**: This is the whole point of the issue — it exercises the short-circuit and close-on-no-work behaviour end to end.

**Independent Test**: Run the issue through the pipeline and observe the final state: remaining stages marked skipped, the issue in Done and closed, and no pull request linked to it.

**Acceptance Scenarios**:

1. **Given** an issue that requests only verification of existing files, **When** a stage emits the no-work-needed signal together with stage completion, **Then** all remaining pipeline stages are marked complete as skipped.
2. **Given** the short-circuit has fired, **When** the engine finishes processing, **Then** the issue is moved to Done and closed.
3. **Given** the short-circuit has fired, **When** the engine finishes processing, **Then** no pull request has been created for the issue.

---

### User Story 2 - README verification is recorded (Priority: P2)

The verification the issue describes is carried out and its outcome is stated, so the "no work needed" decision is justified rather than assumed.

**Why this priority**: It supports the P1 story by grounding the decision, but the e2e outcome does not depend on the verification result.

**Independent Test**: Inspect `README.md` and the repository's license file and compare with the three stated checks.

**Acceptance Scenarios**:

1. **Given** the repository at its current state, **When** the README is inspected, **Then** `README.md` is present.
2. **Given** the README, **When** it is searched for the substring `fabrik-test-alpha`, **Then** the result (present or absent) is reported.
3. **Given** the README, **When** its license section is read, **Then** it points to Apache 2.0, consistent with the repository's `LICENSE` file.

---

### Edge Cases

- The README is found to lack a checked item (see the observation under Assumptions): the verification is reported as such and still requires no change to the repository, since the issue asks only to verify.
- Any later stage produces a code change or opens a PR: this contradicts the issue's intent and would count as a failure of the short-circuit.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The verification confirms that `README.md` exists at the repository root.
- **FR-002**: The verification reports whether `README.md` contains the substring `fabrik-test-alpha`.
- **FR-003**: The verification confirms that `README.md` contains an Apache 2.0 license pointer.
- **FR-004**: The issue results in no code, documentation, or configuration change to the repository.
- **FR-005**: A stage signals no-work-needed (together with stage completion) so that the remaining stages are skipped.
- **FR-006**: After the short-circuit, the issue is in Done and closed, and no pull request is created for it.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue reaches Done and is closed without any manual intervention beyond filing it.
- **SC-002**: Zero pull requests are linked to the issue at the end of the run.
- **SC-003**: Every pipeline stage after the one that signalled no-work-needed is recorded as skipped.
- **SC-004**: The outcome of each of the three README checks is stated explicitly.

## Assumptions

- The test is about the short-circuit behaviour, not about the README content; the README checks are descriptive and do not gate the outcome.
- Observation made while specifying: `README.md` exists and has an Apache 2.0 license section (matching the `LICENSE` file), but at the time of writing it does not contain the substring `fabrik-test-alpha`; it refers to this bed as `fabrik-test-gamma` and its pair as `fabrik-test-delta`. The check is carried over from a bed with a different name, and because the issue says nothing needs changing, no README edit is proposed. That discrepancy is reported, not fixed.
- The `fabrik:yolo` label means this stage completes without waiting for human approval.

## Out of Scope

- Editing `README.md` or `LICENSE`, including adding the `fabrik-test-alpha` text.
- Any change to engine behaviour for the marker (#733) or close-on-no-work (#742).
- Creating a branch, commit with content changes, or pull request for this issue.

## Source References

- #733 — `FABRIK_NO_WORK_NEEDED` marker
- #742 — close-on-no-work behaviour
