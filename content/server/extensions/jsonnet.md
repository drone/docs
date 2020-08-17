---
date: 2000-01-01T00:00:00+00:00
title: Jsonnet
author: bradrydzewski
weight: 30
separator: true
draft: true
---


Drone provides an official extension that enables support for Jsonnet scripting, a data templating language, to configure pipelines. Jsonnet can be used as an alternative to yaml.

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
    --env=PLUGIN_DEBUG=true \
    --env=PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6 \
    --restart=always \
    --name=jsonnet drone/drone-jsonnet
    ```

3. Update your Drone server configuration to include the extension address and the shared secret.

    ```
    DRONE_YAML_ENDPOINT=http://1.2.3.4:3000
    DRONE_YAML_SECRET=bea26a2221fd8090ea38720fc445eca6
    ```
