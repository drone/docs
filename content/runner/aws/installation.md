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
The AWS runner is in BETA and may not be suitable for production workloads. Furthermore this runner is a community effort and is not subject to support services or service level agreements at this time.
</div>

This article explains how to install the AWS runner on Linux. The AWS runner is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone-runner-aws).

## Drone specific Configuration

The AWS runner is configured using environment variables. This article references the below configuration options. See [Configuration]({{< relref "configuration/reference" >}}) for a complete list of configuration options.

- __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.
- __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.
- __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.

## AWS specific Configuration

### AWS EC2 prerequisites

There are some pieces of setup that need to be performed on the AWS side first.

- Set up an access key and access secret [aws secret](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_access-keys.html#Using_CreateAccessKey) which is needed for configuration of the runner.
- Setup up vpc firewall rules for the build instances [ec2 authorizing-access-to-an-instance](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/authorizing-access-to-an-instance.html) We need allow ingress and egress access to port 22. Once complete you will have a security group id, which is needed for configuration of the runner.
- (windows only instance) You need to add the `AdministratorAccess` policy to the IAM role associated with the access key and access secret [IAM](https://console.aws.amazon.com/iamv2/home#/users). You will use the instance profile arn `iam_profile_arn`, in your pipeline.

### AWS EC2 environment variables

Set up the runner by using either an environemt variables or a `.env` file similar to other Drone runners. Below is a list of the AWS specific environment variables.

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

```BASH
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

Before starting the runner please create a pool file. See [Pool file]({{< relref "configuration/pool.md" >}}) for more information.

## AWS pipeline syntax

For information on configuring an AWS pipeline see [AWS pipeline syntax]({{< relref "../../pipeline/aws/overview.md" >}})
