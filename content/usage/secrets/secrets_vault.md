---
title: Vault Secrets
url: vault-secrets

menu:
  usage:
    weight: 5
    identifier: vault-secrets
    parent: usage_secrets
---

{{% alert enterprise %}}
This feature is only available in the [Enterprise Edition](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise edition supports loading secrets from [vault](https://www.vaultproject.io/). Note that vault credentials must be globally configured by an administrator. Vault secrets are declared in your yaml configuration file and loaded at runtime.

# Usage

Example configuration with vault secrets, passing them to the pipeline by name:

```diff
pipeline:
  build:
    image: golang
    commands:
      - go test
      - go build
  publish:
    image: plugins/docker
    repo: octocat/app
+   secrets: [ docker_username, docker_password ]

+secrets:
+  docker_username:
+    path: secret/docker_username
+  docker_password:
+    path: secret/docker_password
```

Add secrets to any valid path in vault using the vault command line utility:

```nohighlight
vault write secret/docker_username value=...
vault write secret/docker_password value=...
```

Then include the secrets paths in your pipeline configuration:

```diff
secrets:
  docker_username:
    path: secret/docker_username
  docker_password:
    path: secret/docker_password
```

# Alternate Names

In some cases the secret names in your vault instance may not match the names expected by the secrets. The secret names can be mapped to the correct values:

```diff
pipeline:
  build:
    image: golang
    commands:
      - go test
      - go build
  publish:
    image: plugins/docker
    repo: octocat/app
    secrets:
+     - source: username
+       target: docker_username
+     - source: password
+       target: docker_password

secrets:
- docker_username:
+ username:
    path: secret/docker_username
- docker_password:
+ password:
    path: secret/docker_password
```

# Restricting Repos

You can restrict access to vault secrets based on repository name using the `repo` attribute. This is a comma-separated list with glob support.

```nohighlight
vault write secret/password value=<value> repo=octocat/spoon-knife,octocat/hello-world
vault write secret/password value=<value> repo=octocat/hello-world
vault write secret/password value=<value> repo=octocat/*
```

# Restricting Events

You can restrict access to vault secrets based on hook event using the `event` attribute. This may be a string or comma-separated list:

```nohighlight
vault write secret/password value=<value> event=push
vault write secret/password value=<value> event=push,tag
vault write secret/password value=<value> event=push,pull_request
```

# Restricting Images

You can restrict access to vault secrets to specific docker images using the `image` attribute. This may be a string or comma-separated list:

```nohighlight
vault write secret/password value=<value> image=plugins/docker
vault write secret/password value=<value> image=plugins/ecr,plugins/s3
```
