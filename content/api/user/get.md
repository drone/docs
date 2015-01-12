+++
weight = 1
date = "2014-11-08T12:54:06-08:00"
title = "Get the authenticated user"

path = "/api/user"
method = "GET"
resource = "user"
+++


### Example Request: 

```bash
curl -X GET https://drone.io/api/user
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
