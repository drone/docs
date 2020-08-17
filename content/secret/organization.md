---
date: 2000-01-01T00:00:00+00:00
title: Per Organization
author: bradrydzewski
weight: 20
disable_toc: true
aliases:
- /configure/secrets/organization/
- /global-secrets/
- /global-registries/
- /setup-global-secrets/
- /setup-global-registry-credentials
description: |
  Store and manage sensitive information per-organization.
---

Organization secrets are used to store and manage sensitive information, such as passwords, tokens, and ssh keys. Storing this information in a secret is considered safer than storing it in your configuration file. Organization secrets can be used by any repository that belongs to the named organization.

<div class="alert alert-no-cloud">
Please note this feature is disabled on Drone Cloud. This feature is only available when self-hosting.
</div>

<div class="alert alert-warn">
The system administrator role is required to create, update or delete organization secrets.
</div>

Create organization secrets using the command line tools:

```
$ drone orgsecret add [organization] [name] [data]
```

```
$ drone orgsecret add octocat docker_password pa55word
```

Source environment variables from named organization secrets:

{{< highlight yaml "linenos=table,hl_lines=8-11" >}}
kind: pipeline
name: default

steps:
- name: build
  image: alpine
  environment:
    USERNAME:
      from_secret: docker_username
    PASSWORD:
      from_secret: docker_password
{{< / highlight >}}

Source plugin settings from named organization secrets:

{{< highlight yaml "linenos=table,hl_lines=9-12" >}}
kind: pipeline
name: default

steps:
- name: build
  image: plugins/docker
  settings:
    repo: octocat/hello-world
    username:
      from_secret: docker_username
    password:
      from_secret: docker_password
{{< / highlight >}}

# Pull Requests

Secrets are not exposed to pull requests by default. This prevents a bad actor from sending a pull request and attempting to expose your secrets. You can override this default behavior, at your own risk, using the `--allow-pull-request` flag.