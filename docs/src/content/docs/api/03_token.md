---
title: "Token"
description: "Recieving token in callback"
---

Recieve an oauth2 token if a valid code is present in url.

## Request

endpoint: /token

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/token -d { "code": "auth_code" }
```

## Response

200 - Success

```json
{
    "access_token": string,
    "expires_in": string,
    "refresh_token": string,
    "token_type": string
}
```

401 - Bad request

```json
"No code in url"
```
