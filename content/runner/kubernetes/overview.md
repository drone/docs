---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /installation/runners/kubernetes/
---

The Kubernetes runner is a standalone service that executes pipelines inside Pods. The Kubernetes runner is very similar to the Docker runner, and should be used when running Drone on Kubernetes.

{{< link "/runner/kubernetes/installation.md" "Install the Kubernetes Runner" >}}

# When to Use?

The Kubernetes runner is a general purpose runner, optimized for projects that can run tests and compile code inside stateless containers. If your goal is to execute pipelines inside your Kubernetes cluster you definitely want to use the Kubernetes runner.

# When to Avoid?

The Kubernetes runner is poorly suited for projects that cannot run tests or compile code inside containers, including projects that target operating systems or architectures not supported by Docker, such as macOS.

The Kubernetes runner is also poorly suited for stateful pipelines that need to store files or folders on the host machine in-between pipeline executions. Kubernetes pipelines are ephemeral and do not mutate the host machine.

# Differences

Kubernetes pipelines are considered experimental and may not be suitable for production use yet. You may experience unexpected differences and behaviors, some of which are detailed below.

* Unlike docker pipelines, the pipeline status is passed to steps by file as opposed to environment variable. Existing plugins may not be compatible with kubernetes pipelines and will need to be patched accordingly. See [how we patched](https://github.com/drone-plugins/drone-slack/commit/cb7ba0bbbaff73c8c14f28431c0ff2866368ef50) the Slack plugin.

* The command line utility does not support linting, formatting or execution of Kubernetes pipelines.

# Known Issues

Below are some issues that can occur when running your pipelines inside Kubernetes. This is due to its inherent behavior. The runner listens for events from Kubernetes and sometimes these events are not sent successfully.  We use the [v0.21.4](https://github.com/kubernetes/client-go/releases/tag/v0.21.4) version of the library, that lines up with the Kubernetes v1.21.4 release. For their support matrix look [here](https://github.com/kubernetes/client-go/#compatibility-matrix).

| Error | Description | Mitigation |
|------|-------------|-------------|
| "unknown container:" | An error that is returned when an unregistered container name is provided | Use a registered container name, or check connectivity to your container registry. |
| "pod is terminated" | An error that is returned when the pod is already terminated. | Check to see if another service has terminated the pod allocate more memory to kubernetes. |
| "kubernetes has failed: container failed to start" | An error returned placeholder container terminates abnormally. The correct container image failed to load. | Usually happens when image doesn't exist, or check connectivity to your container registry.  |
| "kubernetes has failed: container failed to start in timely manner:" | An error when a container fails to run after some predefined time. This can be due to any number of network / registry issues. | You can configure the timeout threshold with the setting `DRONE_ENGINE_CONTAINER_START_TIMEOUT` |
| "aborting due to error" | An error by wait function when some other container in the same pod fails with a "kubernetes has failed" error.| Investigate the issue that caused the original container to fail. |
| Pods stuck in a running state | The Kubernetes runner looks after the lifecycle of the pods, and must not be restarted while pipelines are running. The runner does not attempt to kill pods on start. | These pods need to be manually removed. |

# Our test setup

We used the following criteria for testing the Kubernetes runner inside of GCR:

| PRs/Day |  Nodes with 4 CPU, 8GB RAM,100GB disk | Nodes with 8 CPU, 16GB RAM, 200GB disk |
|---------|------------------------------|------------------------------|
| 100 | 19 - 26  | 11 - 15 |
| 500 | 87 - 121 | 45 - 62 |
| 1000 | 172 - 239 | 89 - 123 |
