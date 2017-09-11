+++
date = "2017-04-15T14:39:04+02:00"
title = "Integrate Vault"
url = "setup-vault-integration"

[menu.install]
  parent = "install_enterprise"
  identifier = "setup-vault-integration"
  weight = 10
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

{{% alert warn %}}
This feature is experimental and should be considered unstable
{{% /alert %}}

The enterprise expansion pack provides support for global environment variables, sourced from a vault server. Update your drone server configuration to include the vault connection details, including the vault server address and token.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
+     VAULT_ADDR=https://vault.mycompany.com:8200
+     VAULT_TOKEN=adb0c2b6-246f-10b3-bd9e-c23a0982043e
```

Please note that you can use any of the following vault configuration parameters:
https://www.vaultproject.io/docs/commands/environment.html

# Token

Please note that you must provide drone with the root token at this time. We are actively working to provide support for periodic tokens and automatic token renewal in a future release.

<!--
# Periodic Tokens

# Root Tokens

# Google IAM

# Amazon IAM
-->

# Usage

Please see the vault usage [documentation]({{< ref "usage/secrets/secrets_vault.md" >}}) which include details usage instructions for creating and consuming vault secrets in your pipelines.
