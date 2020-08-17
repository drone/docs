---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
---

The Nomad runner is a standalone service that executes Docker pipelines using the Nomad scheduler. The Nomad runner is functionally equivalent to the [Docker runner]({{< relref "runner/docker/overview.md" >}}) and shares the same pipeline syntax and execution engine.

# When to Use?

The Nomad runner is functionally equivalent to the Docker runner except it schedules pipelines using Nomad.  Use the Nomad runner if your organization has an existing Nomad cluster and Nomad expertise.
