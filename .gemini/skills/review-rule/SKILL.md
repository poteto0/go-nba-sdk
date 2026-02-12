---
name: review-rule
description: Guidelines for scoring and conducting code reviews.
---

# Review Rule

## Scoring System
-   **Range**: 0-100 points.
-   **Passing Score**: > 95 (as per development-cycle).

## Deductions

### Non-Impactive (-1)
-   **Spell Miss**: Minor typos in comments or non-critical strings. Usually ignored during implementation unless in public API.

### Core Checks (Significant Penalties)
-   **-10**: **Not following Clean Architecture**.
    -   Domain logic leaking into infrastructure.
    -   Incorrect dependency direction.
-   **-10**: **Not following Namespace API**.
    -   Public interfaces not adhering to project standards.
-   **-5**: **Not DRY (Don't Repeat Yourself)**.
    -   Duplicate logic or hardcoded values.
-   **UT Coverage Deductions**:
    -   Check line-level coverage for each modified file using: `just ut-cov <path> && go tool cover -func=coverage.out | grep <filename>`
    -   Deduct points for **each file** based on its line coverage:
        -   `< 95%`: **-1**
        -   `< 90%`: **-3**
        -   `< 85%`: **-4**
        -   `< 80%`: **-5**
        -   `< 75%`: **-10**
    -   **Note**: Existing tests must pass, and new code must be covered.
