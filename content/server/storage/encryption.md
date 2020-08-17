---
date: 2000-01-01T00:00:00+00:00
title: Encryption
author: bradrydzewski
weight: 10
aliases:
- /installation/storage/encryption/
description: |
  Configure database encryption.
---

Drone supports aesgcm encryption of secrets stored in the database. You must enable encryption before any secrets are stored in the database.

1. Create an encryption key:

    ```
    $ openssl rand -hex 16
    0c549fd39ae397333761d2cb0c53c219
    ```

2. Provide the encryption key to the Drone server:

    ```
    DRONE_DATABASE_SECRET=0c549fd39ae397333761d2cb0c53c219
    ```

<div class="alert alert-info">
Database encryption must be configured before secrets are added to the database. You cannot enable encryption if the secrets table has existing records.
</div>
