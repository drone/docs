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
date = 2016-02-21T08:37:46Z
menu = ""
weight = 1

+++

Use the Terraform plugin to apply the infrastructure configuration contained within the repository. The following parameters are used to configure this plugin:

* `plan` - if true, calculates a plan but does __NOT__ apply it.
* `remote` - contains the configuration for the Terraform remote state tracking.
  * `backend` - the Terraform remote state backend to use.
  * `config` - a map of configuration parameters for the remote state backend. Each value is passed as a `-backend-config=<key>=<value>` option.
* `vars` - a map of variables to pass to the Terraform `plan` and `apply` commands. Each value is passed as a `-var
 <key>=<value>` option.
* `ca_cert` - ca cert to add to your environment to allow terraform to use internal/private resources
* `sensitive` (default: `false`) - Whether or not to suppress terraform commands to stdout.

The following is a sample Terraform configuration in your .drone.yml file:

```yaml
deploy:
  terraform:
    plan: false
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

# Advanced Configuration

## CA Certs
You may want to run terraform against internal resources, like an internal
OpenStack deployment.  Usually these resources are signed by an internal
CA Certificate.  You can inject your CA Certificate into the plugin by using
`ca_certs` key as described above.  Below is an example.

```yaml
deploy:
  terraform:
    plan: false
    remote:
      backend: swift
      config:
        path: drone/terraform
    vars:
      app_name: my-project
      app_version: 1.0.0
    ca_cert: |
      -----BEGIN CERTIFICATE-----
      asdfsadf
      asdfsadf
      -----END CERTIFICATE-----
```

## Suppress Sensitive Output
You may be passing sensitive vars to your terraform commands.  If you do not want
the terraform commands to display in your drone logs then set `sensitive` to `true`.
The output from the commands themselves will still display, it just won't show
want command is actually being ran.

```yaml
deploy:
  terraform:
    plan: false
    sensitive: true
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

