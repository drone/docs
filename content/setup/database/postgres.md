+++
weight = 3
date = "2014-11-08T12:53:40-08:00"
title = "Postgres"

[menu.main]
parent = "Database"
+++


You may configure Drone to use a Postgres database. Drone will automatically handle database schema setup and migration. Please alter the following section in `/etc/drone/drone.toml`. This is an example configuration:

```toml
[database]
driver = "postgres"
datasource = "host=127.0.0.1 user=postgres dbname=drone sslmode=disable"
```

## Environment Variables

You may also configure the database using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_DATABASE_DRIVER="postgres"
DRONE_DATABASE_DATASOURCE="host=127.0.0.1 user=postgres dbname=drone sslmode=disable"
```

## Datasource

For more details about how to configure the datasource string please see the official driver documentation:

http://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING
