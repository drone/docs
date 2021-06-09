---
date: 2000-01-01T00:00:00+00:00
title: drone template rm
author: eoinmcafee
weight: 5
aliases:
- /cli-template-rm/
---
This subcommand removes a named organization template. Please note this command requires push privileges to the repository.

```
NAME:
   drone template rm - remove a template

USAGE:
   drone template rm [command options] [namespace] [name]

OPTIONS:
   --name value       template name
   --namespace value  organization name

```
Example usage:
```
$ drone template rm --namespace octocat --name plugin.starlark
```