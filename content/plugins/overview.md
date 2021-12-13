---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
toc: true
weight: 1
related:
  items:
  - name: Example Go Plugin
    path: tutorials/golang.md
  - name: Example Bash Plugin
    path: tutorials/bash.md

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


# Source Code

Plugins automatically have access to the relevant source code and commit for a build, mounted as a volume into the plugin container. The plugin is also started with the current working directory set to the root of the git repository. _The plugin does not need to clone or checkout code; this is handled by Drone._

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

# Plugin Distribution

Plugins are distributed as Docker images. You can publish plugins to any Docker registry, private or public, to share plugins internally with your organization, or publicly with the broader developer community.

# Plugin Registry

The Drone plugin registry is a listing of Open Source plugins created by the Drone community. Want to add your plugin to the registry? _Send us a [pull request](https://github.com/drone/drone-plugin-index) that adds your plugin to the registry website_.

{{< link-to "http://plugins.drone.io" >}}
Browse the Plugin Registry
{{< / link-to >}}

<!-- # Common Questions

## Which plugin should I choose when multiple options exist?

You may encounter plugins in the registry that overlap in functionality (e.g. multiple kubernetes plugins). This is conceptually no different than searching for a packages on npm, and finding multiple npm packages that overlap in functionality. Here are some things to consider when choosing a plugin:

- Which plugins are the most widely used?
- Which plugins are the most actively developed?
- Number of starts, forks, contributors, downloads?
- Are there major open issues? major bugs? -->
