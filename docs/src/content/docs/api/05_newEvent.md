---
title: "New Event"
description: "Create a new event in the calendar"
---

Create a new event in the calendar

## Request

endpoint: /newEvent

method: POST

Sample request:

```bash
curl -X GET http://localhost:8080/leagues/msi -d { "type": "refresh_token" "refresh_token": "your refresh token" }
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
