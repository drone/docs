+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Cloud"
url = "install-for-bitbucket-cloud"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-cloud"
  weight = 5
+++


Drone comes with built-in support for Bitbucket Cloud. To enable Bitbucket Cloud you should configure the Drone container using the following environment variables:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_BITBUCKET=true
+     - DRONE_BITBUCKET_CLIENT=95c0282573633eb25e82
+     - DRONE_BITBUCKET_SECRET=30f5064039e6b359e075
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/drone:{{% version %}}
    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=ws://drone-server:8000/ws/broker
      - DRONE_SECRET=${DRONE_SECRET}
```

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.

DRONE_BITBUCKET=true
: Set to true to enable the Bitbucket driver.

DRONE_BITBUCKET_CLIENT
: Bitbucket oauth2 client id

DRONE_BITBUCKET_SECRET
: Bitbucket oauth2 client secret

# Registration

You must register your application with Bitbucket in order to generate a client and secret. Navigate to your account settings and choose OAuth from the menu, and click Add Consumer.

Please use the Authorization callback URL:

```nohighlight
http://drone.mycompany.com/authorize
```

Please also be sure to check the following permissions:

```nohighlight
Account:Email
Account:Read
Team Membership:Read
Repositories:Read
Webhooks:Read and Write
```

# Missing Features

Merge requests and mercurial repositories are not currently supported. We are interested in patches to include this functionality. If you are interested in contributing to Drone and submitting a patch please [contact us](https://gitter.im/drone/drone).
