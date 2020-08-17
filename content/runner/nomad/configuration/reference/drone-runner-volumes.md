---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RUNNER_VOLUMES
author: bradrydzewski
weight: 1
---

Optional comma separated list. Provides a list of Docker volumes that are mounted into every pipeline step.

```
DRONE_RUNNER_VOLUMES=/path/on/host:/path/in/container
```

In the above example, the path to the left of the colon is the __host machine path__. The path to the right is the path inside your pipeline containers.
