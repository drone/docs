---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: tphoney
weight: 1
---

<div class="alert">
The VM runner is in the Release Candidate phase.
</div>

The VM runner executes a build pipeline on a dedicated VM instance. The instance will be created for each pipeline execution and terminated upon completion.

## When to Use?

The VM runner is recommended for pipelines that require privileged access to a full virtual machine, and are unsuitable for execution inside containers. The VM runner executes pipelines in dedicated, ephemeral virtual machines, and provides secure isolation for un-trusted workloads.

## When to Avoid?

If you want a production ready runner or you do not want to run builds off prem.

## Known Issues / Differences

VM pipelines are considered experimental and may not be suitable for production use yet. You may experience unexpected differences and behaviors, some of which are detailed below.

* Changes between RC4 and latest are documented [here]({{< relref "configuration/migration.md" >}}).
