+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Chown"
url = "api-repo-chown"

[menu.api]
  weight = 4
  identifier = "api-repo-chown"
  parent = "api_repo"
+++

Lets a user assume ownership of a named repository.
Please note this api requires administrative access to the repository.

```text
POST /api/repos/{owner}/{repo}/chown
```
