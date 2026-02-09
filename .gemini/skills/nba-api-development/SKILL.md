---
name: nba-api-development
description: Use when implementing features which connect to nba-api.
---

# NBA-API-Development

## Overview

Cycle of development features which connect to nba-api.

## When to Use

**Always:**

- New features which connect to nba-api

## Development Cycle

### RED: Impl API handler

Analyzing endpoint, referencing swar/nba-api(refs: https://github.com/swar/nba_api).

path: `api/[live|stats|static]/api\_<target_endpoint>.go

```go
func GetTargetEndpoint(api.IProvider, params *types.TargetEndpointParams) types.Response[any] {
    // if needed
    filePath := <params parsing>

    path := constants.<Type>BaseUrl + constants.<TargetPath> + "/" + filePath

    resp, err := provider.Get(path, &constants.<TargetHeader>)
	if err != nil {
		return types.Response[any]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}
	defer resp.Body.Close()

    fmt.Println(resp.Body)

    return types.Response[any]{
        Contents: "hello",
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
```

### RED: write ut

write minimize ut w/o httpmocking

```go
func Test_Get<TargetEndpoint>(t \*testing.T) {
    // Arrange
    provider := newProviderForTest()

    // Act
    result := live.GetBoxScore(provider, &types.BoxScoreParams{
    	GameID: "0022500733",
    })

    // Assert
    assert.NotError(t, result.Error)
    assert.Equal(t, 200, result.StatusCode)

}
```

### RED: run ut

Mainly, it is failing

```bash
$ just ut ./api/[live|stats|static]
```

### GREEN: fix params, endpoint to request success

repeat below till you can get 200 from target endpoint

1. check swar/nba_api & fix endpoint or params
2. run ut

### GREEN: run ut to check raw response from target endpoint

you can get raw response w/ `-v` option.

```bash
$ just ut ./api/[live|stats|static] -v
```

remember on test code

```go
func Test_Get<TargetEndpoint>(t *testing.T) {
    ...
    fmt.Println(resp.Body)
    ...
}
```

### GREEN: add sample response to fixtures

path: `fixtures/samples/[live|stats|static]_<target_endpoint>.go`

add raw response to var `Sample<TargetEndpoint>Response`

### GREEN: define response & write parsers

define response & write parser w/ reference on raw sample response.

If response is perfect json, you can skip this phase.
Because it is enough `encoding/json`.

Mainly of statsAPI, you need to write parser.

```jsonc
{
    ...,
    "resultSets": [
        "name": "SeasonTotalsRegularSeason",
        "headers": [...], // headers
        "rowSet": [
            [...],
        ], // data sets
    ],
}
```

### GREEN: fix handler's response

till you pass ut, fix response type on this phase, too.

```go
func GetTargetEndpoint(api.IProvider, params *types.TargetEndpointParams) types.Response[types.<ResponseType>] {
    // if needed
    filePath := <params parsing>

    path := constants.<Type>BaseUrl + constants.<TargetPath> + "/" + filePath

    resp, err := provider.Get(path, &constants.<TargetHeader>)
	if err != nil {
		return types.Response[types.<ResponseType>]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}
	defer resp.Body.Close()

    /* ! w/o parser case */
    var <data> types.<ResponseType>
	err = internal.ParseResponseTo(resp, &<data>)
	if err != nil {
		return types.Response[types.LiveBoxScoreResponse]{
			StatusCode: resp.StatusCode,
			Error:      err,
		}
	}

    /* w/ parser case */
    rawResp, err := internal.ParseResponse(resp)
	if err != nil {
		return types.Response[types.<ResponseType>]{Error: err}
	}

    contents, err := parser.Parse<Target>Response(rawResp)
	if err != nil {
		return types.Response[types.<ResponseType>]{Error: err}
	}


    return types.Response[types.<ResponseType>]{
        Contents: contents,
		StatusCode: resp.StatusCode,
		Error:      nil,
	}
}
```

### GREEN: check ut pass & fix w/ httpmocking

```go
func Test_Get<TargetEndpoint>(t \*testing.T) {
        ...
        httpmock.Activate(t)
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder(
			"GET",
			<path>,
			httpmock.NewStringResponder(200, samples.Sample<TargetEndpoint>Response),
		)
```

you need to write ut, error case as long as you can, too.

### GREEN: add namespace & wite ut of namespace
