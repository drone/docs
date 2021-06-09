---
date: 2000-01-01T00:00:00+00:00
title: Starklark
author: eoinmcafee
weight: 20
disable_toc: true
aliases:
- /configure/templates/jsonnet/
description: Store and manage starklark templates per-organization.
---
Templates can be used by setting the kind to 'template', the name of the base template to load, and a set of freeform template inputs
in their drone.yml file:
```
kind: template
load: plugin.starlark
data:
  stepName: my_step
  image: my_image
  commands: my_command
```
Example base template:
```
def main(ctx):
  return {
    "kind": "pipeline",
    "name": "build",
    "steps": [
      {
        "name": ctx.input.stepName,
        "image": ctx.input.image,
        "commands": [
            ctx.input.commands
        ]
      }
    ]
  }
```
Create organization templates using the command line tools:
```
$ drone template add [namespace] [name] [data]
```
```
$ drone template add octocat template_name @path_to_file
```