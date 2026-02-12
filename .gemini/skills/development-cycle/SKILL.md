---
name: development-cycle
description: Use for the entire development process including TDD, CI, Review, and Commit
---

# Development Cycle

## Overview

This skill defines the end-to-end development process for all tasks. Adhere strictly to these steps to ensure high quality and consistency.

## The Cycle

1.  **TDD Development**:
    - Always follow the `test-driven-development` skill.
    - Write a failing test before any production code.
2.  **CI Check**:
    - Run `just ci` to ensure all checks (lint, tests, etc.) pass locally.
3.  **Code Review**:
    - Follow the **`review-rule`** skill.
    - Perform a self-review or use the `go-code-review` skill.
    - Output review artifacts to `/.gemini/reviews/<commit-hash|timestamp>.md`.
    - **Score**: Must be **> 95**.
4.  **Planning to Deal with Review**:
    - Follow the **`planning-to-reflect-review`** skill.
    - Analyze the score and feedback.
    - Prioritize fixing high-impact deductions first.
5.  **Repeat**:
    - Iterate steps 1-4 until the review score is **> 95**.
6.  **Write Docs**:
    - Update relevant documentation in `docs/` or `README.md`.
7.  **Ask User's Review**:
    - Present the final state and the review score to the user for confirmation.
8.  **Commit**:
    - Follow the **`commit-rule`** skill.

## Related Skills

- **test-driven-development**: Core TDD rules.
- **review-rule**: Scoring and deductions (-1 to -10).
- **planning-to-reflect-review**: Strategy for addressing feedback.
- **commit-rule**: Gitmoji, issue linking, and co-author attribution.
