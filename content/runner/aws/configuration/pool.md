---
date: 2000-01-01T00:00:00+00:00
title: Pool File
author: tphoney
weight: 2
toc: true
description: |
  Configure AWS pools.
---
# Overview

The pool file sets up the build pools, it allows builds to use a hot AWS instances (you do not have to wait for an instance to spin up). For a deeper explanation of how this works, see the [design](https://github.com/drone/proposal/blob/master/design/01-aws-runner.md) documentation.
If you are moving from an older version eg RC4, some change to setup may be required, they are outlined [here]({< relref "version migration" >}}).

+ `pool.yml` is the default file name.
+ Each pool only has one instance type.
+ There can be multiple pools. Different pipelines can use the same pool
+ You can specify the minimum size of the pool.
+ You can specify the maximum size of the pool.
+ A pool can only be in one region.
+ Changing the pool configuration will mean removing the existing images and restarting the daemon.
+ If the pool is empty, it will trigger an adhoc instance.
+ For **Microsoft Windows** pools it is important to set platform. As seen in the windows example below.

# Recommended AMIs

+ [Windows Server 2019 with containers](https://aws.amazon.com/marketplace/pp/prodview-iehgssex6veoi)
+ [Ubuntu 20.04](https://aws.amazon.com/marketplace/pp/prodview-iftkyuwv2sjxi?sr=0-2&ref_=beagle&applicationId=AWSMPContessa)

Depending on the AMI's you are using, you may need to subscribe to it. We have tested against [Ubuntu 20.04](https://aws.amazon.com/marketplace/pp/prodview-iftkyuwv2sjxi?sr=0-2&ref_=beagle&applicationId=AWSMPContessa) or [Windows 2019 with containers](https://aws.amazon.com/marketplace/pp/prodview-iehgssex6veoi?sr=0-6&ref_=beagle&applicationId=AWSMPContessa).

# Example pool setup

EG, This `pool.yml` file configures 2 pools each with a pool size of 2 and a limit of 4.

```YAML
version: "1"
instances:
  - name: ubuntu-aws
    default: true
    type: amazon
    pool: 1    # total number of warm instances in the pool at all times
    limit: 4   # limit the total number of running servers. If exceeded block or error.
    platform:
      os: linux
      arch: amd64
    spec:
      account:
        region: us-east-2
        availability_zone: us-east-2c
        access_key_id: XXXXXXXXXXXXXXXXXXXXX
        access_key_secret: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
      ami: ami-051197ce9cbb023ea
      size: t2.nano
      network:
        security_groups:
          - XXXXXXXXXXXXXXXX
```

# Pool file sections

## Root section

```yaml
  name          string   # name of the pool, used by pipelines to select the pool
  default       bool     # default pool
  type          string   # amazon, google
  pool          int      # total number of warm instances in the pool at all times
  limit         int      # limit the total number of running servers. If exceeded block or error.
  platform      Platform # explained in section below
  spec          Spec     # explained in section below
```

## Platform

is (this is the same as plaform in other runners) NB windows support is implemented:

```yaml
  os      string
  arch    string
  variant string
  version string
```

## Spec

Cloud specific configuration.

```yaml
  account          Account   # explained in section below
  ami              string    # ami id
  size             string    # t2.nano, m4.large, etc
  tags             []string  # tags to apply to the instance
  iam_profile_arn  string    # arn of the iam profile to apply to the instance
  disk             Disk      # explained in section below
  network          Network   # explained in section below
```

### Account

Contains the AWS account configuration.

```yaml
  access_key_id     string   # access key id
  access_key_secret string   # access key secret
  region            string   # aws region
```

### Disk

contains AWS block information:

```yaml
  size int      # size in GB
  type string   # gp2, io1, standard
  iops string   # iops for io1
```

### Network

contains AWS network information:

```yaml
  vpc                 int       # vpc id
  vpc_security_groups []string  # vpc security groups
  security_groups     []string  # security groups
  subnet_id           string    # subnet id
  private_ip          bool      # assign private ip
```
