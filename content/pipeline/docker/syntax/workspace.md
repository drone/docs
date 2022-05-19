---
date: 2000-01-01T00:00:00+00:00
title: Workspace
author: bradrydzewski
weight: 30
aliases:
- /workspace/
- /questions/how-to-customize-the-workspace/
description: |
  Describes the pipeline workspace and directory structure.
---

Drone automatically creates a temporary volume, known as your workspace, where it clones your repository. The workspace is the current working directory for each step in your pipeline. The default workspace path is `/drone/src`.

Because the workspace is a volume, filesystem changes are persisted between pipeline steps. In other words, individual steps can communicate and share state using the filesystem.

<div class="alert alert-warn">
Workspace volumes are ephemeral. They are created when the pipeline starts and destroyed after the pipeline completes.
</div>

# Customizing the Workspace

You can customize the workspace directory by defining the `workspace` section in your yaml. Here is a basic example:

{{< highlight text "linenos=table,linenostart=1,hl_lines=5-7" >}}
kind: pipeline
type: docker
name: default

workspace:
  path: /drone/src

steps:
- name: backend
  image: golang:latest
  commands:
  - go get
  - go test

- name: frontend
  image: node:latest
  commands:
  - npm install
  - npm run tests
{{< / highlight >}}

This would be equivalent to the following docker commands:

```
$ docker volume create my-named-volume
$ docker run --volume=my-named-volume:/drone/src golang
$ docker run --volume=my-named-volume:/drone/src node
```
