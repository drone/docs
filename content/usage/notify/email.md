+++
weight = 1
date = "2014-11-08T12:54:12-08:00"
title = "Email"

[menu.docs]
parent = "Notify"
+++

You may configure Drone to trigger an email notification when your build completes. This can be configured in the `notify` section of the `drone.yml` file. For example:

```coffeescript
notify:
  email:
    recipients:
      - foo@bar.com
      - baz@bar.com
    on_success: false
    on_failure: true
```

Use `blame` to send an email only to the commit author:

```coffeescript
notify:
  email:
    on_failure: blame
```

Use `change` to send an email only when the status changes:

```coffeescript
notify:
  email:
    on_success: change
    on_failure: change
```
