+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 密文"
url = "zh/release-0.6.0-secrets"
+++

<!--Previously secrets were interpolated in the yaml using `${variable}` syntax. In the past we have attempted alternate approaches to limit secret expose to specific images without success. This release adopts the docker-compose syntax for exposing secrets to your pipeline steps.-->

之前版本密文在 yaml 文件中使用 `${variable}` 语法来进行解析。我们尝试修改这个方法来限制密文只在指定容器中使用密文，但是没有成功。

这个版本使用了 docker-compose 语法来在工作流中使用密文。

<!--The pipeline steps request access to named secrets. If access is granted, secrets are loaded as environment variables. Support for loading secrets from a temporary file are planned.-->

工作流请求使用已命名的密文，如果请求被许可，密文被读取为环境变量。从临时文件读取密文的功能在计划中。

<!--Example configuration:-->

使用示例：

```diff
pipeline:
  notify:
    image: plugins/slack
    channel: drone
-   webhook: ${slack_webhook}
+   secrets: [ slack_webhook ]
```

<!--Secret interpolation is being removed in favor of this docker-compose syntax. This implementation improves upon prior attempts and eliminates previously encountered [design flaws](https://github.com/drone/drone/issues/1888).-->

密文解析语法已经被移除，请使用 docker-compose 语法。这个实现改进了之前的尝试，并解决了[之前的设计问题](https://github.com/drone/drone/issues/1888)。

<!--# Commands-->

# 命令

<!--The command line utility can be used to manage secrets. These commands have slightly changed, using explicit flags for providing the secret name and value.-->

这个命令行工具可以用来管理密文。这个命令的语法有稍微的改变，使用显式的标签来指定密文的名字和数值。

```text
drone secret add <repo> \
  --image=<image> \
  --name=<secret> \
  --value=<value>
```

<!--Example command:-->

示例命令：

```text
drone secret add octocat/hello-world \
  --image=plugins/slack \
  --name=slack_webhook \
  --value=https://hooks.slack.com/services/...
```

<!--# Matching-->

# 匹配

<!--Prior implementations had very literal matching logic. The matching logic has been simplified to match the literal image name and ignore tags. Wildcards are no longer supported. Please update your secrets accordingly.-->

之前的实现使用了文本匹配的逻辑。新的匹配逻辑被简化为只匹配镜像名，忽略标签。Wildcards `*` 匹配不再被支持，请更新对应密文设置。

```diff
drone secret update octocat/hello-world \
- --image=plugins/slack:* \
+ --image=plugins/slack \
  --name=slack_webhook \
  --value=https://hooks.slack.com/services/...
```

<!--# Plugins-->

# 插件

<!--Secrets are passed to pipeline steps (including plugins) as environment variables. Plugin authors must update their plugins to source secrets from specific environment variables.-->

密文被传递到工作流的对应步骤（包括插件）作为环境变量。插件作者需要更新插件了从指定环境变量来使用密文。

<!--Example slack plugin changes:-->

Slack 插件变更示例：

```diff
cli.StringFlag{
  Name:   "webhook",
  Usage:  "slack webhook url",
- EnvVar: "PLUGIN_WEBHOOK",
+ EnvVar: "SLACK_WEBHOOK,PLUGIN_WEBHOOK",
},
```

<!--Example yaml configuration:-->

yaml 配置示例：

```diff
pipeline:
  notify:
    image: plugins/slack
    channel: drone
+   secrets: [ slack_webhook ]
```

<!--Please note that plugin authors need to document the names of the environment variables from which they will source these values. Users will need to create secrets using these exact environment variable names.-->

插件作者需要说明使用的环境变量名，用户需要生成对应指定的环境变量的密文。

<!--# Global and Organization Secrets-->

# 全局和组织密文

<!--Global and organization secrets are removed from Drone core and are delegated to plugins, using the new pluggable [secret backend](https://github.com/drone/drone/issues/1997). Commands for managing global and organization secrets have been removed accordingly.-->

全局和组织密文已被 Drone core 移除，对应功能被移动到插件中，使用新的可插拔密文[后端服务](https://github.com/drone/drone/issues/1997)。对应命令行工具命令也已经被移除。

```diff
drone secret add <repo> <name> <value>
- drone org secret add <org> <name> <value>
- drone global secret add <name> <value>
```

<!--This decision was made because organization secrets were not fully integrated and only supported GitHub (GitLab, Gitea, Gogs, Bitbucket, and Stash were not fully supported). There were also conflicting requirements from organizations regarding access control and storage, including desire to use third party secret management tools (e.g. Vault).-->

这个改变是因为组织密文只被 GitHub 项目使用 (GitLab, Gitea, Gogs, Bitbucket 和 Stash 没有全面支持组织密文)。这个功能也和组织的访问权限控制，储存，第三方密文管理工具（比如 Vault）相冲突。

<!--This change reflects the fact that Drone is moving toward a more focused core, where advanced functionality and customization is delegated to plugins. With pluggable secret backends, organizations have fine grained control over where secrets are stored and which builds have access to those secrets.-->

这个改变也反映了 Drone 这在向 core 核心靠拢。高级功能和自定义功能将以插件的形式呈现。使用可插拔的密文后端服务，组织可以很好地控制密文的储存以及构建访问密文。
