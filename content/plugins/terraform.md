+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Terraform"
description = "Execute Terraform plan and apply"
user = "drone-plugins"
repo = "drone-terraform"
image = "plugins/drone-terraform"
tags = ["terraform"]
categories = "deploy"
draft = false
date = 2016-02-13T09:02:07Z
menu = ""
weight = 1

+++

Use the Terraform plugin to apply the infrastructure configuration contained within the repository. The following parameters are used to configure this plugin:

* `dry_run` - if true, calculates a plan but does __NOT__ apply it.
* `remote` - contains the configuration for the Terraform remote state tracking.
  * `backend` - the Terraform remote state backend to use.
  * `config` - a map of configuration parameters for the remote state backend. Each value is passed as a `-backend-config=<key>=<value>` option.
* `vars` - a map of variables to pass to the Terraform `plan` and `apply` commands. Each value is passed as a `-var <key>=<value>` option.

The following is a sample Terraform configuration in your .drone.yml file:

```yaml
deploy:
  terraform:
    dry_run: false
    remote:
      backend: S3
      config:
        bucket: my-terraform-config-bucket
        key: tf-states/my-project
        region: us-east-1
      vars:
        app_name: my-project
        app_version: 1.0.0
```

