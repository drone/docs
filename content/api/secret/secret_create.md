+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Create"
url = "api-secret-create"

[menu.api]
  weight = 1
  identifier = "api-secret-create"
  parent = "api_secret"
+++

Create a new repository secret.
Please note this api requires write access to the repository.

```text
POST /api/repos/{owner}/{repo}/secrets
```

Example Request Body:

```json
{
  "Name": "docker_username",
  "Value": "octocat",
  "Event": [
    "push",
    "tags"
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
    "tags"
  ]
}
```
