---
date: 2000-01-01T00:00:00+00:00
title: Azure
author: tphoney
weight: 100
---

<div class="alert">
The Azure driver is in the Beta phase.
</div>

# Overview

**By default it will create a pool with a max size of 2 running Ubuntu 18.04. The pools is called `testpool`.**

Azure specific configuration in a pool file.

## Authentication

By default we require client_id, client_secret subscription_id and tenant_id which is needed for create an the instance.

You can create an application or use an existing one [here](https://portal.azure.com/#view/Microsoft_AAD_RegisteredApps/ApplicationsListBlade). On your application `Overview` page you can view the client id. Then you can add a new secret under `certificates and secrets`.

Find your tenant_id by going to the Azure portal [here](https://portal.azure.com/#view/Microsoft_AAD_IAM/TenantPropertiesBlade).

Find your subsription_id by going to the Azure portal and selecting your subscription [here](https://portal.azure.com/#view/Microsoft_Azure_Billing/SubscriptionsBlade). You will then need to add permissions to your application under IAM. You will need to add the Owner role to your application. To do this, click on your subscription, then click the  IAM link on the left hand side, then add a new role. Select your application (not user), Assign the owner role, save.

# Pool Spec

Cloud specific configuration.

{{< highlight yaml "linenos=table" >}}
account             account           # explained in section below
resource_group      string            # resource group name
location            string            # location
vm_id               string            # vm id
root_directory      string            # root directory
user_data           string            # user data to apply to the instance
user_data_key       string            # path to user data script
user_data_path      string            # path to user data script
image               image             # explained in section below
size                string            # size of the instance
zones               []string          # zones to use
tags                map[string]string # tags to apply to the instance
security_group_name string            # security group name
{{< / highlight >}}

More information on user_data and user_data_path can be found [custom cloud-init]({{< relref "../../configuration/cloud-init.md" >}})

## account

Contains the Azure account configuration.

{{< highlight yaml "linenos=table" >}}
subscription_id string # subscription id
client_id       string # client id
client_secret   string # client secret
tenant_id       string # tenant id
{{< / highlight >}}

## image

Contains Azure base image configuration:

{{< highlight yaml "linenos=table" >}}
publisher string # publisher
offer     string # offer
sku       string # sku
version   string # version
username  string # username
password  string # password
{{< / highlight >}}

# Recommended images

## [Ubuntu 18.04](https://az-vm-image.info/?cmd=--all+--publisher+Canonical+--sku+18_04-lts)

This is the default image for the runner.

```yaml
       image:
         username: XXXXXXX
         password: XXXXXXX
         publisher: Canonical
         offer: UbuntuServer
         sku: 18.04-LTS
         version: latest
```

## [Windows Server 2019 with containers](https://az-vm-image.info/?cmd=--all+--publisher+microsoftwindowsserver+--sku+containers)

You must use virtual machines with container support enabled.

NB: be sure to set the platform to windows

  ```yaml
version: "1"
instances:
- name: ubuntu-azure
  default: true
  type: azure
  platform:
    os: windows
  spec:
    account:
      client_id: XXXXXXX
      client_secret: XXXXXXX
      subscription_id: XXXXXXX
      tenant_id: XXXXXXX
    location: eastus2
    size: Standard_F2s
    zones: ["1"]
    tags:
      tagName: tag
    image:
      username: XXXXXXX
      password: XXXXXXX
      publisher: MicrosoftWindowsServer
      offer: WindowsServer
      sku: 2019-Datacenter-with-Containers
      version: latest
```

Azure does not support running custom scripts for windows out of the box (linux has cloudinit). We use extensions to do this, this unfortunately adds to the spin up time.

# Example pool setup

EG, This `pool.yml` file configures a pool with a pool size of 1 and a limit of 4.

{{< highlight yaml "linenos=table" >}}
version: "1"
instances:

- name: ubuntu-azure
    default: true
    type: azure
    pool: 1    # total number of warm instances in the pool at all times
    limit: 4   # limit the total number of running servers. If exceeded block or error.
    platform:
      os: linux
      arch: amd64
    spec:
      account:
      client_id: XXXXXXX
      client_secret: XXXXXXX
      subscription_id: XXXXXXX
      tenant_id: XXXXXXXXX
      location: eastus2
      size : Standard_F2s
      tags:
        tagName: tag
      image:
        username: azureuser
        password: password123!
        publisher: Canonical
        offer: UbuntuServer
        sku: 18.04-LTS
        version: latest
{{< / highlight >}}
