+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Start"
url = "api-build-start"

[menu.api]
  weight = 1
  identifier = "api-build-start"
  parent = "api_build"
+++

Restart the specified build.
Please note this api requires read and write access to the repository and the request parameter `{build}` is not the build id but the build number.

```text
POST /api/repos/{owner}/{repo}/builds/{build}
```

Example Response Body:

```json
{
  "id": 1,
  "number": 1,
  "parent": 0,
  "event": "push",
  "status": "pending",
  "error": "",
  "enqueued_at": 1516175264,
  "created_at": 1516175264,
  "started_at": 0,
  "finished_at": 0,
  "deploy_to": "",
  "commit": "62126a02ffea3dabd7789e5c5407553490973665",
  "branch": "master",
  "ref": "refs/heads/master",
  "refspec": "",
  "remote": "https://github.com/octocat/hello-world.git",
  "title": "",
  "message": "New line at end of file. --Signed off by Spaceghost",
  "timestamp": 0,
  "sender": "Spaceghost",
  "author": "Spaceghost",
  "author_avatar": "https://avatars0.githubusercontent.com/u/251370?v=3",
  "author_email": "octocat@github.com",
  "link_url": "https://github.com/octocat/hello-world/commit/762941318ee16e59dabbacb1b4049eec22f0d303",
  "signed": false,
  "verified": true,
  "reviewed_by": "",
  "reviewed_at": 0,
  "procs": [
    {
      "id": 2740,
      "build_id": 732,
      "pid": 1,
      "ppid": 0,
      "pgid": 1,
      "name": "",
      "state": "pending",
      "exit_code": 0
    },
    {
      "id": 2741,
      "build_id": 732,
      "pid": 2,
      "ppid": 1,
      "pgid": 2,
      "name": "clone",
      "state": "pending",
      "exit_code": 0
    },
    {
      "id": 2742,
      "build_id": 732,
      "pid": 3,
      "ppid": 1,
      "pgid": 3,
      "name": "test",
      "state": "pending",
      "exit_code": 0
    }
  ]
}
```
