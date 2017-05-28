+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 克隆步骤的变化"
url = "zh/release-0.6.0-cloning"
+++

<!--# Overview-->

# 概述

<!--The clone section is removed from the pipeline and has its own section in the yaml. Please note this change is experimental and the naming could be adjusted in the future.-->

克隆步骤从 pipeline 中移除，成为一个独立的章节。这是一个实验性的变更，在以后的版本中可能还有变化。

```diff
+clone:
+  git:
+    image: plugins/git
+    depth: 50

pipeline:
- clone:
-   image: plugins/git
-   depth: 50
```

<!--Please note that if you are using a plugin to cache your git repository you will need to adjust your pipeline configuration accordingly-->

如果您使用插件来缓存 git 仓库，您需要调整您的工作流配置。

```diff
+clone:
+  from_cache:
+    image: drillster/drone-volume-cache
+    restore: true
+  git:
+    image: plugins/git
+    depth: 50

pipeline:
- restore-cache:
-   image: drillster/drone-volume-cache
-   restore: true

- clone:
-   image: plugins/git
-   depth: 50
```
