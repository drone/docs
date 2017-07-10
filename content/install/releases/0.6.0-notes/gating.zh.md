+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Gated Builds 门禁构建"
url = "zh/release-0.6.0-gating"
+++

<!--Drone adds __experimental__ support for gating your builds. You can enable gating for your repository in the user interface or with the following command:-->

Drone 添加了 Gated Builds 门禁构建的实验性支持。可以使用命令行工具来启用门禁：

```text
drone repo update <repo> --gated=true
```

<!--The default gating logic is very basic. If the repository is public, and the hook is a pull request, and the user is not the repository admin, then the pull request is blocked pending approval. This default logic is intended to help mitigate bad actors sending harmful pull requests to open source projects.-->

默认的门禁逻辑很简单。如果仓库是一个公开的仓库，webhook 是 pull request，则 pull request 会在构建前被拦截，等待许可后才执行构建。这可以帮助拦截开源项目遇到的恶意 pull request。 

<!--# Customization-->

# 自定义

<!--The default gating logic is not going to be useful for many installations. We fully expect most organizations will need to customize this logic. To enable customization we have made gating pluggable. This means you will be able to supply your own custom logic, tailored to your organization's specific needs.-->

默认的门禁逻辑对大多数安装意义不大。大多数组织可能需要自定义这个逻辑。门禁功能是可插拔的，这意味着您可以使用适合于对应项目的自定义的逻辑。

<!--Please note that we do not want to embed a complex approval engine inside of Drone. Nor do we want to embed your organizations custom approval logic in Drone. So please use plugins and please do not send pull requests that attempt to alter the default gating logic.-->

我们不想在 Drone 中引入复杂的审查流程。我们也不想在 Drone 中引入您的组织中的自定义审查逻辑。所以请使用插件，并不要发送 pull request 来修改默认的门禁逻辑。

# 未来的变化

<!--# What's Next-->

<!--* Builtin logic should not require approval if pull request is sender has write access
* Ability to permanently add users to a repository whitelist
* Ability to permanently add users to a repository blacklist-->
