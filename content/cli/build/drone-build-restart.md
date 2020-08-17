---
date: 2000-01-01T00:00:00+00:00
title: drone build restart
author: bradrydzewski
weight: 1
aliases:
- /cli-build-start/
---

This subcommand restarts an existing build. When the build is restarted, it is assigned a new, unique build number. Please note this command requires push privileges to the repository.

```
NAME:
   drone build restart - restart a build

USAGE:
   drone build restart [command options] <repo/name> [build]

OPTIONS:
   --param value, -p value  custom parameters to be injected into the job environment.
```

Example usage, restarts a build by build number:

```
$ drone build restart octocat/hello-world 42
```

Example usage, with custom parameters:

```
$ drone build restart octocat/hello-world 42 \
  --param=foo=bar \
  --param=baz=qux \
```
