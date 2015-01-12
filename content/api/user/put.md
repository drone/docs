+++
weight = 2
date = "2014-11-08T12:54:06-08:00"
title = "Update the authenticated user"

path = "/api/user"
method = "PUT"
resource = "user"
+++

Updates the currently authenticated user profile. This is currently limited to the user's full name
and email address.

### Example Request: 

```bash
curl -X PUT https://drone.io/api/user
```

```json
{
  "name": "Foo Bar",
  "email": "foo.bar@gmail.com",
}
```

### Example Response (200 OK):

```json
{
  "remote": "github.com",
  "login": "foo",
  "name": "Foo Bar",
  "email": "foo.bar@gmail.com",
  "gravatar": "8c58a0be77aa441bb8f8595b7f1b4e87",
  "admin": false,
  "active": true,
  "created_at": 1413339968,
  "updated_at": 1415047663,
  "synced_at": 1415047663
}
```