# Review: Unstaged Skills Update

## Summary
Review of `AGENTS.md` and new skill definitions (`commit-rule`, `review-rule`, `planning-to-reflect-review`, `development-cycle`).

## Score: 94 / 100

## Deductions

### Core Checks
-   **[-5] Not DRY**: `AGENTS.md` duplicates the detailed rules for commits and reviews that are now defined in their respective `SKILL.md` files. This creates a maintenance burden where updates must happen in two places.

### Non-Impactive
-   **[-1] Spelling**: `AGENTS.md` contains "gemi" instead of "Gemini" and "nummber" instead of "number".

## Recommendations
1.  **Refactor `AGENTS.md`**: Remove the detailed "Commit Rule" and "Review Rule" sections. The "Available Skills" section is sufficient to point users/agents to the right place.
2.  **Fix Typos**: Correct "gemi" and "nummber" if retaining any text.
