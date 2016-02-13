+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "ECS"
description = "Deploy or update a project on AWS ECS"
user = "drone-plugins"
repo = "drone-ecs"
image = "plugins/drone-ecs"
tags = ["aws", "paas"]
categories = "deploy"
draft = false
date = 2016-02-13T09:01:44Z
menu = ""
weight = 1

+++

Use this plugin for deploying a docker container application to AWS EC2 Container Service (ECS).  

### Settings

* `access_key` - AWS access key ID, MUST be an IAM user with the AmazonEC2ContainerServiceFullAccess policy attached
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `service` - Name of the service in the cluster, **MUST** be created already in ECS
* `family` - Family name of the task definition to create or update with a new revision
* `image_name`, Container image to use, do not include the tag here
* `image_tag` - Tag of the image to use, defaults to latest
* `port_mappings` - Port mappings from host to container, format is `hostPort containerPort`, protocol is automatically set to TransportProtocol
* `memory`, Amount of memory to assign to the container, defaults to 128
* `environment_variables` - List of Environment Variables to be passed to the container, format is `NAME=VALUE`

## Example

```yaml
deploy:
  ecs:
    image: plugins/drone-ecs

    region: eu-west-1
    access_key: $$ACCESS_KEY_ID
    secret_key: $$SECRET_ACCESS_KEY
    family: my-ecs-task
    image_name: namespace/repo
    image_tag: latest
    service: my-ecs-service
    environment_variables:
      - DATABASE_URI=$$MY_DATABASE_URI
    port_mappings:
      - 80 9000
    memory: 128
```

