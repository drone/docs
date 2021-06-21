---
date: 2000-01-01T00:00:00+00:00
title: drone log
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

View a log:

```
drone log view <repo/name> <build> <stage> <step>
```

Example usage:

```
drone log view octocat/hello-world 19 1 3
```

Purge a log:

```
drone log purge <repo/name> <build> <stage> <step>
```

Example usage:

```
drone log purge octocat/hello-world 19 1 3
```
