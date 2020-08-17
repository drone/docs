---
date: 2000-01-01T00:00:00+00:00
title: DRONE_STASH_PRIVATE_KEY
author: bradrydzewski
---

String value configures the path to your Bitbucket Server private key file. Note that this file needs to also be mounted into the Drone server container as a volume.

```
DRONE_STASH_PRIVATE_KEY=/etc/bitbucket/key.pem
```
