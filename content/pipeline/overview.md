---
date: 2000-01-01T00:00:00+00:00
title: Overview
author: bradrydzewski
weight: 1
aliases:
- /configure/overview/
- /configure/pipeline/overview/
- /getting-started/
- /pipelines/
- /user-guide/
- /usage/overview/
- /0.5/usage/repository/configuration/
- /config
description: |
  Overview of Pipeline configuration.
---

Pipelines help you automate steps in your software delivery process, such as initiating code builds, running automated tests, and deploying to a staging or production environment.

Pipeline execution is triggered by a source code repository. A change in code triggers a webhook to Drone which runs the corresponding pipeline. Other common triggers include automatically scheduled or user-initiated workflows.

Pipelines are configured by placing a `.drone.yml` file in the root of your git repository. The yaml syntax is designed to be easy to read and expressive so that anyone viewing the repository can understand the workflow.

Example pipeline configuration:

{{< highlight yaml "linenos=table" >}}
---
kind: pipeline
type: docker
name: default

steps:
- name: backend
  image: golang
  commands:
  - go build
  - go test

- name: frontend
  image: node
  commands:
  - npm install
  - npm run test

...
{{< / highlight >}}

Drone supports different types of pipelines, each optimized for different use cases and runtime environments:

<!-- Drone supports different types of pipelines, each optimized for different use cases and runtime environments. [Docker]({{< relref "pipeline/docker/overview" >}}) pipelines, for example, execute pipeline steps inside isolated Docker containers. [Exec]({{< relref "pipeline/exec/overview" >}}) pipelines, on the other hand, execute pipelines steps directly on the host machine without isolation. -->


{{< link "pipeline/docker/overview" "Docker Pipelines" >}}
{{< link "pipeline/kubernetes/overview" "Kubernetes Pipelines" >}}
{{< link "pipeline/exec/overview" "Exec Pipelines" >}}
{{< link "pipeline/ssh/overview" "SSH Pipelines" >}}
{{< link "pipeline/digitalocean/overview" "Digital Ocean Pipelines" >}}
{{< link "pipeline/macstadium/overview" "MacStadium Pipelines" >}}


<!-- # Webhooks

Pipelines are triggered by webhooks sent from your source control management system (e.g. GitHub) every time you push code to your repository. When you activate a repository in the Drone user interface, Drone automatically registers a webhook with your source control management system. -->
<!-- 
# Pipelines
Drone supports different types of pipeline execution environments, where each type has its own custom yaml specification. The kind and type attributes define the type of pipeline and target execution environment.

{{< highlight text "linenos=table,hl_lines=2-3" >}}
---
kind: pipeline
type: docker
name: default
{{< / highlight >}}

## Docker Pipelines

Docker pipelines execute pipeline commands inside isolated Docker containers. Jump to the Docker pipeline documentation to learn more.



## Kubernetes Pipelines

Kubernetes pipelines execute pipeline commands inside pods, where each pipeline step is represented by a container in the pod. Jump to the Kubernetes pipeline documentation to learn more.

{{< link "pipeline/kubernetes/overview" "Kubernetes Pipelines" >}}

## Exec Pipelines

Exec pipelines execute pipeline commands directly on the host machine using the default shell. Jump to the exec pipeline documentation to learn more.

{{< link "pipeline/exec/overview" "Exec Pipelines" >}}

## SSH Pipelines

SSH pipelines execute pipeline commands on a remote machine using the SSH protocol. Jump to the ssh pipeline documentation to learn more.

{{< link "pipeline/ssh/overview" "SSH Pipelines" >}} -->
