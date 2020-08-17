---
date: 2000-01-01T00:00:00+00:00
title: Administrators
author: bradrydzewski
weight: 50
aliases:
- /user-admins
- /manage/user/admins/
- /admin/user-admins/
- /administration/user/admins

description: |
  Learn how to manage administrative accounts.
---

The administrator role has the highest permission level in the system. An administrator can use all features and has access to all system data, user data, and repository data.

# Capabilities

Admins have a number of special capabilities. Among them are the following:

* __Manage Accounts__
  : Admins have the ability to create and delete user accounts from the system using the command line tools. Admins are also capable of granting or revoking administrative permissions to other accounts.

* __Edit Repository Details__
  : Admins can modify repository settings, including cron jobs and secrets. Admins cannot view secrets. Admins also have the ability to increase or decrease the individual repository timeout (the default timeout is 60 minutes).

* __Edit Repository Trusted Flag__
  : Admins can enable or disable trusted mode for a repository. If trusted mode is enabled, the repository pipelines have access to privileged capabilities, including the ability to start privileged containers and mount host machine volumes.

* __System Endpoints__
  : Admins have access to restricted API endpoints, including system metrics, queue management and user management endpoints.

Admins cannot perform the following tasks:

* __Enable Any Repository__
  : Admins cannot enable repositories unless they have admin access to the repository in the source control management system (e.g. github). Admin access to the repository is required in order to register webhooks.

# Create the Primary Admin

When you configure the Drone server you can create the initial administrative account by passing the below environment variable, which defines the account username (e.g. github handle) and admin flag set to true.

```
DRONE_USER_CREATE=username:octocat,admin:true
```

If you need to grant the primary administrative role to an existing user, you can provide an existing username. Drone will update the account and grant administrator role on server restart.

# Create Additional Admins
You can create administrator accounts using the command line tools. Please see the command line tools [documentation]({{< relref "/cli/install.md" >}}) for installation instructions.

* Create a new administrator account:

    ```
    $ drone user add octocat --admin
    ```

* Or grant the administrator role to existing accounts:

    ```
    $ drone user update octocat --admin
    ```