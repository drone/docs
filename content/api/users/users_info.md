+++
date = "2000-01-01T00:00:00+00:02"
title = "User Info"
url = "api-users-info"

[menu.api]
  weight = 4
  identifier = "api-users-info"
  parent = "api_users"
+++

Returns information about the named registered user.
Please note this api requires administrative privileges.

```text
GET /api/users/{login}
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
