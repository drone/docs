+++
date = "2017-04-15T14:39:04+02:00"
title = "发送到 Slack Messages"
url = "zh/trigger-slack-notifications"

[menu.usage]
  weight = 1
  identifier = "trigger-slack-notifications-zh"
  parent = "usage_notifications"
+++

<!--This guide provides instructions for integrating slack plugins into your build pipeline using the slack plugin. Please see the plugin registry for a list of all notification plugins and detailed documentation.-->

这个文档介绍了将 Slack 插件整合到构建工作流中的步骤。请查看插件市场了解所有通知插件以及详细文档。

<!--Example configuration for sending Slack notification:-->

发送 Slack 通知的配置示例：

```yaml
pipeline:
  slack:
    image: plugins/slack
    channel: dev
    webhook: https://hooks.slack.com/services/...
```

<!--Note that the pipeline exits immediately on failure. This means your slack notification step is not executed if a previous step in the pipeline fails. You can override this default behavior with the below yaml configuration.-->

注意，工作流在遇到错误时会立即退出，这意味着如果在 Slack 插件之前的步骤发生了错误， Slack 插件将不会执行。可以使用下面的 yaml 配置来修改这个默认行为。

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
    webhook: https://hooks.slack.com/services/...
+   when:
+     status: [ success, failure ]
```


# 使用密钥

<!--In the previous example we store the webhook directly in the yaml configuration file in plain text. To avoid storing sensitive configuration parameters in your yaml file we recommend using the secret store.-->

在之前的 yaml 配置中，我们将 webhook 以明文的方式保存。为了避免在 yaml 文件中保存敏感信息，我们推荐使用密钥的方式保存对应信息。

<!--Example command adds the webhook to the secret store:-->

将 webhook 添加为密钥。

```text
drone secret add \
  -repository <repo> \
  -name slack_webhook \
  -image plugins/slack \
  -value https://hooks.slack.com/services/...
```

<!--Example configuration requests access to the webhook from the secret store:-->

访问密钥获得对应 webhook 的示例：

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
-   webhook: https://hooks.slack.com/services/...
+   secrets: [ slack_webhook ]
```

<!--Secrets are not exposed to pull requests by default for security reasons. If you would like to send Slack notifications for pull requests you can override this default.-->

考虑到安全因素，默认合并请求（pull request）不能访问密钥。如果您希望为合并请求到构建能够发送 Slack 通知，您可以修改这一默认设置。

```diff
drone secret add \
  -repository <repo> \
  -name slack_webhook \
  -image plugins/slack \
+ -event push \
+ -event pull_request \
  -value https://hooks.slack.com/services/...
```
