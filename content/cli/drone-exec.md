---
date: 2000-01-01T00:00:00+00:00
title: drone exec
author: bradrydzewski
weight: 4
aliases:
- /cli-exec
- /cli/drone-exec
---

This subcommand executes a local build.

```
$ drone exec --help
NAME:
   drone exec - execute a local build

USAGE:
   drone exec [command options] [path/to/.drone.yml]

OPTIONS:
   --pipeline value        name of the pipeline to execute
   --include value         name of steps to include
   --exclude value         name of steps to exclude
   --resume-at value       name of start to resume at
   --trusted               build is trusted
   --timeout value         build timeout (default: 1h0m0s)
   --secret-file value     secret variable file
   --env-file value        environment variable file
```

# Quick Start

Please see our quick start guide to learn how to execute your pipeline on your local machine using the command line tools:

{{< link "quickstart/cli" "Quick Start" >}}
