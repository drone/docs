---
date: 2022-09-29T00:00:00+00:00
title: DRONE_WORKSPACE
author: Sherry3
---

Drone automatically creates a temporary volume, known as your workspace, where it clones your repository. The workspace is the current working directory for each step in your pipeline.

Because the workspace is a volume, filesystem changes are persisted between pipeline steps. In other words, individual steps can communicate and share state using the workspace filesystem.

It is supported by following pipelines: -
1. docker
2. kubernetes
3. exec
4. ssh
5. digitalocean
