---
date: 2000-01-01T00:00:00+00:00
title: Yaml
author: eoinmcafee
weight: 20
disable_toc: true
aliases:
- /configure/templates/yaml/
description: Store and manage yaml templates per-organization.
---
Templates can be used by setting the kind to 'template', the name of the base template to load, and a set of freeform template inputs
in their drone.yml file:
```
kind: template
load: plugin.yaml
data:
  name: name
  image: image
  commands: commands
```
Example base template:
```
kind: pipeline
type: docker
name: default
steps:
   - name: {{ .input.name }}
     image: {{ .input.image }}
     commands:
        - {{ .input.commands }}
```
Predefined variable list can be found at: https://docs.drone.io/template/variables

Create organization templates using the command line tools:
```
$ drone template add [namespace] [name] [data]
```
```
$ drone template add octocat template_name @path_to_file
```