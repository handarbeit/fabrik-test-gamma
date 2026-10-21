# Feature Specification: Improve the README

**Feature Branch**: `fabrik/issue-684`
**Created**: 2026-10-21
**Status**: Draft
**Input**: User description: "Make the README \"better\"."

## Background

This issue is filed by the e2e harness to verify the `FABRIK_BLOCKED_ON_INPUT` pause and resume behaviour. The request is deliberately ambiguous. It should make the Specify stage stop and ask for clarification instead of proceeding to Research on guesses. The engine should then pause the issue with `fabrik:paused` and `fabrik:awaiting-input`, and the harness resumes it by posting a clarifying comment.

Regression coverage: #733 (the `FABRIK_BLOCKED_ON_INPUT` marker) and ed46b7fc (`fabrik:awaiting-input` clears on `FABRIK_STAGE_COMPLETE`).

The repository's `README.md` is about 250 lines. It opens with a long run of HTML comment markers from earlier e2e convergence-race runs. "Better" could mean any of several things, such as readability, accuracy or completeness. The request does not say which.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - README improvement along an agreed dimension (Priority: P1)

A maintainer wants the README to be better in a specific, agreed way. The way has to be stated before work starts.

**Why this priority**: Without a definition of "better", no requirement can be written or tested.

**Independent Test**: Once the open questions are answered, a reviewer can compare the README before and after and check each change against the agreed criteria.

**Acceptance Scenarios**:

1. **Given** the agreed definition of "better", **When** the README is revised, **Then** every change maps to a stated criterion.
2. **Given** content outside the agreed scope, **When** the README is revised, **Then** that content is left unchanged.

---

### Edge Cases

- The leading HTML comment markers may be machine-written test fixtures that other tooling depends on. They must not be removed without confirmation.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The README MUST be improved along the dimension(s) the requester confirms [NEEDS CLARIFICATION: which dimension(s) of "better" apply?]
- **FR-002**: The change MUST NOT remove or alter the HTML comment markers at the top of the README unless the requester explicitly confirms [NEEDS CLARIFICATION: may the markers be removed or moved?]

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: A reviewer can point to each README change and the agreed criterion it satisfies.
- **SC-002**: No content outside the agreed scope is modified.

## Assumptions

- Only `README.md` is in scope, not other documentation.
- The README's existing factual content stays accurate and is not rewritten for style alone.

## Out of Scope *(optional)*

- Changes to code, tests, or documentation other than `README.md`.
