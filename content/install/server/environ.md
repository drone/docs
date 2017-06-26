+++
date = "2017-04-15T14:39:04+02:00"
title = "Global Environment"
url = "configure-global-environment"

[menu.install]
  identifier = "configure-global-environment"
  parent = "install_server"
  weight = 7
+++

{{% alert enterprise %}}
This feature is only available in the [Enterprise Edition](https://drone.io/enterprise/)
{{% /alert %}}

The enterprise edition supports global environment variables, sourced from a yaml file on your server. You should mount the environment file into your container and specify the path to the file in your configuration.

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
    volumes:
      - /var/lib/drone:/var/lib/drone/
+     - /etc/drone-env.yml:/etc/drone-env.yml
    restart: always
    environment:
+     DRONE_GLOBAL_ENVIRON=/etc/drone-env.yml
```

Example environment file:

```nohighlight
- name: GITHUB_USERNAME
  value: octocat
- name: GITHUB_PASSWORD
  value: correct-horse-batter-staple
```

# Restricting Access

You can restrict access to global environment variables based on repository name using the `repos` attribute. This is defined as an array list with glob support.

```
- name: GITHUB_USERNAME
  value: octocat
  repos: [ octocat/hello-world, github/* ]
- name: GITHUB_PASSWORD
  value: correct-horse-batter-staple
  repos: [ octocat/hello-world, github/* ]
```
