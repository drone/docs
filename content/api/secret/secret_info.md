+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret Info"
url = "api-secret-info"

[menu.api]
  weight = 4
  identifier = "api-secret-info"
  parent = "api_secret"
+++

Returns the repository secret.
Please note this api requires write access to the repository,
and the request parameter `{secret}` is not the secret's id but secret name.

```text
GET /api/repos/{owner}/{repo}/secrets/{secret}
```

Example Response Body:

```json
{
  "id": 1,
  "name": "sample_secret_name",
  "image": null,
  "event": [
    "push",
    "tag",
    "deployment"
  ]
}
```
