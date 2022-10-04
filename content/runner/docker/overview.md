---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /runners/docker
- /runners/docker/install
- /installation/runners/docker/
---

The Docker runner is a daemon that executes pipelines steps inside ephemeral Docker containers. You can install a single Docker runner, or install the Docker runner on multiple machines to create your own build cluster.

{{< link "/runner/docker/installation/linux.md" "Install the Runner for Linux" >}}
{{< link "/runner/docker/installation/windows.md" "Install the Runner for Windows" >}}

# When to Use?

The Docker runner is a general purpose runner, optimized for projects that can run tests and compile code inside stateless containers. If you are new to Drone, you probably want to [get started]({{< relref "installation/_index.md" >}}) with the Docker runner.

# When to Avoid?

The Docker runner is poorly suited for projects that cannot run tests or compile code inside containers, including projects that target operating systems or architectures not supported by Docker, such as macOS.

The Docker runner is also poorly suited for stateful pipelines that need to store files or folders on the host machine in-between pipeline executions. Docker pipelines are ephemeral and do not mutate the host machine.
