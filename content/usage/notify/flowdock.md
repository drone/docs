+++
weight = 5
date = "2014-11-08T12:54:12-08:00"
title = "Flowdock"

[menu.docs]
parent = "Notify"
+++

Flowdock notifications

```coffeescript
notify:
  flowdock:
    source: Drone
    token: c90f0e53d4973302fe0b6159488423d6
    tags: ci
    on_started: false
    on_success: true
    on_failure: true
```