+++
date = "2000-01-01T00:00:00+00:02"
title = "Current User Info"
url = "api-user-info"

[menu.api]
  weight = 1
  identifier = "api-user-info"
  parent = "user"
+++

Returns the currently authenticated user.

```text
GET /api/user
```

Example Response Body:

```json
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}
```
