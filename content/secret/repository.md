---
date: 2000-01-01T00:00:00+00:00
title: Per Repository
author: bradrydzewski
weight: 5
disable_toc: true
aliases:
- /configure/secrets/
- /configure/secrets/repository/
- /usage/secret-guide/
- /manage-secrets/
- /user-guide/secrets/pre-repository
- /user-guide/secrets/per-repository
description: |
  Store and manage sensitive information per-repository.
---

Repository secrets are used to store and manage sensitive information, such as passwords, tokens, and ssh keys. Storing this information in a secret is considered safer than storing it in your configuration file in plain text.

<div class="alert alert-info">
Push access or administrator access is required to create, update or delete repository secrets.
</div>

Manage repository secrets from the repository settings screen:

![Repository Secrets](/screenshots/repository_secrets.png)

Source environment variables from named secrets:

```yaml {linenos=table,hl_lines=["8-11"],linenostart=1}
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
  commands:
    - echo "You can call the secrets like this examples below."
    - echo $USERNAME
    - echo $PASSWORD
    - echo "In both cases, and for security reasons, you will see asteriks '*******' instead the value under the echo command."
```

Source plugin settings from named secrets:

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
  commands:
    - echo "You can call the secrets like this examples below."
    - echo $PLUGIN_USERNAME
    - echo $PLUGIN_PASSWORD
    - echo "In both cases, and for security reasons, you will see asteriks '*******' instead the value under the echo command."
{{< / highlight >}}

# Pull Requests

Secrets are not exposed to pull requests by default. This prevents a bad actor from sending a pull request and attempting to expose your secrets. You can override this default behavior, at your own risk, by checking _Allow Pull Requests_ when you create your secret.
