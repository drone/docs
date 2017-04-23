+++
date = "2017-04-15T14:39:04+02:00"
title = "Plugins"
url = "using-plugins"

[menu.usage]
  weight = 7
  identifier = "plugins"
  parent = "usage_concepts"
+++

Plugins are Docker containers that perform pre-defined tasks and are configured as steps in your pipeline. Plugins can be used to deploy code, publish artifacts, send notification, and more.

Example pipeline using the Docker and Slack plugins:

```yaml
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

  publish:
    image: plugins/docker
    repo: foo/bar
    tags: latest

  notify:
    image: plugins/slack
    channel: dev
```

# Plugin Isolation

Plugins are executed in Docker containers and are isolated from the other steps in your build pipeline. Plugins do share the build workspace, mounted as a volume, and therefore have access to your source tree.

# Plugin Marketplace

Plugins are packaged and distributed as Docker containers. They are conceptually similar to software libraries (think npm) and can be published and shared with the community. You can find a list of available plugins at [http://plugins.drone.io](http://plugins.drone.io).
