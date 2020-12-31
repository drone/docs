+++
date = "2000-01-01T00:00:00+00:02"
title = "Cron Trigger"
description = "Endpoint to trigger a cron task"
+++

Trigger an existing cron task.
Please note this api requires write access to the repository.

```
POST /api/repos/{owner}/{repo}/cron/{name}
```

Example Response Body:

```json {linenos=table}
{
  "id": 100207,
  "repo_id": 296163,
  "trigger": "@cron",
  "number": 42,
  "status": "pending",
  "event": "cron",
  "action": "",
  "link": "https://github.com/octocat/hello-world/compare/e3320539a4c0...9fc1ad6ebf12",
  "timestamp": 0,
  "message": "updated README",
  "before": "",
  "after": "9fc1ad6ebf12462f3f9773003e26b4c6f54a772e",
  "ref": "refs/heads/master",
  "source_repo": "",
  "source": "",
  "target": "master",
  "author_login": "octocat",
  "author_name": "The Octocat",
  "author_email": "octocat@github.com",
  "author_avatar": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "sender": "octocat",
  "cron": "weekly-update-check",
  "started": 0,
  "finished": 0,
  "created": 1609336029,
  "updated": 1609336029,
  "version": 1
}
```