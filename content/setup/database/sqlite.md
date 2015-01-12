+++
weight = 1
date = "2014-11-08T12:53:40-08:00"
title = "SQLite"

[menu.main]
parent = "Database"
+++

SQLite is the default database driver. Drone will automatically create a database file at `/var/lib/drone/drone.sqlite` and will handle database schema setup and migration.

If you would like to customize the SQLite configuration you may alter the following section in `/etc/drone/drone.toml`. This is an example configuration:

```toml
[database]
driver = "sqlite3"
datasource = "/var/lib/drone/drone.sqlite"
```

## Environment Variables

You may also configure the database using environment variables:

```bash
DRONE_DATABASE_DRIVER="sqlite3"
DRONE_DATABASE_DATASOURCE="/var/lib/drone/drone.sqlite"
```
