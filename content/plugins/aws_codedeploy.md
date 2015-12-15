+++
title = "AWS CodeDeploy"
description = "Deploys or updates a project on AWS CodeDeploy"
user = "drone-plugins"
repo = "drone-aws-codedeploy"
image = "plugins/drone-aws-codedeploy"
tags = ["aws", "codedeploy"]
categories = "deploy"
draft = false
date = 2015-12-15T06:53:09Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to AWS CodeDeploy. You can override
the default configuration with the following parameters:

* `access_key_id` - AWS access key ID
* `secret_access_key` - AWS secret access key
* `region` - AWS availability zone

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  aws_codedeploy:
    access_key_id:
    secret_access_key:
    region:
```

