+++
date = "2017-11-18T20:13:01+04:00"
title = "Automatic Build Cancellation"
url = "auto-cancel-builds"

[menu.install]
  identifier = "auto-cancel-builds"
  parent = "install_server"
  weight = 11
+++

Drone supports automatically cancelling builds in either the running or pending state for the same branch.

**NOTE**: This is an experimental feature that may cause issues if you have an existing installation

```diff
services:
  image: drone/drone:{{% version %}}
  ports:
    - 80:80
    - 443:443
  volumes:
    - /var/lib/drone:/var/lib/drone/
  restart: always
  environment:
    - DRONE_OPEN=true
+   - DRONE_EXPERIMENTAL_AUTO_CANCEL_PENDING=true
+   - DRONE_EXPERIMENTAL_AUTO_CANCEL_RUNNING=true
```

# Pending Build Cancellation

If the previous build for a branch is in the pending state, it will automatically be killed in the same way a zombie build is cleaned up. This does not replace that queued event with the new one and your new build will be added to the end of the queue.

# Running Build Cancellation

If the previous build for a branch is running, it will automatically be killed in the same way a build is cancelled from the UI. As with pending builds, your new build will be added at the end of the build queue.
