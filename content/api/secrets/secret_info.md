+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Info"
description = "Endpoint to find a repository secret"
+++

Returns the repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

```
GET /api/repos/{owner}/{repo}/secrets/{secret}
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
