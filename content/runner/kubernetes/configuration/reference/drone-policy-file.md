---
date: 2000-01-01T00:00:00+00:00
title: DRONE_POLICY_FILE
author: bradrydzewski
weight: 1
---

Optional string value. Provides the path to the [policy file]({{< relref "runner/kubernetes/configuration/policies.md" >}}) that should be loaded on startup. _Note that the policy file needs to be mounted into the runner container._

```
DRONE_POLICY_FILE=/path/to/policy.yml
```
