---
date: 2017-04-15T14:39:04+02:00
title: "Error: Insufficient privileges"
url: error-insufficient-privileges
---

These error messages occur when the Yaml sets restricted configuration parameters, which are disabled by default for security reasons.

```
ERROR: Insufficient privileges to use privileged mode
ERROR: Insufficient privileges to override shm_size
ERROR: Insufficient privileges to use custom dns
ERROR: Insufficient privileges to use dns_search
ERROR: Insufficient privileges to use devices
ERROR: Insufficient privileges to use extra_hosts
ERROR: Insufficient privileges to override the network
ERROR: Insufficient privileges to disable oom_kill
ERROR: Insufficient privileges to use volumes
ERROR: Insufficient privileges to use volumes_from
```

The above errors can be resolved by having an administrator grant your repository elevated privileges, by toggling the trusted flag in the repository settings.

{{% alert error %}}
Trusted mode should only be enabled in trusted, private environments. This flag could otherwise allow an attacker to compromise your host system.
{{% /alert %}}
