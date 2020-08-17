---
date: 2000-01-01T00:00:00+00:00
title: DRONE_CRON_INTERVAL
author: bradrydzewski
aliases:
- /installation/reference/drone-cron-interval/
---

Optional duration value. Configures the interval at which the cron scheduler is run. The cron scheduler is not meant to be accurate and batches pending jobs every hour by default.

Process pending cron jobs every hour:

```
DRONE_CRON_INTERVAL=1h
```

Process pending cron jobs every five minutes:
```
DRONE_CRON_INTERVAL=5m
```

_Please note even when you reduce the cron interval, you should not expect high levels of accuracy. Our primary design goal was to create a safe, efficient cron scheduler that prevents users from overloading the system; this comes at the expense of accuracy._