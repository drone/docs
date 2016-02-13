+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Gitter"
description = "Send build status notifications via Gitter"
user = "drone-plugins"
repo = "drone-gitter"
image = "plugins/drone-gitter"
tags = ["chat", "messaging", "gitter"]
categories = "notify"
draft = false
date = 2016-02-13T08:59:15Z
menu = ""
weight = 1

+++

Use this plugin for sending build status notifications via Gitter. The status
updates are displayed in a room's activity feed. You can override the default
configuration with the following parameters:

* `webhook` - A single or a list of webhooks

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  gitter:
    webhook: https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
```

### Multiple Channels

In some cases you want to send notifications to multiple different channels
to. In that case you can simply provide a list of webhooks.

Example configuration that sends multiple message:

```yaml
notify:
  gitter:
    webhook:
     - https://webhooks.gitter.im/e/91e06797227ae5dbe6ec
     - https://webhooks.gitter.im/e/27a2e6ece5db91e06797
```

