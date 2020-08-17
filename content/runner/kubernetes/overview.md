---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /installation/runners/kubernetes/
---

<div class="alert">
The Kubernetes runner is in Beta and may not be suitable for production workloads. Furthermore this runner is a community effort and is not subject to support services or service level agreements at this time.
</div>

The Kubernetes runner is a standalone service that executes pipelines inside Pods. The Kubernetes runner is very similar to the Docker runner, and should be used when running Drone on Kubernetes.

{{< link "/runner/kubernetes/installation.md" "Install the Kubernetes Runner" >}}

# When to Use?

The Kubernetes runner is a general purpose runner, optimized for projects that can run tests and compile code inside stateless containers. If your goal is to execute pipelines inside your Kubernetes cluster you definitely want to use the Kubernetes runner.

# When to Avoid?

The Kubernetes runner is poorly suited for projects that cannot run tests or compile code inside containers, including projects that target operating systems or architectures not supported by Docker, such as macOS.

The Kubernetes runner is also poorly suited for stateful pipelines that need to store files or folders on the host machine in-between pipeline executions. Kubernetes pipelines are ephemeral and do not mutate the host machine.

# Known Issues / Differences

Kubernetes pipelines are considered experimental and may not be suitable for production use yet. You may experience unexpected differences and behaviors, some of which are detailed below.

* Unlike docker pipelines, the pipeline status is passed to steps by file as opposed to environment variable. Existing plugins may not be compatible with kubernetes pipelines and will need to be patched accordingly. See [how we patched](https://github.com/drone-plugins/drone-slack/commit/cb7ba0bbbaff73c8c14f28431c0ff2866368ef50) the Slack plugin.

* The command line utility does not support linting, formatting or execution of Kubernetes pipelines.

* The Kubernetes runner must not be restarted while pipelines are running. If you stop or restart the runner while a pipeline is running, it will be stuck in a running state, and the associated Pod must be manually removed.
