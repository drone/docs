---
date: 2000-01-01T00:00:00+00:00
title: DRONE_WEBHOOK_SECRET
author: bradrydzewski
aliases:
- /installation/reference/drone-webhook-secret/
---

Shared secret used to create an [http-signature](https://tools.ietf.org/html/draft-cavage-http-signatures). The webhook recipient can use the shared secret to verify request authenticity.

```
DRONE_WEBHOOK_SECRET=correct-horse-battery-staple
```