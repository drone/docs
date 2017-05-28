+++
date = "2017-04-15T14:39:04+02:00"
title = "Configure Lets Encrypt"
url = "zh/configure-lets-encrypt"

[menu.install]
  identifier = "configure-lets-encrypt-zh"
  parent = "install_server"
  weight = 8
+++

<!--Drone supports automated ssl configuration and updates using let's encrypt. You can enable let's encrypt by making the following modifications to your server configuration:-->

Drone 支持自动 SSL 配置和升级 let's encrypt。你需要修改服务器配置来启动 let's encrypt：

```diff
services:
  drone-server:
    image: drone/drone:{{% version %}}
    ports:
+     - 80:80
+     - 443:443
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_GITHUB=true
      - DRONE_GITHUB_CLIENT=${DRONE_GITHUB_CLIENT}
      - DRONE_GITHUB_SECRET=${DRONE_GITHUB_SECRET}
      - DRONE_SECRET=${DRONE_SECRET}
+     - DRONE_LETS_ENCRYPT=true
```

<!--Note that Drone uses the hostname from the `DRONE_HOST` environment variable when requesting certificates. For example, if `DRONE_HOST=https://foo.com` the certificate is requested for `foo.com`.-->

注意 Drone 使用 `DRONE_HOST` 环境变量来请求对应 hostname 的证书。比如 `DRONE_HOST=https://foo.com` 将请求 `foo.com` 的 SSL 证书。

 <!-- by author Once enabled you can visit your website at both the http and the https address. There are no immediate plans to redirect from http to https, but it may be considered for a future release. -->


<!--# Certificate Cache-->

# 证书缓存

<!--Drone writes the certificates to the below directory:-->

Drone 将证书写道以下目录

```
/var/lib/drone/golang-autocert
```

<!--# Certificate Updates-->

# 证书更新

<!--Drone uses the official Go acme library which will handle certificate upgrades. There should be no addition configuration or management required.-->

Drone 使用官方 Go acme 仓库来处理证书升级。这里不需要手工更新。
