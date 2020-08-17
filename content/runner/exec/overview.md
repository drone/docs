---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /installation/runners/exec/
- /runners/exec
- /runners/exec/configure
---

<div class="alert">
The Exec runner is in Beta and may not be suitable for production workloads.
</div>

The Exec runner is a daemon that executes build pipelines directly on the host machine without isolation, using the default shell. This runner is not suitable for un-trusted workloads for security reasons.

{{< link "/runner/exec/installation/linux.md" "Install the Exec Runner for Linux" >}}
{{< link "/runner/exec/installation/macos.md" "Install the Exec Runner for MacOS" >}}
{{< link "/runner/exec/installation/windows.md" "Install the Exec Runner for Windows" >}}

# When to Use?

The Exec runner is recommended for workloads that cannot be executed inside containers, or cannot easily be adapted to container-based pipelines. This runner is especially useful for operating systems and architectures that do not support containers, such as macOS.

# When to Avoid?

The Exec runner is not suitable for untrusted workloads due to lack of isolation. _We recommend using the Docker runner for all projects by default, and only using the Exec runner for the subset of projects requiring host machine access._
