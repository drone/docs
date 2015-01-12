+++
weight = 12
date = "2014-11-08T12:54:06-08:00"
title = "Update a Repository"

path = "/api/repos/:repo"
method = "PUT"
resource = "repos"
+++

Updates the configuration for the specified repository. This behaves more like a `PATCH` method in that all fields are optional. Note that `privileged` and `timeout` fields can only be updated by a system administrator.

### Example Request:

```bash
curl -X PUT https://drone.io/api/repos/github.com/foo/bar
```

```json
{
  "privileged": true,
  "post_commits": true,
  "pull_requests": true,
  "timeout": 900,
}
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

## Key Management:

Note that you can also modify the default Public and Private RSA key pair used to clone your repository, and used by the SSH and Rsync plugins. Drone automatically generates a default for each repository, so pleaase be careful when overriding.

```bash
curl -X PUT https://drone.io/api/repos/github.com/foo/bar
```

```json
{
  "public_key": "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQ ...",
  "private_key": "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQCqGK ...",
}
```