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

**By default the runner does not need a pool file. It can use environment variables to set up a simple pool in memory. However if you want a more complex configuration or multiple pools follow the below quide.**

The pool file sets up the build pools that instantiates hot instances (builds do not wait for an instance to spin up). You can have multiple pools, each with a different configuration or that even use different cloud providers.

+ `pool.yml` is the default file name.
+ Each pool only has one instance type.
+ There can be multiple pools. Different pipelines can use the same pool
+ You can specify the minimum size of the pool.
+ You can specify the maximum size of the pool.
+ A pool can only be in one region.
+ Changing the pool configuration will mean removing the existing images and restarting the daemon.
+ If the pool is empty, it will trigger an adhoc instance.
+ For **Microsoft Windows** pools it is important to set platform. As seen in the windows example below.

# Pool file sections

## Root section

{{< highlight yaml "linenos=table" >}}
  name          string   # name of the pool, used by pipelines to select the pool
  default       bool     # default pool
  type          string   # amazon, google
  pool          int      # total number of warm instances in the pool at all times
  limit         int      # limit the total number of running servers. If exceeded block or error.
  platform      Platform # explained in section below
  spec          Spec     # explained in section below
{{< / highlight >}}

## Platform

is (this is the same as plaform in other runners) NB windows support is implemented:

{{< highlight yaml "linenos=table" >}}
  os      string
  arch    string
  variant string
  version string
{{< / highlight >}}

## Spec

This is where we confiure the cloud provider specific configuration. There are a number of different providers.

+ [Amazon]({{< relref "amazon.md" >}}) for examples and more information.
+ Google

## Using a pool file

Below is an example of using a pool file with the docker command. We can use a config folder that contains the necessary configuration files.

{{< highlight bash "linenos=table" >}}
ls  /path/on/host/config/
.drone_pool.yml
.env
{{< / highlight >}}

The below command creates a container and starts the runner.

{{< highlight bash "linenos=table" >}}
docker run -d \
  --volume=/path/on/host/config:/config/ \
  --publish=3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-aws /config/.env /config/pool.yml
{{< / highlight >}}
