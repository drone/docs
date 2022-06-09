---
date: 2000-01-01T00:00:00+00:00
title: Google
author: eoinmcafee
weight: 100
---

<div class="alert">
The Google provider is in the Beta phase.
</div>

<div class="alert">
Linux instance support only currently
</div>

The goal of this document is to give you enough technical specifics to configure and run google instances. When properly configured, it will automatically provision and terminate Google Compute Engine instances based on your Drone server’s build volume.

# Prerequisites

## Create Google Credentials file

If your vm runner is running in a GCP instance then create `application_default_credentials.json` by executing next command in the GCP instance:
```
gcloud auth application-default login
```

And credentials will be saved in `/home/$(whoami)/.config/gcloud/application_default_credentials.json`.

Otherwise, [create service account](https://cloud.google.com/iam/docs/creating-managing-service-accounts), [generate credentials file](https://cloud.google.com/iam/docs/creating-managing-service-account-keys) and save it somewhere in the machine where plan to run the vm runner.

# Recommended Images

+ [Ubuntu 18.04 LTS (Bionic)](https://console.cloud.google.com/marketplace/product/ubuntu-os-cloud/ubuntu-bionic) 

to find images to use on google compute engine, use the following command:

```bash
gcloud compute images list
```

A valid image will look like `projects/{PROJECT}/global/images/{IMAGE}` eg `projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131`.

# Example Pool File

{{< highlight bash "linenos=table" >}}
version: "1"
instances:
  - name: ubuntu-gcp
    default: true
    type: google
    pool: 1
    limit:
    platform:
      os: linux
      arch: amd64
    spec:
      account:
        project_id: xxxxxxxxxxx
        json_path: path/to/credentials.json
      image: projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131
      machine_type: e2-small
      zone:
        - europe-west1-b
{{< / highlight >}}
