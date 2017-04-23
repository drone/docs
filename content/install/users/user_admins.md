+++
date = "2017-04-15T14:39:04+02:00"
title = "User Admins"
url = "user-admins"

[menu.install]
  weight = 3
  identifier = "user-admins"
  parent = "install_access"
+++

You can grant administrative privileges to users by providing an enumerated user list, separated by a comma, using the designated environment variable.

```diff
services:
  drone-server:
    image: drone/drone:0.6
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```

Please note that the usernames are case-sensitive and must match the exact values as returned from your version control system (e.g. GitHub).
