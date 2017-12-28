+++
date = "2000-01-01T00:00:00+00:02"
title = "User Update"
url = "api-users-update"

[menu.api]
  weight = 2
  identifier = "api-users-update"
  parent = "api_users"
+++

Updates the specified user.
Please note this api requires administrative privileges.

```text
PATCH /api/users/{login}
```

Example Request Body:

```json
{
  "email": "octocat@github.com",
  "admin": false,
  "active": true
}
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
