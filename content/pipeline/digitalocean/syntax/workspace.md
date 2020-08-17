---
date: 2000-01-01T00:00:00+00:00
title: Workspace
author: bradrydzewski
weight: 2
toc: false
description: |
  Describes the pipeline workspace and directory structure.
---

Drone automatically creates a temporary directory, known as your workspace, where it clones your repository. The workspace is the current working directory for each step in your pipeline.

Workspace path on Posix systems:

```
/tmp/drone-${RANDOM}/drone/src
```
