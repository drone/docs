---
title: Example using Postgres
url: postgres-example

menu:
  usage:
    weight: 5
    identifier: postgres_example
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

If you are still unable to connect to the Postgres container, please make sure you are using the service name as the hostname.

```diff
- psql -U root -d test -h 127.0.0.1:5432
+ psql -U root -d test -h tcp://database:5432
```

# Working Example

This is a fully functioning example that demonstrates launching the Postgres service and then connecting with the service from the pipeline. Instead of waiting a fixed amount of time before attempting to query the database (as in the previous example) the pipeline container polls the postgres service every second until it is available.

```yaml
pipeline:
  ping:
    image: postgres
    commands:
      # wait for postgres service to become available
      - |
        until psql -U postgres -d test -h database \
         -c "SELECT 1;" >/dev/null 2>&1; do sleep 1; done
      # query the database
      - |
        psql -U postgres -d test -h database \
          -c "SELECT * FROM pg_catalog.pg_tables;"

services:
  database:
    image: postgres
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_DB=test
```
