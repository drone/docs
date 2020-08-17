---
date: 2000-01-01T00:00:00+00:00
title: drone jsonnet
author: bradrydzewski
weight: 5
toc: false
---

This subcommand converts a jsonnet configuration file to a yaml configuration file.

```
$ drone jsonnet --help
NAME:
   drone jsonnet - generate .drone.yml from jsonnet

USAGE:
   drone jsonnet [command options] [path/to/.drone.jsonnet]

OPTIONS:
   --source value  Source file (default: ".drone.jsonnet")
   --target value  target file (default: ".drone.yml")
   --stream        Write output as a YAML stream.
   --format        Write output as formatted YAML
   --stdout        Write output to stdout
   --string        Expect a string, manifest as plain text
```

Example prints the yaml file to stdout

```
$ drone jsonnet --stdout
```

Example prints a multi-document yaml to stdout

```
$ drone jsonnet --stdout --stream
```
