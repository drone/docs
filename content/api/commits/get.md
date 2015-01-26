+++
weight = 15
date = "2014-11-08T12:54:06-08:00"
title = "Get a Commit"
draft = false

path = "/api/repos/:repo/branches/:branch/commits/:commit"
method = "GET"
resource = "commits"
+++

Gets the commit matching the specified repository, branch and sha.

### Example Request: 

```bash
curl -X GET https://drone.io/api/repos/github.com/foo/bar/branches/master/commits/d5e5797
```

### Example Response (200 OK):

```json
{
  "id": 328,
  "status": "Success",
  "started_at": 1421029827,
  "finished_at": 1421030070,
  "duration": 243,
  "sha": "d5e5797934b80e0203d0f47b207c4848cfc834bf",
  "branch": "master",
  "pull_request": "",
  "author": "john.smith@gmail.com",
  "gravatar": "8c58a0be77ee441bb8f8595b7f1b4e87",
  "timestamp": "2015-01-11T18:30:26-08:00",
  "message": "Merge pull request #800",
  "created_at": 1421029827,
  "updated_at": 1421030070
}
```