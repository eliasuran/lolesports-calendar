---
title: "Auth"
description: "How to authorize and get your tokens"
---

Recieve an oauth2 token used to interact with google calendar. After making the request, a browser will popup, asking you to authorize with your google account. 
After authorizing with google, another browser will popup telling you the status of your request.

## Request

endpoint: /auth

method: GET

Sample request:

```bash
curl -X GET http://localhost:8080/auth
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
"Bad request"
```
