+++
date = "2017-04-15T14:39:04+02:00"
title = "User Registration"
url = "user-registration"

[menu.install]
  weight = 1
  identifier = "user-registration"
  parent = "install_access"
+++

Drone provides multiple configurations for open or limited access to the system. This section describes different options for user registration and access.

# Open Registration

Open registration is only recommended for secure installations on a private network. This configuration allows anyone to self-register and login to the system.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     - DRONE_OPEN= true
```

# Restricted Registration

Restricted registration is the recommended configuration. This configuration allows members of white-listed organizations to self-register and login to the system.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     DRONE_OPEN: true
+     DRONE_ORGS: dogpatch,dolores
```

# Closed Registration

Closed registration is enabled by default. Closed registration requires an administrator to manually register users using the command line utility. When using closed registration it is __imperative__ you provide a list of administrators that are able to login and manage accounts.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
-     - DRONE_OPEN=true
+     - DRONE_OPEN=false
+     - DRONE_ADMIN=janedoe,johnsmith
```

You can then manually grant users access using the command line utility:

```nohighlight
drone user add <username>
```
