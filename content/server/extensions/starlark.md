---
date: 2000-01-01T00:00:00+00:00
title: Starlark
author: bradrydzewski
weight: 10
separator: true
toc: false

description: |
  Install the Starlark configuration extension.
---

Drone provides an official extension that enables support for [Starlark](https://docs.bazel.build/versions/master/skylark/language.html), a configuration language inspired by Python, to build pipelines. Starlark provides a powerful alternative to traditional yaml configurations.

<!-- If your yaml configuration files are growing increasing complex, you may want to give Starlark a shot. -->

# Installation

1. Create a shared secret.

    ```
    $ openssl rand -hex 16
    bea26a2221fd8090ea38720fc445eca6
    ```

2. Download and run the extension.

    ```
    $ docker run -d \
    --publish=3000:3000 \
    --env=DRONE_DEBUG=true \
    --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
    --restart=always \
    --name=starlark drone/drone-convert-starlark
    ```

3. Update your Drone server configuration to include the extension address and the shared secret.

    ```
    DRONE_CONVERT_PLUGIN_ENDPOINT=http://1.2.3.4:3000
    DRONE_CONVERT_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
    ```

# Verification

You can verify the extension is configured and is processing requests using the command line utility.

1. Provide the command line utility with the extension endpoint and secret.

    ```
    $ export DRONE_CONVERT_ENDPOINT=http://1.2.3.4:3000
    $ export DRONE_CONVERT_SECRET=bea26a2221fd8090ea38720fc445eca6
    ```

2. Use the command line utility to convert a Starlark script:

    ```
    $ drone plugins convert path/to/.drone.star
    ```
