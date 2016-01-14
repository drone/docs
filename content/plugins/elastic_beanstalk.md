+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Elastic Beanstalk"
description = "Deploys or updates a project on AWS Elastic Beanstalk"
user = "drone-plugins"
repo = "drone-elastic-beanstalk"
image = "plugins/drone-elastic-beanstalk"
tags = ["aws", "paas"]
categories = "deploy"
draft = false
date = 2016-01-14T18:46:05Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to AWS Elastic Beanstalk. You can
override the default configuration with the following parameters:

* `access_key_id` - AWS access key ID
* `secret_access_key` - AWS secret access key
* `region` - AWS availability zone

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  elastic_beanstalk:
    access_key_id:
    secret_access_key:
    region:
```

