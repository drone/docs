+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Update"
url = "api-secret-update"

[menu.api]
  weight = 2
  identifier = "api-secret-update"
  parent = "api_secret"
+++

Updates the specified repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

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
  "image": null,
  "event": [
    "push",
    "pull_request"
  ]
}
```
