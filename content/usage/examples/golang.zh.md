---
title: Golang 示例项目
url: zh/golang-example

menu:
  usage:
    weight: 1
    identifier: golang_example-zh
    parent: usage_examples
---

Example Yaml configuration for a project written in the Go programming language:

```yaml
workspace:
  base: /go
  path: src/github.com/octocat/hello-world

pipeline:
  build:
    image: golang:latest
    commands:
      - go get ./...
      - go test
      - go build
```

Using the official Go image:

```diff
pipeline:
  build:
+   image: golang:latest
    commands:
      - go get
      - go test
```

Using the build matrix to test multiple Go versions:

```diff
pipeline:
  build:
-   image: golang:latest
+   image: golang:${GO_VERSION}
    commands:
      - go get
      - go test

+matrix:
+ GO_VERSION:
+   - latest
+   - "1.7"
+   - "1.6"
```

# What is the workspace?

The workspace defines the shared volume and working directory for your build. We recommend projects customize the workspace to use the desired GOPATH.

```diff
+workspace:
+  base: /go
+  path: src/github.com/octocat/hello-world

pipeline:
  build:
    image: golang:latest
    commands:
      - go get
      - go test
```

# What is the workspace base?

The base defines a shared volume that is available to all pipeline steps. This ensures your source code, dependencies and compiled binaries are persisted and shared between steps.

```diff
workspace:
+ base: /go
  path: src/github.com/octocat/hello-world

pipeline:
  deps:
    image: golang:latest
    commands:
      - go get
      - go test
  build:
    image: node:latest
    commands:
      - go build
```

# What is the workspace path

The path defines the working directory of your build. This is where your code is cloned and will be the default working directory of every step in your build process. The path must be relative and is combined with your base path.

```diff
workspace:
  base: /go
+ path: src/github.com/octocat/hello-world
```

```text
git clone https://github.com/octocat/hello-world \
  /go/src/github.com/octocat/hello-world
```
