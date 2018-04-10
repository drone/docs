+++
date = "2017-07-23T16:38:54+08:00"
title = "Coding"
url = "install-for-coding"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-coding"
  weight = 7
+++

Drone comes with built-in support for Coding. To enable Coding you should configure the Drone container using the following environment variables:

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
      - 80:8000
      - 9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_CODING=true
+     - DRONE_CODING_CLIENT=${DRONE_CODING_CLIENT}
+     - DRONE_CODING_SECRET=${DRONE_CODING_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}

  drone-agent:
    image: drone/agent:{{% version %}}
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=drone-server:9000
      - DRONE_SECRET=${DRONE_SECRET}
```

# Registration

Register your application with Coding to create your client id and secret.
Navigate to Account(账户)/Applications(应用管理), then click New(添加应用), and you will see Application Setting(应用设置) page.
It is very important the callback URL(回调地址) matches your http(s) scheme and hostname exactly with `/authorize` as the path.

<a href="images/coding_oauth.png" target="_blank"><img src="images/coding_oauth.png" alt="coding oauth setup"></a>

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.

DRONE_CODING=true
: Set to true to enable the Coding driver.

DRONE_CODING_URL=`https://coding.net`
: Coding server address. The default value is for platform edition. For enterprise edition, set to `https://{company}.coding.net`.

DRONE_CODING_CLIENT
: Coding oauth2 client id.

DRONE_CODING_SECRET
: Coding oauth2 client secret.

DRONE_CODING_SCOPE=user,project,project:depot
: Comma-separated Coding oauth scope.

DRONE_CODING_GIT_MACHINE=git.coding.net
: Coding git server hostname. The default value is for platform edition. For enterprise edition, set to `e.coding.net`.

DRONE_CODING_GIT_USERNAME
: Optional. Use a single machine account username to clone all repositories.

DRONE_CODING_GIT_PASSWORD
: Optional. Use a single machine account password to clone all repositories.

DRONE_CODING_SKIP_VERIFY=false
: Set to true to disable SSL verification.
