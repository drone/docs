+++
date = "2000-01-01T00:00:00+00:02"
title = "Registry Update"
url = "api-registry-update"

[menu.api]
  weight = 2
  identifier = "api-registry-update"
  parent = "api_registry"
+++

Updates a registry.
Please note this api requires write access to the repository,
and the request parameter `{registry}` is the registry address.

```text
PATCH /api/repos/{owner}/{repo}/registry/{registry}
```
