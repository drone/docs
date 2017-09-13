---
date: 2016-01-01T00:00:00+00:00
title: "Plugin Overview"
url: plugin-overview
weight: 1

menu:
  plugin:
    parent: pipeline
    identifier: plugin-overview
---

Plugins are Docker containers that perform pre-defined tasks and are configured as steps in your pipeline. Plugins can be used to deploy code, publish artifacts, send notification, and more.

Example pipeline using the Docker and Slack plugins:

```yaml
pipeline:
  backend:
    image: golang
    commands:
      - go get
      - go build
      - go test

  docker:
    image: plugins/docker
    username: kevinbacon
    password: pa55word
    email: kevin.bacon@mail.com
    repo: foo/bar
    tags: latest

  notify:
    image: plugins/slack
    channel: developers
    username: drone
```
