+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Update"
description = "Endpoint to update a repository secret"
+++

Updates the specified repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

```
PATCH /api/repos/{owner}/{repo}/secrets/{secret}
```

Example Request Body:

```json {linenos=table}
{
  "data": "octocat",
  "pull_request": true
}
```

Example Response Body:

```json {linenos=table}
{
  "id": 1,
  "repo_id": 42,
  "name": "docker_username",
  "data": "octocat",
  "pull_request": true
}
```
