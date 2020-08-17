---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
---

The Digital Ocean runner executes a build pipeline on a dedicated droplet using the SSH protocol. The droplet is created for each pipeline execution and terminated upon completion.

{{< link "/runner/digitalocean/installation.md" "Install the Digital Ocean Runner" >}}

# When to Use?

The Digital Ocean runner is recommended for pipelines that require privileged access to a full virtual machine, and are unsuitable for execution inside containers. The Digital Ocean runner executes pipelines in dedicated, ephemeral virtual machines, and provides secure isolation for un-trusted workloads.

__tldr;__ if your pipeline is unsafe to run inside a container, you may want to consider the Digital Ocean runner.
