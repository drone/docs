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

# Restricting Access

You can restrict access to global secrets based on repository name using the `repos` attribute. This is defined as an array list with glob support.

```
- name: docker_username
  value: octocat
  repos: [ octocat/hello-world, github/* ]
- name: docker_password
  value: correct-horse-battery-staple
  repos: [ octocat/hello-world, github/* ]
```
