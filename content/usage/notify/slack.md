+++
weight = 3
date = "2014-11-08T12:54:12-08:00"
title = "Slack"

[menu.docs]
parent = "Notify"
+++

To configure Slack Notifications, you first need to create a new Incoming WebHook.
Once you have the WebHook URL, you add the following to your `drone.yml`

```coffeescript
notify:
  slack:
    webhook_url: 'https://hooks.slack.com/services/...'
    username: 'drone'
    channel: '#general'
    on_started: false
    on_success: true
    on_failure: true
```
> Note: `username` and `channel` will override the settings that you specified when creating the WebHook with Slack.
