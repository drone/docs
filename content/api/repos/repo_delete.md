+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Disable"
description = "Endpoint to disable a repository"
+++

Permanently deletes a repository. It cannot be undone.
Please note this api requires administrative access to the repository,
and repository's secrets and builds aren't deleted.

```
DELETE /api/repos/{owner}/{repo}
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
    "active": false,
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
