---
date: 2000-01-01T00:00:00+00:00
title: drone build ls
author: bradrydzewski
weight: 1
aliases:
- /cli-build-ls/
---

This subcommand prints a list of recent builds to the console.

```
NAME:
   drone build ls - show build history

USAGE:
   drone build ls [command options] <repo/name>

OPTIONS:
   --branch value  branch filter
   --event value   event filter
   --status value  status filter
   --limit value   limit the list size (default: 25)
   --page value    page number (default: 1)
```

Example usage:

```
$ drone build ls octocat/hello-world
```