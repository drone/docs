---
date: 2000-01-01T00:00:00+00:00
title: Logging
author: bradrydzewski
weight: 2

toc: true
aliases:
- /runners/exec/logging
description: |
  Configure trace and debug logging.
---

The exec runner is installed as a service and writes logs to a file on the host machine. The log file location must be configured before you start the service.

On Posix (root):

```
DRONE_LOG_FILE=/var/log/drone-runner-exec/log.txt
```

On Posix (non-root):

```
DRONE_LOG_FILE=~/.drone-runner-exec/log.txt
```

On windows:

```
DRONE_LOG_FILE=C:\Drone\drone-runner-exec\log.txt
```

# Log Rotation

The exec runner supports automatic rotation and compression of log files. You can configure the maximum file size in megabytes, number of days to retain, and number of files to retain.

```
DRONE_LOG_FILE_MAX_SIZE=10
DRONE_LOG_FILE_MAX_AGE=30
DRONE_LOG_FILE_MAX_BACKUPS=7
```

# Log Levels

The exec runner is configured to log runtime events. You can enable debug or trace level logs to get detailed information on the flow through the system.

```
DRONE_DEBUG=true
DRONE_TRACE=true
```

# Request Logging

You can enable http request logging to get detailed http communication details between the runner and the server. _This generates significant output and should only be enabled to troubleshoot communication issues._

```
DRONE_RPC_DUMP_HTTP=true
DRONE_RPC_DUMP_HTTP_BODY=true
```