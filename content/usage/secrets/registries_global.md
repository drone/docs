---
title: Global Registries
url: global-registries

menu:
  usage:
    weight: 4
    identifier: global-registries
    parent: usage_secrets
---

{{% alert enterprise %}}
This feature is only available in the [Enterprise Edition](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise edition supports global registry credentials, sourced from a yaml file on your server. You should mount the registry credentials file into your container and specify the path to the file in your configuration.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
+     - /etc/drone-registry.yml:/etc/drone-registry.yml
    restart: always
    environment:
+     DRONE_GLOBAL_REGISTRY=/etc/drone-registry.yml
```

Example registry credentials file:

```nohighlight
- address: docker.io
  username: octocat
  password: correct-horse-batter-staple
- address: gcr.io
  username: _json_key
  password: |
    {
      "private_key_id": "...",
      "private_key": "...",
      "client_email": "...",
      "client_id": "...",
      "type": "..."
    }
```

# Restricting Access

Currently, global registry credentials do not support any attribute-based constraints (repo, images, events).

This is because registry credentials are internal-only to drone, and unlike secrets, are never exposed to the build.
