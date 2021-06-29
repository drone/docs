---
date: 2000-01-01T00:00:00+00:00
title: Bitbucket Server
author: bradrydzewski
steps: true
toc: true
weight: 5
aliases:
- /installation/providers/bitbucket-cloud
- /install-for-bitbucket-server
- /admin/setup-bitbuket-server/
- /installation/bitbucket-server
- /installation/bitbucket-server/single-machine
- /installation/bitbucket-server/multi-machine
- /installation/bitbucket-server/kubernetes
- /installation/providers/bitbucket-server/
description: |
  Installation and configure for use with Bitbucket Server
---

This article explains how to install the Drone server for Bitbucket Server, formerly known as Atlassian Stash. The server is packaged as a minimal Docker image distributed on [DockerHub](https://hub.docker.com/r/drone/drone).

# Preparation

## Create a Personal Access Token

Create a personal access token that is capable of cloning all repositories in the system. The token and associated username are used for all clone operations. We recommend creating a machine account for this purpose.

Navigate to the _Personal Access Tokens_ page in the account settings, and click the _Create Token_ button.

![Token List](/screenshots/stash_token_list.png)

Create the personal access token. The creation form should indicate pull and clone access as pictured below. Click the _Create_ button and copy the generated token.

![Token Create](/screenshots/stash_token_create.png)

## Create a Key Pair

Create a key pair on your server. The key pair is used to setup an authentication provide with Bitbucket and authorize API access.

Generate the private key:

```
$ openssl genrsa -out /etc/bitbucket/key.pem 1024
Generating RSA private key, 1024 bit long modulus
....................................++++++
..........++++++
e is 65537 (0x10001)
```

Generate a public key:

```
$ openssl rsa \
  -in /etc/bitbucket/key.pem \
  -pubout >> /etc/bitbucket/key.pub
```

## Create an Application Link

Create a Bitbucket Application Link. The link will provide a Consumer ID and Private Key used to authorize access to Bitbucket resources. The Bitbucket application creation process is convoluted and error prone. Please bear with us.

Navigate the administrator panel and click the _Application Links_ settings page. Enter your Drone server URL and click _Create New Link_.

![stash_application_link](/screenshots/stash_application_link.png)

Please fill out the form using the values specified below. Once complete click _Continue_ to create your application.

* Set the application name to _Drone_
* Set the application type to _Generic Application_
* Set the provider name to _Drone_
* Set the consumer key to _OauthKey_
* Set the shared secret to any random alphanumeric value
* Set the request token url to your Drone server URL
* Set the access token url to your Drone server URL
* Set the authorize token url to your Drone server URL

![stash_application_link_create](/screenshots/stash_application_link_create.png)

Once the application is created it needs to be edited so that we can configure the _Incoming Authentication_. Please fill out the form using the values specified below and save your changes.

* Set the consumer key to _OauthKey_
* Set the consumer name to _Drone_
* Paste the contents of `/etc/bitbucket/key.pub` in the public key textarea
* Leave _Consumer Callback_ empty
* Leave _Allow 2-Legged Oauth_ unchecked

![stash_application_link_edit](/screenshots/stash_application_link_edit.png)

Congratulations, you have made it through the most painful part of the installation. With luck, everything will work as expected and you will never have to do this again.

## Create a Shared Secret
Create a shared secret to authenticate communication between runners and your central Drone server.

You can use openssl to generate a shared secret:

```
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

# Download

The Drone server is distributed as a lightweight Docker image. The image is self-contained and does not have any external dependencies.

```
$ docker pull drone/drone:2
```

# Configuration

The Drone server is configured using environment variables. This article references a subset of configuration options, defined below. See [Configuration]({{< relref "../reference/_index.md" >}}) for a complete list of configuration options.

* __DRONE_GIT_USERNAME__
  : Required string value set to username associated with the Personal Account token. This username is used to authenticate and clone all private repositories.
* __DRONE_GIT_PASSWORD__
  : Required string value set to your Personal Account Token. The token is used to authenticate and clone all private repositories.
* __DRONE_GIT_ALWAYS_AUTH__
  : Optional boolean value configures Drone to authenticate when cloning public repositories. This should only be enabled when using GitHub Enterprise with private mode enable.
* __DRONE_STASH_CONSUMER_KEY__
  : Required string value configures your Bitbucket Server consumer key.
* __DRONE_STASH_PRIVATE_KEY__
  : Required string value configures the path to your Bitbucket Server private key file. Note that this file needs to also be mounted into the Drone server container as a volume.
* __DRONE_STASH_SERVER__
  : Required string value provides the Bitbucket Server address. For example `https://bitbucket.company.com`
* __DRONE_RPC_SECRET__
  : Required string value provides the shared secret generated in the previous step. This is used to authenticate the rpc connection between the server and runners. The server and runner must be provided the same secret value.
* __DRONE_SERVER_HOST__
  : Required string value provides your external hostname or IP address. If using an IP address you may include the port.
* __DRONE_SERVER_PROTO__
  : Required string value provides your external protocol scheme. This value should be set to http or https. This field defaults to https if you configure ssl or acme.

# Start the Server

The server container can be started with the below command. The container is configured through environment variables. _Remember to replace the placeholder values below with the appropriate values._

{{< highlight handlebars "linenos=table" >}}
docker run \
  --volume=/var/lib/drone:/data \
  --volume=/etc/bitbucket/key.pem:/etc/bitbucket/key.pem \
  --env=DRONE_GIT_PASSWORD={{DRONE_GIT_PASSWORD}} \
  --env=DRONE_GIT_USERNAME={{DRONE_GIT_USERNAME}} \
  --env=DRONE_GIT_ALWAYS_AUTH=false \
  --env=DRONE_STASH_SERVER={{DRONE_STASH_SERVER}} \
  --env=DRONE_STASH_CONSUMER_KEY=OauthKey \
  --env=DRONE_STASH_PRIVATE_KEY=/etc/bitbucket/key.pem \
  --env=DRONE_SERVER_HOST={{DRONE_SERVER_HOST}} \
  --env=DRONE_SERVER_PROTO={{DRONE_SERVER_PROTO}} \
  --env=DRONE_RPC_SECRET={{DRONE_RPC_SECRET}} \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  drone/drone:2
{{< / highlight >}}

# Install Runners

Once your server is up and running you will need to install runners to execute your build pipelines. See our runner installation documentation for detailed installation instructions. 

{{< link "/runner/overview" "Install Runners" >}}
