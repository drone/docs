---
date: 2000-01-01T00:00:00+00:00
title: DRONE_RPC_SECRET
author: bradrydzewski
aliases:
- /installation/reference/drone-rpc-secret/
---

Required literal value provides the Drone shared secret. This is used to authenticate the RPC connection to the server. The server and runners must be provided the same secret value.

```
DRONE_RPC_SECRET=correct-horse-batter-staple
```