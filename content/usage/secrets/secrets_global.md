---
title: Global Secrets
url: global-secrets

menu:
  usage:
    weight: 4
    identifier: global-secrets
    parent: usage_secrets
---

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global secrets variables, sourced from a yaml file on your server. You should mount the secrets file into your container and specify the path to the file in your configuration.

# Configuration

Please see [the installation notes]({{<relref "install/extensions/global_secrets.md">}}) to configure global secrets file and its restrictions.

# Precedence

{{% alert info %}}
Secrets configured via *global secrets* take precedence over secrets configured in individual repositories.
{{% /alert %}}

# Usage

Usage notes about [secrets]({{<relref "registries.md">}}) apply for global secrets as well.
