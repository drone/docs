---
date: 2000-01-01T00:00:00+00:00
title: drone cron enable
author: bradrydzewski
weight: 1
---

This subcommand enables a named cron job that was previously disabled. Please note this command requires administrative privileges to the repository.

```
NAME:
   drone cron enable - enable cron jobs

USAGE:
   drone cron enable [repo/name] [cronjob]
```


Example usage:

```
drone cron enable octocat/hello-world nightly
```