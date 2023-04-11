+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Disable"
description = "Endpoint to disable a repository"
+++

Disables a named repository in Drone and disables the webhook in the source control management system (e.g. GitHub). Please note this api requires administrative access to the repository.

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

As mentioned above, this endpoint disables but does not delete the repository. Use the `remove` query parameter to delete the repository from the database. _Please note the repository may be automatically re-added to the database when you perform your next sync operation, if the repository still exists in your source control management system._

```
DELETE /api/repos/{owner}/{repo}?remove=true
```
