+++
weight = 16
date = "2014-11-08T12:54:06-08:00"
title = "Rebuild a Commit"
draft = false

path = "/api/repos/:repo/branches/:branch/commits/:commit"
method = "POST"
resource = "commits"
+++

Re-executes the build for the commit matching the specified repository, branch and sha.

### Example Request: 

```bash
curl -X POST https://drone.io/api/repos/github.com/foo/bar/branches/master/commits/d5e5797
```

### Example Response:

```nohighlight
Status: 200 OK
```

### From the Command Line: 

```bash
drone restart github.com/foo/bar master d5e5797
```