---
title: Example using Mongo
url: mongodb-example

menu:
  usage:
    weight: 3
    identifier: mongodb_example
    parent: usage_examples
---

Example Yaml configuration for a project with a MongoDB service dependency. Note that the mongo service will be available at `mongo:27017`.

```yaml
pipeline:
  ping:
    image: mongo:3.0
    commands:
      - sleep 15
      - mongo --host mongo --eval "{ ping: 1 }"

services:
  mongo:
    image: mongo:3.0
    command: [ --smallfiles ]
```

If you are unable to connect to the mongo container please make sure you are giving mongodb adequate time to initialize and begin accepting connections.

```diff
pipeline:
  ping:
    image: mongo:3.0
    commands:
+     - sleep 15
      - mongo --host mongo --eval "{ ping: 1 }"

services:
  mongo:
    image: mongo:3.0
    command: [ --smallfiles ]
```

If you are still unable to connect to the mongo container, please make sure you are using the service name as the hostname.

```diff
- mongo --host localhost --eval "{ ping: 1 }"
+ mongo --host mongo     --eval "{ ping: 1 }"
```
