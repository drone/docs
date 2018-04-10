---
title: Global Registry Credentials
url: global-registries

menu:
  usage:
    weight: 2
    identifier: global-registries
    parent: usage_secrets
---

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global registry credentials, sourced from a yaml file on your server. You should mount the credentials file into your container and specify the path to the file in your configuration.

# Configuration

Please see [the installation notes]({{<relref "install/extensions/global_registries.md">}}) to configure global registries file and its restrictions.

# Precedence

{{% alert info %}}
Credentials configured via *individual repositories* take precedence over credentials configured in global registries.
{{% /alert %}}

{{% alert warn %}}
In versions < `0.8.4`, credentials configured via *global registries* take precedence over credentials configured in individual repositories.
{{% /alert %}}

# Usage

Usage notes about [registry credentials]({{<relref "registries.md">}}) apply for global registry credentials as well.
