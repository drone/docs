+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Slack"
description = "Send build status notifications via Slack"
user = "drone-plugins"
repo = "drone-slack"
image = "plugins/drone-slack"
tags = ["chat", "messaging"]
categories = "notify"
draft = false
date = 2016-02-13T08:58:10Z
menu = ""
weight = 1

+++


Use the Slack plugin to send a message to your Slack channel when a build completes. You will need to supply Drone with an [Incoming Webhook URL](https://my.slack.com/services/new/incoming-webhook).

The following parameters are used to configuration the notification:

Parameter     | Description
--------------|----------------------------
`webhook_url` | json payloads are sent here
`channel`     | messages sent to the above webhook are posted here
`recipient`   | alternatively you can send it to a specific user
`username`    | choose the username this integration will post as

Example configuration:

```yaml
---
notify:
  slack:
    webhook_url: https://hooks.slack.com/services/...
    channel: dev
    username: drone
```

Example configuration using template to customize the message:

```yaml
---
notify:
  slack:
    webhook_url: https://hooks.slack.com/services/...
    channel: dev
    username: drone
    template: >
      build #{{ build.number }} finished with a {{ build.status }} status
```

