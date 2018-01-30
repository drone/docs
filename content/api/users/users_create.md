+++
date = "2000-01-01T00:00:00+00:02"
title = "User Create"
url = "api-users-create"

[menu.api]
  weight = 1
  identifier = "api-users-create"
  parent = "api_users"
+++

Creates a user.
Please note this api requires administrative privileges.

```text
POST /api/users
```

Example Request Body:

```json
{
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
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
