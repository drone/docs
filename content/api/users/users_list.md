+++
date = "2000-01-01T00:00:00+00:02"
title = "User List"
description = "Endpoint to list user accounts"
+++

Returns a list of all registered users.
Please note this api requires administrative privileges.

```
GET /api/users
```

Example Response Body:

```json {linenos=table}
[
  {
    "id": 1,
    "login": "octocat",
    "email": "octocat@github.com",
    "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
    "active": true,
    "admin": false,
    "machine": false,
    "syncing": false,
    "synced": 1564096971,
    "created": 1564096971,
    "updated": 1564096971,
    "last_login": 1564096971
  }
]
```
