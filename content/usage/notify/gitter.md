+++
weight = 2
date = "2014-11-08T12:54:12-08:00"
title = "Gitter"

[menu.docs]
parent = "Notify"
+++

You may configure Drone to trigger a message on [Gitter](https://gitter.im) when your build completes. This can be configured in the notify section of the `.drone.yml` file. For example:

```coffeescript
notify:
  gitter:
    room_id: cf23db8155e6700caccc
    token: c90f0e53d4973302fe0b6159488423d6
    on_started: false
    on_success: true
    on_failure: true
```

> ## Help Wanted
> 
> We could acheive better integration with Gitter by adding Drone as an official service. Please see [issue #569](https://github.com/drone/drone/issues/569) for more details. You will be our hero!