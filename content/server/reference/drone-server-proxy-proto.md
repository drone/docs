---
date: 2000-01-01T00:00:00+00:00
title: DRONE_SERVER_PROXY_PROTO
author: bradrydzewski
aliases:
- /installation/reference/drone-server-proxy-proto/
---

Optional string value used to create webhooks that are routed through an alternate proxy server. The target use case for this setting is when your server is behind a firewall and you need GitHub webhooks to route through a public proxy.

```
DRONE_SERVER_PROXY_PROTO=https
```
