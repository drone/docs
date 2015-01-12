+++
weight = 2
date = "2014-11-08T12:53:40-08:00"
title = "MySQL"

[menu.main]
parent = "Database"
+++

You may configure Drone to use a MySQL database. Drone will automatically handle database schema setup and migration. Please alter the following section in `/etc/drone/drone.toml`. This is an example configuration:

```toml
[database]
driver = "mysql"
datasource = "root@tcp(127.0.0.1:3306)/drone"
```

## Environment Variables

You may also configure the database using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_DATABASE_DRIVER="mysql"
DRONE_DATABASE_DATASOURCE="root@tcp(127.0.0.1:3306)/drone"
```

## Datasource

For more details about how to configure the datasource string please see the official driver documentation:

https://github.com/go-sql-driver/mysql#dsn-data-source-name
