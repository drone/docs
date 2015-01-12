+++
weight = 13
date = "2014-11-08T12:54:06-08:00"
title = "Disable a Repository"

path = "/api/repos/:repo"
method = "DELETE"
resource = "repos"
+++

Disables a repository and discards any future build requests. Note that this will only disable builds. The repository and build history will remain in tact.

### Example Request: 

```bash
curl -X DELETE https://drone.io/api/repos/github.com/foo/bar
```

### Example Response:

```nohighlight
Status: 204 No Content
```

### From the Command Line: 

```bash
drone disable github.com/foo/bar
```