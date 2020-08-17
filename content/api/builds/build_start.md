+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Restart"
description = "Endpoint to re-start a build"
+++

Restart the specified build.
Please note this api requires read and write access to the repository and the request parameter `{build}` is not the build id but the build number.

```
POST /api/repos/{owner}/{repo}/builds/{build}
```

Example Response Body:

```json {linenos=table}
{
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "pending",
    "event": "pull_request",
    "action": "sync",
    "link": "https://github.com/octoat/hello-world/compare/e3320539a4c0...9fc1ad6ebf12",
    "message": "updated README",
    "before": "e3320539a4c03ccfda992641646deb67d8bf98f3",
    "after": "9fc1ad6ebf12462f3f9773003e26b4c6f54a772e",
    "ref": "refs/heads/master",
    "source_repo": "spaceghost/hello-world",
    "source": "develop",
    "target": "master",
    "author_login": "octocat",
    "author_name": "The Octocat",
    "author_email": "octocat@github.com",
    "author_avatar": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
    "sender": "bradrydzewski",
    "started": 0,
    "finished": 0,
    "created": 1564085874,
    "updated": 1564085874,
    "version": 1,
    "stages": [
        {
            "id": 199937,
            "repo_id": 296163,
            "build_id": 100207,
            "number": 1,
            "name": "default",
            "kind": "pipeline",
            "type": "docker",
            "status": "pending",
            "errignore": false,
            "exit_code": 0,
            "machine": "15e89c0f84f1",
            "os": "linux",
            "arch": "amd64",
            "started": 0,
            "stopped": 0,
            "created": 1564085874,
            "updated": 1564086343,
            "version": 1,
            "on_success": true,
            "on_failure": false
        }
    ]
}
```
