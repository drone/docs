+++
weight = 6
date = "2014-11-08T12:54:12-08:00"
title = "IRC"

[menu.docs]
parent = "Notify"
+++

IRC notifications

```coffeescript
notify:
  irc:
    channel: foo
    nick: bar
    server: "irc.someserver.com:6667"
    on_started: true
    on_success: true
    on_failure: true
```
