---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RESOURCE_REQUEST_MEMORY
author: bradrydzewski
weight: 1
---

Optional integer, sets the default memory resource request for all pipeline steps. This value maps to the `spec.containers[].resources.requests.memory` field.

```
DRONE_RESOURCE_REQUEST_MEMORY=500MiB
```