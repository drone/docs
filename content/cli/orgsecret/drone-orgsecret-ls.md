---
date: 2000-01-01T00:00:00+00:00
title: drone orgsecret ls
author: bradrydzewski
weight: 5
---

This subcommand prints a list of organization secrets to the console. Please note this command requires system administrator privileges.

```
NAME:
   drone orgsecret ls - list secrets

USAGE:
   drone orgsecret ls [command options] [arguments...]

OPTIONS:
   --filter value  filter output by organization
```

Example usage:

```
$ drone orgsecret ls
```

Example usage, filter the results by organization:

```
$ drone orgsecret ls --filter=acme
```