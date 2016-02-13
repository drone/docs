+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Tutum"
description = "Deploy or update a project on Tutum"
user = "drone-plugins"
repo = "drone-tutum"
image = "plugins/drone-tutum"
tags = ["docker", "image", "container"]
categories = "deploy"
draft = false
date = 2016-02-13T08:59:40Z
menu = ""
weight = 1

+++

Use the tutum plugin to redeploy or update a service in [tutum](https://tutum.co).

The following parameters are used to configure this plugin:

- `username` - username
- `api_key` - tutum api key
- `service` - name of tutum service to act on
- `docker_image` - new image to assign to service, including tag (`drone/drone:latest`)
- `redeploy` - should the service be redeployed after updating
- `reuse_volumes` - whether or not to reuse volumes when redeploying (defaults to `true`)

The following is a sample Tutum configuration in your `.drone.yml` file:

```yaml
deploy:
  tutum:
    username: somebody
    api_key: abcdefg1234567
    service: drone
    image: drone/drone:latest
    redeploy: true
    reuse_volumes: true
```

Note that if your `service` is part of a stack, you should use the notation `servicename.stackname` as this will make sure that the found service is part of the correct stack.

