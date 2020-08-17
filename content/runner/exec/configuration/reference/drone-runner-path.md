---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RUNNER_PATH
author: bradrydzewski
weight: 1
---

Optional string value. Sets the PATH variable for all pipeline steps. This may be required if the pipeline shell cannot find your commands, and you receive command not found errors.

```
DRONE_RUNNER_PATH=/usr/local/bin:/usr/bin:/usr/sbin:/sbin
```

