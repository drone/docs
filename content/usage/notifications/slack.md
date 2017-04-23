+++
date = "2017-04-15T14:39:04+02:00"
title = "Sending Slack Messages"
url = "trigger-slack-notifications"

[menu.usage]
  weight = 1
  identifier = "trigger-slack-notifications"
  parent = "usage_notifications"
+++

This guide provides instructions for integrating slack plugins into your build pipeline using the slack plugin. Please see the plugin registry for a list of all notification plugins and detailed documentation.

Example configuration for sending Slack notification:

```yaml
pipeline:
  slack:
    image: plugins/slack
    channel: dev
    webhook: https://hooks.slack.com/services/...
```

Note that the pipeline exits immediately on failure. This means your slack notification step is not executed if a previous step in the pipeline fails. You can override this default behavior with the below yaml configuration.

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
    webhook: https://hooks.slack.com/services/...
+   when:
+     status: [ success, failure ]
```


# Using Secrets

In the previous example we store the webhook directly in the yaml configuration file in plain text. To avoid storing sensitive configuration parameters in your yaml file we recommend using the secret store.

Example command adds the webhook to the secret store:

```text
drone secret add \
  -repository <repo> \
  -name slack_webhook \
  -image plugins/slack \
  -value https://hooks.slack.com/services/...
```

Example configuration requests access to the webhook from the secret store:

```diff
pipeline:
  slack:
    image: plugins/slack
    channel: dev
-   webhook: https://hooks.slack.com/services/...
+   secrets: [ slack_webhook ]
```

Secrets are not exposed to pull requests by default for security reasons. If you would like to send Slack notifications for pull requests you can override this default.

```diff
drone secret add \
  -repository <repo> \
  -name slack_webhook \
  -image plugins/slack \
+ -event push \
+ -event pull_request \
  -value https://hooks.slack.com/services/...
```
