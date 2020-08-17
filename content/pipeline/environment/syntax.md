---
date: 2000-01-01T00:00:00+00:00
title: Syntax
author: bradrydzewski
weight: 10
toc: true
aliases:
- /environment/
description: |
  Configure pipeline environment variables.
---

Drone provides the ability to define environment variables scoped to individual build steps. _The examples in this section showcase Docker pipelines, however, the syntax is shared across all pipeline types._

Example step with custom environment variables:

{{< highlight text "linenos=table,hl_lines=10-12" >}}
kind: pipeline
type: docker
name: default

steps:
- name: build
  commands:
  - go build
  - go test
  environment:
    GOOS: linux
    GOARCH: amd64
{{< / highlight >}}

Drone automatically injects environment variables containing repository and commit metadata into each pipeline step. See the environment [Reference]({{< relref "reference/_index.md" >}}) for a full list of injected variables.

# Per Pipeline

Drone supports global environment variables per pipeline. Globally defined variables are automatically injected into to every pipeline step.

<div class="alert">
Note this feature is only available to Docker pipelines at this time. It is not available to Kubernetes pipelines or other pipeline types.
</div>

{{< highlight text "linenos=table,hl_lines=5-7" >}}
kind: pipeline
type: docker
name: default

environment:
  GOOS: linux
  GOARCH: amd64

steps:
- name: build
  commands:
  - go build

- name: test
  commands:
  - go test
{{< / highlight >}} 

# From Secrets

<div class="alert">
Pipelines execute directly on the host without isolation. Any process running on the host may be able to intercept your secrets. You should only expose secrets on a trusted server.
</div>

Drone provides the ability to source environment variables from secrets. In the below example we provide the username and password as environment variables to the step.

{{< highlight text "linenos=table,linenostart=5,hl_lines=8-11" >}}
steps:
- name: build
  commands:
  - docker login -u $USERNAME -p $PASSWORD
  - docker build -t hello-world .
  - docker push hello-world
  environment:
    PASSWORD:
      from_secret: password
    USERNAME:
      from_secret: username
{{< / highlight >}}

# Common Problems

Parameter expansion is subject to pre-processing _before_ the yaml is parsed. If you do not want the system to evaluate an expression it must be escaped.

{{< highlight text "linenos=table,hl_lines=5,linenostart=5" >}}
steps:
- name: build
  commands:
  - echo $GOOS
  - echo $${GOARCH}
  - go build
  - go test
{{< / highlight >}}

Also note the environment section cannot expand environment variables or evaluate shell expressions. If you need to construct variables it should be done in the commands section.

Bad:

{{< highlight text "linenos=table,hl_lines=3-4,linenostart=5" >}}
steps:
- name: build
  environment:
    GOPATH: ${HOME}/golang
  commands:
  - go build
  - go test
{{< / highlight >}}

Good:

{{< highlight text "linenos=table,hl_lines=4,linenostart=5" >}}
steps:
- name: build
  commands:
  - export GOPATH=$HOME/golang
  - go build
  - go test
{{< / highlight >}}
