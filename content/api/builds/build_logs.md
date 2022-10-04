+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Logs"
description = "Endpoint to find build logs"
+++

Please note this api requires read access to the repository.

```
GET /api/repos/{owner}/{repo}/builds/{build}/logs/{stage}/{step}
```

Example Response Body:

```json {linenos=table}
[
  {
    "pos": 0,
    "out": "Initialized empty Git repository in /drone/src/.git/\n",
    "time": 5
  },
  {
    "pos": 1,
    "out": "+ git fetch origin +refs/heads/master:\n",
    "time": 5
  },
  {
    "pos": 2,
    "out": "From https://github.com/octocat/hello-world.git\n",
    "time": 5
  },
  {
    "pos": 3,
    "out": " * branch            master     -> FETCH_HEAD\n",
    "time": 5
  },
  {
    "pos": 4,
    "out": " * [new branch]      master     -> origin/master\n",
    "time": 5
  },
  {
    "pos": 5,
    "out": "+ git reset --hard -q 62126a02ffea3dabd7789e5c5407553490973665\n",
    "time": 6
  }
]
```
