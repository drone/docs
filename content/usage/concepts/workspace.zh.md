+++
date = "2017-04-15T14:39:04+02:00"
title = "Workspace 工作区"
url = "zh/workspace"

[menu.usage]
  weight = 2
  identifier = "workspace-zh"
  parent = "usage_concepts"
+++

<!--The workspace defines the shared volume and working directory shared by all pipeline steps. The default workspace matches the below pattern, based on your repository url.-->

工作区去定义了所有工作流步骤共享的容器空间和目录。默认的工作区的目录和仓库 URL 相匹配，如下面的例子所示： 

```
/drone/src/github.com/octocat/hello-world
```

<!--The workspace can be customized using the workspace block in the Yaml file:-->

可以使用 Yaml 文件的 `workspace` 区块来自定义工作区。

```diff
+workspace:
+  base: /go
+  path: src/github.com/octocat/hello-world

pipeline:
  build:
    image: golang:latest
    commands:
      - go get
      - go test
```

<!--The base attribute defines a shared base volume available to all pipeline steps. This ensures your source code, dependencies and compiled binaries are persisted and shared between steps.-->

base 属性定义了所有工作流步骤共享的基础容器空间。基础容器空间保证了代码、依赖和编译的二进制文件能够在各步骤间持久化和共享。

```diff
workspace:
+ base: /go
  path: src/github.com/octocat/hello-world

pipeline:
  deps:
    image: golang:latest
    commands:
      - go get
      - go test
  build:
    image: node:latest
    commands:
      - go build
```

<!--This would be equivalent to the following docker commands:-->

上面的步骤将和下面的 docker 命令类似：

```
docker volume create my-named-volume

docker run --volume=my-named-volume:/go golang:latest
docker run --volume=my-named-volume:/go node:latest
```

<!--The path attribute defines the working directory of your build. This is where your code is cloned and will be the default working directory of every step in your build process. The path must be relative and is combined with your base path.-->

path 属性定义了构建的工作目录。这是代码被克隆到的目录，也将是每一个构建步骤的默认工作目录。这个路径必须是基于 base 路径的相对路径。

```diff
workspace:
  base: /go
+ path: src/github.com/octocat/hello-world
```

```text
git clone https://github.com/octocat/hello-world \
  /go/src/github.com/octocat/hello-world
```
