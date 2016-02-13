+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "CodeDeploy"
description = "Deploy or update a project on AWS CodeDeploy"
user = "drone-plugins"
repo = "drone-codedeploy"
image = "plugins/drone-codedeploy"
tags = ["aws", "codedeploy"]
categories = "deploy"
draft = false
date = 2016-02-13T09:01:13Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to CodeDeploy. You can override the
default configuration with the following parameters:

* `access_key` - AWS access key ID
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `application` - Application name, defaults to repo name
* `deployment_group` - Name of the deployment group
* `deployment_config` - Name of the deployment config, optional
* `description` - A description about the deployment, optional
* `ignore_stop_failures` - Causes the ApplicationStop deployment lifecycle
  event to fail to a specific instance, defaults to `false`
* `revision_type` - Revision type, defaults to `GitHub`, can be set to `S3`
* `bundle_type` - File type of the application for `S3` revision type
* `bucket_name` - Bucket for `S3` revision type
* `bucket_key` - Key for `S3` revision type
* `bucket_etag` - ETag for `S3` revision type, optional
* `bucket_version` - Version for `S3` revision type, optional

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  codedeploy:
    access_key: 970d28f4dd477bc184fbd10b376de753
    secret_key: 9c5785d3ece6a9cdefa42eb99b58986f9095ff1c
    region: us-east-1
    deployment_group: my-deployment
    ignore_stop_failures: true
```

