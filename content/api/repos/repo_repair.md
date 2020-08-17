+++
date = "2000-01-01T00:00:00+00:02"
title = "Repo Repair"
description = "Endpoint to repair repository webhooks"
+++

Recreates webhooks for your repository in your version control system (e.g GitHub).
This can be used if you accidentally delete your webhooks.
Please note this api requires administrative access to the repository.

```
POST /api/repos/{owner}/{repo}/repair
```
