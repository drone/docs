---
date: 2000-01-01T00:00:00+00:00
title: drone cron add
author: bradrydzewski
weight: 1
---

This subcommand registers a new cron job with the system. Please note this command requires administrative privileges to the repository.

```
NAME:
   drone cron add - adds a cronjob

USAGE:
   drone cron add [command options] [repo/name] [cronjob] [cronexpr]

OPTIONS:
   --branch value  branch name (default: "master")
```

Example usage:

```
drone cron add octocat/hello-world nightly "0 0 0 * * *"
```