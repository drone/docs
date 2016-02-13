+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "S3 Sync"
description = "Sync a directory with an Amazon S3 Bucket"
user = "drone-plugins"
repo = "drone-s3-sync"
image = "plugins/drone-s3-sync"
tags = ["publish", "artifacts", "amazon", "aws", "s3"]
categories = "publish"
draft = false
date = 2016-02-13T08:59:18Z
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
* `content_type` - override default mime-types to use this value
* `metadata` - set custom metadata
* `redirects` - targets that should redirect elsewhere

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

Both `acl` and `content_type` can be passed as a string value to apply to all files, or as a map to apply to a subset of files.

For example:

```yaml
publish:
  s3_sync:
    acl:
      "public/*": public-read
      "private/*": private
    content_type:
      ".svg": image/svg+xml
    region: "us-east-1"
    bucket: "my-bucket.s3-website-us-east-1.amazonaws.com"
    access_key: "970d28f4dd477bc184fbd10b376de753"
    secret_key: "9c5785d3ece6a9cdefa42eb99b58986f9095ff1c"
    source: folder/to/archive
    target: /target/location
    delete: true
```

In the case of `acl` the key of the map is a glob. If there are no matches in your settings for a given file, the default is `"private"`.

The `content_type` field the key is an extension including the leading dot `.`. If you want to set a content type for files with no extension, set the key to the empty string `""`. If there are no matches for the `content_type` of any file, one will automatically be determined for you.

The `metadata` field can be set as either an object where the keys are the metadata headers:

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
    metadata:
      Cache-Control: "max-age: 10000"
```

Or you can specify metadata for file patterns by using a glob:

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
    metadata:
      "*.png":
        Cache-Control: "max-age: 10000000"
      "*.html":
        Cache-Control: "max-age: 1000"
```

Additionally, you can specify redirect targets for files that don't exist by using the `redirects` key:

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
    redirects:
      some/missing/file: /somewhere/that/actually/exists
```

