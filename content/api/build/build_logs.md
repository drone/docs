+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Logs"
url = "api-build-logs"

[menu.api]
  weight = 5
  identifier = "api-build-logs"
  parent = "api_build"
+++

Please note this api requires read access to the repository.

```text
GET /api/repos/{owner}/{repo}/logs/{log}/{pid}
```

Example Response Body:

```json

[
  {
    "proc": "clone",
    "pos": 0,
    "out": "+ git init\n"
  },
  {
    "proc": "clone",
    "pos": 1,
    "out": "Initialized empty Git repository in /drone/src/github.com/octocat/hello-world/.git/\n"
  },
  {
    "proc": "clone",
    "pos": 2,
    "out": "+ git remote add origin https://github.com/octocat/hello-world.git\n"
  },
  {
    "proc": "clone",
    "pos": 3,
    "out": "+ git fetch --no-tags origin +refs/heads/master:\n"
  },
  {
    "proc": "clone",
    "pos": 4,
    "out": "From https://github.com/octocat/hello-world\n"
  },
  {
    "proc": "clone",
    "pos": 5,
    "out": " * branch            master     -> FETCH_HEAD\n"
  },
  {
    "proc": "clone",
    "pos": 6,
    "out": " * [new branch]      master     -> origin/master\n"
  },
  {
    "proc": "clone",
    "pos": 7,
    "out": "+ git reset --hard -q 62126a02ffea3dabd7789e5c5407553490973665\n"
  },
  {
    "proc": "clone",
    "pos": 8,
    "out": "+ git submodule update --init --recursive\n"
  }
]
```
