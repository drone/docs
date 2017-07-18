+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Update"
url = "api-secret-update"

[menu.api]
  weight = 2
  identifier = "api-secret-update"
  parent = "api_secret"
+++


Update a repository secret.

```text
PATCH /api/repos/{owner}/{repo}/secrets/{secret}
```

Example Request Body:

```json
{
  "value": "octocat",
  "event": [
    "push",
    "pull_request"
  ]
}
```

Example Response Body:

```json
{
  "id": 1,
  "name": "docker_username",
  "event": [
    "push",
    "pull_request"
  ]
}
```
