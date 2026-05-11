---
title: Draft Board
---

# Draft Board

Get the draft board for a specific season.

## Endpoint

`client.Draft.GetDraftBoard(params)`

## Parameters

| Parameter | Type | Required | Description |
| :--- | :--- | :--- | :--- |
| LeagueID | string | No | Default: "00" |
| Season | string | Yes | Format: "YYYY" (e.g., "2023") |

## Response

Returns `types.DraftBoardResponse`.
