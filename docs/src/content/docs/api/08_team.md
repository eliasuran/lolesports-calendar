---
title: "Team"
description: "Get data about a specific team"
---

Get data about a specific team

## Request

endpoint: /team/{id}

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/team/fnc
```

## Response

200 - Success

```json
{
    "Name": string,
    "Short": string,
    "Image": string,
}
```

401 - Bad request

```json
"Bad request"
```
