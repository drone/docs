---
date: 2000-01-01T00:00:00+00:00
title: drone secret update
author: bradrydzewski
weight: 5
aliases:
- /cli-secret-update/
---

This subcommand updates a named repository secret. Please note this command requires push privileges to the repository.

```
NAME:
   drone secret update - update a secret

USAGE:
   drone secret update [command options] [repo/name]

OPTIONS:
   --repository value            repository name (e.g. octocat/hello-world)
   --name value                  secret name
   --data value                  secret value
   --allow-pull-request          permit read access to pull requests
```