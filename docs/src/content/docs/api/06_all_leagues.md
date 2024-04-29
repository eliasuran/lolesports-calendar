---
title: "All Leagues"
description: "Get an array containing all leagues and teams"
---

Get an array of all leagues containing info about the league and all it's teams

## Request

endpoint: /all_leagues

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/all_leagues
```

## Response

200 - Success

```json
[
    {
        "ID": string,
        "Name": string,
        "Logo": string,
        "Teams": [
            {
                "Name": string,
                "Short": string,
                "Image": string,
            }
        ]
    }
]
```

401 - Bad request

```json
"Bad request"
```
