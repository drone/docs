+++
weight = 1
date = "2014-11-08T12:53:08-08:00"
title = "Image"

[menu.docs]
parent = "Build"
+++

Your build environment is defined by a Docker image. This means you can use any of the thousands of images in the [Docker index](https://index.docker.io) as your build environment, or even create your own.

Use the `image` key to specify the Docker image:

```yaml
image: google/dart
```

The above example will retrieve the `latest` Docker image from the index. Drone supports fully qualified image names, allowing you to specify a specific image tag or revision:

```yaml
image: google/dart:1.8.3
```

> Please note that an image must contain `git` in order to clone your repository.

## Official Language Stacks

You can use the official [language stacks](http://blog.docker.com/2014/09/docker-hub-official-repos-announcing-language-stacks/) by providing the fully qualified image name:

```yaml
image: library/java:openjdk-8
```

> Please note that most language stacks use [buildpack-deps](https://github.com/docker-library/buildpack-deps) as their base image which includes useful dependencies such as git, mercurial, curl and more.