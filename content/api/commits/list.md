+++
weight = 14
date = "2014-11-08T12:54:06-08:00"
title = "List Commits"

path = "/api/repos/:repo/commits"
method = "GET"
resource = "commits"
+++

Gets a list of recent commits for the specified repository. The repository parameter is required and must be the canonical repository name, for example, `github.com/foo/bar` or `bitbucket.org/baz/bar`.

### Example Request: 

```bash
curl -X GET https://drone.io/api/repos/github.com/foo/bar/commits
```

### Example Response (200 OK):

```json
[
  {
    "id": 328,
    "status": "Success",
    "started_at": 1421029827,
    "finished_at": 1421030070,
    "duration": 243,
    "sha": "68e6d530fbdc8ac8d3c656a04a322aba2045a760",
    "branch": "master",
    "pull_request": "",
    "author": "john.smith@gmail.com",
    "gravatar": "8c58a0be77ee441bb8f8595b7f1b4e87",
    "timestamp": "2015-01-11T18:30:26-08:00",
    "message": "Merge pull request #800",
    "created_at": 1421029827,
    "updated_at": 1421030070
  },
  {
    "id": 327,
    "status": "Success",
    "started_at": 1421029603,
    "finished_at": 1421029813,
    "duration": 210,
    "sha": "9f2849d5adf2aa47745e1b7b9f76f1fcca1ebdb3",
    "branch": "master",
    "pull_request": "800",
    "author": "john.smith@gmail.com",
    "gravatar": "0bb3b7d4301b9e1335c6910f86362194",
    "timestamp": "2015-01-12 02:21:54.959678775 +0000 UTC",
    "message": "Update the Readme",
    "created_at": 1421029315,
    "updated_at": 1421029813
  }
]
```