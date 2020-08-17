---
date: 2000-01-01T00:00:00+00:00
title: Install On Windows
title_in_sidebar: On Windows
author: bradrydzewski
weight: 1
steps: true
toc: true
aliases:
- /runners/exec/windows
description: |
  Install the runner on Windows Server.
---

This article explains how to install the exec runner on Windows. The exec runner is packaged in binary format and distributed as a Github [release](https://github.com/drone-runners/drone-runner-exec/releases).

# Download

Download and unpack the binary.

```
https://github.com/drone-runners/drone-runner-exec/releases/latest/download/drone_runner_exec_windows_amd64.tar.gz
```

# Configuration

The exec runner is configured using an environment variable file. The environment file should be stored at the following path: 

```
C:\Drone\drone-runner-exec\config
```

This file should use the syntax `<variable>=value` which sets the variable to the given value and `#` for comments. Please note this is not a bash file. Bash syntax and Bash expressions are not supported.

```
DRONE_RPC_PROTO=https
DRONE_RPC_HOST=drone.company.com
DRONE_RPC_SECRET=super-duper-secret
```

This article references the below configuration options. See [Configuration]({{< relref "../configuration/reference" >}}) for a complete list of configuration options.

* __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.

* __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.

* __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

# Logging

The exec runner writes logs to a file on the host machine. The log file location should be configured in the environment file before you start the service.

```
DRONE_LOG_FILE=C:\Drone\drone-runner-exec\log.txt
```

The log file directory must be created before you start the service:

```
$ mkdir C:\Drone
$ mkdir C:\Drone\drone-runner-exec
```

# Installation

Install and start the service.

```
$ drone-runner-exec.exe service install
$ drone-runner-exec.exe service start
```

# Verification

Inspect the logs and verify the runner successfully established a connection with the Drone server.

```
$ get-content -Path C:\Drone\drone-runner-exec\log.txt

INFO[0000] starting the server
INFO[0000] successfully pinged the remote server
```