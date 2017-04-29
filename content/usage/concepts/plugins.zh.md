+++
date = "2017-04-15T14:39:04+02:00"
title = "Plugins 插件"
url = "zh/using-plugins"

[menu.usage]
  weight = 7
  identifier = "plugins-zh"
  parent = "usage_concepts"
+++

<!--Plugins are Docker containers that perform pre-defined tasks and are configured as steps in your pipeline. Plugins can be used to deploy code, publish artifacts, send notification, and more.-->

插件是执行预定义任务的容器，它们在工作流中被配置为步骤（steps）。插件可以用来部署代码，发布构建结果，发送通知以及部署其他更多的功能。

<!--Example pipeline using the Docker and Slack plugins:-->

Docker 和 Slack 插件工作流示例：

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

# 插件隔离

<!--Plugins are executed in Docker containers and are isolated from the other steps in your build pipeline. Plugins do share the build workspace, mounted as a volume, and therefore have access to your source tree.-->

插件在 Docker 容器中执行，它们与工作流中的其他步骤相互隔离。注意，插件挂载和共享当前工作区，因此它将可以访问对应源代码。

# 插件市场

<!--Plugins are packaged and distributed as Docker containers. They are conceptually similar to software libraries (think npm) and can be published and shared with the community. You can find a list of available plugins at [http://plugins.drone.io](http://plugins.drone.io).-->

插件被打包和发布为 Docker 容器，它们在概念上和软件库类似（比如 npm），可以被在社区中发布和共享。可以在 [http://plugins.drone.io](http://plugins.drone.io) 找到一系列可用的插件。
