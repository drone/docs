---
date: 2000-01-01T00:00:00+00:00
title: drone cron disable
author: bradrydzewski
weight: 1
---

This subcommand disables the named cron job. Please note this command requires administrative privileges to the repository.


```
NAME:
   drone cron disable - disable cron jobs

USAGE:
   drone cron disable [repo/name] [cronjob]
```


Example usage:

```
drone cron disable octocat/hello-world nightly
```