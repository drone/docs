---
date: 2000-01-01T00:00:00+00:00
title: Logging
author: tphoney
weight: 2
description: |
  Configure trace and debug logging.
---

The AWS runner writes logs to stdout and stderr. You can enable debug or trace level logs to get detailed information on the flow through the system.

```
DRONE_DEBUG=true
DRONE_TRACE=true
```

You can enable http request logging to get detailed http communication details between the runner and the server. _This generates significant output and should only be enabled to troubleshoot communication issues._

```
DRONE_RPC_DUMP_HTTP=true
DRONE_RPC_DUMP_HTTP_BODY=true
```
