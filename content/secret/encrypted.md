---
date: 2000-01-01T00:00:00+00:00
title: Encrypted
author: bradrydzewski
weight: 100
disable_toc: true
aliases:
- /configure/secrets/encrypted/
- /user-guide/secrets/encrypted/
description: |
  Store sensitive information as encrypted strings in your configuration file.
---

Encrypted secrets are used to store sensitive information, such as passwords, tokens, and ssh keys directly in your configuration file as an encrypted string. Each secret is represented as a yaml document in your configuration file.

You can use the command line tools to encrypt secrets. Each secret is encrypted with a per-repository encryption key using aesgcm. This key never leaves the server environment.

Example command to encrypt the secret:

```
$ drone encrypt <repository> <secret>
```

```
$ drone encrypt secret octocat/hello-world top-secret-password
hl3v+FODjduX0UpXBHgYzPzVTppQblg51CVgCbgDk4U=
```

Example configuration with encrypted secrets:

{{< highlight yaml "linenos=table,hl_lines=8-9 12-14" >}}
kind: pipeline
name: default

steps:
- name: build
  image: alpine
  environment:
    USERNAME:
      from_secret: username

---
kind: secret
name: username
data: hl3v+FODjduX0UpXBHgYzPzVTppQblg51CVgCbgDk4U=

...
{{< / highlight >}}

# Pull Requests

Secrets are not exposed to pull requests that originate from forks. This prevents a bad actor from sending a pull request and attempting to expose your secrets.
