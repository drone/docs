+++
weight = 3
date = "2014-11-08T12:54:12-08:00"
title = "Slack"

[menu.docs]
parent = "Notify"
+++

Slack notifications

```coffeescript
notify:
  slack:
    username: foo
    token: c90f0e53d4973302fe0b6159488423d6
    team: bar
    channel: baz
    on_started: false
    on_success: true
    on_failure: true
```
