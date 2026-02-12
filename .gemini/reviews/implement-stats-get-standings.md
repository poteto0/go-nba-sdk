# Review: Implement Stats.GetStandings (IST Standings)

## Summary
Implemented the `GetStandings` method (referencing `iststandings` endpoint) in the `Stats` namespace.

## Score: 100 / 100

## Core Checks
-   **[Pass] Clean Architecture**: Followed the established pattern: Handler -> Internal Parser -> Types.
-   **[Pass] Namespace API**: Added `GetStandings` to `IStatsNamespace` and `StatsNamespace`.
-   **[Pass] DRY**: Reused `internal` response parsing and `toInt`/`toFloat` helpers in the parser.
-   **[Pass] UT Coverage**: 
    -   Handler test with `httpmock` in `api/stats/api_ist_standings_test.go`.
    -   Namespace test in `namespace/stats_namespace_test.go`.

## Non-Impactive
-   **[Pass] Spelling**: No issues.

## Recommendations
-   **Proceed to Commit**: The feature is correctly implemented and tested according to project standards.
