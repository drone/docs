+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Create"
description = "Endpoint to create a new build"
+++

Create a build using the latest commit for the specified branch. _Please note the resulting build is created with event type `custom`._

* Creates a build using the latest commit to the default branch:
  ```
  POST /api/repos/{namespace}/{name}/builds
  ```

* Creates a build using the latest commit to the named branch:
   ```
   POST /api/repos/{namespace}/{name}/builds?branch={branch}
   ```

* Creates a build using the named branch and commit sha:
   ```
   POST /api/repos/{namespace}/{name}/builds?branch={branch}&commit={commit}
   ```

* Create a build with custom parameters. These parameters will be available to your pipeline steps as environment variables.

  ```
  POST /api/repos/{namespace}/{name}/builds?branch={branch}&{key=value}
  ```

Example Response Body:

```json {linenos=table}
{
    "id": 100207,
    "repo_id": 296163,
    "number": 42,
    "status": "pending",
    "event": "custom",
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

# Pipeline Configuration

When you trigger a build using the API the resulting build is created with event type `custom`.  If you define conditional pipelines or conditional steps based on event type you may need to update accordingly.

Example pipeline condition includes the custom event type:

```yaml  {linenos=table, hl_lines=["5"]}
trigger:
  event:
  - push
  - pull_request
  - custom
```


Example pipeline condition includes the custom event type:

```yaml  {linenos=table, hl_lines=["5"]}
when:
  event:
  - push
  - pull_request
  - custom
```
