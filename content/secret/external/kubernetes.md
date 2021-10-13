---
date: 2000-01-01T00:00:00+00:00
title: Kubernetes
author: bradrydzewski
weight: 60
toc: true
aliases:
- /configure/secrets/external/kubernetes/
description: |
  Retrieve sensitive information from Kubernetes secrets.
---

The Kubernetes Secret resource secures, stores, and controls access to tokens, passwords, certificates, and other secrets in modern computing. The Kubernetes Secrets <a href="/content/runner/extensions/kube.md">extension</a> provides your pipeline with access to Kubernetes secrets.

<div class="alert alert-info">
Kubernetes Secrets integration is provided by an extension and is only available if your system administrator has installed this <a href="/content/runner/extensions/kube.md">extension</a>.
</div>

# Creating Secrets

Create a secret resource using the Kubernetes yaml configuration language, and persist to your cluster using `kubectl`. In the below example we store the Docker username and password.

|    apiVersion   |    kind         |     type       |        Data:                    |     metadata:      |
|    :----:       |    :----:       |     :----:     |       :----:                    |     :----:         |
|     v1          |    Secret       |     Opaque     |   username: **YWRtaW4=**        | name: **docker**   |
|                 |                 |                | password: **MWYyZDFlMmU2N2Rm**  |                    |

# Accessing Secrets

Once our secrets are stored in Kubernetes, we can update our yaml configuration file to request access to our secrets. First we define a secret resource in our yaml for each external secret. We include the path to the secret, and the name or key of value we want to retrieve:

<div display=flex>

|    kind         |     name         |                   steps                  |
|    :----:       |     :----:       |                   :----:                 |
|    pipeline     |     default      |     name: build <br/> image: alphine     |


|    kind         |     name         |                   get                     |
|    :----:       |     :----:       |                   :----:                  |
|    secret       |     username     |     name: docker <br/> name: username     |


|    kind         |     name         |                   get                     |
|    :----:       |     :----:       |                   :----:                  |
|    secret       |     password     |     name: docker <br/> name: password     |

</div>
  
We can then reference the named secrets in our pipeline:

{{< highlight text "linenos=table,hl_lines=8-11" >}}
kind: pipeline
name: default

steps:
- name: build
  image: alpine
  environment:
    USERNAME:
      from_secret: username
    PASSWORD:
      from_secret: password

---
kind: secret
name: username
get:
  path: docker
  name: username

---
kind: secret
name: password
get:
  path: docker
  name: password

...
{{< / highlight >}}

# Limiting Access

Secrets are available to all repositories and all build events by default. We strongly recommend that you limit access to secrets by repository and build events. This can be done by adding special annotations:

{{< highlight text "linenos=table,hl_lines=10-11" >}}
apiVersion: v1
kind: Secret
type: Opaque
data:
  username: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm
metadata:
  name: docker
  annotations:
    X-Drone-Repos: octocat/*
    X-Drone-Events: push,tag
{{< / highlight >}}

## Limit By Repository

Use the `X-Drone-Repos` key to limit which repositories can access your secret. The value is a comma-separate list of glob patterns. If a repository name matches at least one of the patterns, it is granted access to the secret.

Limit access to a single repository:

{{< highlight text "linenos=table,linenostart=7,hl_lines=4" >}}
metadata:
  name: docker
  annotations:
    X-Drone-Repos: octocat/hello-world
{{< / highlight >}}

Limit access to all repositories in an organization:

{{< highlight text "linenos=table,linenostart=7,hl_lines=4" >}}
metadata:
  name: docker
  annotations:
    X-Drone-Repos: octocat/*
{{< / highlight >}}

Limit access to multiple repositories or organizations:

{{< highlight text "linenos=table,linenostart=7,hl_lines=4" >}}
metadata:
  name: docker
  annotations:
    X-Drone-Repos: octocat/*,spaceghost/*
{{< / highlight >}}

## Limit By Event

Use the `X-Drone-Events` key to limit which build events can access your secret. The value is a comma-separate list of events. If a build matches at least one of the events, it is granted access to the secret.

Limit access to push and tag events:

{{< highlight text "linenos=table,linenostart=7,hl_lines=4" >}}
metadata:
  name: docker
  annotations:
    X-Drone-Events: push,tag
{{< / highlight >}}

You can combine annotations to limit by repository and event:

{{< highlight text "linenos=table,linenostart=7,hl_lines=4-5" >}}
metadata:
  name: docker
  annotations:
    X-Drone-Repos: octocat/*
    X-Drone-Events: push,tag
{{< / highlight >}}
