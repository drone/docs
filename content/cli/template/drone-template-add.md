---
date: 2000-01-01T00:00:00+00:00
title: drone template add
author: eoinmcafee
weight: 5
aliases:
- /cli-template-add/
---

This subcommand creates a new organization template. Please note this command requires push privileges to the repository.

```
NAME:
   drone template add - adds a template

USAGE:
   drone template add [command options] [namespace] [name] [data]

OPTIONS:
   --name value       template name
   --namespace value  organization namespace
   --data value       template file data
```

Example usage, creates a template:

```
$ drone template add octocat plugin.starklark @path_to_file
```
