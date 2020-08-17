---
date: 2000-01-01T00:00:00+00:00
title: DRONE_TASK_MEMORY
author: bradrydzewski
weight: 1
---

Optional integer value that assigns the tasks default memory resources. _Please note that this does not limit the memory resources for the pipeline; it limits the compute resources for the Pipeline runner task, which is a controller process that is separate from the pipeline processes._

```
DRONE_TASK_MEMORY=500
```
