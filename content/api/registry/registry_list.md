+++
date = "2000-01-01T00:00:00+00:02"
title = "Registry List"
url = "api-registry-list"

[menu.api]
  weight = 4
  identifier = "api-registry-list"
  parent = "api_registry"
+++

Returns a list of all registries.
Please note this api requires write access to the repository.

```text
GET /api/repos/{owner}/{repo}/registry
```

Example Response Body:

```json
[
  {
    "id": 1,
    "address": "docker.io",
    "username": "octocat",
    "password": "",
    "email": "",
    "token": ""
  }
]
```
