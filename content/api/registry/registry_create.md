+++
date = "2000-01-01T00:00:00+00:02"
title = "Registry Create"
url = "api-registry-add"

[menu.api]
  weight = 1
  identifier = "api-registry-add"
  parent = "api_registry"
+++

Adds a registry.
Please note this api requires write access to the repository.

```text
POST /api/repos/{owner}/{repo}/registry
```
