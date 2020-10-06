---
date: 2000-01-01T00:00:00+00:00
title: Vault
title_in_tree: From Vault
author: bradrydzewski
weight: 50
toc: true
aliases:
- /configure/secrets/external/overview/
- /configure/secrets/external/vault/
- /vault-secrets/
description: |
  Retrieve sensitive information from Vault.
---

Vault secures, stores, and tightly controls access to tokens, passwords, certificates, keys, and other secrets in modern computing. The Vault extension provides your pipeline with access to Vault secrets.

<div class="alert alert-no-cloud">
Please note this feature is disabled on Drone Cloud. This feature is only available when self-hosting.
</div>

<div class="alert alert-info">
Vault integration is provided by an extension and is only available if you system administrator has installed the extension.
</div>


# Creating Secrets

Use the Vault command line tools to write secrets to the store. In the below example we store the Docker username and password.

{{< highlight text "linenos=table,hl_lines=2-3" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple
{{< / highlight >}}

# Accessing Secrets

Once our secrets are stored in Vault, we can update our yaml configuration file to request access to our secrets. First we define a secret resource in our yaml for each external secret. We include the path to the secret, and the name or key of value we want to retrieve:

{{< highlight text "linenos=table,hl_lines=10-21" >}}
---
kind: pipeline
name: default

steps:
- name: build
  image: alpine

---
kind: secret
name: username
get:
  path: secrets/data/docker
  name: username

---
kind: secret
name: password
get:
  path: secrets/data/docker
  name: password
...
{{< / highlight >}}

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
  path: secrets/data/docker
  name: username

---
kind: secret
name: password
get:
  path: secrets/data/docker
  name: password

...
{{< / highlight >}}

# Limiting Access

Secrets are available to all repositories and all build events by default. We strongly recommend that you limit access to secrets by repository, build events and branch. This can be done by adding special properties:

{{< highlight text "linenos=table,hl_lines=4-5" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-events=push,tag \
    x-drone-repos=octocat/*,spaceghost/* \
    x-drone-branches=master,development
{{< / highlight >}}

## Limit By Repository

Use the `X-Drone-Repos` key to limit which repositories can access your secret. The value is a comma-separate list of glob patterns. If a repository name matches at least one of the patterns, it is granted access to the secret.

Limit access to a single repository:

{{< highlight text "linenos=table,hl_lines=3-4" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-repos=octocat/hello-world
{{< / highlight >}}

Limit access to all repositories in an organization:

{{< highlight text "linenos=table,hl_lines=4" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-repos=octocat/*
{{< / highlight >}}

Limit access to multiple repositories or organizations:

{{< highlight text "linenos=table,hl_lines=4" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-repos=octocat/*,spaceghost/*
{{< / highlight >}}

## Limit By Event

Use the `X-Drone-Events` key to limit which build events can access your secret. The value is a comma-separate list of events. If a build matches at least one of the events, it is granted access to the secret.

Limit access to push and tag events:

{{< highlight text "linenos=table,hl_lines=4" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-events=push,tag
{{< / highlight >}}

You can combine annotations to limit by repository and event:

{{< highlight text "linenos=table,hl_lines=4-5" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-events=push,tag \
    x-drone-repos=octocat/*,spaceghost/*
{{< / highlight >}}

## Limit By Branch

Use the `X-Drone-Branches` key to limit which branch can access your secret. The value is a comma-separate list of branches. If a build matches at least one of the branch, it is granted access to the secret.

Limit access to master and development branches:

{{< highlight text "linenos=table,hl_lines=4" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-branches=master,development
{{< / highlight >}}

You can combine annotations whatever you need to limit by repository, event and branches:

{{< highlight text "linenos=table,hl_lines=4-5" >}}
$ vault kv put secret/docker \
    username=octocat \
    password=correct-horse-battery-staple \
    x-drone-events=push,tag \
    x-drone-repos=octocat/*,spaceghost/* \
    x-drone-branches=master,development
{{< / highlight >}}
