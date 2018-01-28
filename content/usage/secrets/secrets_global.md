---
title: Global Secrets
url: global-secrets

menu:
  usage:
    weight: 3
    identifier: global-secrets
    parent: usage_secrets
---

{{% alert enterprise %}}
This feature is only available in the [Enterprise Edition](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise edition supports global secrets, sourced from a yaml file on your server. You should mount the secret file into your container and specify the path to the secret file in your configuration.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
+     - /etc/drone-secrets.yml:/etc/drone-secrets.yml
    restart: always
    environment:
+     DRONE_GLOBAL_SECRETS=/etc/drone-secrets.yml
```

Example secrets file:

```nohighlight
- name: docker_username
  value: octocat
- name: docker_password
  value: correct-horse-batter-staple
```

{{% alert info %}}
Secrets configured via global secrets take precedence over secrets configured in individual repositories.
{{% /alert %}}

# Restricting Access

Restrict access to global secrets based on repository name using the `repos` attribute. This is defined as an array list with glob support.

```
- name: docker_username
  value: octocat
  repos: [ octocat/hello-world, github/* ]
- name: docker_password
  value: correct-horse-battery-staple
  repos: [ octocat/hello-world, github/* ]
```

Restrict access to global secrets based on image name using the `images` attribute. This is defined as an array list with glob support.

```
- name: docker_username
  value: octocat
  images: [ plugins/docker, plugins/* ]
- name: docker_password
  value: correct-horse-battery-staple
  images: [ plugins/docker:latest, plugins/ecr:* ]
```

Restrict access to global secrets based on event name using the `events` attribute.

```
- name: docker_username
  value: octocat
  events: [ push, pull_request ]
- name: docker_password
  value: correct-horse-battery-staple
  events: [ push, tag ]
```

Any combination of restrictions can be combined.

```
- name: docker_username
  value: octocat
  repos: [ octocat/hello-world, github/* ]
  events: [ push, tag ]
  images: [ plugins/* ]
- name: docker_password
  value: correct-horse-battery-staple
  repos: [ octocat/hello-world, github/* ]
  images: [ plugins/docker ]
```

Currently, global secrets does not support `status` as an attribute-based usage restriction.
