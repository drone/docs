---
date: 2000-01-01T00:00:00+00:00
title: DRONE_USER_FILTER
author: bradrydzewski
aliases:
- /reference/server/drone-user-filter/
- /installation/reference/drone-user-filter/
---

Optional comma-separated list of accounts. Registration is limited to users in this list, or users that are members of organizations included in this list.

```
DRONE_USER_FILTER=octocat,spacheghost,github
```

If a user attempts to register and they are not a named user and not a member of a named organization, they will receive the following error:

> _Login Failed. User must be a member of an approved organization._
