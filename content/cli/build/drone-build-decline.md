---
date: 2000-01-01T00:00:00+00:00
title: drone build decline
author: bradrydzewski
weight: 1
aliases:
- /cli-build-decline/
---

This subcommand declines a pipeline step pending manual approval, causing the overall pipeline to fail. Please note this command requires admin privileges to the repository.

```
NAME:
   drone build decline - decline a build stage

USAGE:
   drone build decline <repo/name> <build> <stage>
```

Example usage, declines the named stage:

```
$ drone build decline octocat/hello-world 42 default
```
