---
title: Postgres 示例项目
url: zh/postgres-example

menu:
  usage:
    weight: 5
    identifier: postgres_example-zh
    parent: usage_examples
---

Example Yaml configuration for a project with a Postgres service dependency. Note that the postgres service will be accessible at `database:5432`.

```yaml
pipeline:
  test:
    image: golang
    commands:
      - go get
      - go test

services:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_DB=test
```

The official Postgres image provides environment variables used at startup to create the default username, password, database and more. Please see the official image documentation for more details.

```diff
services:
  database:
    image: postgres
    environment:
+     - POSTGRES_USER=postgres
+     - POSTGRES_DB=test
```

If you are unable to connect to the Postgres container please make sure you are giving Postgres adequate time to initialize and begin accepting connections.

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
    image: postgres
```

If you are still unable to connect to the mysql container, please make sure you using container name as the address.

```diff
- psql -U root -d test -h 127.0.0.1:5432
+ psql -U root -d test -h tcp://database:5432
```
