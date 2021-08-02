---
date: 2000-01-01T00:00:00+00:00
title: Docker
toc: true
author: bradrydzewski
description: |
  Plugin to build and publish Docker images using Docker-in-Docker.
---

The Docker plugin can be used to build and publish images to the Docker registry using Docker-in-Docker. The following pipeline configuration uses the Docker plugin to build and publish Docker images.

* Example configuration builds and publishes an image using a Dockerfile in the root of your git repository.  The docker context also defaults to the root of your git repository.

  ```yaml  {linenos=table}
  kind: pipeline
  name: default

  steps:
  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: foo/bar
      tags:
      - latest
  ```

* Example configuration using multiple tags:

  ```yaml {linenos=table, linenostart=5, hl_lines=["8-11"]}
  steps:
  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: foo/bar
      tags:
        - latest
        - '1.0.1'
        - '1.0'
  ```

* Example configuration uses a `.tags` file, a comma-separate list of tags, to dynamically configure tags. The plugin automatically loads this file if it exists.

  ```yaml {linenos=table, linenostart=5, hl_lines=[7]}
  steps:
  - name: build
    image: golang
    commands:
      - go build
      - go test
      - echo -n "5.2.6,5.2.4" > .tags

  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: foo/bar
  ```

* Example configuration using build arguments:

  ```yaml {linenos=table, linenostart=5, hl_lines=[8,9]}
  steps:
  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: foo/bar
      build_args:
        - HTTP_PROXY=http://yourproxy.com
  ```

* Example configuration passing secrets as build arguments:

  ```yaml {linenos=table, linenostart=5, hl_lines=[5,6,11,12]}
  steps:
    - name: build-image
      image: plugins/docker
      environment:
        TOKEN:
          from_secret: one_of_the_tokens
      settings:
        username: kevinbacon
        password: pa55word
        repo: foo/bar
        build_args_from_env:
          - TOKEN
  ```

* Example configuration using alternate Dockerfile:

  ```yaml {linenos=table, linenostart=4, hl_lines=[8]}
  steps:
  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: foo/bar
      dockerfile: path/to/Dockerfile
  ```

* Example configuration using a custom registry:

  ```yaml {linenos=table, linenostart=4, hl_lines=[8]}
  steps:
  - name: docker  
    image: plugins/docker
    settings:
      username: kevinbacon
      password: pa55word
      repo: index.company.com/foo/bar
      registry: index.company.com
  ```

# Secrets

It is considered best practice to source sensitive configuration parameters from secrets. The following example demonstrates how you can source the username and password from secrets:

```yaml {linenos=table, hl_lines=["8-11"]}
kind: pipeline
name: default

steps:
- name: docker  
  image: plugins/docker
  settings:
    username:
      from_secret: docker_username
    password:
      from_secret: docker_password
    repo: foo/bar
    tags:
    - latest
```


# Autotag

The Docker plugin can be configured to automatically tag your images. When this feature is enabled and the event type is tag, the plugin will automatically tag the image using the standard major, minor, release convention. For example:

- `v1.0.0` produces docker tags `1`, `1.0`, `1.0.0`
- `v1.0.0-rc.1` produces docker tags `1.0.0-rc.1`

When the event type is push and the target branch is your default branch (e.g. master) the plugin will automatically tag the image as `latest`. All other event types and branches are ignored.

Example configuration:

```yaml  {linenos=table}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: foo/bar
    auto_tag: true
    username: kevinbacon
    password: pa55word
```

Example configuration with tag suffix:

```yaml  {linenos=table}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: foo/bar
    auto_tag: true
    auto_tag_suffix: linux-amd64
    username: kevinbacon
    password: pa55word
```

Please note that auto-tagging is intentionally simple and opinionated. We are not accepting pull requests at this time to further customize the logic.

# Parameters

* `registry`
  : authenticates to this registry

* `username`
  : authenticates with this username

* `password`
  : authenticates with this password

* `repo`
  : repository name for the image

* `tags`
  : repository tag for the image

* `dockerfile`
  : dockerfile to be used, defaults to Dockerfile

* `dry_run`
  : boolean if the docker image should be pushed at the end

* `purge`
  : boolean if cleanup of the docker image should be done at the end

* `context`
  : the context path to use, defaults to root of the git repo

* `target`
  : the build target to use, must be defined in the docker file

* `force_tag=false`
  : replace existing matched image tags

* `insecure=false`
  : enable insecure communication to this registry

* `mirror`
  : use a mirror registry instead of pulling images directly from the central Hub

* `bip=false`
  : use for pass bridge ip

* `custom_dns`
  : set custom dns servers for the container

* `custom_dns_search`
  : docker daemon dns search domains

* `storage_driver`
  : supports aufs, overlay or vfs drivers

* `storage_path`
  : docker daemon storage path

* `build_args`
  : pass custom arguments to docker build

* `build_args_from_env`
  : pass the envvars as custom arguments to docker build

* `auto_tag=false`
  : generate tag names automatically based on git branch and git tag

* `auto_tag_suffix`
  : generate tag names with this suffix

* `debug, launch_debug`
  : launch the docker daemon in verbose debug mode

* `mtu`
  : docker daemon custom mtu setting

* `ipv6`
  : docker daemon IPv6 networking

* `experimental`
  : docker daemon Experimental mode

* `daemon_off`
  : don’t start the docker daemon

* `cache_from`
  : images to consider as cache sources

* `squash`
  : squash the layers at build time

* `pull_image`
  : force pull base image at build time

* `compress`
  : compress the build context using gzip

* `custom_labels`
  : additional k=v labels

* `label_schema`
  : label-schema labels

* `email`
  : docker email

* `no_cache`
  : do not use cached intermediate containers

* `add_host`
  : additional host:IP mapping

# Common Problems

## Missing Username or Password

If the registry username or password are missing you will see the below entry in your pipeline logs. _The most common root cause for this issue is incorrectly configuration secrets._

```
Registry credentials or Docker config not provided. Guest mode enabled.
```

## Invalid Username or Password

If the registry username and password are invalid you will see the below entry in your pipeline logs. _The most common root cause for this issue is copy / paste problems when creating secrets, such as newlines or special characters being accidentally copied with the secret string._

```
level=fatal msg="Error authenticating: exit status 1"
```

## Insecure Registry

If the registry is self-hosted and if the registry uses plain http or self-signed certificates, the docker client will be unable to establish a secure connection to the registry. This will cause the docker login to fail with the following error:

```
level=fatal msg="Error authenticating: exit status 1"
```

This can be solved with the following plugin configuration:

```yaml  {linenos=table, hl_lines=["8"]}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: foo/bar
    username: kevinbacon
    password: pa55word
    insecure: true
```

## Using Volumes

If you attempt to mount a volume into the plugin you will see the below entry in your pipeline logs. _The docker plugin restricts mounting volumes for security reasons._

```
level=fatal msg="Error authenticating: exit status 1"
```

This can be resolved by removing mouted volumes or by configuring the plugin step to run in privileged mode:

```yaml  {linenos=table, hl_lines=["4"]}
steps:
- name: docker  
  image: plugins/docker
  privileged: true
  settings:
    repo: foo/bar
    username: kevinbacon
    password: pa55word
  volumes:
  - name: temp
    path: /tmp
```

## Docker In Docker Issues

If a docker daemon cannot be started inside the plugin container you will see the below entry in your pipeline logs. _The most common root cause for this issue is security software (selinux, apparmor, etc) preventing nested containerization._

```
level=fatal msg="Error authenticating: exit status 1"
```

This can be further triaged with the following plugin configuration:

```yaml  {linenos=table, hl_lines=["8"]}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: foo/bar
    username: kevinbacon
    password: pa55word
    debug: true
```

If docker fails you will see the below entry in your pipeline logs:

```
failed to start daemon: Error initializing network controller: error obtaining controller instance: failed to create NAT chain DOCKER: iptables failed: iptables -t nat -N DOCKER: iptables v1.8.3 (legacy): can’t initialize iptables table `nat’: Permission denied (you must be root)
```

If docker succeeds you will see the below entry in your pipeline logs:

```
time=“2021-01-21T14:15:59.634657433Z” level=info msg=“Daemon has completed initialization”
time=“2021-01-21T14:15:59.661299094Z” level=info msg=“API listen on /var/run/docker.sock”
```

<!-- ## Insufficient Privileges

## Incorrect Registry
 -->

