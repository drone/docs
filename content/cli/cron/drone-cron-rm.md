---
date: 2000-01-01T00:00:00+00:00
title: drone cron rm
author: tphoney
weight: 1
---

This subcommand deletes a named cron job. Please note this command requires administrative privileges to the repository.

```
NAME:
   drone cron rm - deletes cron jobs

USAGE:
   drone cron rm [repo/name] [cronjob]
```


Example usage:

```
drone cron rm octocat/hello-world nightly
```
