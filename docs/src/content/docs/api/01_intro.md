---
title: "Intro"
description: "How to get started with the lolesports-calendar api"
---

## Auth

lolesports-calendar's api uses oauth to interact with google calendar.

Every request that interacts with google calendar in some way, needs either a refresh token, or an access token, and a specified type.
In order to get your tokens, you can use the <a href="02_auth">auth</a> endpoints.

## Refresh token vs Access token

Access tokens only last 1 hour before expiring, as stated in the expired field returned from a successful authorization. To get around this,
you can use refresh tokens, which dont run out.

When making requests, add the type: "your token type" field to your request body, in addition to the token in this format: "token_type": "token".
When using access token, the expiry field also has to be present.

Example body for request with access token in js
```js
body: JSON.stringify({
    type: "access_token", 
    token: "1234",
    expiry: "expiry date"
})
```
