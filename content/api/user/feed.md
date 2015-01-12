+++
weight = 3
date = "2014-11-08T12:54:06-08:00"
title = "List your recent commits"

path = "/api/user/feed"
method = "GET"
resource = "user"
+++

List recent commits for the authenticated user. **NOTE** this endpoint is expected to change in
the near future to work more like a feed.

### Example Request:

```bash
curl -X GET https://drone.io/api/user/feed
```

### Example Response (200 OK):

```json
[
  {
    "remote": "github.com",
    "host": "github.com",
    "owner": "foo",
    "name": "bar",
    "status": "Success",
    "started_at": 1415379751,
    "finished_at": 1415379938,
    "duration": 187,
    "sha": "a4cc2e2949f1912b83ef9e925ce2ce19f7d12e83",
    "branch": "docker-client",
    "pull_request": "678",
    "author": "foo",
    "gravatar": "60db20036235b833778e757d48501953",
    "timestamp": "2014-11-07 17:02:31.334424773 +0000 UTC",
    "message": "Return ErrInternalServer when docker server return 500 code",
    "created_at": 1415379751,
    "updated_at": 1415379938
  },
  {
    "remote": "github.com",
    "host": "github.com",
    "owner": "bar",
    "name": "baz",
    "status": "Failure",
    "started_at": 1413522048,
    "finished_at": 1413522068,
    "duration": 20,
    "sha": "eda8570122a35a308605f948364ef0df3235c175",
    "branch": "master",
    "author": "foo@bar.com",
    "gravatar": "8c58a0be77ee441bb8f8595b7f1b4e87",
    "timestamp": "2014-10-16T22:00:48-07:00",
    "message": "Update .drone.yml",
    "created_at": 1413522048,
    "updated_at": 1413522068
  }
]
```