---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
related:
  items:
  - name: Example Go Plugin
    path: golang.md
  - name: Example Bash Plugin
    path: bash.md

aliases:
- /plugin-overview/
- /plugin-guidelines/
---

Plugins are Docker containers that perform pre-defined tasks and are configured as steps in your pipeline. Plugins can be used to deploy code, publish artifacts, send notification, and more.

Example pipeline using the Docker and Slack plugins:

```yaml  {linenos=table}
kind: pipeline
type: docker
name: default

steps:
- name: build
  image: golang
  commands:
  - go get
  - go test
  - go build

- name: publish
  image: plugins/docker
  settings:
    username: kevinbacon
    password: pa55word
    repo: foo/bar
    tags:
    - 1.0.0
    - 1.0

- name: notify
  image: plugins/slack
  settings:
    channel: developers
    username: drone
```

_Plugins are just Docker containers which means you can write plugins in any programming language that runs inside a container. You can even create plugins using simple bash scripting._

# Plugin Inputs

Plugins parameters are defined in the settings section of the pipeline step and are passed to the plugin container as environment variables. The environment variables are prefixed to prevent naming conflicts.

* Example plugin configuration.
   ```
   - name: publish
     image: plugins/docker
     settings:
       username: kevinbacon
       password: pa55word
       repo: foo/bar
       tags:
       - 1.0.0
       - 1.0
   ```

* Example plugin variables passed to the container.
  ```
  PLUGIN_USERNAME=kevinbacon
  PLUGIN_PASSWORD=pa55word
  PLUGIN_REPO=foo/bar
  PLUGIN_TAGS=1.0.0,1.0
  ```

Plugin parameters can be any primitive type or array of primitive types. Arrays are converted to comma-separate strings.


<!-- Plugins parameters are defined in the settings section of your configuration:

```yaml  {linenos=table,linenostart=13}
- name: publish
  image: plugins/docker
  settings:
    username: kevinbacon
    password: pa55word
    repo: foo/bar
    tags:
    - 1.0.0
    - 1.0
```

Plugin parameters are passed to the plugin as environment variables and are prefixed to prevent naming conflicts. Example using the parameters from the previous example:

```
docker run \
-e PLUGIN_USERNAME=kevinbacon
-e PLUGIN_PASSWORD=kevinbacon
-e PLUGIN_REPO=foo/bar
-e PLUGIN_TAGS=1.0.0,1.0
```

Plugin parameters can be any primitive type or array of primitive types. Arrays are converted to comma-separate strings. -->






