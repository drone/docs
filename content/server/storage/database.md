---
date: 2000-01-01T00:00:00+00:00
title: Database
author: bradrydzewski
weight: 1
aliases:
- /database-settings
- /installation/storage/database/
- /administration/server/database
description: |
  Configure database storage.
---


Drone requires the use of a database backend for persistence. Drone uses an embedded __sqlite__ database by default, which does not require any additional configuration. This article provides a basic overview of alternate database configurations.

# Postgres

Drone supports postgres 9.6 and higher as the database engine. The below example demonstrates postgres database configuration. Please reference the official driver [documentation](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING) for connection string configuration details.

```
DRONE_DATABASE_DRIVER=postgres
DRONE_DATABASE_DATASOURCE=postgres://root:password@1.2.3.4:5432/postgres?sslmode=disable
```

# MySQL

<div class="alert alert-warn">
We strongly recommend using postgres instead of mysql. The system has been optimized for features not found in mysql.
</div>

Drone supports mysql 5.6 and higher as the database engine. The below example demonstrates mysql database configuration. Please reference the official driver [documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for connection string configuration details.

```
DRONE_DATABASE_DRIVER=mysql
DRONE_DATABASE_DATASOURCE=root:password@tcp(1.2.3.4:3306)/drone?parseTime=true
```
