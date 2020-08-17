---
date: 2000-01-01T00:00:00+00:00
title: DRONE_LICENSE
author: bradrydzewski
aliases:
- /reference/server/drone-license/
- /installation/reference/drone-license/
---

Optional string value provides the filepath of the Drone Enterprise license key. This is used to unlock the Drone Enterprise edition.

```
DRONE_LICENSE=/etc/drone.key
```

If you are running the Drone server in a Docker container you will need to mount the license key as a volume:

```
$ docker run \
  --volume=/path/on/host/drone.key:/etc/drone.key
```
