---
date: 2000-01-01T00:00:00+00:00
title: drone sign
author: bradrydzewski
weight: 4
---

This subcommand signs your yaml configuration file to prevent tampering. See [Signatures]({{< relref "/signature/_index.md" >}}) for complete instructions.

```
NAME:
   drone sign - sign the yaml file

USAGE:
   drone sign [command options] <source>

OPTIONS:
   --save  save result to source
```

Example usage:

```
drone sign --save octocat/hello-world 
```
