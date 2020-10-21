---
date: 2000-01-01T00:00:00+00:00
title: DRONE_CLEANUP_DEADLINE_RUNNING
author: bradrydzewski
aliases:
- /installation/reference/drone-cleanup-deadline-running/
---


Optional duration value. Configures the interval after which a running job will be killed by the reaper. The default value is 24 hours.

Kill running jobs older than 24 hours:

```
DRONE_CLEANUP_DEADLINE_RUNNING=24h
```
