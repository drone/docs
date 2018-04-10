+++
date = "2017-04-15T14:39:04+02:00"
title = "User Admins"
url = "user-admins"

[menu.install]
  weight = 2
  identifier = "user-admins"
  parent = "install_access"
+++

You can grant administrative privileges to users by providing an enumerated user list, separated by a comma, using the designated environment variable.

{{% alert %}}
Admins can perform operations over any repository in Drone.
{{% /alert %}}

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_ADMIN=janedoe,johnsmith
```

Please note that the usernames are case-sensitive and must match the exact values as returned from your version control system (e.g. GitHub).

# Access Privileged Endpoints

View the server's metrics endpoint:

```nohighlight
<DRONE_HOST>/metrics
```

View the server's queue endpoint:

```nohighlight
<DRONE_HOST>/api/info/queue
```

View the server's pending and running builds endpoint:

```nohighlight
<DRONE_HOST>/api/builds
```

# User management

[Approving]({{< ref "install/users/user_access.md">}}) users for closed registrations and [managing]({{< ref "install/users/user_management.md">}}) users.
