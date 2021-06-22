---
date: 2000-01-01T00:00:00+00:00
title: drone template ls
author: eoinmcafee
weight: 5
aliases:
- /cli-template-ls/
---

This subcommand prints a list of organization templates to the console. Please note this command requires push privileges to the repository.

```
NAME:
   drone template ls - list templates

USAGE:
   drone template ls [command options] [namespace]

OPTIONS:
   --repository value namespace (e.g. octocat)
```

Example usage:

```
$ drone template ls octocat
```