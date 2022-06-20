---
date: 2000-01-01T00:00:00+00:00
title: Digital Ocean
author: tphoney
weight: 100
---

<div class="alert">
The Digital ocean driver is in the 'beta' phase.
</div>

# Overview

**By default it will create a pool with a max size of 2 running `Ubuntu 18.04` with `s-2vcpu-4gb` in `nyc1`. The pool is called `testpool`.**

Digital Ocean specific configuration in a pool file.

## Authentication

Create a Digital Ocean Personal Access Token

You need to create a Digital Ocean Personal Access token that you can provide to the runner. The runner will use the token to authorize API requests. For further documentaiton look [here](https://docs.digitalocean.com/reference/api/create-personal-access-token/)

## Firewall rules

By default it will create the necessary firewall. It is named "harness runner".

# Pool Spec

Digital Ocean specific configuration.

{{< highlight yaml "linenos=table" >}}
  account         Account   # explained in section below
  image           string    # the image ID of a public or private image or the slug identifier for a public image 
  size            string    # the slug identifier for the size, `https://www.digitalocean.com/community/questions/how-do-i-work-out-the-size-slug-string-for-a-given-droplet-size`
  firewall_id     string    # use a specific firewall (fingerprint)
  ssh_keys        []string  # list of keys to add (fingerprint)
  tags            []string  # tags to apply to the instance
  user_data       string    # user data to apply to the instance
  user_data_path  string    # path to user data script
{{< / highlight >}}

More information on user_data and user_data_path can be found [custom cloud-init]({{< relref "../../configuration/cloud-init.md" >}})

## Account

Contains the Digital ocean account configuration.

{{< highlight yaml "linenos=table" >}}
  pat     string  # personal access token, for more information see above
  region  string  # region `https://docs.digitalocean.com/products/platform/availability-matrix/`
{{< / highlight >}}

# Example pool setup

EG, This `pool.yml` file configures 2 pools each with a pool size of 2 and a limit of 4.

{{< highlight yaml "linenos=table" >}}
version: "1"
instances:
  - name: ubuntu
    default: true
    type: digitalocean
    pool: 1    # total number of warm instances in the pool at all times
    limit: 4   # limit the total number of running servers. If exceeded block or error.
    platform:
      os: linux
      arch: amd64
    spec:
      account:
        region: nyc2
      image: docker-18-04
      size: s-2vcpu-8gb
{{< / highlight >}}
