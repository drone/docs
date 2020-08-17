---
date: 2000-01-01T00:00:00+00:00
title: Workspace
author: bradrydzewski
weight: 1
toc: false
description: |
  Describes the pipeline workspace and directory structure.
---

Drone automatically creates a temporary directory, known as your workspace, where it clones your repository. The workspace is the current working directory for each step in your pipeline.

Workspace path on Posix systems:

```
/tmp/drone-${RANDOM}/drone/src
```

Workspace path on Windows systems:

```
C:\windows\drone-%RANDOM%\drone\src
```

<div class="alert">
Note the workspace is ephemeral. It is created when the pipeline starts and is destroyed after the pipeline completes.
</div>