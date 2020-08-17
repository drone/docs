---
date: 2000-01-01T00:00:00+00:00
title: Signatures
author: bradrydzewski
weight: 140
toc: true
aliases:
- /configure/signature/
- /user-guide/signature
- /gated-builds/
description: |
  Sign your configuration file to prevent tampering.
---

You can optionally sign your configuration file to verify authenticity and prevent tampering. This is useful if your repository is public and you need to prevent unauthorized changes to your configuration.

If a user modifies the configuration and signature verification fails, the pipeline is blocked pending manual approval by an authorized user with write or administrative access to the repository.

# Enforcing Signatures

To enforce signature verification you need to enable Protected mode for your repository. Navigate to your repository _Settings_ screen and check the _Protected_ checkbox.

# Storing Signatures

Signatures are stored in the Yaml configuration file as a `signature` resource. The signature resource provides an hmac signature of your configuration.

{{< highlight text "linenos=table,hl_lines=14-16" >}}
---
kind: pipeline
type: docker
name: default

steps:
- name: build
  image: golang
  commands:
  - go build
  - go test

---
kind: signature
hmac: F10E2821BBBEA527EA02200352313BC059445190

...
{{< / highlight >}}

# Calculating Signatures

The contents of each yaml resource, excluding any existing signature resources, are signed using a 256-bit secret key. The secret key is unique per-repository, and never leaves the Drone server.

# Creating Signatures

The signature is created using the Drone command line utility. This command makes an authenticated request to the Drone server, posting your yaml configuration file, to calculate and return the hmac signature.

Example:

```
$ drone sign octocat/hello-world --save
```

<div class="alert">
Please note that you must re-generate the signature any time the configuration file is modified.
</div>
