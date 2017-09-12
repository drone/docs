+++
date = "2017-04-15T14:39:04+02:00"
title = "Custom Policies"
url = "custom-access-policies"

[menu.install]
  parent = "install_enterprise"
  identifier = "custom-access-policies"
  weight = 12
+++

{{% alert enterprise %}}
This feature is only available with the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

{{% alert warn %}}
This feature is experimental and should be considered unstable
{{% /alert %}}

The enterprise expansion pack will include support for advanced access control policies to limit access and usage.


# `DRONE_REPO_WHITELIST`

The repository whitelist is an optional parameter used to limit which repositories can be activated within the system. In order to activate a repository it must match at least one pattern in the whitelist.

```
DRONE_REPO_WHITELIST=Microsoct/*,twitter/*
```
