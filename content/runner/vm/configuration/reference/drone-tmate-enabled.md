---
date: 2000-01-01T00:00:00+00:00
title: DRONE_TMATE_ENABLED
author: tphoney
weight: 1
---

Optional boolean value. Enables remote ssh access to your pipeline instance using the tmate. Both the hosted service and self-hosted services are supported. This feature is disabled by default. _This feature only works on linux ie where platform's os is set to linux._

```bash
DRONE_TMATE_ENABLED=true
```

Note that you can also configure a self-hosted tmate server using the below configuration parameters. Please see the official [tmate documentation](https://tmate.io/) to learn more about self-hosting a tmate server.

```bash
DRONE_TMATE_ENABLED=true
DRONE_TMATE_HOST=tmate.company.com
DRONE_TMATE_PORT=2200
DRONE_TMATE_FINGERPRINT_RSA=SHA256:iL3StSCmPU+7p2IoD8y0huMXRVFIZyGFZa8r+lO3U5I
DRONE_TMATE_FINGERPRINT_ED25519=SHA256:gXLaN8IUxUMmlm/xu7M2NEFMlbUr5UORUgMi86Kh+tI
```
