---
date: 2000-01-01T00:00:00+00:00
title: Support
author: bradrydzewski
weight: 3
---

Plugins are currently maintained and supported by the community.  _We are currently compiling a list of plugins that are certified by Harness and will update this page in the coming weeks. Certified plugins will be subject to Harness support and service level agreements._

# Contributing

Many plugins in the Drone ecosystem are open source and can be forked and modified to meet your unique business needs. We highly recommend forking plugins and sharing your changes with the community, either by publishing your changes or by submitting pull requests to the upstream repository.

If you need to use the forked version of your plugin, you can publish your fork to Dockerhub, and modify you pipeline configuration to use the forked image. For example:

{{<highlight diff "linenos=table" >}}
kind: pipeline
type: docker
name: default

steps:
  - name: slack
-   image: plugins/slack
+   image: my-fork/slack
    settings:
      channel: dev
      webhook:
        from_secret: slack_webhook
{{< / highlight >}}
