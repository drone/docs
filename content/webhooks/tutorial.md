---
date: 2000-01-01T00:00:00+00:00
title: Create a Webhook Receiver
title_in_sidebar: Starter Project
author: bradrydzewski
steps: true
weight: 10
separator: true
---

The goal of this tutorial is to build a custom webhook receiver using our starter project template as the base. A webhook receiver is a small http server that receives and processes webhooks.

This tutorial assumes you have a working knowledge of the Go programming language. You will also need to install [boilr](https://github.com/tmrts/boilr), a scaffolding system used to generate new code projects from templates.

# Create the Project

Drone provides scaffolding for creating a new webhook receiver from a project template. We will bootstrap our project using this template.

1. Install the webhook template.
   ```
   $ boilr template download drone/boilr-webhook drone-webhook
   ```

2. Create a new webhook project. 
   ```
   $ boilr template use drone-webhook my-webhook-receiver
   ```

# Customize the Code

The generated project creates a simple server capable of receiving, authenticating and parsing requests. You only need to modify the `plugin.Deliver` function to include your custom processing logic.

```Go  {linenos=table}
package plugin

import (
	"context"

	"github.com/drone/drone-go/plugin/webhook"
)

// New returns a new webhook extension.
func New() webhook.Plugin {
	return &plugin{}
}

type plugin struct {
}

func (p *plugin) Deliver(ctx context.Context, req *webhook.Request) error {
	// insert your custom logic here
	return nil
}
```

# Build Your Docker Image

The generated project includes a `Dockefile` that you can use to build and publish images. _You will need to update the Dockerfile entrypoint._

1. Build your binary file:
   ```
   $ GOOS=linux CGO_ENABLED=0 go build
   ```

2. Build and publish your Docker image:
   ```
   $ docker build -t my-webhook-receiver .
   $ docker push my-webhook-receiver
   ```

# Start the Webhook Receiver

1. Create a shared secret used to sign the http request:
   ```
   $ openssl rand -hex 16
   bea26a2221fd8090ea38720fc445eca6
   ```

2. Start the webhook receiver.
   ```
   $ docker run \
   -p 3000:3000 \
   -e DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
   my-webhook-receiver
   ```


# Enable Webhooks

Once the webhook receiver is running you will need to provide your Drone server with the webhook endpoint and secret:

```
DRONE_WEBHOOK_ENDPOINT=http://...
DRONE_WEBHOOK_SECRET=bea26a2221fd8090ea38720fc445eca6
```