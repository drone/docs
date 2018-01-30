+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Repair"
url = "api-repo-repair"

[menu.api]
  weight = 5
  identifier = "api-repo-repair"
  parent = "api_repo"
+++

Recreates webhooks for your repository in your version control system (e.g GitHub).
This can be used if you accidentally delete your webhooks.
Please note this api requires administrative access to the repository.

```text
POST /api/repos/{owner}/{repo}/repair
```
