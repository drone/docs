+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Create"
url = "api-repo-create"

[menu.api]
  weight = 1
  identifier = "api-repo-create"
  parent = "api_repo"
+++

Registers a named repository with Drone.
Please note this api requires administrative access to the repository.

```text
POST /api/repos/{owner}/{name}
```

Example Response Body:

```json
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
```
