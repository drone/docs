+++
date = "2017-04-15T14:39:04+02:00"
title = "Deployments 部署"
url = "zh/deployments"

[menu.usage]
  weight = 8
  identifier = "deployment-zh"
  parent = "usage_concepts"
+++

<!--Drone provides the ability to trigger deployments. When you trigger a deployment your pipeline is executed with event type `deployment`. You can use the event type and target environment to limit step execution.-->

您可以触发 Drone 来部署项目。部署项目在工作流中的事件类型（event type）是 `deployment`。您可以使用事件类型（event type）或者目标环境变量（target environment）来限制执行的步数。

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

  publish:
    image: plugins/docker
    registry: registry.heroku.com
    repo: registry.heroku.com/my-staging-app/web
    when:
+     event: deployment
+     environment: staging

  publish_to_prod:
    image: plugins/docker
    registry: registry.heroku.com
    repo: registry.heroku.com/my-production-app/web
    when:
+     event: deployment
+     environment: production
```

<!--The above example demonstrates how we can configure pipeline steps to only execute when the deployment matches a specific target environment.-->

上面的例子展示了如何设置工作流来只在特定的目标环境中部署项目，比如 production 生产环境。

# 触发部署

<!--Deployments are triggered from the command line utility. They are triggered from an existing build. This is conceptually similar to promoting builds.-->

可以使用命令行从已有的构建结果中触发部署。这个行为与使用构建结果（promoting builds）的概念类似。

```text
drone deploy <repo> <build> <environment>
```

<!--Promote the specified build number to your staging environment:-->

在准生产环境（staging environment）中使用某一序号的构建结果。

```text
drone deploy octocat/hello-world 24 staging
```

<!--Promote the specified build number to your production environment:-->

在生产环境（production environment）中使用某一序号的构建结果。


```text
drone deploy octocat/hello-world 24 production
```
