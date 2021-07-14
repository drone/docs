---
date: 2000-01-01T00:00:00+00:00
title: drone secret add
author: bradrydzewski
weight: 5
aliases:
- /cli-secret-add/
---

This subcommand creates a new repository secret. Please note this command requires push privileges to the repository.

```
NAME:
   drone secret add - adds a secret

USAGE:
   drone secret add [command options] [repo/name]

OPTIONS:
   --repository value            repository name (e.g. octocat/hello-world)
   --name value                  secret name
   --data value                  secret value
   --allow-pull-request          permit read access to pull requests
   --allow-push-on-pull-request  permit write access to pull requests (e.g. allow docker push)
```

Example usage, creates a secret:

```
$ drone secret add --name my_token --data e72e16c7e42f29 octocat/hello-world
```

Example usage, creates a secret from a file:

```
$ drone secret add --name my_token --data @/path/to/secret.json octocat/hello-world
```
