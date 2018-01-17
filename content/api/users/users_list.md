+++
date = "2000-01-01T00:00:00+00:02"
title = "User List"
url = "api-users-list"

[menu.api]
  weight = 5
  identifier = "api-users-list"
  parent = "api_users"
+++

Returns a list of all registered users.
Please note this api requires administrative privileges.

```text
GET /api/users
```

Example Response Body:

```json
[{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}
]
```
