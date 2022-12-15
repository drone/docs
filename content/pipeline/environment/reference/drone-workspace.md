---
date: 2000-01-01T00:00:00+00:00
title: DRONE_WORKSPACE
author: Sherry3
---

Drone's working directory for a pipeline is known as the `DRONE_WORKSPACE`. This is where it clones your repository. It is the working directory for each step in your pipeline and this folder is shared / persisted throughout the lifecycle of the pipeline.

In other words, individual steps can communicate and share state using the  `DRONE_WORKSPACE` directory structure.

It is a physical location for the following runners:

- Exec runner
- SSH runner
- Digitalocean runner
- AWS / VM runner * it is also shared as a mount for docker container steps

NB The `DRONE_WORKSPACE` is a volume for the following Docker runner
