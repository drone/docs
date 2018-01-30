+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Update"
url = "api-repo-update"

[menu.api]
  weight = 2
  identifier = "api-repo-update"
  parent = "api_repo"
+++

Updates a named repository.
Please note this api requires administrative access to the repository.

```text
PATCH /api/repos/{owner}/{repo}
```

Example Request Body:

```json
{
  "timeout": 60,
  "private": false,
  "trusted": false,
  "allow_pr": true,
  "allow_push": true,
  "allow_deploys": false,
  "allow_tags": false
}
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
