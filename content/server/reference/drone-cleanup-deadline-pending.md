---
date: 2000-01-01T00:00:00+00:00
title: DRONE_CLEANUP_DEADLINE_PENDING
author: bradrydzewski
aliases:
- /installation/reference/drone-cleanup-deadline-pending/
---


Optional duration value. Configures the interval after which a pending job will be killed by the reaper. The default value is 24 hours.

Kill pending jobs older than 24 hours:

```
DRONE_CLEANUP_DEADLINE_PENDING=24h
```
