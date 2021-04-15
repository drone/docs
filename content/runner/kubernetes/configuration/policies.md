---
date: 2000-01-01T00:00:00+00:00
title: Policies
author: bradrydzewski
weight: 2
description: |
  Configure policies to apply pipeline defaults.
---

The Drone policy file gives you the ability to define policies that set and enforce pipeline values. For example, this gives you the ability to set namespace, tolerations and more based on organization, repository and other matching criteria.

Example policy file:

```yaml {linenos=table}
---
kind: policy
name: octocat

match:
  repo:
  - "octocat/*"
  - "octocat/hello-world"

metadata:
  namespace: octocat

resources:
  request:
    cpu: 1
    memory: 512MiB
  limit:
    cpu: 4
    memory: 1GiB

node_selector:
  disktype: ssd

---
kind: policy
name: default

metadata:
  namespace: default
```

The policy file must be mounted into your runner container and you must provide the runner the location of the policy file.  See the [policy configuration parameter]({{< relref "runner/kubernetes/configuration/reference/drone-policy-file.md" >}}) for configuration instructions.

# Multiple Policies

You can define multiple policies in the policy file.  The `match` section is used to match the policy the pipeline. The first matching policy is applied to the pipeline.

```yaml {linenos=table, hl_lines=["5-8"]}
---
kind: policy
name: octocat

match:
  repo:
  - "octocat/*"
  - "octocat/hello-world"

metadata:
  namespace: octocat

---
kind: policy
name: default

metadata:
  namespace: default
```

# Default Policies

You can optionally define a default policy in the policy file, named accordingly.  The default policy is applied if no other policy matches the pipeline.

```yaml {linenos=table, hl_lines=["14-18"]}
---
kind: policy
name: octocat

match:
  repo:
  - "octocat/*"
  - "octocat/hello-world"

metadata:
  namespace: octocat

---
kind: policy
name: default

metadata:
  namespace: default
```

# File Format

* `kind`
  : The kind attribute defines the kind of object.
* `name`
  : The name attribute defines a name for your policy.
* `metadata`
  : The metadata section defines metadata attached to the pipeline pod.
  * `namespace`
    : The namespace attribute defines the namespace in which the pipeline pod is created. This takes precedence over the value defined in the yaml.
  * `annotations`
    : The annotations attribute defines a set of arbitrary key / value pairs that are attached to the pipeline pod. These are appended to existing annotations that are defined in the yaml and take precedence on conflict.
  * `labels`
    : The annotations attribute defines a set of arbitrary key / value pairs that are attached to the pipeline pod. These are appended to existing labels that are defined in the yaml and take precedence on conflict.
* `resources`
  : The resource attribute defines resource requirements and limits for pipeline steps.
  * `request`
    : The request section defines resource requirements used when the scheduler defines which node to place the pipeline pod on.
    * `cpu`
      : The cpu attribute defines cpu requirements.
    * `memory`
      : The memory attribute defines memory requirements.
  * `limit`
    : The limit section defines container resource limits applied to each pipeline step.
    * `cpu`
      : The cpu attribute defines cpu limits.
    * `memory`
      : The memory attribute defines memory limits.
* `service_account`
  : The service_account attribute defines the kubernetes service account used to create the pipeline pod. This takes precedence over the value defined in the yaml.
* `node_selector`
  : The node_selector attribute defines a set of key / value pairs used to route pipeline pods to matching nodes. This takes precedence over the values defined in the yaml.
* `tolerations`
  : The tolerations section defines and applies [tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) to pipeline pods to schedule onto nodes with matching taints.
  * `effect`
    : The effect attribute defines the taint effect.
  * `key`
    : The key attribute defines the toleration key.
  * `operator`
    : The key attribute defines the toleration operator.
  * `toleration_seconds`
    : The key attribute defines the toleration seconds.
  * `value`
    : The key attribute defines the toleration value.

# Examples

* Example policy sets the default service account:

  ```yaml {linenos=table, hl_lines=["4"]}
  kind: policy
  name: default

  service_account: drone
  ```

* Example policy sets the default service account for matching pipelines:

  ```yaml {linenos=table, hl_lines=["6-9"]}
  kind: policy
  name: default

  service_account: drone

  match:
    repo:
    - "octocat/*"
    - "octocat/hello-world"
  ```

* Example policy sets the default namespace:

  ```yaml {linenos=table, hl_lines=["4-5"]}
  kind: policy
  name: default

  metadata:
    namespace: default

  match:
    repo:
    - "octocat/*"
    - "octocat/hello-world"
  ```

* Example policy sets the default resource limits:

  ```yaml {linenos=table, hl_lines=["7-13"]}
  kind: policy
  name: default

  metadata:
    namespace: default

  resources:
    request:
      cpu: 1
      memory: 512MiB
    limit:
      cpu: 4
      memory: 1GiB
  ```

* Example policy sets the default node selection:

  ```yaml {linenos=table, hl_lines=["7-8"]}
  kind: policy
  name: default

  metadata:
    namespace: default

  node_selector:
    disktype: ssd
  ```

* Example policy sets the default metadata:

  ```yaml {linenos=table, hl_lines=["4-11"]}
  kind: policy
  name: default

  metadata:
    namespace: default
    labels:
      keyA: valueA
      keyB: valueB
    annotations:
      keyA: valueA
      keyB: valueB
  ```
