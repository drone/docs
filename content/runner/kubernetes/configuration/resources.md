---
date: 2021-07-28T12:16:38+00:00
title: Resources
author: marko-gacesa
weight: 2
description: |
  Managing Kubernetes resources.
---

The Kubernetes runner supports management of Kubernetes resources.

# Types of Resources

Two types of resources are supported:
* Memory
* CPU

The values for CPU resource are given as an integer and represent number of millicores.
Memory resource values are in bytes, but are typically used with a unit, for example: `Mi` or `Gi`.

For example, to request 2 cores and 640 megabytes of memory use `cpu=2000` and `memory=640Mi`. 

# Resource Requests

Resource request tells Kubernetes how much resources a pipeline needs, so it can correctly select a node on which to run a pod. So, values apply to an entire pod on which a pipeline is executed on.

Internally Kubernetes runner will split the values, so that every container gets approximately the same amount. In total values for all containers will sum to the requested value.
Note that this is only initial value for a container. It's allowed to use more.

It can be set with:
* Environment variables: `DRONE_RESOURCE_REQUEST_CPU` and `DRONE_RESOURCE_REQUEST_MEMORY`
* YAML, pipeline `resources` section (overrides environment variables)
* Policy (overrides all other values)

Example:

```
DRONE_RESOURCE_REQUEST_CPU=2000
DRONE_RESOURCE_REQUEST_MEMORY=640Mi
```

# Resource Limits

Containers can be limited to use a resource only up to a certain value.

Limits can be set with:
* Environment variables: `DRONE_RESOURCE_LIMIT_CPU` and `DRONE_RESOURCE_LIMIT_MEMORY`
* Policy (overrides environment variables)
* YAML, `resources` section defined in each step (can only be less than env/policy)

Note that unlike resource requests, a resource limit is set for each container individually. In addition to that, in YAML limits can be set for each step separately (but can't be higher than the value set by env variable or a policy).

Example:

```
DRONE_RESOURCE_LIMIT_CPU=1500
DRONE_RESOURCE_LIMIT_MEMORY=512Mi
```

# Minimum Values

There is a way to guarantee that each container will get at least a minimum value of a resource when the runner splits a resource request between containers in a pod.

Minimum values can be set with environment variables `DRONE_RESOURCE_MIN_REQUEST_CPU` and `DRONE_RESOURCE_MIN_REQUEST_MEMORY`.

Default values for these are:
* `DRONE_RESOURCE_MIN_REQUEST_CPU=1`
* `DRONE_RESOURCE_MIN_REQUEST_MEMORY=4Mi`

Note that setting a value too high can result that the sum of resource requests for each container is higher than the pod resource request.

# Example Yaml

```yaml {linenos=table}
kind: pipeline
type: kubernetes

# resource requests are defined on the pipeline level
resources:
  requests:
    cpu: 2000
    memory: 2000MiB

steps:
  - name: en
    image: alpine
    commands:
      - echo hello
    resources:
      # resource limits are defined for each step
      limits:
        cpu: 1000
        memory: 1000MiB

  - name: es
    image: alpine
    commands:
      - echo hola

  - name: fr
    image: alpine
    commands:
      - echo bonjour
```
