---
date: 2000-01-01T00:00:00+00:00
title: DRONE_USER_CREATE
author: bradrydzewski
aliases:
- /reference/server/drone-user-create/
- /installation/reference/drone-user-create/
---

Optional user account that should be created on startup. This should be used to seed the system with an administrative account. It can be a real account (i.e. a real GitHub user) or it can be a machine account.

```
DRONE_USER_CREATE=username:octocat,machine:false,admin:true,token:55f24eb3d61ef6ac5e83d550178638dc
```

Providing a token is required for machine accounts, and must be 32 bytes. You can generate a random 32-byte token with the following command:

```
$ openssl rand -hex 16
55f24eb3d61ef6ac5e83d550178638dc
```