# Feature Specification: Make the README "better"

**Feature Branch**: `fabrik/issue-156`
**Created**: 2026-10-17
**Status**: Draft
**Input**: User description: "## Goal

Verify FABRIK_BLOCKED_ON_INPUT pause + resume.

## Deliberately ambiguous

Make the README \"better\".

That is all. Specify should immediately notice this is wildly under-specified and emit FABRIK_BLOCKED_ON_INPUT asking for clarification (instead of proceeding to Research with assumptions). The engine should then pause the issue with `fabrik:paused` and `fabrik:awaiting-input`.

This issue is filed by the e2e harness, which will post a clarifying comment to resume.

Regression coverage:
- #733 / FABRIK_BLOCKED_ON_INPUT marker
- ed46b7fc — fabrik:awaiting-input clears on FABRIK_STAGE_COMPLETE"

## Background

This issue is an end-to-end check of the `FABRIK_BLOCKED_ON_INPUT` pause/resume path. The request ("Make the README \"better\"") is deliberately under-specified. The Specify stage is expected to ask for clarification and pause the issue (`fabrik:paused` + `fabrik:awaiting-input`) rather than guess and move on to Research. The e2e harness will answer with a comment, which resumes the issue.

Observation: the current `README.md` is about 250 lines. It begins with many `convergence-race-*` HTML comment markers, which look like artifacts left by other e2e scenarios. "Better" could mean cleaning those up, fixing content or structure, or something else entirely. The request gives no basis for choosing.

## User Scenarios & Testing *(mandatory)*

### User Story 1 - Improve the README along a defined dimension (Priority: P1)

A maintainer or reader of the repository gets a README that is better in the specific way the requester defines in the answers to the open questions.

**Why this priority**: This is the whole purpose of the issue. It cannot be scoped until "better" is defined.

**Independent Test**: A reviewer can compare the README before and after and check each change against the agreed criteria.

**Acceptance Scenarios**:

1. **Given** the criteria for "better" have been answered, **When** the README is updated, **Then** every change maps to one of those criteria and nothing outside them is changed.

---

### Edge Cases

- Existing HTML comment markers in the README may be relied on by other e2e tests, so removing them could break those tests.

## Requirements *(mandatory)*

### Functional Requirements

- **FR-001**: The README MUST be improved along the dimension(s) of "better" the requester specifies. [NEEDS CLARIFICATION: which dimension(s) of "better" apply? See Open Questions.]
- **FR-002**: The README's existing `<!-- convergence-race-* -->` HTML comment markers MUST be handled as the requester specifies. [NEEDS CLARIFICATION: keep or remove them?]

## Success Criteria *(mandatory)*

### Measurable Outcomes

- **SC-001**: Each agreed improvement criterion can be checked, with a yes or no, against the updated README.

## Assumptions

- Nothing in the README is changed until the criteria are clarified.

## Out of Scope *(optional)*

- Changes to code, tests, or any file other than the README, unless the requester says otherwise.
