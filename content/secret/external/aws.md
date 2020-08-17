---
date: 2000-01-01T00:00:00+00:00
title: AWS Secrets
author: bradrydzewski
weight: 70
toc: true
aliases:
- /configure/secrets/external/aws/
description: |
  Retrieve sensitive information from AWS Secrets Manager.
---

The AWS Secrets Manager secures, stores, and controls access to tokens, passwords, certificates, and other secrets in modern computing. The AWS Secrets Manager extension provides your pipeline with access to AWS secrets.

<div class="alert alert-no-cloud">
Please note this feature is disabled on Drone Cloud. This feature is only available when self-hosting.
</div>

<div class="alert alert-info">
AWS Secrets integration is provided by an extension and is only available if you system administrator has installed the extension.
</div>

# Creating Secrets

Create a secret from the AWS console. In the below example we store the Docker username and password.

![AWS Secrets](/screenshots/aws_secrets_manager_add_1.png)
![AWS Secrets](/screenshots/aws_secrets_manager_add_2.png)


# Accessing Secrets

Once our secrets are stored in AWS, we can update our yaml configuration file to request access to our secrets. First we define a secret resource in our yaml for each external secret. We include the path to the secret, and the name or key of value we want to retrieve:

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
  path: prod/docker
  name: username

---
kind: secret
name: password
get:
  path: prod/docker
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
  path: prod/docker
  name: username

---
kind: secret
name: password
get:
  path: prod/docker
  name: password

...
{{< / highlight >}}
