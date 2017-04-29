+++
date = "2017-04-15T14:39:04+02:00"
title = "Webhooks 网络钩子"
url = "zh/hooks"

[menu.usage]
  weight = 1
  identifier = "hooks-zh"
  parent = "usage_concepts"
+++

<!--When you activate your repository Drone automatically add webhooks to your version control system (e.g. GitHub). There is no manual configuration required.-->

当一个仓库被激活，Drone 会自动将 webhooks 添加到对应的版本控制系统中（比如 GitHub）。不需要手动设置。

<!--Webhooks are used to trigger pipeline executions. When you push code to your repository, open a pull request, or create a tag, your version control system will automatically send a webhook to Drone which will in turn trigger pipeline execution.-->

Webhooks 被用来触发工作流执行。当代码被推送到仓库，当新建一个合并请求（pull request）时，或者当新建一个标签时，版本控制系统将会自动发送一个 webhook 请求到 Drone，这将会触发工作流执行。

<!-- # Recreate Webhooks

Drone provides the ability to recreate webhooks, in case they were accidentally removed or altered, using the command line utility.

```text
drone repo repair <repo>
drone repo repair octocat/hello-world
``` -->

# 跳过提交

<!--Drone gives the ability to skip individual commits by adding `[CI SKIP]` to the commit message. Note this is case-insensitive.-->

可以通过添加 `[CI SKIP]` 到提交信息（commit message）中来让 Drone 跳过某个提交。注意，这里的文本是大小写不敏感的。

```diff
git commit -m "updated README [CI SKIP]"
```

# 跳过分支

<!--Drone gives the ability to skip commits based on the target branch. The below example will skip a commit when the target branch is not master.-->

Drone 可以忽略某个分支上的提交。下面的例子将会跳过不是 master 分支的提交。

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches: master
```

<!--Example matching multiple target branches:-->

匹配多个目标分支的例子：

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches: [ master, develop ]
```

<!--Example uses glob matching:-->

批量匹配的例子：

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches: [ master, feature/* ]
```

包含分支的例子：

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches:
+  include: [ master, feature/* ]
```

忽略分支的例子：

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

+branches:
+  exclude: [ develop, feature/* ]
```
