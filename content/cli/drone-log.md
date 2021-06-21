---
date: 2000-01-01T00:00:00+00:00
title: drone encrypt
author: tphoney
weight: 4
separator: true
---

The command views or purges logs. 

```
NAME:
   drone log - manage logs

USAGE:
   drone log [global options] command [command options] [arguments...]

COMMANDS:
     purge  purge a log
     view   display the step logs

GLOBAL OPTIONS:
   --help, -h  show help
```

Example usage, view a log:

```
drone log view octocat/hello-world 1
```

Example usage, purge a log:

```
drone log purge octocat/hello-world
```
