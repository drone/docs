---
date: 2000-01-01T00:00:00+00:00
title: DRONE_DATABASE_DATASOURCE
author: bradrydzewski
aliases:
- /installation/reference/drone-database-datasource/
---

Optional string value. Configures the database connection string. The default value is the path of the embedded sqlite database file.

```
DRONE_DATABASE_DATASOURCE=/data/database.sqlite
```

Example mysql connection string (below). See the [official driver](https://github.com/go-sql-driver/mysql#dsn-data-source-name) documentation for more connection string details.

```
DRONE_DATABASE_DATASOURCE=root:password@tcp(1.2.3.4:3306)/drone?parseTime=true
```

Example postgres connection string (below). See the [official driver](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING) documentation for more connection string details.

```
DRONE_DATABASE_DATASOURCE=postgres://root:password@1.2.3.4:5432/postgres?sslmode=disable
```