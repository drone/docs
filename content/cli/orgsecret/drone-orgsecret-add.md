---
date: 2000-01-01T00:00:00+00:00
title: drone orgsecret add
author: bradrydzewski
weight: 5
---

This subcommand creates a new organization secret. Please note this command requires system administrator privileges.

```
NAME:
   drone orgsecret add - adds a secret

USAGE:
   drone orgsecret add [command options] [organization] [name] [data]

OPTIONS:
   --allow-pull-request    permit read access to pull requests
```

Example usage, creates a secret:

```
$ drone orgsecret add acme my_token e72e16c7e42f29
```

Example usage, expose the secret to pull requests:

```
$ drone orgsecret add acme my_token e72e16c7e42f29 --allow-pull-request
```

Example usage, creates a secret from a file:

```
$ drone orgsecret add acme my_token @/path/to/secret.json
```
