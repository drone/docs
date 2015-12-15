+++
title = "Flowdock"
description = "Send build status notifications via Flowdock"
user = "drone-plugins"
repo = "drone-flowdock"
image = "plugins/drone-flowdock"
tags = ["flowdock"]
categories = "notify"
draft = false
date = 2015-12-15T06:50:56Z
menu = ""
weight = 1

+++

Use this plugin for sending build status notifications via Flowdock. You can
override the default configuration with the following parameters:

* `` -

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  flowdock:
```

