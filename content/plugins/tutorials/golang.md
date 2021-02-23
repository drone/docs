---
date: 2000-01-01T00:00:00+00:00
title: Example Go Plugin
author: bradrydzewski
weight: 10
aliases:
- /creating-custom-plugins-golang/
- /plugins/golang
---
This guide provides a brief tutorial for creating a webhook plugin, using the Go programming language, to trigger http requests during the build pipeline. The below example demonstrates how we might configure a webhook plugin in the Yaml file:

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

Create a Go program that makes an http request using the Yaml configuration parameters, which are passed to the program as environment variables in uppercase, prefixed with PLUGIN_. Here is more information on [plugin inputs](https://docs.drone.io/plugins/overview/#plugin-inputs).

```go  {linenos=table}
package main

import (
  "net/http"
  "os"
)

func main() {
  body := strings.NewReader(
    os.GetEnv("PLUGIN_BODY"),
  )

  req, err := http.NewRequest(
    os.GetEnv("PLUGIN_METHOD"),
    os.GetEnv("PLUGIN_URL"),
    body,
  )
  if err != nil {
    os.Exit(1)
  }
}
```

Compile your binary on the host machine for the target platform. Compiling on the host machine and adding the binary to the image is considered a best practice because it reduces the overall image size.

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o webhook
```

Create a Dockerfile that adds your compiled binary to the image, and configures the image to run your binary as the main entrypoint.

```dockerfile  {linenos=table}
FROM alpine
ADD webhook /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/webhook
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

# Plugin templates

To make it easier to create plugins we use templates to make them. We use the tool [Boilr](https://raw.githubusercontent.com/tmrts/boilr) to make the plugin. 

The default plugin template is [here](https://github.com/drone/boilr-plugin) and there are more complex community versions available too.
