---
date: 2000-01-01T00:00:00+00:00
title: Substitution
author: bradrydzewski
weight: 20
toc: true
aliases:
- /substitution/
- /usage/variables
description: |
  Learn how Drone emulates bash string substitution.
---

Drone provides the ability to expand, or _substitute_, repository and build metadata to facilitate dynamic pipeline configurations.

* Example commit substitution:
  ```
  kind: pipeline
  type: docker
  name: default

  steps:
  - name: publish
    image: plugins/docker
    settings:
      tags: ${DRONE_COMMIT}
      repo: octocat/hello-world
  ```

* Example tag substitution:

  ```
  steps:
  - name: publish
    image: plugins/docker
    settings:
      tags: ${DRONE_TAG}
      repo: octocat/hello-world
  ```

Please see the environment [Reference]({{< relref "./reference/_index.md" >}}) for a list of parameters that can be used for substitution. _Please note that some parameters in this list are unavailable for substitution, such as step name, step number, as well as parameters that store statuses and timestamps._

# String Operations

Drone provides partial emulation for bash string operations. This can be used to manipulate string values prior to substitution.

* Example variable substitution with substring:
  ```
  steps:
  - name: publish
    image: plugins/docker
    settings:
      tags: ${DRONE_COMMIT_SHA:0:8}
      repo: octocat/hello-world
  ```

* Example variable substitution strips v prefix from v1.0.0:
  ```
  steps:
  - name: publish
    image: plugins/docker
    settings:
      tags: ${DRONE_TAG##v}
      repo: octocat/hello-world
  ```
  
* Example variable substitution replaces ```/``` with ```-```:
  ```
  steps:
  - name: publish
    image: plugins/docker
    settings:
      tags: ${DRONE_BRANCH/\//-}
      repo: octocat/hello-world
  ```

Drone emulates the below string operations. _Drone makes a best-effort to emulate these operations however we do not promise perfect emulation._


```
${parameter^}
${parameter^^}
${parameter,}
${parameter,,}
${parameter:position}
${parameter:position:length}
${parameter#substring}
${parameter##substring}
${parameter%substring}
${parameter%%substring}
${parameter/substring/replacement}
${parameter//substring/replacement}
${parameter/#substring/replacement}
${parameter/%substring/replacement}
${#parameter}
${parameter=default}
${parameter:=default}
${parameter:-default}
```


<!-- * `${parameter^}`
* `${parameter^^}`
* `${parameter,}`
* `${parameter,,}`
* `${parameter:position}`
* `${parameter:position:length}`
* `${parameter#substring}`
* `${parameter##substring}`
* `${parameter%substring}`
* `${parameter%%substring}`
* `${parameter/substring/replacement}`
* `${parameter//substring/replacement}`
* `${parameter/#substring/replacement}`
* `${parameter/%substring/replacement}`
* `${#parameter}`
* `${parameter=default}`
* `${parameter:=default}`
* `${parameter:-default}` -->

<!--
OPERATION	        | DESC
--------------------|---
`${param}`          | parameter substitution
`${param,}`         | parameter substitution with lowercase first char
`${param,,}`        | parameter substitution with lowercase
`${param^}`         | parameter substitution with uppercase first char
`${param^^}`        | parameter substitution with uppercase
`${param:pos}`      | parameter substitution with substring
`${param:pos:len}`  | parameter substitution with substring and length
`${param=default}`  | parameter substitution with default
`${param##prefix}`  | parameter substitution with prefix removal
`${param%%suffix}`  | parameter substitution with suffix removal
`${param/old/new}`  | parameter substitution with find and replace
-->

# Escaping

Parameter expressions are evaluated _before_ the yaml is parsed. If you do not want the system to evaluate an expression it must be escaped.

```yaml {linenos=table, hl_lines=["9"]}
kind: pipeline
type: docker
name: default

steps:
- name: build
  commands:
  - echo $GOOS
  - echo $${GOARCH}
  - go build
  - go test
```

# Common Problems

Parameter substitution occurs _before_ the yaml is parsed. If the substitution results in an invalid yaml file you will receive a parsing error:

```
yaml: unmarshal errors:
cannot unmarshal !!map into string
```

This can be resolved by quoting parameters to ensure special / reserved characters are escaped:

```yaml {linenos=table, hl_lines=["10"]}
kind: pipeline
type: docker
name: default

steps:
- name: notify
  image: plugins/slack
  settings:
    channel: team
    message: "${DRONE_COMMIT}"
```
