+++
weight = 12
date = "2015-02-11T08:01:50+01:00"
title = "Deactivate a Repository"

path = "/api/repos/:repo/deactivate"
method = "POST"
resource = "repos"
+++

Deactivates a repository and discards any future build requests. Note that this will 
only disable builds. The repository and build history will remain in tact.

### Example Request: 

```bash
curl -X POST https://drone.io/api/repos/github.com/foo/bar/deactivate
```

### Example Response:

```json
{
  "remote": "github.com",
  "host": "github.com",
  "owner": "foo",
  "name": "bar",
  "url": "https:\/\/github.com\/foo\/bar",
  "clone_url": "git@github.com:bradrydzewski\/ami-utils.git",
  "git_url": "git:\/\/github.com\/foo\/bar.git",
  "ssh_url": "git@github.com:foo\/bar.git",
  "active": false,
  "private": true,
  "privileged": false,
  "post_commits": false,
  "pull_requests": false,
  "timeout": 900,
  "created_at": 1413339970,
  "updated_at": 1413339970
}
```

### From the Command Line: 

```bash
drone disable github.com/foo/bar
```
