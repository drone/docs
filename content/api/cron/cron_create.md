+++
date = "2000-01-01T00:00:00+00:02"
title = "Cron Create"
description = "Endpoint to create cron tasks"
+++

Create a new cron job. Please note this api requires write access to the repository.

```
POST /api/repos/{owner}/{repo}/cron
```

Example Request Body:

```json {linenos=table}
{
  "name": "nightly",
  "expr": "0 0 1 * * *",
  "branch": "master"
}
```

Example Response Body:

```json {linenos=table}
{
  "id": 1,
  "repo_id": 42,
  "name": "nightly",
  "expr": "0 0 1 * * *",
  "branch": "master",
  "next": 1564120128,
  "pref": 1564116480,
  "created": 1564096971,
  "updated": 1564096971,
}
```
