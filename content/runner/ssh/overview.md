---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /runners/ssh
---

The SSH runner executes pipeline commands on a static, remote server using the SSH protocol. The pipeline commands are executed directly on the remote server without isolation, using the default shell.

This runner is not suitable for un-trusted workloads for security reasons, nor is it recommended for use with public repositories. Do not use this runner if you do not trust your contributors.

{{< link "/runner/ssh/installation.md" "Install the SSH Runner" >}}

# Known Issues

Please ensure you are using openssh 7.9 or higher on your servers. Prior versions of openssh do not support the signal command and do not signal remote processes. As a result, when a pipeline is canceled or times out, the runner cannot terminate remote pipeline processes.
