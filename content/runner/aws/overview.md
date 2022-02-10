---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
weight: 1
---

<div class="alert">
The AWS runner is in Alpha and may not be suitable for production workloads. Furthermore this runner is a community effort and is not subject to support services or service level agreements at this time.
</div>

The AWS runner executes a build pipeline on a dedicated EC2 instance. The instance is created for each pipeline execution and terminated upon completion.

{{< link "/runner/aws/installation.md" "Install the AWS Runner" >}}

## When to Use?

The AWS runner is recommended for pipelines that require privileged access to a full virtual machine, and are unsuitable for execution inside containers. The AWS runner executes pipelines in dedicated, ephemeral virtual machines, and provides secure isolation for un-trusted workloads.

## When to Avoid?

If you want a production ready runner or you do not want to run builds off prem.

## Known Issues / Differences

AWS pipelines are considered experimental and may not be suitable for production use yet. You may experience unexpected differences and behaviors, some of which are detailed below.

* Not all of the syntax from the docker run is supported, we are adding more as we go.
