# Feature Specification: Verify README.md contents (e2e no-work-needed)

**Feature Branch**: `fabrik/issue-245`
**Created**: 2026-10-17
**Status**: Specified
**Input**: User description: "## Goal\n\nVerify that the FABRIK_NO_WORK_NEEDED short-circuit closes the issue without creating a PR.\n\n## What this asks for\n\n`README.md` in this repo already exists and contains a description of the test bed. Please verify that the README.md file is present, contains the substring \"`+`fabrik-test-alpha`+`\", and has the Apache 2.0 LICENSE pointer.\n\n## Hint to Plan\n\nThis issue is purely a verification request. There is no code change to make. The Plan stage should recognise that and emit FABRIK_NO_WORK_NEEDED on a line of its own (instead of decomposing or proceeding to Implement).\n\nThis is an e2e regression test for #733 (the marker) and #742 (close-on-no-work)."

## Background

This is an end-to-end regression scenario for the pipeline's "no work needed" short-circuit (#733, the marker; #742, close-on-no-work). It asks for a verification of the repository's `README.md` rather than any change. The expected outcome is that the pipeline recognises that nothing needs to be built, skips the remaining stages without creating a PR, and moves the issue to Done.

Observations from the current repository state (a read of the file, not an implementation review):

- `README.md` is present at the repository root.
- It has a `## License` section reading "Apache 2.0 (matches Fabrik upstream)", and a `LICENSE` file exists at the root.
- It does not contain the literal string `fabrik-test-alpha`. This repository is `fabrik-test-gamma`; its README names `fabrik-test-delta` as its paired bed. The substring in the request is also rendered with stray `+` characters around it ("`+`fabrik-test-alpha`+`"), which looks like a templating artifact.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Pipeline closes a verification-only issue without a PR (Priority: P1)

An operator files an issue that asks only for a verification of existing repository state. The pipeline determines that no code or documentation change is required and closes the issue without opening a PR.

**Why this priority**: This is the whole point of the scenario. It is the behaviour under regression test.

**Independent Test**: Run the issue through the pipeline with `fabrik:yolo`. Confirm that a later stage (expected to be Plan) signals that no work is needed, the remaining stages are marked skipped, the issue ends in Done and closed, and no PR was created.

**Acceptance Scenarios**:

1. **Given** an issue that asks only to verify existing file contents, **When** the pipeline reaches the stage that evaluates whether work is needed, **Then** it signals no work needed and does not decompose the issue or proceed to Implement.
2. **Given** the no-work-needed signal, **When** the engine processes it, **Then** all remaining stages are marked complete with "skipped" comments and the issue moves to Done and is closed.
3. **Given** the same run, **When** it finishes, **Then** no pull request exists for this issue.

---

### User Story 2 - Report the README verification result (Priority: P2)

The operator wants to know whether `README.md` meets the three stated conditions, so that the closing record is accurate.

**Why this priority**: The request names these checks explicitly, but the outcome the test cares about is the short-circuit, not the check result.

**Independent Test**: Read `README.md` and `LICENSE` and compare against the three conditions.

**Acceptance Scenarios**:

1. **Given** the repository at the issue's base, **When** `README.md` is checked, **Then** the file is confirmed present.
2. **Given** the same README, **When** it is checked for the Apache 2.0 license pointer, **Then** the `## License` section stating Apache 2.0 is confirmed, along with the root `LICENSE` file.
3. **Given** the same README, **When** it is checked for `fabrik-test-alpha`, **Then** the result is reported as it is, which is that the substring is not present in this repository's README. The repository is not altered to make the check pass.

---

### Edge Cases

- The `fabrik-test-alpha` substring check fails against the current README. The scenario still expects the no-work-needed outcome, so this result must not trigger an Implement stage or a PR.
- The substring in the request is rendered with stray `+` characters. The intended string is read as `fabrik-test-alpha`.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The pipeline MUST treat this issue as verification-only and make no changes to repository files.
- **FR-002**: The pipeline MUST signal that no work is needed (expected at Plan, per the issue's hint) rather than decomposing the issue or proceeding to Implement.
- **FR-003**: On a no-work-needed signal, the engine MUST mark all remaining stages complete with "skipped" comments, move the issue to Done and close it.
- **FR-004**: No pull request MUST be created for this issue.
- **FR-005**: The verification result for the three README conditions (present, contains `fabrik-test-alpha`, Apache 2.0 license pointer) MUST be reported accurately, including any condition that does not hold.

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: The issue ends in Done and closed with zero pull requests linked to it.
- **SC-002**: Every stage after the one that signalled no work needed carries a "skipped" comment.
- **SC-003**: The recorded verification result states, for each of the three README conditions, whether it holds.

## Assumptions

- The intent of the scenario is to exercise the no-work-needed short-circuit. The README substring check is incidental, and a failed check does not change the expected pipeline outcome.
- The string in the request is `fabrik-test-alpha`. The surrounding `+` characters are a templating artifact.
- The Specify stage does not itself signal no work needed. The issue's hint assigns that to Plan, so Specify only writes the spec and completes.
- The mismatch between `fabrik-test-alpha` and this repository (`fabrik-test-gamma`) is left as is. Editing the README to add the string would contradict the "no code change" statement.

## Out of Scope

- Editing `README.md` or `LICENSE`.
- Changing the engine's no-work-needed handling, which this scenario tests but does not modify.
- Creating a PR, branch changes or follow-up issues.

## Source References

- Issue #245 body (above).
- `README.md` (`## License` section) and `LICENSE` in the repository root.
- ADR/issue references named in the request: #733 (marker), #742 (close-on-no-work).
