+++
date = "2017-04-15T14:39:04+02:00"
title = "Global Secrets File"
url = "setup-global-secrets"

[menu.install]
  parent = "install_enterprise"
  identifier = "setup-global-secrets"
  weight = 1
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise expansion pack](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise expansion pack provides support for global secrets variables, sourced from a yaml file on your server. You should mount the secrets file into your container and specify the path to the file in your configuration.

Example server configuration:

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

Both restrictions can be combined.

```
- name: docker_username
  value: octocat
  repos: [ octocat/hello-world, github/* ]
  images: [ plugins/* ]
- name: docker_password
  value: correct-horse-battery-staple
  repos: [ octocat/hello-world, github/* ]
  images: [ plugins/docker ]
```
