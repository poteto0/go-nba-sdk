# Agents & Development Cycle

## Development Cycle

Always follow the `development-cycle` skill for all development tasks.

1.  **TDD Development**: Write tests first (`test-driven-development`).
2.  **CI Check**: Run `just ci`.
3.  **Code Review**: Score must be **> 95** (`review-rule`). Output to `/.gemini/reviews/<commit-hash>.md`.
4.  **Planning to Reflect Review**: Deal with the most impact minus first (`planning-to-reflect-review`). Ignore spell miss.
5.  **Repeat**: Until pass.
6.  **Write Docs**
7.  **Ask User's Review**
8.  **Commit**: Use gitmoji and co-author trailer (`commit-rule`).

## Available Skills

-   **test-driven-development**: Core TDD rules.
-   **development-cycle**: Full process from TDD to Commit.
-   **go-code-review**: Idiomatic Go review.
-   **nba-api-development**: Specifics for NBA API integration.
-   **commit-rule**: Gitmoji, issue linking, and co-author attribution.
-   **review-rule**: Scoring and deductions (-1 to -10).
-   **planning-to-reflect-review**: Strategy for addressing feedback.