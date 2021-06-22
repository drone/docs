---
date: 2000-01-01T00:00:00+00:00
title: Templates
author: eoinmcafee
weight: 130
header: true
aliases:
- /user-guide/templates
---

# _Why use templates?_

Large organizations can have difficulty maintaining large numbers of build configuration files. Very often those configuration files are very similar or almost identical across projects. Some changes need to be made across all projects, for example needing to change the version docker image.

To make this simpler Drone has build templates that can be shared across projects. A project can use a template, and provide project specific information to alter the build.

This will simplify configuration management at organizations that have large numbers of similar configuration files; it will reduce complexity of yaml files in individual repositories and will simplify management of configuration across repositories.

Here are the articles in this section:

{{< cards >}}