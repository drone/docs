+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Users"
weight = 5
toc = true

[menu.main]
	parent="installation"
+++

# Overview

Drone provides multiple configurations for open or limited access to the system. This section describes these configuration and overall user management capabilities.

# Open Registration

Open registration is only recommended for secure installations on a private network. This configuration allows anyone to self-register and login to the system.

```
DRONE_OPEN=true
```

# Restricted Registration

Restricted registration is the recommended configuration. This configuration allows users of white-listed organizations to self-register and login to the system.

```
DRONE_ORGS=drone,github
DRONE_OPEN=true
```

# Closed Registration

Closed registration is enabled by default. Closed registration requires an administrator to manually register users using the command line utilities. See the next section of this guide for more details.

```
DRONE_OPEN=false
```

# User Management

You can manually manage user registration using the command line utility. Please see the command line documentation for installing and configuring the command line utility.

Use the `ls` command to list all active users:

```
drone user ls
```

Use the `add` command to add users to the system by login:

```
drone user add octocat
```

Use the `rm` command to remove users from the system by login:

```
drone user rm octocat
```

# Admin Users

You can configure administrative using the below settings. Note that you can configure one or many administrators, separated by a comma.

```
DRONE_ADMIN=octocat,bradrydzewski
```

<!--
You can grant all users administrative access using the below setting. Please note that that should only be enabled when you trust all Drone users with root access to your servers.

```
DRONE_ADMIN_ALL=true
```
-->
