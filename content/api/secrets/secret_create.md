+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Create"
description = "Endpoint to create a repository secret"
+++

Create a new repository secret.
Please note this api requires write access to the repository.

```
POST /api/repos/{owner}/{repo}/secrets
```

Example Request Body:

```json {linenos=table}
{
  "name": "docker_username",
  "data": "octocat",
  "pull_request": false
}
```

Example Response Body:

```json {linenos=table}
{
  "id": 1,
  "repo_id": 42,
  "name": "docker_username",
  "pull_request": false
}
```
