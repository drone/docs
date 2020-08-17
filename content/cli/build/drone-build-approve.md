---
date: 2000-01-01T00:00:00+00:00
title: drone build approve
author: bradrydzewski
weight: 1
aliases:
- /cli-build-approve/
---

This subcommand approves a pipeline step pending manual approval. Please note this command requires admin privileges to the repository.

```
NAME:
   drone build approve - approve a build stage

USAGE:
   drone build approve <repo/name> <build> <stage>
```

Example usage, approves the named stage:

```
$ drone build approve octocat/hello-world 42 default
```
