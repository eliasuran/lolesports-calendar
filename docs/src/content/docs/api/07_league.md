---
title: "League"
description: "Get data about a specific league"
---

Get data about a specific league

## Request

endpoint: /league/{id}

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/league/lec
```

## Response

200 - Success

```json
{
   "ID": string,
   "Name": string,
   "Teams": [
        {
            "Name": string,
            "Short": string,
            "Image": string,
        }
   ]
}
```

401 - Bad request

```json
"Bad request"
```
