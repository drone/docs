+++
date = "2017-04-15T14:39:04+02:00"
title = "Environment"
url = "environment"

[menu.usage]
  weight = 3
  identifier = "environment"
  parent = "config"
+++

Drone provides the ability to define environment variables scoped to individual build steps. Example pipeline step with custom environment variables:

```diff
pipeline:
  build:
    image: golang
+   environment:
+     - CGO=0
+     - GOOS=linux
+     - GOARCH=amd64
    commands:
      - go build
      - go test
```

Please note that the environment section is not able to expand environment variables. If you need to expand variables they should be exported in the commands section.

```diff
pipeline:
  build:
    image: golang
-   environment:
-     - PATH=$PATH:/go
    commands:
+     - export PATH=$PATH:/go
      - go build
      - go test
```

Please be warned that `${variable}` expressions are subject to pre-processing. If you do not want the pre-processor to evaluate your expression it must be escaped:

```diff
pipeline:
  build:
    image: golang
    commands:
-     - export PATH=${PATH}:/go
+     - export PATH=$${PATH}:/go
      - go build
      - go test
```

<!--
# String Substitution

Drone provides the ability to substitute environment variables at runtime. This gives us the ability to use dynamic build or commit details in our pipeline configuration.

Example commit substitution:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: ${DRONE_COMMIT_SHA}
```

Example tag substitution:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: ${DRONE_TAG}
```

# String Operations

Drone also emulates bash string operations. This gives us the ability to manipulate the strings prior to substitution. Example use cases might include substring and stripping prefix or suffix values.

OPERATION             | DESC
----------------------|---------------------------------------------------------
`${param}`            | parameter substitution
`${param,}`           | parameter substitution with lowercase first char
`${param,,}`          | parameter substitution with lowercase
`${param^}`           | parameter substitution with uppercase first char
`${param^^}`          | parameter substitution with uppercase
`${param:pos}`        | parameter substitution with substring
`${param:pos:len}`    | parameter substitution with substring and length
`${param=default}`    | parameter substitution with default
`${param##prefix}`    | parameter substitution with prefix removal
`${param%%suffix}`    | parameter substitution with suffix removal
`${param/old/new}`    | parameter substitution with find and replace

Example variable substitution with substring:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: ${DRONE_COMMIT_SHA:0:8}
```

Example variable substitution strips `v` prefix from `v.1.0.0`:

```diff
pipeline:
  docker:
    image: plugins/docker
+   tags: ${DRONE_TAG##v}
``` -->
