---
date: 2000-01-01T00:00:00+00:00
title: Pool File
author: tphoney
weight: 2
description: |
  Configure AWS pools.
---

This sets up the build pools, it allows builds to use a hot AWS instances (you do not have to wait for an instance to spin up). For a deeper explanation of how this works, see the [design](https://github.com/drone/proposal/blob/master/design/01-aws-runner.md) documentation.

+ `.drone_pool.yml` is the default file name.
+ Each pool only has one instance type.
+ There can be multiple pools. Different pipelines can use the same pool
+ You can specify the size of the pool.
+ A pool can only be in one region.
+ Changing the pool configuration will mean removing the existing images and restarting the daemon.
+ If the pool is empty, it will trigger an adhoc instance.

#### Here are pool settings settings

A pool has:

```yaml
  name          string
  max_pool_size int
  platform      Platform
  account       Account
  instance      Instance
```

where Platform is(this is the same as plaform in other runners):

```yaml
  os      string
  arch    string
  variant string
  version string
```

where Account is (over-rides env settings):

```yaml
  access_key_id     string
  access_key_secret string
  region            string
```

where Instance contains aws instance specific ima:

```yaml
  ami              string
  tags             []string
  iam_profile_arn  string
  type             string
  disk             Disk
  network          Network
```

We recommend the following AMIs:

+ [Windows Server 2019 with containers](https://aws.amazon.com/marketplace/pp/prodview-iehgssex6veoi)
+ [Ubuntu 20.04](https://aws.amazon.com/marketplace/pp/prodview-iftkyuwv2sjxi?sr=0-2&ref_=beagle&applicationId=AWSMPContessa)

Depending on the AMI's you are using, you may need to subscribe to it. We have tested against [Ubuntu 20.04](https://aws.amazon.com/marketplace/pp/prodview-iftkyuwv2sjxi?sr=0-2&ref_=beagle&applicationId=AWSMPContessa) or [Windows 2019 with containers](https://aws.amazon.com/marketplace/pp/prodview-iehgssex6veoi?sr=0-6&ref_=beagle&applicationId=AWSMPContessa).

where Disk contains aws block information:

```yaml
  size int
  type string
  iops string
```

where Network contains aws network information:

```yaml
  vpc                 int
  vpc_security_groups []string
  security_groups     []string
  subnet_id           string
  private_ip          bool
```

EG, This `.drone_pool.yml` file configures 2 pools each with a maximum pool size of 1

```YAML
name: common
max_pool_size: 1

account:
  region: us-east-2

instance:
# ubuntu 18.04 ohio
  ami: ami-051197ce9cbb023ea
  type: t2.nano
  network:
    security_groups:
      - sg-0f5aaeb48d35162a4
```
