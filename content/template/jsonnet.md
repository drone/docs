---
date: 2000-01-01T00:00:00+00:00
title: Jsonnet
author: eoinmcafee
weight: 20
disable_toc: true
aliases:
- /configure/templates/jsonnet/
description: Store and manage jsonnet templates per-organization.
---
Templates can be used by setting the kind to 'template', the name of the base template to load, and a set of freeform template inputs
in their drone.yml file:
```
kind: template
load: plugin.jsonnet
data:
  stepName: my_step
  image: my_image
  commands: my_command
```
Example base template:
```
local stepName = std.extVar("input.stepName");
local image = std.extVar("input.image");
local commands = std.extVar("input.commands");
{
  "kind": "pipeline",
  "type": "docker",
  "name": "default",
  "steps": [
    {
      "name": stepName,
      "image": image,
      "commands": [
        commands
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