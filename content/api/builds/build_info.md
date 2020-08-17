+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Info"
description = "Endpoint to find a build"
+++

Returns the specified repository build.
Please note this api requires read access to the repository and the request parameter `{build}` is not the build id but the build number.

```
GET /api/repos/{owner}/{repo}/builds/{build}
```

Example Response Body:

```json {linenos=table}
{
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "success",
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
    "started": 1564085874,
    "finished": 1564086343,
    "created": 1564085874,
    "updated": 1564085874,
    "version": 3,
    "stages": [
        {
            "id": 199937,
            "repo_id": 296163,
            "build_id": 100207,
            "number": 1,
            "name": "default",
            "kind": "pipeline",
            "type": "docker",
            "status": "success",
            "errignore": false,
            "exit_code": 0,
            "machine": "15e89c0f84f1",
            "os": "linux",
            "arch": "amd64",
            "started": 1564085874,
            "stopped": 1564086343,
            "created": 1564085874,
            "updated": 1564086343,
            "version": 4,
            "on_success": true,
            "on_failure": false,
            "steps": [
                {
                    "id": 575525,
                    "step_id": 199937,
                    "number": 1,
                    "name": "clone",
                    "status": "success",
                    "exit_code": 0,
                    "started": 1564085876,
                    "stopped": 1564085901,
                    "version": 4
                },
                {
                    "id": 575526,
                    "step_id": 199937,
                    "number": 2,
                    "name": "test",
                    "status": "success",
                    "exit_code": 0,
                    "started": 1564085901,
                    "stopped": 1564085994,
                    "version": 4
                }
            ]
        }
    ]
}
```
