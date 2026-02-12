---
name: planning-to-reflect-review
description: Strategy for addressing code review feedback and improving code quality.
---

# Planning to Reflect Review

## Strategy

When reviewing feedback from a code review:

1.  **Analyze the Score**: Check the final score and the breakdown of deductions.
2.  **Prioritize Impact**:
    -   **Start with the Most Impactful Minus**: Address the items with the largest deductions first (e.g., -10 for Clean Arch violations).
    -   **Fix Core Issues**: Ensure structural and architectural integrity is solid.
3.  **Spell Miss Policy**:
    -   Almost always ignore spell miss checks (-1), unless they affect public APIs or break functionality.
    -   Focus on code correctness and maintainability over minor typos in comments.

## Goal

-   Reflect feedback to improve the code.
-   Raise the review score to **> 95**.
