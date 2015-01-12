+++
weight = 7
date = "2014-11-08T12:54:12-08:00"
title = "Webhook"

[menu.docs]
parent = "Notify"
+++

Webhook notifications

```coffeescript
notify:
  webhook:
    urls:
      - "http://foo.com/hook/url"
    on_success: true
    on_failure: true
```
