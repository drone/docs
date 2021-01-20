---
date: 2000-01-01T00:00:00+00:00
title: DRONE_NETRC_CLONE_ONLY
author: bradrydzewski
weight: 1
---

Optional boolean value. Configures the runner to only inject the clone credentials into the clone step. _Please note that clone credentials are injected into all pipeline steps by default, but only for private repositories. Clone credentials are never injected into pipelines for public repositories._

```
DRONE_NETRC_CLONE_ONLY=true
```

