+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "OpsWorks"
description = "Deploy or update a project on AWS OpsWorks"
user = "drone-plugins"
repo = "drone-opsworks"
image = "plugins/drone-opsworks"
tags = ["aws", "opsworks"]
categories = "deploy"
draft = false
date = 2016-02-13T09:01:09Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to OpsWorks. You can override the
default configuration with the following parameters:

* `access_key` - AWS access key ID
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `stack_id` - The ID of the stack to deploy
* `app_id` - The ID of the application to deploy
* `command` - The deployment command
* `arguments` - A nested list of command arguments
* `comment` - A comment for the deployment, optional
* `custom_json` - It is used to override the default configuration
* `instances` - A list of instance IDs for the deploy targets

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  opsworks:
    access_key: 970d28f4dd477bc184fbd10b376de753
    secret_key: 9c5785d3ece6a9cdefa42eb99b58986f9095ff1c
    region: us-east-1
    stack_id: my-stack
    app_id: my-app
    command: deploy
    arguments:
      arg_name1:
        - value1
        - value2
      arg_name2:
        - value1
        - value2
    instances:
      - instance1
      - instance2
      - instance3
```

