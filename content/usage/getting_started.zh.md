---
title:  入门
url: zh/getting-started

next_steps:
  - file: usage/concepts/pipelines.zh.md
    name: Pipeline 配置
  - file: usage/concepts/services.zh.md
    name: Services 服务配置
  - file: usage/publishing/docker.zh.md
    name: 发布到 Dockerhub
  - file: usage/deployments/kubernetes.zh.md
    name: 部署到 Kubernetes

menu:
  usage:
    weight: 1
    identifier: getting-started-zh
    parent: usage_overview
---

<!--Welcome to the Drone community. This document briefly explains the process for activating and configuring a continuous delivery pipeline.-->

欢迎来到 Drone 社区。这个页面简单地介绍了启动和配置一个持续交付工作流（continuous delivery pipeline）的步骤。

# 启动

<!--To activate your project navigate to your account settings. You will see a list of repositories which can be activated with a simple toggle. When you activate your repository, Drone automatically adds webhooks to your version control system (e.g. GitHub).-->

访问账户设置页面（account settings）来激活您的项目。您将看到一列可以使用摇杆（toggle）激活的仓库。当您激活了一个仓库，Drone 将自动添加 webhooks 到对应的版本控制系统中（比如 GitHub）。

<!--Webhooks are used to trigger pipeline executions. When you push code to your repository, open a pull request, or create a tag, your version control system will automatically send a webhook to Drone which will in turn trigger pipeline execution.-->

Webhooks 是用来激活工作流运行的。当您将代码推送到仓库时，当新开启一个合并请求（pull request），或者当新建一个标签（tag）时，您的版本控制系统将会自动发送一个 webhook 给 Drone，Drone 紧接着触发工作流运行。

![repository list](/images/drone_repo_list.png)

# 配置

<!--To configure you pipeline you should place a `.drone.yml` file in the root of your repository. The .drone.yml file is used to define your pipeline steps. It is a superset of the widely used docker-compose file format.-->

在仓库的根目录放置 `.drone.yml` 文件来配置您的工作流。`.drone.yml` 文件用来定义您的工作流文件。它是被广泛使用的 docker-compose 文件的一个超集。

<!--Example pipeline configuration:-->

工作流配置的例子：

```yaml
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go build
      - go test

services:
  postgres:
    image: postgres:9.4.5
    environment:
      - POSTGRES_USER=myapp
```

<!--Example pipeline configuration with multiple, serial steps:-->

包含多个步骤的工作流配置的例子：

```yaml
pipeline:
  backend:
    image: golang
    commands:
      - go get
      - go build
      - go test

  frontend:
    image: node:6
    commands:
      - npm install
      - npm test

  notify:
    image: plugins/slack
    channel: developers
    username: drone
```

# 执行

<!--To trigger your first pipeline execution you can push code to your repository, open a pull request, or push a tag. Any of these events triggers a webhook from your version control system and execute your pipeline.-->

您可以通过推送代码到您的仓库，开启一个合并请求，推送标签来激活工作流运行。任一这类事件都将会激活您的版本控制系统的 webhook，同时启动工作流。

<!--You can view your pipeline execution in realtime in the user interface.-->

您可以在用户界面中实时地看到工作流执行的情况。

![running build](/images/drone_build_running.png)
