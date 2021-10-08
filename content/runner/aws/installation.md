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
The AWS runner is in Alpha and may not be suitable for production workloads. Furthermore this runner is a community effort and is not subject to support services or service level agreements at this time.
</div>

This article explains how to install the AWS runner on Linux. The AWS runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-aws).

# Download

Install Docker and pull the public image:

```
docker pull drone/drone-runner-aws
```

NB It is recommended to use a tagged version of the image.

# Drone specific Configuration

The AWS runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

- __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.
- __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.
- __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

# AWS specific Configuration

## AWS EC2 prerequisites

There are some pieces of setup that need to be performed on the AWS side first.

- Set up an access key and access secret [aws secret](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_CreateAccessKey) which is needed for configuration of the runner.
- Setup up vpc firewall rules for the build instances [ec2 authorizing-access-to-an-instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/authorizing-access-to-an-instance.html) We need allow ingress and egress access to port 22. Once complete you will have a security group id, which is needed for configuration of the runner.
- (windows only instance) You need to add the `AdministratorAccess` policy to the IAM role associated with the access key and access secret [IAM](https://console.aws.amazon.com/iamv2/home#/users). You will use the instance profile arn `iam_profile_arn`, in your pipeline.

## AWS EC2 environment variables

Set up the runner by using either an environment variables or a `.env` file similar to other Drone runners. Below is a list of the AWS specific environment variables.

- __DRONE_SETTINGS_AWS_ACCESS_KEY_ID__
  : AWS access key id, obtained above.
- __DRONE_SETTINGS_AWS_ACCESS_KEY_SECRET__
  : AWS access key secret, obtained above.
- __DRONE_SETTINGS_AWS_REGION__
  : AWS region
- __DRONE_SETTINGS_PRIVATE_KEY_FILE__
  : Private key file for the EC2 instances. You can generate a public and private key using [ssh-keygen](https://ssh.com/ssh/keygen)
- __DRONE_SETTINGS_PUBLIC_KEY_FILE__
  : Public key file for the EC2 instances
- __DRONE_SETTINGS_REUSE_POOL__
  : Reuse existing ec2 instances on restart of the runner.

## Example AWS Runner configuration `.env` file

```
DRONE_RPC_HOST=localhost:8080
DRONE_RPC_PROTO=http
DRONE_RPC_SECRET=bea26a2221fd8090ea38720fc445eca6
DRONE_SETTINGS_AWS_ACCESS_KEY_ID=XXXXXX
DRONE_SETTINGS_AWS_ACCESS_KEY_SECRET=XXXXX
DRONE_SETTINGS_AWS_REGION=us-east-2
DRONE_SETTINGS_PRIVATE_KEY_FILE=/config/private.key
DRONE_SETTINGS_PUBLIC_KEY_FILE=/config/public.key
```

## Pool File

The AWS runner requires a pool file, this describes the number and type of AWS instances to create in a hot swappable pool. For example:

{{< highlight yaml "linenos=table,hl_lines=5-7" >}}
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
      - sg-0b8f8f8f8f8f8f8f8
{{< / highlight >}}

See [Pool file]({{< relref "configuration/pool.md" >}}) for more detailed information.

# Installation

We can use a config folder that contains the necessary configuration files.

```
ls  /path/on/host/config/
.drone_pool.yml
.env
private.key
public.key
```

The below command creates a container and starts the runner.

```
docker run -d \
  --volume=/path/on/host/config:/config/ \
  --publish=3000:3000 \
  --restart always \
  --name runner \
  drone/drone-runner-aws /config/.env /config/.drone_pool.yml
```

# Verification

Use the docker logs command to view the logs and verify the runner successfully established a connection with the Drone server.

```
docker logs runner

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
```

# AWS pipeline syntax

For information on configuring an AWS pipeline see [AWS pipeline syntax]({{< relref "../../pipeline/aws/overview.md" >}})
