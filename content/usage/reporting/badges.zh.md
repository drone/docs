+++
date = "2000-01-01T00:00:00+00:00"
title = "Badges 标签"
url = "zh/badges"

[menu.usage]
  weight = 5
  identifier = "badges-zh"
  parent = "usage_reports"
+++

<!--Drone has integrated support for repository status badges. These badges can be added to your website or project readme file to display the status of your code.-->

Drone 支持显示仓库状态 Badge 标签。这些标签可以被添加到网站或者项目的 readme 中来显示代码的状态。

Badge 路径（endpoint）:

```text
<scheme>://<hostname>/api/badges/<owner>/<repo>/status.svg
```

<!--The status badge displays the status for the latest build to your default branch (e.g. master). You can customize the branch by adding the `branch` query parameter.-->

默认 badge 标签显示了默认分支（比如 master）上的最新构建状态。可以通过添加 `?branch=<branch>` 来自定义分支。

```diff
-<scheme>://<hostname>/api/badges/<owner>/<repo>/status.svg
+<scheme>://<hostname>/api/badges/<owner>/<repo>/status.svg?branch=<branch>
```

<!--Please note status badges do not include pull request results, since the status of a pull request does not provide an accurate representation of your repository state.-->

请注意 badge 标签不包括合并请求（pull request）的构建结果，因为合并请求并不代表仓库的准确状态。
