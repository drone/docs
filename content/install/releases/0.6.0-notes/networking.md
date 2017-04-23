+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Networking"
url = "release-0.6.0-networking"
+++

Drone now aligns with docker-compose and bridge networking with hostnames. This means services are no longer accessible on `localhost` or `127.0.0.1` and are instead available using their hostname as defined in the yaml configuration file.

```diff
pipeline:
  build:
    image: golang
    commands:
-     - mysql -h localhost:3306
+     - mysql -h mysql:3306

services:
  mysql:
    image: mysql
```

There are no immediate plans to continue support for pod networking (e.g. localhost) unless natively supported by Docker. I have opened [this issue](https://github.com/docker/docker/issues/26024) with the Docker project to request native pod networking. Please considering voicing your support if this capability is important to you or your organization.
