---
date: 2000-01-01T00:00:00+00:00
title: DRONE_CALVER
author: bradrydzewski
---

If the git tag is a valid [calendar version](https://calver.org/) string, the system provides the tag as a calendar version string.

```
DRONE_CALVER=19.1.0-beta.20190318
```

The calendar version string is also available in the following formats:

```
DRONE_CALVER_SHORT=19.1.0
DRONE_CALVER_MAJOR_MINOR=19.1
DRONE_CALVER_MAJOR=19
DRONE_CALVER_MINOR=1
DRONE_CALVER_MICRO=0
DRONE_CALVER_MODIFIER=beta.20190318
```
