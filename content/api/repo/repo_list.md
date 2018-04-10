+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo List"
url = "api-repo-list"

[menu.api]
  weight = 7
  identifier = "api-repo-list"
  parent = "api_repo"
+++

Returns repositories which are registered to Drone.

```text
GET /api/user/repos
```

Example Response Body:

```json
[
  {
    "id": 1,
    "scm": "git",
    "owner": "octocat",
    "name": "hello-world",
    "full_name": "octocat/hello-world",
    "avatar_url": "https://avatars.githubusercontent.com/u/2181346?v=3",
    "link_url": "https://github.com/octocat/hello-world",
    "clone_url": "https://github.com/octocat/hello-world.git",
    "default_branch": "master",
    "timeout": 60,
    "private": false,
    "trusted": false,
    "allow_pr": true,
    "allow_push": true,
    "allow_deploys": false,
    "allow_tags": false,
    "visibility": "public",
    "gated": false,
    "active": true,
    "last_build": 0,
    "config_file": ".drone.yml"
  }
]
```
