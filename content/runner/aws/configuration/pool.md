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
+ You can specify the minimum size of the pool.
+ A pool can only be in one region.
+ Changing the pool configuration will mean removing the existing images and restarting the daemon.
+ If the pool is empty, it will trigger an adhoc instance.
+ For Microsoft windows pools it is important to set platform. As seen in the windows example below.

#### Here are pool settings settings

A pool has:

```yaml
  name          string
  min_pool_size int
  platform      Platform
  account       Account
  instance      Instance
```

where Platform is (this is the same as plaform in other runners) NB windows support is implemented:

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

where Instance contains AWS instance specific ima:

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

where Disk contains AWS block information:

```yaml
  size int
  type string
  iops string
```

where Network contains AWS network information:

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
min_pool_size: 1

account:
  region: us-east-2

instance:
# ubuntu 18.04 ohio
  ami: ami-051197ce9cbb023ea
  type: t2.nano
  network:
    security_groups:
      - sg-0b8f8f8f8f8f8f8f8

---

name: windows 2019
min_pool_size: 1

platform:
 os: windows

account:
  region: us-east-2
  access_key_id: <super-secret>
  access_key_secret: <super-secret>

instance:
  ami: ami-0840994b9b4c03cb1
  type: t2.medium
  network:
    security_groups:
      - sg-00000000
  tags:
    cat: dog
```
