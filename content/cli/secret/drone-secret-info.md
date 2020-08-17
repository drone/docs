---
date: 2000-01-01T00:00:00+00:00
title: drone secret info
author: bradrydzewski
weight: 5
aliases:
- /cli-secret-info/
---

This subcommand prints the secret metadata to the console. Please note this command requires push privileges to the repository.

```
NAME:
   drone secret info - display secret info

USAGE:
   drone secret info [command options] [repo/name]

OPTIONS:
   --repository value  repository name (e.g. octocat/hello-world)
   --name value        secret name
```

Example usage:

```
$ drone secret info octocat/hello-world my_token
```