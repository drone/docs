---
title: Example using Redis
url: redis-example

menu:
  usage:
    weight: 3
    identifier: redis_example
    parent: usage_examples
---

Example Yaml configuration for a project with a Redis service dependency. Note that the redis service will be available at `cache:6379`.

```yaml
pipeline:
  build:
    image: redis
    commands:
      - redis-cli -h redis ping
      - redis-cli -h redis set FOO bar
      - redis-cli -h redis get FOO

services:
  redis:
    image: redis
```

If you are unable to connect to the redis container please make sure you are giving redis adequate time to initialize and begin accepting connections.

```diff
pipeline:
  build:
    image: redis:3.0
    commands:
+     - sleep 5
      - redis-cli -h redis ping
      - redis-cli -h redis set HELLO hello
      - redis-cli -h redis get HELLO

services:
  redis:
    image: redis:3.0
```
