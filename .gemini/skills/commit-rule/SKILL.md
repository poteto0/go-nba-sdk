---
name: commit-rule
description: Guidelines for creating standardized commit messages with gitmoji and issue tracking.
---

# Commit Rule

## Overview
All commits must follow a strict format to ensure traceability and standardization.

## Rules

1.  **Gitmoji**: Use gitmojis aggressively to categorize the commit. Use the raw emoji character (e.g., ✨), not the shortcode (e.g., :sparkles:) unless the environment requires it, and NEVER use brackets.
2.  **Issue Linking**:
    -   Extract the issue number from the branch name (e.g., `user/#123/feature-x` -> `#123`).
    -   Include `refs: #<issue-number>` in the message.
3.  **Co-Author Attribution**:
    -   Always include the co-author trailer for the AI assistant.
    -   `Co-Authored-By: gemini-cli <218195315+gemini-cli@users.noreply.github.com>`

## Format Template

```bash
git commit -m "<gitmoji>: <message>. refs: #<issue-number>" -m "Co-Authored-By: gemini-cli <218195315+gemini-cli@users.noreply.github.com>"
```

**Do NOT use brackets `[]` around the gitmoji.**

## Examples



-   `✨: Add user login feature. refs: #42`

-   `🐛: Fix null pointer exception in auth. refs: #101`

-   `♻️: Refactor database connection logic. refs: #88`

-   `🤖: Add new AI skill or update agent configuration. refs: #26`
