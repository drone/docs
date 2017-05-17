---
title: Node 示例项目
url: zh/node-example

menu:
  usage:
    weight: 2
    identifier: node_example-zh
    parent: usage_examples
---

Example Yaml configuration for a project written in JavaScript using Node:

```diff
pipeline:
  build:
    image: node:latest
    commands:
      - npm install
      - npm run test
```

Using the official Node images:

```diff
pipeline:
  build:
+   image: node:latest
    commands:
      - npm install
      - npm run lint
      - npm run test
```

Using the official NPM plugin to publish packages:

```diff
pipeline:
  build:
    image: node:latest
    commands:
      - npm install
      - npm run lint
      - npm run test
  publish:
+   image: plugins/npm
+   when:
+     branch: master
```

Using the build matrix to test multiple node versions:

```diff
pipeline:
  build:
-   image: node:latest
+   image: node:${NODE_VERSION}
    commands:
      - npm install
      - npm run lint
      - npm run test

+matrix:
+ NODE_VERSION:
+   - latest
+   - "7"
+   - "6"
+   - "4"
```
