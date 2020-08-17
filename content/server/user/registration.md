---
date: 2000-01-01T00:00:00+00:00
title: Registration
author: bradrydzewski
weight: 10
aliases:
- /user-registration
- /custom-access-policies/
- /administration/user/registration

description: |
  Learn how to disable open registration.
---

Drone registration is open by default. This means any user can register an account and use the system. If your server is public facing it is considered best practice to limit registration.

* Limits registration to named users:
  ```
  DRONE_USER_FILTER=octocat,spaceghost
  ```

* Limit registration to members of approved organizations:
    ```
    DRONE_USER_FILTER=google,kubernetes
    ```

* Limit registration to named users and members of approved organizations:
    ```
    DRONE_USER_FILTER=octocat,spaceghost,google,kubernetes
    ```