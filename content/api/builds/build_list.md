+++
date = "2000-01-01T00:00:00+00:02"
title = "Build List"
description = "Endpoint to list builds"
+++

Returns recent builds for the repository based on name.
Please note this api requires read access to the repository.

```
GET /api/repos/{owner}/{repo}/builds
```

Example Response Body:

```json {linenos=table}
[
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
      "version": 3
  }
]
```
