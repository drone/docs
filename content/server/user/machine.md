---
date: 2000-01-01T00:00:00+00:00
title: Machine
author: bradrydzewski
weight: 60
aliases:
- /manage/user/machine/
- /administration/user/machine
description: |
  Learn how and when to create machine accounts.
---

The machine account is intended to be used exclusively for automation and integrations. Since this account won’t be used by a human, and will not have a corresponding account in your source code management system, it is referred to as a machine user.

# Create Accounts

You can manage machine accounts using the command line utility. Please see the command line tools [documentation]({{< relref "/cli/install.md" >}}) for installation instructions.

1. Generate a token for your machine user. The machine account can use this token for API access.
    ```
    $ openssl rand -hex 16
      fe8c402a51e6629aa1f43a4234afee81
    ```

2. Create a machine account with username prometheus:
    ```
    $ drone user add prometheus \
      --machine \
      --token=fe8c402a51e6629aa1f43a4234afee81
    ```

3. Create the machine account with administrative privileges:
    ```
    $ drone user add prometheus \
      --admin \
      --machine \
      --token=fe8c402a51e6629aa1f43a4234afee81
    ```
