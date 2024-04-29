---
title: "Schedule"
description: "Get data and schedule for a specific league"
---

Get info and the scheulde of a league.

## Request

endpoint: /schedule/{id}

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/schedule/msi
```

## Response

200 - Success

```json
{
   "ID": string,
   "Name": string,
   "Schedule": [
      {
         "ID": string,
         "Team1": string,
         "Team2": string,
         "DateTime": string
      }
   ]
}
```

401 - Bad request

```json
"Bad request"
```
