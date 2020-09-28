---
date: 2000-01-01T00:00:00+00:00
title: Example Bash Plugin
author: bradrydzewski
weight: 11
aliases:
- /0.5/custom-plugins-in-bash/
- /creating-custom-plugins-bash/
- /plugins/bash

---
This provides a brief tutorial for creating a Drone webhook plugin, using simple shell scripting, to make an http requests during the build pipeline. The below example demonstrates how we might configure a webhook plugin in the Yaml file:

```yaml  {linenos=table}
kind: pipeline
type: docker
name: default

steps:
- name: webhook
  image: acme/webhook
  settings:
    url: http://hook.acme.com
    method: post
    body: |
      hello world
```

Create a simple shell script that invokes curl using the plugin settings defined in the Yaml, which are passed to the script as environment variables in uppercase and prefixed with `PLUGIN_`.

```bash  {linenos=table}
#!/bin/sh

curl \
  -X ${PLUGIN_METHOD} \
  -d ${PLUGIN_BODY} \
  ${PLUGIN_URL}
```

Create a Dockerfile that adds your shell script to the image, and configures the image to execute your shell script as the main entrypoint.

```dockerfile  {linenos=table}
FROM alpine
ADD script.sh /bin/
RUN chmod +x /bin/script.sh
RUN apk -Uuv add curl ca-certificates
ENTRYPOINT /bin/script.sh
```

Build and publish your plugin to the Docker registry. Once published your plugin can be shared with the broader Drone community.

```
$ docker build -t acme/webhook .
$ docker push acme/webhook
```

Execute your plugin locally from the command line to verify it is working:

```
$ docker run --rm \
  -e PLUGIN_METHOD=post \
  -e PLUGIN_URL=http://hook.acme.com \
  -e PLUGIN_BODY=hello \
  acme/webhook
```