---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RESOURCE_LIMIT_CPU
author: bradrydzewski
weight: 1
---

Optional integer, sets the default millicpu resource limits for all pipeline containers. This value maps to the `spec.containers[].resources.limits.cpu` field.

```
DRONE_RESOURCE_LIMIT_CPU=1000
```