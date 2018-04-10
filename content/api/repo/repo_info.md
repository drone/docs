+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Info"
url = "api-repo-info"

[menu.api]
  weight = 6
  identifier = "api-repo-info"
  parent = "api_repo"
+++

Retrieves the details of a repository.
Please note this api requires read access to the repository.

```text
GET /api/repos/{owner}/{repo}
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
