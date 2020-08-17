---
date: 2000-01-01T00:00:00+00:00
title: DRONE_NAMESPACE_DEFAULT
author: bradrydzewski
weight: 1
---

Optional string value provides the default namespace used to create pipeline resources. The default namespace is used if the no namespace is configured in the yaml.

```
DRONE_NAMESPACE_DEFAULT=default
```

You can optionally configure the runner to create a namespace for each pipeline execution which gets deleted when the pipeline completes. This can be configured using this exact value:

```
DRONE_NAMESPACE_DEFAULT=drone-
```