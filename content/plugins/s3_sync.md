+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "S3 Sync"
description = "Syncs a directory with an Amazon S3 Bucket"
user = "drone-plugins"
repo = "drone-s3-sync"
image = "plugins/drone-s3-sync"
tags = ["publish", "artifacts", "amazon", "aws", "s3"]
categories = "publish"
draft = false
date = 2016-01-19T23:01:49Z
menu = ""
weight = 1

+++

Use the S3 sync plugin to synchronize files and folders with an Amazon S3 bucket. The following parameters are used to configure this plugin:

* `access_key` - amazon key
* `secret_key` - amazon secret key
* `bucket` - bucket name
* `region` - bucket region (`us-east-1`, `eu-west-1`, etc)
* `acl` - access to files that are uploaded (`private`, `public-read`, etc)
* `source` - location of folder to sync
* `target` - target folder in your S3 bucket
* `delete` - deletes files in the target not found in the source
* `include` - don't exclude files that match the specified pattern
* `exclude` - exclude files that match the specified pattern
* `content_type` - override default mime-tpyes to use this value

The following is a sample S3 configuration in your .drone.yml file:

```yaml
publish:
  s3_sync:
    acl: public-read
    region: "us-east-1"
    bucket: "my-bucket.s3-website-us-east-1.amazonaws.com"
    access_key: "970d28f4dd477bc184fbd10b376de753"
    secret_key: "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c"
    source: folder/to/archive
    target: /target/location
    delete: true
```

