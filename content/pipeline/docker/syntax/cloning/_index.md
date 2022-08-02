---
date: 2000-01-01T00:00:00+00:00
title: Cloning
author: bradrydzewski
weight: 40
toc: true
aliases:
- /cloning/
- /user-guide/pipeline/cloning/
description: |
  Configure the pipeline to use custom clone logic.
---

Drone automatically clones your repository before executing your pipeline steps. No special configuration is required. In some cases, however, you may need to customize, override or disable the default clone behavior.

# The `--depth` flag

The default clone configuration does use the `--depth` flag. You can enforce a clone depth by declaring a `clone` block and adding the `depth` attribute:

{{< highlight text "linenos=table,hl_lines=5-6" >}}
kind: pipeline
type: docker
name: default

clone:
  depth: 50

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}

# The `--tags` flag

The default clone configuration does not use the `--tags` flag. If you would like to fetch tags you should handle this as a step in your pipeline. For example:

{{< highlight text "linenos=table,hl_lines=6-9" >}}
kind: pipeline
type: docker
name: default

steps:
- name: fetch
  image: alpine/git
  commands:
  - git fetch --tags

- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}


# The `--recursive` flag

The default clone behavior does not use the `--recursive` flag and does not fetch submodules. If you would like to fetch submodules you should handle this as a step in your pipeline. For example:

{{< highlight text "linenos=table,hl_lines=6-9" >}}
kind: pipeline
type: docker
name: default

steps:
- name: submodules
  image: alpine/git
  commands:
  - git submodule update --init --recursive

- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}

# retries

By default, the clone step will fail if any of the git commands fail. This can be altered by setting a number of retries. When this is set, the failing command will be rerun the given number of times, and the step will only be marked as failure if none of the tries are successful. For example:


{{< highlight text "linenos=table,hl_lines=5-6" >}}
kind: pipeline
type: docker
name: default

clone:
  retries: 3

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}

# Custom Logic

The default clone behavior can be disabled and custom clone logic implemented, when necessary. In the following example we implement custom clone commands as a pipeline step:

{{< highlight text "linenos=table,hl_lines=9-13 5-6" >}}
kind: pipeline
type: docker
name: default

clone:
  disable: true

steps:
- name: clone
  image: alpine/git
  commands:
  - git clone https://github.com/octocat/hello-world.git .
  - git checkout $DRONE_COMMIT

- name: build
  image: golang
  commands:
  - go build
  - go test
{{< / highlight >}}

# Authentication

If your repository is private or requires authentication to clone, Drone automatically injects the clone credentials into your pipeline environment. Drone uses the oauth2 token associated with the repository to authenticate. Please read the below article to learn more about authentication.

{{< link "/content/pipeline/docker/syntax/cloning/auth.md" >}}
