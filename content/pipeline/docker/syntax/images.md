---
date: 2000-01-01T00:00:00+00:00
title: Images
author: bradrydzewski
weight: 80
toc: true
aliases:
- /manage-registry-credentials/
- /images/
description: |
  Configure pipeline images.
---

Pipeline steps are defined as a series of Docker containers. Each step must therefore define the Docker image used to create the container.

```yaml {linenos=table, hl_lines=["7"]}
kind: pipeline
type: docker
name: default

steps:
- name: build
  image: golang:1.12
  commands:
  - go build
  - go test
```

Drone supports any valid Docker image from any Docker registry:

```
image: golang
image: golang:1.7
image: library/golang:1.7
image: index.docker.io/library/golang
image: index.docker.io/library/golang:1.7
image: docker.company.com/golang
```

# Pulling Images

If the image does not exist in the local cache, Drone instructs Docker to pull the image automatically. You will never need to manually pull images.

If the image is tagged with :latest either explicitly or implicitly, Drone attempts to pull the newest version of the image from the remote registry, even if the image exists in the local cache.


To only pull the image if not found in the local cache:

```yaml {linenos=table, linenostart=15, hl_lines=["3"]}
steps:
- name: build
  pull: if-not-exists
  image: golang
```

To always pull the newest version of the image:

```yaml {linenos=table, linenostart=15, hl_lines=["3"]}
steps:
- name: build
  pull: always
  image: golang:1.12
```

To never pull the image and always use the image in the local cache:

```yaml {linenos=table, linenostart=15, hl_lines=["3"]}
steps:
- name: build
  pull: never
  image: golang:1.12
```

# Pulling Private Images

If the image is private you need to provide Drone with docker credentials, sourced from a secret. You can manage secrets in your repository settings screen in the web interface.

First create a secret that includes your Docker credentials in the format of a Docker config.json file. This file may provide credentials for multiple registries.

```yaml {linenos=table}
{
    "auths": {
        "docker.io": {
            "auth": "4452D71687B6BC2C9389C3..."
        }
    }
}
```

Next, define which secrets should be used to pull private images using the `image_pull_secrets` attribute:

```yaml {linenos=table, linenostart=5, hl_lines=["8-9"]}
steps:
- name: build
  image: registry.internal.company.com/golang:1.12
  commands:
  - go build
  - go test

image_pull_secrets:
- dockerconfig
```

<div class="alert alert-info">
If you want to pull private images from Amazon Elastic Container Registry (ECR) you will need to install a registry credential plugin.
</div>


## Google Container Registry Issues

If the config.json has entries for both `gcr.io` and `https://gcr.io` you should prune the file and remove the entry with the `https://` prefix otherwise you will intermittently receive the following error message:

> default: Error response from daemon: unauthorized: You don't have the needed permissions to perform this operation, and you may have invalid credentials. To authenticate your request, follow the steps in: https://cloud.google.com/container-registry/docs/advanced-authentication

# Image Caching Behavior

The Docker daemon caches all images that is pulls, including private images. Docker does not restrict the use of cached images. An image already in the local cache can be used by any pipeline.

It is therefore possible for a pipeline to pull a private image that is cached by Docker, and for another pipeline to use this image from the cache without having credentials. Keep this in mind when pulling private images in a shared or public environment.
