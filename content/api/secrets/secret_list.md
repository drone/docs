+++
date = "2000-01-01T00:00:00+00:02"
title = "Secret List"
description = "Endpoint to list repository secrets"
+++

Returns the repository secret list.
Please note this api requires write access to the repository.

```
GET /api/repos/{owner}/{repo}/secrets
```

Example Response Body:

```json {linenos=table}
[
  {
    "id": 1,
    "repo_id": 42,
    "name": "docker_username",
    "data": "octocat",
    "pull_request": false
  }
]
```
