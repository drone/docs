+++
date = "2000-01-01T00:00:00+00:02"
title = "Registry Info"
url = "api-registry-info"

[menu.api]
  weight = 4
  identifier = "api-registry-info"
  parent = "api_registry"
+++

Returns a registry.
Please note this api requires write access to the repository,
and the request parameter `{registry}` is the registry address.

```text
GET /api/repos/{owner}/{repo}/registry/{registry}
```

Example Response Body:

```json
{
  "id": 1,
  "address": "docker.io",
  "username": "octocat",
  "password": "",
  "email": "",
  "token": ""
}
```
