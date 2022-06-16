---
date: 2000-01-01T00:00:00+00:00
title: Installation
author: tphoney
weight: 2
steps: true
toc: true
description: |
  Install the runner.
---

<div class="alert">
The VM runner is in the Release Candidate phase.
</div>

This article explains how to install the runner on Linux. The VM runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-aws).

## Drivers

This runner provides support for multiple drivers, please see the links at the bottom of the page or the section "drivers" for their specific setup information.

By default we use the Amazon ec2 driver.

# Download

Install Docker and pull the public image:

{{< highlight bash "linenos=table" >}}
docker pull drone/drone-runner-aws
{{< / highlight >}}

# Configuration

The VM runner can be configured using environment variables. This article references the below configuration options. See [Reference]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

- __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.
- __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.
- __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.
- __DRONE_MIN_POOL_SIZE__ **(optional)**
  : provides the minimum size of the pool. The pool will not shrink below this size. The default is 1.
- __DRONE_MAX_POOL_SIZE__ **(optional)**
  : provides the maximum size of the pool. The pool will not grow above this size. The default is 2.

## Amazon EC2 prerequisites

By default the runner requires two additional Amazon specific environment variables. For more information look at the amazon specific configuration options for [Amazon]({{< relref "drivers/amazon" >}})

- __AWS_ACCESS_KEY_ID__
  : provides the access key id for your AWS account. This must have permissions to create and run EC2 instances.

- __AWS_SECRET_ACCESS_KEY__
  : provides the secret access key for your AWS account. [aws secret](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_CreateAccessKey).

- __AWS_DEFAULT_REGION__ **(optional)**
  : provides the default region for your AWS account. Defaults to us-east-2.

# Installation

{{< highlight bash "linenos=table" >}}
docker run --detach \
  --env=DRONE_RPC_PROTO=https \
  --env=DRONE_RPC_HOST=drone.company.com \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_RUNNER_CAPACITY=2 \
  --env=DRONE_RUNNER_NAME=my-first-runner \
  --env=AWS_ACCESS_KEY_ID=your-access-key \
  --env=AWS_ACCESS_KEY_SECRET=your-access-secret \
  --publish=3000:3000 \
  --restart=always \
  --name=runner \
  drone/drone-runner-aws
{{< / highlight >}}

This will start 2 Ubuntu instances in the us-east-2 region in a pool called "testpool".

# Verification

Use the docker logs command to view the logs and verify the runner successfully established a connection with the Drone server.

{{< highlight bash "linenos=table" >}}
$ docker logs runner

level=info msg="daemon: starting the server" addr=":3000"
level=info msg="daemon: successfully connected to aws"
level=info msg="daemon: successfully pinged the remote drone server"
level=debug msg="check pool" pool=ubuntu
level=info msg="daemon: polling the remote drone server" capacity=2 endpoint="http://172.21.97.69:8080" kind=pipeline type=aws
level=debug msg="poller: request stage from remote server" thread=2
level=debug msg="poller: request stage from remote server" thread=1
level=debug msg="check pool" pool="windows 2019"
level=debug msg="provision: creating instance" adhoc=false ami=ami-0840994b9b4c03cb1 pool="windows 2019"
level=debug msg="instance create" image=ami-0840994b9b4c03cb1 region=us-east-2 size=t2.medium
level=info msg="instance create success" id=i-08bb839ae0fc19524 image=ami-0840994b9b4c03cb1 region=us-east-2 size=t2.medium
level=debug msg="check instance network" image=ami-0840994b9b4c03cb1 name=i-08bb839ae0fc19524 region=us-east-2 size=t2.medium
level=debug msg="check instance network" image=ami-0840994b9b4c03cb1 name=i-08bb839ae0fc19524 region=us-east-2 size=t2.medium
level=debug msg="instance network ready" id=i-08bb839ae0fc19524 image=ami-0840994b9b4c03cb1 ip=18.119.101.233 region=us-east-2 size=t2.medium
level=info msg="provision: created the instance" adhoc=false ami=ami-0840994b9b4c03cb1 id=i-08bb839ae0fc19524 ip=18.119.101.233 pool="windows 2019" time(seconds)=61.8857165
level=debug msg="provision: Instance responding" adhoc=false ami=ami-0840994b9b4c03cb1 id=i-08bb839ae0fc19524 ip=18.119.101.233 pool="windows 2019"
level=debug msg="provision: complete" adhoc=false ami=ami-0840994b9b4c03cb1 id=i-08bb839ae0fc19524 ip=18.119.101.233 pool="windows 2019"
level=info msg="buildPools: created instance windows 2019 i-08bb839ae0fc19524 18.119.101.233"
level=info msg="daemon: pool created"
{{< / highlight >}}

If the runner is unable to create an instance in Amazon, we have a setup command to help check the AWS configuration, see [Amazon Setup]({{< relref "drivers/amazon/" >}}).

# Advanced Configuration

For more information on advanced runner configuration options, see [Reference]({{< relref "configuration/reference" >}}).

For more information on advanced pool configuration options, see [Pool]({{< relref "configuration/pool.md" >}}).

More information on using the [Amazon driver]({{< relref "drivers/amazon/" >}}).

More information on using the [Anka driver]({{< relref "drivers/anka/_index.md" >}}).

More information on using the [Digital Ocean driver]({{< relref "drivers/digitalocean/_index.md" >}}).

More information on using the [Google driver]({{< relref "drivers/google/_index.md" >}}).
