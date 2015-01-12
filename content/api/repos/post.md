+++
weight = 11
date = "2014-11-08T12:54:06-08:00"
title = "Activate a Repository"

path = "/api/repos/:repo"
method = "POST"
resource = "repos"
+++

Activates the specified repository. A post-commit hook and ssh key (private repositories only)
is added to the repository in the remote system (i.e. GitHub).

Attempts to re-activate an already active repository. This is helpful if the post-commit hooks
or keys were accidentally deleted and need to be restored.


### Example Request: 

```bash
curl -X POST https://drone.io/api/repos/github.com/foo/bar
```

### Example Response (201 Created):

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
  "active": true,
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
drone enable github.com/foo/bar
```