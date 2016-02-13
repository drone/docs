+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Elastic Beanstalk"
description = "Deploy or update a project on AWS Elastic Beanstalk"
user = "drone-plugins"
repo = "drone-elastic-beanstalk"
image = "plugins/drone-elastic-beanstalk"
tags = ["aws", "paas"]
categories = "deploy"
draft = false
date = 2016-02-13T09:00:25Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to AWS Elastic Beanstalk. You can
override the default configuration with the following parameters:

* `access_key` - AWS access key ID
* `secret_key` - AWS secret access key
* `region` - AWS availability zone
* `version_label` - A label identifying this version
* `application` - Application name, defaults to repo name
* `description` - A description about the deployment, optional
* `auto_create` - Automatically create the application, defaults to `false`
* `process` - Preprocess and validate the manifest, defaults to `false`
* `bucket_name` - Bucket for `S3` source bundle
* `bucket_key` - Key for `S3` source bundle
* `environment_update` - Flag whether to update ElasticBeansTalk environment with the new version
* `environment_name` - Environment Name (optional), if update_environment true

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  elastic_beanstalk:
    access_key: 970d28f4dd477bc184fbd10b376de753
    secret_key: 9c5785d3ece6a9cdefa42eb99b58986f9095ff1c
    region: us-east-1
    version_label: v1
    description: Deployed with DroneCI
    auto_create: true
    bucket_name: my-bucket-name
    bucket_key: 970d28f4dd477bc184fbd10b376de753
```

