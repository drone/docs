+++
date = "2017-04-15T14:39:04+02:00"
title = "Pipelines 工作流"
url = "zh/pipelines"

[menu.usage]
  weight = 4
  identifier = "pipelines-zh"
  parent = "usage_concepts"
+++

<!--The pipeline section defines a list of steps to build, test and deploy your code. Pipeline steps are executed serially, in the order in which they are defined. If a step returns a non-zero exit code, the pipeline immediately aborts and returns a failure status.-->

工作流部分定义了一系列步骤，包括代码构建，代码测试和代码部署。工作流根据各个步骤的定义位置按顺序执行。如果一个步骤返回了非 0 的退出代码（non-zero exit code），工作流将立即停止并返回一个错误状态。

<!--Example pipeline:-->

示例工作流

```yaml
pipeline:
  backend:
    image: golang
    commands:
      - go build
      - go test
  frontend:
    image: node
    commands:
      - npm install
      - npm run test
      - npm run build
```

<!--In the above example we define two pipeline steps, `frontend` and `backend`. The names of these steps are completely arbitrary.-->

在上面的示例中，我们定义了两个工作流步骤，`frontend` 和 `backend`。这两个步骤的名字可以是任意指定的。

# 构建步骤

<!--Build steps are steps in your pipeline that execute arbitrary commands inside the specified docker container. The commands are executed using the workspace as the working directory.-->

构建步骤是工作流在 Docker 容器中执行的任意命令。这些命令将工作区（workspace）作为工作路径。

```diff
pipeline:
  backend:
    image: golang
    commands:
+     - go build
+     - go test
```

<!--There is no magic here. The above commands are converted to a simple shell script. The commands in the above example are roughly converted to the below script:-->

这里没有什么魔法。上面的命令（commands）将会被转换成简单的 Shell 脚本。上面的命令被大致转换为下面的脚本：

```diff
#!/bin/sh
set -e

go get
go build
go test
```

<!--The above shell script is then executed as the docker entrypoint. The below docker command is an (incomplete) example of how the script is executed:-->

上面的脚本在之后会被作为 Docker 的入口。下面的 docker 命令，作为一个不完整的例子，展示了这个脚本是如何被执行的。

```
docker run --entrypoint=build.sh golang
```

<!--Please note that only build steps can define commands. You cannot use commands with plugins or services.-->

{{% alert info %}}
请注意，只有构建步骤可以定义命令，您不能使用插件（plugins）或者服务（services）来定义命令。
{{% /alert %}}

# 并行执行

<!--Drone supports parallel step execution for same-machine fan-in and fan-out. Parallel steps are configured using the `group` attribute. This instructs the pipeline runner to execute the named group in parallel.-->

Drone 支持在同一个示例上并行执行多个步骤。使用 `group` 属性来配置并行步骤，这将让工作流执行者（pipeline runner）并行执行指定的命令。

<!--Example parallel configuration:-->

并行执行配置示例：

```diff
pipeline:
  backend:
+   group: build
    image: golang
    commands:
      - go build
      - go test
  frontend:
+   group: build
    image: node
    commands:
      - npm install
      - npm run test
      - npm run build
  publish:
    image: plugins/docker
    repo: octocat/hello-world
```

<!--In the above example, the `frontend` and `backend` steps are executed in parallel. The pipeline runner will not execute the `publish` step until the group completes.-->

在上面到例子中，`frontend` 和 `backend`将并行执行。在这两组任务完成之前，`publish` 步骤将不会执行。

# 条件执行

<!--Drone gives you the ability to conditionally limit the execution of steps at runtime. The below example limits execution of Slack plugin steps based on branch:-->

Drone 可以有条件地执行步骤。下面的例子限制了使用 Slack 插件的 Git 分支：

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
+   when:
+     branch: master
```

# 故障执行

<!--Drone uses the container exit code to determine the success or failure status of a build. Non-zero exit codes fail the build and cause the pipeline to immediately exit.-->

Drone 使用容器退出代码来决定一个构建到成功或者失败。非 0 退出代码（Non-zero exit codes）使构建失败，同时立即退出当前工作流。

<!--There are use cases for executing pipeline steps on failure, such as sending notifications for failed builds. Use the status constraint to override the default behavior and execute steps even when the build status is failure:-->

有的时候需要在构建失败时，执行特定的工作流步骤。可以使用状态条件限制（status constraint）来修改构建失败时的默认行为和执行步骤。

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
+   when:
+     status: [ success, failure ]
```
