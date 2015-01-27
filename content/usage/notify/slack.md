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

---

## Create a WebHook

In order to create a webhook you'll need to login to your Slack account and navigate to:
https://my.slack.com/services/new/incoming-webhook

Then follow the intructions to generate a new token:

![Slack Setup](/img/slack_setup.png)

> ## Help Wanted
> 
> We could acheive better integration with Slack by becoming an official Slack service. Please see [issue #854](https://github.com/drone/drone/issues/854) for more details. You will be our hero!