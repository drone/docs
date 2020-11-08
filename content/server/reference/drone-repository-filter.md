---
date: 2000-01-01T00:00:00+00:00
title: DRONE_REPOSITORY_FILTER
author: bradrydzewski
aliases:
- /installation/reference/drone-repository-filter/
---

Optional comma-separated list of accounts, used to limit which repositories are syncronized between your source control management system and Drone. _Note that this variable must be set before your first login. Setting this variable after having already authenticated and syncronized your account has no effect._

```
DRONE_REPOSITORY_FILTER=octocat,spacheghost
```
