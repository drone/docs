---
date: 2000-01-01T00:00:00+00:00
title: Management
author: bradrydzewski
weight: 20
aliases:
- /user-management
- /manage/user/management/
description: |
  Learn how to manage user accounts.
---

The [administrator]({{< relref "admin.md" >}}) is responsible for managing user accounts. This includes the ability to create and delete accounts, and to grant the administrator role to other accounts.

You can manage users using the command line utility. Please see the command line tools [documentation]({{< relref "/cli/install.md" >}}) for installation instructions.

* Example command to list active users:

    ```
    $ drone user ls
    ```

* Example command to add a user to the system:

    ```
    $ drone user add octocat
    ```

* Example command to remove a user from the system:

    ```
    $ drone user rm octocat
    ```

* Example command to grant the administrator role to a user:

    ```
    $ drone user update octocat --admin
    ```