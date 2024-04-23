---
title: "Auth"
description: "How to authorize and get your tokens"
---

Recieve an oauth2 token used to interact with google calendar. Making a request to auth will return a url. Visit the url and follow the auth process.
After authorizing with google, you will be redirected to /callback with your token. See <a href="/api/callback">callback</a> for info about the callback endpoint.

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
"url"
```

401 - Bad request

```json
"Bad request"
```
