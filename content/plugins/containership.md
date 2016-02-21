+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "ContainerShip"
description = "Deploy or update a project on ContainerShip"
user = "drone-plugins"
repo = "drone-containership"
image = "plugins/drone-containership"
tags = ["containership", "paas"]
categories = "deploy"
draft = false
date = 2016-02-21T08:37:52Z
menu = ""
weight = 1

+++

Use this plugin for deploying an application to [ContainerShip](https://containership.io).
You can override the default configuration with the following parameters:

- `organization` - Your ContainerShip organization
- `api_key` - Your ContainerShip organization API Key
- `cluster_id` - ID of your ContainerShip cluster
- `application` - Name of the application, defaults to repo name
- `image` - Docker image to use, including tag (`MyOrg/MyImage:latest`)

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  containership:
    organization: MyContainerShipOrganization
    api_key: abcdef1234567890
    cluster_id: abcdef1234567890
    application: my-app-name
    image: MyOrg/MyImage:latest
```

