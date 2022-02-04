---
date: 2000-01-01T00:00:00+00:00
title: Common Docker Problems
author: bradrydzewski
weight: 20000
hidden: true
toc: true
---

This page can be used to triage issues with docker plugins, which include the gcr and ecr plugins. When troubleshooting, please ensure you are testing with the latest release of our plugins.  _If you have forked or mirrored our plugins, please ensure you are testing with the official image instead of the fork or mirror._

# Required Data

In order to help troubleshoot issues, the engineering team requires _all_ of the below information. The engineering team will be unable to assist until _all_ of the below information is provided.

- the full yaml configuration
- the full runner configuration
- the result of `drone build info <repository> <build>` for the failing build
- the result of `drone orgsecret ls <organization>`
- the results of `drone secret ls <repository>`
- the results of `drone secret info <repository> <secret>` for relevant secrets
- the full logs, copied from the user interface, for the failing plugin step


# Problems Starting Docker

The docker plugin uses docker-in-docker to build and publish images; this means the plugin starts a docker daemon inside its own container.  If the plugin fails to start the docker daemon the plugin will fail and you will see the following entry in the pipeline step logs:

```text  {linenos=table}
level=fatal msg="Error authenticating: exit status 1"
```

If you are running the latest version of the plugin, you will see the following entry in the pipeline step logs:

```text  {linenos=table}
Unable to reach Docker Daemon after 15 attempts
```

If docker is unable to start you should enable debug logging:

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

## Security Software

There have been cases where apparmor or selinux prevent docker-in-docker from working properly. If you are experiencing issues you may want to check your linux security poligies. If you recently modified your security policies you may need to roll back.

## Docker Regression

There have been regressions with docker that prevent docker-in-docker from working properly. If you are experiencing issues you may want to ensure you have a stable, long term release of docker installed on your host machine.  If you recently upgraded Docker and are having issues with the plugin you should rollback.

## Restricted Volumes

The docker plugin does not start correctly if you try to mount volumes. If you are mounting volumes into the plugin step the volume mounts must be removed.

```yaml  {linenos=table, hl_lines=["9-11"]}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: foo/bar
    username: kevinbacon
    password: pa55word
    debug: true
  volumes:
  - name: temp
    path: /tmp
```

## Restricted Environment Variables

The docker plugin does not start correctly if you set any of the below environment variables. If you are setting any of the below variables, in your pipeline or as global environment variables, they must be removed.

```text  {linenos=table}
XDG_RUNTIME_DIR
DOCKER_OPTS
DOCKER_HOST
PATH
HOME
```

## Forks or Mirrors

The docker plugins are white-listed to automatically start with privileged mode enabled.  If you are using a fork or mirror of the image, you need to add the [image](https://docs.drone.io/runner/docker/configuration/reference/drone-runner-privileged-images) to this whitelist.


# Problems with Login

If the plugin is unable to login to the remote docker registry you will see the below entry in your pipeline logs.

```text  {linenos=table}
level=fatal msg="Error authenticating: exit status 1"
```

If you provided the Docker plugin with registry credentials, such as username and password or token, you should see one of the below entries in your pipeline step logs. If you do not see entries in your pipeline logs it means credentials were not provided as expected.

```text  {linenos=table}
Detected registry credentials
Detected registry credentials and registry credentials file
```

The below log entry _may_ also appear if credentials are not provided:

```text  {linenos=table}
Registry credentials or Docker config not provided. Guest mode enabled.
```


## Missing Username and Password

If the registry username and password are missing you will see the below entry in your pipeline logs. Please see our guide to common problems users encounter when using secrets <https://community.harness.io/t/problems-with-secrets/10662>.

```text  {linenos=table}
level=fatal msg="Error authenticating: exit status 1"
```

## Missing Registry Address

If no registry address is provided the plugin attempts to login to dockerhub.  If you are using a private registry you must include the registry address in your configuration.

```yaml  {linenos=table, hl_lines=["5"]}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: quay.io/foo/bar
    registry: quay.io
    username: kevinbacon
    password: pa55word
```


## Insecure Registries

If the registry is self-hosted and if the registry uses plain http or self-signed certificates, the docker client will be unable to establish a secure connection to the registry. This will cause the docker login to fail with the following error:

```text  {linenos=table}
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

## Invalid Username or Password

If the registry username and password are invalid you will see the below entry in your pipeline logs. _The most common root cause for this issue is copy / paste problems when creating secrets, such as newlines or special characters being accidentally copied with the secret string._

```text  {linenos=table}
Detected registry credentials
Error response from daemon: Get https://registry-1.docker.io/v2/: unauthorized: incorrect username or password
level=fatal msg="Error authenticating: exit status 1"
```

## Clashing Username or Password

Be careful when using global secrets, global environment variables, ir organization secrets in conjunction with repository secrets. If you define multiple secrets with the same name they can clash and cause unexpected behavior.


# Problems with Configuration

## Missing registry address

The docker plugin assumes you are publishing images to dockerhub.  If you are publishing an image to a third party registry you must use the fully qualified image name, which includes the registry hostname.


```yaml  {linenos=table, hl_lines=["6"]}
steps:
- name: docker  
  image: plugins/docker
  settings:
    repo: quay.io/foo/bar
    registry: quay.io
    username: kevinbacon
    password: pa55word
```

## Missing zone or region

The ecr plugin assumes the default zone is us-east-1. If the target docker registry is not located in us-east-1 you must include the zone in your plugin configuration.

```yaml  {linenos=table, hl_lines=["6"]}
steps:
- name: docker  
  image: plugins/ecr
  settings:
    repo: xxxxxxxxxxxx.dkr.ecr.us-west-1.amazonaws.com/demo
    zone: us-west-1
```
