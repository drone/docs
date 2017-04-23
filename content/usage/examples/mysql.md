+++
date = "2017-04-15T14:39:04+02:00"
title = "Example using Mysql"
url = "mysql-example"

[menu.usage]
  Parent = "examples"
  weight = 4
  identifier = "mysql-example"
+++

Example Yaml configuration for a project with a MySQL service dependency. Note that the postgres service will be accessiable at `database:3306`.

```yaml
pipeline:
  test:
    image: golang
    commands:
      - go get
      - go test

services:
  database:
    image: mysql
    environment:
      - MYSQL_DATABASE=test
      - MYSQL_ALLOW_EMPTY_PASSWORD=yes
```

The official mysql image provides environment variables used at startup to create the default username, password, database and more. Please see the official image documentation for more details.

```diff
services:
  database:
    image: mysql
    environment:
+     - MYSQL_DATABASE=test
+     - MYSQL_ALLOW_EMPTY_PASSWORD=yes
```

If you need to start the mysql container with additional runtime options you can override the entrypoint and command arguments.

```diff
services:
  database:
    image: mysql
+   entrypoint: [ "mysqld" ]
+   command: [ "--character-set-server=utf8mb4" ]
```

If you are unable to connect to the mysql container please make sure you are giving mysql adequate time to initialize and begin accepting connections.

```diff
pipeline:
  test:
    image: golang
    commands:
+     - sleep 15
      - go get
      - go test

services:
  database:
    image: mysql
```

If you are still unable to connect to the mysql container, please make sure you using container name as the address.

```diff
- mysql -u root -h 127.0.0.1:3306
+ mysql -u root -h tcp://database:3306
```
