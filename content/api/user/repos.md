+++
weight = 4
date = "2014-11-08T12:54:06-08:00"
title = "List your repositories"

path = "/api/user/repos"
method = "GET"
resource = "user"
+++

List repositories for the authenticated user, including repositories owned
by organizations which the user can access.

### Example Request:

```bash
curl -X GET https://drone.io/api/user/repos
```

### Example Response:

```json
[
  {
    "remote": "github.com",
    "host": "github.com",
    "owner": "foo",
    "name": "bar",
    "url": "https:\/\/github.com\/foo\/bar",
    "clone_url": "git@github.com:bradrydzewski\/ami-utils.git",
    "git_url": "git:\/\/github.com\/foo\/bar.git",
    "ssh_url": "git@github.com:foo\/bar.git",
    "active": false,
    "private": true,
    "privileged": false,
    "post_commits": false,
    "pull_requests": false,
    "timeout": 900,
    "created_at": 1413339970,
    "updated_at": 1413339970
  }
]
```

### From the Command Line

```bash
drone repos
drone repos --all
```