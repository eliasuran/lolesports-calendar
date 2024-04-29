---
title: "Active Leagues"
description: "Get an array containing all active leagues"
---

Get an array of all active leagues

## Request

endpoint: /active_leagues

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/leagues
```

## Response

200 - Success

```json
["league"]
```

401 - Bad request

```json
"Bad request"
```
