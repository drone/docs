---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
weight: 1
aliases:
- /runners
- /installation/runners/aws/
---

<div class="alert">
The AWS runner is in BETA and may not be suitable for production workloads. Furthermore this runner is a community effort and is not subject to support services or service level agreements at this time.
</div>

The AWS runner is a standalone service that executes pipelines on AWS EC2 instances. The AWS runner is similar to the exec runner in that you can run builds on a linux or Windows system. It also has the advantage of being very similar to the Docker runner, so you can run plugins and use containers.

{{< link "/runner/aws/installation.md" "Install the AWS Runner" >}}

## When to Use?

The Kubernetes runner is a general purpose runner, optimized for projects that can run tests and compile code on the EC2 instance. Or you want to run containerised steps and drone plugins.

It also supports instance pooling, whereby we provision systems on startup that are ready for an incoming build. This significantly decreases wait time for end users.

## When to Avoid?

If you want a production ready runner or you do not want to run builds off prem.

## Known Issues / Differences

AWS pipelines are considered experimental and may not be suitable for production use yet. You may experience unexpected differences and behaviors, some of which are detailed below.

* Not all of the syntax from the docker run is supported, we are adding more as we go.
* We do not support customisation of the cloudinit file that is used to instantiate the instance for now.
