---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RESOURCE_LIMIT_MEMORY
author: bradrydzewski
weight: 1
---

Optional integer, sets the default memory limits for all pipeline containers. This value maps to the `spec.containers[].resources.limits.memory` field.

```
DRONE_RESOURCE_LIMIT_MEMORY=500MiB
```