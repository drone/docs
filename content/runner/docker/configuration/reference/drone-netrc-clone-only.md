---
date: 2000-01-01T00:00:00+00:00
title: DRONE_NETRC_CLONE_ONLY
author: bradrydzewski
weight: 1
---

Optional boolean value. Configures the runner to only inject the clone credentials into the clone step.

<div class="alert">
Please note that Drone injects clone credentials into all pipeline steps if the repository is private or requires authentication to clone; Drone never injects credentials into pipeline steps if the repository is public.
</div>

```
DRONE_NETRC_CLONE_ONLY=true
```
