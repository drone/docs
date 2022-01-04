---
date: 2000-01-01T00:00:00+00:00
title: C#
title_in_header: Example C# Pipeline
author: LGinC
weight: 1
toc: false
---

# C# Example

This guide covers configuring continuous integration pipelines for C# projects. If you're new to Drone please read our Tutorial and build configuration guides first.

# Build and Test

In the below example we demonstrate a pipeline that executes `dotnet publish` and `dotnet test` commands. These commands are executed inside a Docker container, downloaded at runtime from DockerHub.

```
kind: pipeline
name: default

steps:
- name: test
  image: mcr.microsoft.com/dotnet/sdk:6.0
  commands:
  - dotnet build
  - dotnet test
```

Please note that you can use any Docker image in your pipeline from any Docker registry. You can use the official [dotnet sdk](https://hub.docker.com/_/microsoft-dotnet-sdk/) or [aspnet](https://hub.docker.com/_/microsoft-dotnet-aspnet/) images, or your can bring your own.
