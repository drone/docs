---
date: 2000-01-01T00:00:00+00:00
title: drone cron exec
author: tphoney
weight: 1
---

This subcommand executes a named cron job. Please note this command requires administrative privileges to the repository.

```
NAME:
   drone cron exec - triggers cron job

USAGE:
   drone cron exec [repo/name] [cronjob]
```


Example usage:

```
drone cron exec octocat/hello-world nightly
```
