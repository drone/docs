---
date: 2000-01-01T00:00:00+00:00
title: drone build stop
author: bradrydzewski
weight: 1
aliases:
- /cli-build-stop/
- /cli-build-kill/
---

This subcommand stops a pending or running build. Please note this command requires push privileges to the repository.

```
NAME:
   drone build stop - stop a build

USAGE:
   drone build stop <repo/name> [build]
```

Example usage, cancels a build by build number:

```
$ drone build stop octocat/hello-world 42
```