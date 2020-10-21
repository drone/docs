---
date: 2000-01-01T00:00:00+00:00
title: DRONE_CLEANUP_INTERVAL
author: bradrydzewski
aliases:
- /installation/reference/drone-cleanup-interval/
---

Optional duration value. Configures the interval at which the reaper is run. The reaper finds and kills zombie jobs that are permanently stuck in a pending or running state every 24 hours by default.

Run reaper every 24 hours:

```
DRONE_CLEANUP_INTERVAL=24h
```

Run reaper every hour:
```
DRONE_CLEANUP_INTERVAL=60m
```
