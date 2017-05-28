+++
date = "2017-04-15T14:39:04+02:00"
title = "Bitbucket Server"
url = "zh/install-for-bitbucket-server"

[menu.install]
  parent = "install_integrations"
  identifier = "install-for-bitucket-server-zh"
  weight = 6
+++

<!--Drone comes with experimental support for Bitbucket Server, formerly known as Atlassian Stash. To enable Bitbucket Server you should configure the Drone container using the following environment variables:-->

Drone 实验性支持 Bitbucket Server，之前又被称为 Atlassian Stash。使用下面的环境变量来配置使用 Bitbucket Server。

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
+     - DRONE_STASH=true
+     - DRONE_STASH_GIT_USERNAME=foo
+     - DRONE_STASH_GIT_PASSWORD=bar
+     - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
+     - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
    volumes:
+     - /path/to/key.pem:/path/to/key.pem

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

<!--# Private Key File-->

# Private Key 私钥文件

<!--The OAuth process in Bitbucket server requires a private and a public RSA certificate. This is how you create the private RSA certificate.-->

Bitbucket server OAuth 要求使用一个 RSA 私钥和公钥。下面的步骤用来生成 RSA 私钥凭据：

```nohighlight
openssl genrsa -out /etc/bitbucket/key.pem 1024
```

这个步骤将私钥凭据储存在 `key.pem`。下面的命令用来生成公钥 RSA 凭据，并将其储存在 `key.pub` 中。

<!--This stores the private RSA certificate in `key.pem`. The next command generates the public RSA certificate and stores it in `key.pub`.-->

```nohighlight
openssl rsa -in /etc/bitbucket/key.pem -pubout >> /etc/bitbucket/key.pub
```

<!--
Please note that the private key file can be mounted into your Drone conatiner at runtime or as an environment variable

Private key file mounted into your Drone container at runtime as a volume.-->

私钥文件可以在在运行时绑定在 Drone 容器上作为一个空间（volume）来使用，或者作为一个环境变量。

私钥文件可以在在运行时绑定在 Drone 容器上的示例：

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA=/etc/bitbucket/key.pem
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
+  volumes:
+     - /etc/bitbucket/key.pem:/etc/bitbucket/key.pem
```

<!--Private key as environment variable-->

私钥作为环境变量的示例：

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
    - DRONE_OPEN=true
    - DRONE_HOST=${DRONE_HOST}
      - DRONE_STASH=true
      - DRONE_STASH_GIT_USERNAME=foo
      - DRONE_STASH_GIT_PASSWORD=bar
      - DRONE_STASH_CONSUMER_KEY=95c0282573633eb25e82
+     - DRONE_STASH_CONSUMER_RSA_STRING=contentOfPemKeyAsString
      - DRONE_STASH_URL=http://stash.mycompany.com
      - DRONE_SECRET=${DRONE_SECRET}
```

<!--# Service Account-->

# 服务账户

<!--Drone users `git+https` to clone repositories, however, Bitbucket Server does not currently support cloning repositories with oauth token. To work around this limitation, you must create a service account and provide the username and password to Drone. This service account will be used to authenticate and clone private repositories.-->

Drone 使用 `git+https` 来克隆仓库，Bitbucket Server 不支持 使用 oauth token 来克隆仓库。为了绕过这个限制，需要先生成一个服务账户（service account），同时将用户名和密码提供给 Drone。这个服务账户将会被用来克隆仓库。

<!--# Registration-->

# 注册应用程序

<!--You must register your application with Bitbucket Server in order to generate a consumer key. Navigate to your account settings and choose Applications from the menu, and click Register new application. Now copy & paste the text value from `/etc/bitbucket/key.pub` into the `Public Key` in the incoming link part of the application registration.-->

在 Bitbucket Server 上注册一个应用，并生成 consumer key。访问账户设置（account settings）页面，选择 Applications 页面，点击 Register new application 。在应用注册的 incoming link 部分，复制 `/etc/bitbucket/key.pub` 的对应文本到 `Public Key`。

使用下面的认证回调 URL（Authorization callback URL），请修改域名为自定义域名： `http://drone.mycompany.com/authorize`

<!--Please use http://drone.mycompany.com/authorize as the Authorization callback URL.-->


<!--# Configuration-->

# 配置

<!--This is a full list of configuration options. Please note that many of these options use default configuration values that should work for the majority of installations.-->

下面是所有的配置选项。一般来说，使用默认配置可以满足绝大部分的安装需求：


<!--
DRONE_STASH=true
: Set to true to enable the Bitbucket Server (Stash) driver.

DRONE_STASH_URL
: Bitbucket Server address.

DRONE_STASH_CONSUMER_KEY
: Bitbucket Server oauth1 consumer key

DRONE_STASH_CONSUMER_RSA
: Bitbucket Server oauth1 private key file

DRONE_STASH_CONSUMER_RSA_STRING
: Bibucket Server oauth1 private key as a string

DRONE_STASH_GIT_USERNAME
: Machine account username used to clone repositories.

DRONE_STASH_GIT_PASSWORD
: Machine account password used to clone repositories.-->

DRONE_STASH=true
: true 使用 Bitbucket Server (Stash)

DRONE_STASH_URL
: Bitbucket Server 地址

DRONE_STASH_CONSUMER_KEY
: Bitbucket Server oauth1 consumer key

DRONE_STASH_CONSUMER_RSA
: Bitbucket Server oauth1 private key 文件

DRONE_STASH_CONSUMER_RSA_STRING
: Bibucket Server oauth1 private key 字符串

DRONE_STASH_GIT_USERNAME
: 用来克隆仓库的服务账户的账户名

DRONE_STASH_GIT_PASSWORD
: 用来克隆仓库的服务账户的密码