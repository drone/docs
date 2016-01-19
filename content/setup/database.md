+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Database"
weight = 3
menu = "installation"
toc = true
+++

Drone supports several relational databases. If you don't have a strong preference, SQLite is a great starting point. Once you have configured your database, continue onward to [Create and Run a Drone Container]({{< relref "setup/overview.md#create-and-run-a-drone-container" >}}) 

# SQLite

Drone uses an embedded SQLite as the **default** database with zero configuration required. In order to customize the SQLite database configuration you should specify the following environment variables:

```
DATABASE_DRIVER=sqlite3
DATABASE_CONFIG=/var/lib/drone/drone.sqlite
```

# Postgres

Configure a Postgres database backend:

```
DATABASE_DRIVER=postgres
DATABASE_CONFIG=postgres://root:pa55word@127.0.0.1:5432/postgres?sslmode=disable
```

See the official postgres connection string [documentation](http://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING) for a complete set of configuration options and examples.

# MySQL

Configure a Mysql database backend:

```
DATABASE_DRIVER="mysql"
DATABASE_CONFIG="root:pa55word@tcp(localhost:3306)/drone?parseTime=true"
```
See the driver [documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for a complete set of configuration options and examples.

Note that Drone does not automatically create the mysql database. You should use the mysql command line utility or your preferred management console to create the database:

```
mysql -P 3306 --protocol=tcp -u root -e 'create database if not exists drone;'
```
