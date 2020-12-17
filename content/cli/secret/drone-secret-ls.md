---
date: 2000-01-01T00:00:00+00:00
title: drone secret ls
author: bradrydzewski
weight: 5
aliases:
- /cli-secret-ls/
---

This subcommand prints a list of repository secrets to the console. Please note this command requires push privileges to the repository.

```
NAME:
   drone secret ls - list secrets

USAGE:
   drone secret ls [command options] [repo/name]

OPTIONS:
   --repository value  repository name (e.g. octocat/hello-world)
```

Example usage:

```
$ drone secret ls octocat/hello-world
```
