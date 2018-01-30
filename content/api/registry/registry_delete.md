+++
date = "2000-01-01T00:00:00+00:02"
title = "Registry Delete"
url = "api-registry-delete"

[menu.api]
  weight = 3
  identifier = "api-registry-delete"
  parent = "api_registry"
+++

Removes a registry.
Please note this api requires write access to the repository,
and the request parameter `{registry}` is the registry address.

```text
DELETE /api/repos/{owner}/{repo}/registry/{registry}
```
