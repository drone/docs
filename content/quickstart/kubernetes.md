---
date: 2000-01-01T00:00:00+00:00
title: Kubernetes Pipelines
title_in_header: Kubernetes Pipelines
author: bradrydzewski
weight: 200
hidden: true
steps: true
icon: /icons/kubernetes.svg
description: |
  Execute pipelines steps inside Kubernetes as pods.

---

<div class="alert">
Please note kubernetes pipelines are disabled on Drone Cloud. This feature is only available when self-hosting.
</div>

This tutorial will help you create and execute a simple Kubernetes pipeline. Please see our pipeline [documentation]({{< relref "/pipeline/overview.md" >}}) for detailed usage instructions.

# Authenticate

First, navigate to your Drone server URL in your browser. If you are not already authenticated, Drone will redirect you to GitHub to login.

After login you are redirected back to your Drone dashboard. _If this is your first time using Drone your dashboard will be empty for a few seconds while Drone synchronizes your repository list with GitHub._

# Enable your Repository

Next, search for your repository and click the _Enable_ button. Clicking the enable button adds a webhook to your repository to notify Drone every time you push code. _Please note you must have admin privileges to the repository to enable._

# Configure your Pipeline

Next, you need to configure a pipeline by creating a `.drone.yml` file to the root of your git repository. In this file we define a series of steps that are executed every time a webhook is received.

```yaml {linenos=table}
kind: pipeline
type: kubernetes
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
  : The type attribute defines the type of pipeline. This example defines a Kubernetes pipeline where each pipeline step is executed inside a Pod. Drone supports [different types]({{< relref "pipeline/overview.md" >}}) of pipeline execution environments.
* `name`
  : The name attribute defines a name for your pipeline. You can define one or many pipelines for your project.
* `steps`
  : The steps section defines an array of pipeline steps that are executed serially. If any step in the pipeline fails, the pipeline exits immediately.
  * `name`
    : The name attribute defines the name of the pipeline step.
  * `image`
    : The image attribute defines a container image in which the shell commands are executed. You can use any image in your pipeline from any registry, including private registries.
  * `commands`
    : The commands attribute defines a list of shell commands that are executed inside a container as the entrypoint. If any command returns a non-zero exit code the pipeline step fails.

Please see our pipeline [documentation]({{< relref "/pipeline/kubernetes/overview.md" >}}) for a full list of configuration options.

## Additional Examples

* You can add multiple steps to your pipeline:

    ```yaml
    kind: pipeline
    type: kubernetes
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

* You can conditionally [limit]({{< relref "pipeline/docker/syntax/conditions.md" >}}) pipeline steps to execute based on branch or webhook events:

    ```yaml  {linenos=table, hl_lines=["15-17"]}
    kind: pipeline
    type: kubernetes
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

    ```yaml {linenos=table, hl_lines=["5-20"]}
    kind: pipeline
    type: kubernetes
    name: en

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo hello world

    ---
    kind: pipeline
    type: kubernetes
    name: fr

    steps:
    - name: greeting
      image: alpine
      commands:
      - echo bonjour monde
    ```

* You can use any [image]({{< relref "pipeline/docker/syntax/images.md" >}}) from any container registry:

    ```yaml
    kind: pipeline
    type: kubernetes
    name: default

    steps:
    - name: test
      image: golang:1.13
      commands:
      - go build
      - go test -v
    ```

# Execute your Pipeline

The final step is to commit your `.drone.yml` to your repository and push your changes. When you push code, GitHub sends a webhook to Drone which in turn executes your pipeline.
