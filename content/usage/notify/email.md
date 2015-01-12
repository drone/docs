+++
weight = 1
date = "2014-11-08T12:54:12-08:00"
title = "Email"

[menu.docs]
parent = "Notify"
+++

Email notifications

```coffeescript
notify:
  email:
    recipients:
      - foo@bar.com
      - baz@bar.com
    on_success: true
    on_failure: true
```
