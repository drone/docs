---
date: 2000-01-01T00:00:00+00:00
title: drone repo info
author: bradrydzewski
weight: 5
aliases:
- /cli-repository-info
- /cli/user/drone-repo-info
toc: false
---

This subcommand prints the named repository details.

```
NAME:
   drone repo info - show repository details

USAGE:
   drone repo info [command options] [arguments...]
```

Example usage:

```
$ drone repo info octocat/hello-world

Owner: octocat
Repo: hello-world
Type: git
Private: false
Remote: https://github.com/octocat/hello-world.git
```

Format the output using a custom Go template:

```
$ drone repo info --format="{{ .Link }}" octocat/hello-world

https://github.com/octocat/hello-world.git
```