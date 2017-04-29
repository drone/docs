+++
date = "2017-04-15T14:39:04+02:00"
title = "Cloning 克隆"
url = "zh/cloning"

[menu.usage]
  weight = 3
  identifier = "cloning-zh"
  parent = "usage_concepts"
+++

<!--Drone automatically configures a default clone step if not explicitly defined. You can manually configure the clone step in your pipeline for customization:-->

您可以在工作流中手动配置 clone 步骤。如果没有定义具体的 clone 步骤，Drone 会自动配置默认的 clone 步骤。

```diff
+clone:
+  git:
+    image: plugins/git

pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test
```

<!--Example configuration to override depth:-->

修改 clone 深度的配置示例

```diff
clone:
  git:
    image: plugins/git
+   depth: 50
```

<!--Example configuration to use a custom clone plugin:-->

使用自定义 clone 插件的配置示例

```diff
clone:
  git:
+   image: octocat/custom-git-plugin
```
