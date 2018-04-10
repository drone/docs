+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret List"
url = "api-secret-list"

[menu.api]
  weight = 5
  identifier = "api-secret-list"
  parent = "api_secret"
+++

Returns the repository secret list.
Please note this api requires write access to the repository.

```text
GET /api/repos/{owner}/{repo}/secrets
```

Example Response Body:

```json
[
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
]
```
