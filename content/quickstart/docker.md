---
date: 2000-01-01T00:00:00+00:00
title: Docker Pipelines
title_in_header: Docker Pipelines
author: bradrydzewski
weight: 100
hidden: true
steps: true
icon: /icons/docker.svg
description: |
  Execute pipelines steps inside isolated containers.
---

<!-- This tutorial will help you create and execute a simple Docker pipeline. This tutorial is programming language agnostic. Pipelines are executed inside Docker containers, which means Drone supports any language that runs inside a container. -->

This tutorial will help you create and execute a simple Docker pipeline. Please see our pipeline [documentation]({{< relref "/pipeline/overview.md" >}}) for detailed usage instructions.

# Authenticate

First, navigate to your Drone server URL in your browser. If you are not already authenticated, Drone will redirect you to GitHub to login.

After login you are redirected back to your Drone dashboard. _If this is your first time using Drone your dashboard will be empty for a few seconds while Drone synchronizes your repository list with GitHub._

# Enable your Repository

Next, search for your repository and click the _Enable_ button. Clicking the enable button adds a webhook to your repository to notify Drone every time you push code. _Please note you must have admin privileges to the repository to enable._


# Configure your Pipeline

Next, you need to configure a pipeline by creating a `.drone.yml` file to the root of your git repository. In this file we define a series of steps that are executed every time a webhook is received.

```yaml {linenos=table}
kind: pipeline
type: docker
name: default

steps:
- name: greeting
  image: alpine
  commands:
  - echo hello
  - echo world
```

Here is a quick overview of the variables used in this example:

* `kind`
  : The kind attribute defines the kind of object. This example defines a pipeline object. Other kinds of object are [secret]({{< relref "secret/_index.md" >}}) and [signature]({{< relref "signature/_index.md" >}}) objects.
* `type`
  : The type attribute defines the type of pipeline. This example defines a Docker pipeline where each pipeline step is executed inside a Docker container. Drone supports [different types]({{< relref "pipeline/overview.md" >}}) of pipeline execution environments.
* `name`
  : The name attribute defines a name for your pipeline. You can define one or many pipelines for your project.
* `steps`
  : The steps section defines an array of pipeline steps that are executed serially. If any step in the pipeline fails, the pipeline exits immediately.
  * `name`
    : The name attribute defines the name of the pipeline step.
  * `image`
    : The image attribute defines a Docker image in which the shell commands are executed. You can use any Docker image in your pipeline from any Docker registry, including private registries.
  * `commands`
    : The commands attribute defines a list of shell commands that are executed inside the Docker container as the container entrypoint. If any command returns a non-zero exit code the pipeline step fails.

Please see our pipeline [documentation]({{< relref "/pipeline/docker/overview.md" >}}) for a full list of configuration options.



## Additional Examples

* You can add multiple steps to your pipeline:
    ```yaml  {linenos=table}
    kind: pipeline
    type: docker
    name: greeting

    steps:
    - name: en
      image: alpine
      commands:
      - echo hello world

    - name: fr
      image: alpine
      commands:
      - echo bonjour monde
    ```

* You can conditionally [limit]({{< relref "pipeline/docker/syntax/conditions.md" >}}) step execution:
    ```yaml  {linenos=table, hl_lines=["15-17"]}
    kind: pipeline
    type: docker
    name: greeting

    steps:
    - name: en
      image: alpine
      commands:
      - echo hello world

    - name: fr
      image: alpine
      commands:
      - echo bonjour monde
      when:
        branch:
        - develop
    ```

* You can even define [multiple pipelines]({{< relref "pipeline/configuration.md#multiple-pipelines" >}}):
    ```yaml  {linenos=table}
    kind: pipeline
    type: docker
    name: en

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo hello world

    ---
    kind: pipeline
    type: docker
    name: fr

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo bonjour monde
    ```

* You can conditionally [limit]({{< relref "pipeline/docker/syntax/trigger.md" >}}) pipeline execution:
    ```yaml  {linenos=table hl_lines=["11-13", "26-28"]}
    kind: pipeline
    type: docker
    name: en

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo hello world

    trigger:
      event:
      - push

    ---
    kind: pipeline
    type: docker
    name: fr

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo bonjour monde
    
    trigger:
      event:
      - pull_request
    ```

* You can use any [image]({{< relref "pipeline/docker/syntax/images.md" >}}) from any docker registry:
    ```yaml  {linenos=table, hl_lines=[7]}
    kind: pipeline
    type: docker
    name: default

    steps:
    - name: test
      image: gcr.io/library/golang
      commands:
      - go build
      - go test -v
    ```

* You can define [service containers]({{< relref "pipeline/docker/syntax/services.md" >}}) for integration tests:
    ```yaml  {linenos=table, hl_lines=["12-14"]}
    kind: pipeline
    type: docker
    name: default

    steps:
    - name: test
      image: golang:1.13
      commands:
      - go build
      - go test -v

    services:
    - name: redis
      image: redis
    ```

* You can use [plugins]({{< relref "plugins/_index.md" >}}) to integrate with third party systems and perform common tasks, such as notify, publish or deploy software.
    ```yaml  {linenos=table, hl_lines=["12-16"]}
    kind: pipeline
    type: docker
    name: default

    steps:
    - name: test
      image: golang:1.13
      commands:
      - go build
      - go test -v

    - name: notify
      image: plugins/slack
      settings:
        channel: dev
        webhook: https://hooks.slack.com/services/...
    ```

* You can also source sensitive parameters from [secrets]({{< relref "secret/_index.md" >}}):
    ```yaml  {linenos=table, hl_lines=["16-17"]}
    kind: pipeline
    type: docker
    name: default

    steps:
    - name: test
      image: golang:1.13
      commands:
      - go build
      - go test -v

    - name: notify
      image: plugins/slack
      settings:
        channel: dev
        webhook:
          from_secret: endpoint
    ```

# Execute your Pipeline

The final step is to commit your `.drone.yml` to your repository and push your changes. When you push code, GitHub sends a webhook to Drone which in turn executes your pipeline.
