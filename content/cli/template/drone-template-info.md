---
date: 2000-01-01T00:00:00+00:00
title: drone template info
author: eoinmcafee
weight: 5
aliases:
- /cli-template-info/
---

This subcommand prints the template metadata to the console. Please note this command requires push privileges to the repository.

```
NAME:
   drone template info - display template info

USAGE:
   drone template info [command options] [namespace] [name]

OPTIONS:
   --namespace value  organization namespace (e.g. octocat)
   --name value       template name
```

Example usage:

```
$ drone template info --name plugin.starklark --namespace octocat
```