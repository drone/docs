---
date: 2000-01-01T00:00:00+00:00
title: DRONE_NODE_SELECTOR_DEFAULT
author: bradrydzewski
weight: 1
---


Optional string map. Provides a set of default node_selector values used when creating the pipeline pod. The default node_selector values are used if node_selector is undefined in the yaml.

```
DRONE_NODE_SELECTOR_DEFAULT=disktype:ssd,memory:highmem
```
