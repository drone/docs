---
date: 2000-01-01T00:00:00+00:00
title: drone secret rm
author: bradrydzewski
weight: 5
aliases:
- /cli-secret-rm/
---

This subcommand removes a named repository secret. Please note this command requires push privileges to the repository.

```
NAME:
   drone secret rm - remove a secret

USAGE:
   drone secret rm [command options] [repo/name]

OPTIONS:
   --repository value  repository name (e.g. octocat/hello-world)
   --name value        secret name
```

Example usage:

```
$ drone secret rm octocat/hello-world token
```