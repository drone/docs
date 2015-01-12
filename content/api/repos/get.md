+++
weight = 10
date = "2014-11-08T12:54:06-08:00"
title = "Get a Repository"

path = "/api/repos/:repo"
method = "GET"
resource = "repos"
+++

Gets the specified repository. The repository parameter is required and must be the
canonical repository name, for example, `github.com/foo/bar` or `bitbucket.org/baz/bar`.

### Example Request: 

```bash
curl -X GET https://drone.io/api/repos/github.com/foo/bar
```

### Example Response (200 OK):

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