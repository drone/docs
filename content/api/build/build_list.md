+++
date = "2000-01-01T00:00:00+00:02"
title = "Build List"
url = "api-build-list"

[menu.api]
  weight = 4
  identifier = "api-build-list"
  parent = "api_build"
+++

Returns recent builds for the repository based on name.
Please note this api requires read access to the repository.

```text
GET /api/repos/{owner}/{repo}/builds
```

Example Response Body:

```json
[
  {
    "id": 1,
    "number": 1,
    "parent": 0,
    "event": "push",
    "status": "success",
    "created_at": 1443677151,
    "enqueued_at": 1443677151,
    "started_at": 1443677151,
    "finished_at": 1443677255,
    "deploy_to": "",
    "commit": "2deb7e0d0cbac357eeb110c8a2f2f32ce037e0d5",
    "branch": "master",
    "ref": "refs/heads/master",
    "remote": "https://github.com/octocat/hello-world.git",
    "title": "",
    "message": "New line at end of file. --Signed off by Spaceghost",
    "timestamp": 1443677255,
    "sender": "Spaceghost",
    "author": "Spaceghost",
    "author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
    "author_email": "octocat@github.com",
    "link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
    "signed": false,
    "verified": true,
    "reviewed_by": "",
    "reviewed_at": 0
  }
]
```
