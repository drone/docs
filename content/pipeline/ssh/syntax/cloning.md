---
date: 2000-01-01T00:00:00+00:00
title: Cloning
author: bradrydzewski
weight: 3
toc: true
description: |
  Configure the pipeline to use custom clone logic.
---

Drone automatically clones your repository before executing your pipeline steps. No special configuration is required. In some cases, however, you may need to customize, override or disable the default clone behavior.

# The `--depth` flag

The default clone configuration does use the `--depth` flag. You can enforce a clone depth by declaring a `clone` block and adding the `depth` attribute:

```yaml {linenos=table, hl_lines=["11-12"]}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

clone:
  depth: 50

steps:
- name: build
  commands:
  - go build
  - go test
```

# The `--tags` flag

The default clone configuration does not use the `--tags` flag. If you would like to fetch tags you should handle this as a step in your pipeline. For example:

```yaml {linenos=table, hl_lines=["12-14"]}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

steps:
- name: fetch
  commands:
  - git fetch --tags

- name: build
  commands:
  - go build
  - go test
```


# The `--recursive` flag

The default clone behavior does not use the `--recursive` flag and does not fetch submodules. If you would like to fetch submodules you should handle this as a step in your pipeline. For example:

```yaml {linenos=table, hl_lines=["12-14"]}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

steps:
- name: submodules
  commands:
  - git submodule update --recursive --remote

- name: build
  commands:
  - go build
  - go test
```

# Custom Logic

The default clone behavior can be disabled and custom clone logic implemented, when necessary. In the following example we implement custom clone commands as a pipeline step:

```yaml {linenos=table, hl_lines=["11-12", "15-18"]}
kind: pipeline
type: ssh
name: default

server:
  host: 1.2.3.4
  user: root
  password:
    from_secret: password

clone:
  disable: true

steps:
- name: clone
  commands:
  - git clone https://github.com/octocat/hello-world.git .
  - git checkout $DRONE_COMMIT

- name: build
  commands:
  - go build
  - go test
```
