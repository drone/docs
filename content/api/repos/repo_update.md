+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Update"
description = "Endpoint to update a repository"
+++

Updates a named repository. Please note this api requires administrative access to the repository.

```
PATCH /api/repos/{owner}/{repo}
```

Example Request Body:

```json {linenos=table}
{
  "config_path": ".drone.yml",
  "protected": false,
  "trusted": false,
  "timeout": 60,
  "visibility": "internal"
}
```

Example Response Body:

```json {linenos=table}
{
    "id": 42,
    "uid": "16607898",
    "user_id": 2,
    "namespace": "octocat",
    "name": "octocat",
    "slug": "octocat/hello-world",
    "scm": "git",
    "git_http_url": "https://github.com/octocat/hello-world.git",
    "git_ssh_url": "git@github.com:octocat/hello-world.git",
    "link": "https://github.com/octocat/hello-world",
    "default_branch": "master",
    "private": false,
    "visibility": "public",
    "active": true,
    "config_path": ".drone.yml",
    "trusted": false,
    "protected": false,
    "ignore_forks": false,
    "ignore_pull_requests": false,
    "timeout": 60,
    "counter": 185,
    "synced": 1542160374,
    "created": 1542160374,
    "updated": 1542160374,
    "version": 187
}
```
