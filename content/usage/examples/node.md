+++
date = "2017-04-15T14:39:04+02:00"
title = "Example Node project"
url = "node-example"

[menu.usage]
  Parent = "examples"
  weight = 2
  identifier = "node-example"
+++

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
