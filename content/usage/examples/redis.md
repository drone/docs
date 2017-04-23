+++
date = "2017-04-15T14:39:04+02:00"
title = "Example using Redis"
url = "redis-example"

[menu.usage]
  Parent = "examples"
  weight = 3
  identifier = "redis-example"
+++

Example Yaml configuration for a project with a Redis service dependency. Note that the redis service will be available at `cache:6379`.

```yaml
pipeline:
  test:
    image: golang
    commands:
      - go get
      - go test

services:
  cache:
    image: redis
```

If you are unable to connect to the redis container please make sure you are giving redis adequate time to initialize and begin accepting connections.

```diff
pipeline:
  test:
    image: golang
    commands:
+     - sleep 15
      - go get
      - go test

services:
  cache:
    image: redis
```
