+++
date = "2017-04-15T14:39:04+02:00"
title = "Setup with Caddy"
url = "zh/setup-with-caddy"

[menu.install]
  identifier = "setup-with-caddy-zh"
  parent = "install_server"
  weight = 4
+++

<!--This guide provides a brief overview for installing Drone server behind the [Caddy webserver](https://caddyserver.com/). This is an example caddyfile proxy configuration:-->

这个指南概述了将 Drone 安装在 [Caddy](https://caddyserver.com/) 上。下面是一个示例配置：


```nohighlight
drone.mycomopany.com {
    proxy / localhost:8000 {
        websocket
        transparent
    }
}
```

<!--You must configure the proxy to enable websocket upgrades:-->

您需要启用 websocket upgrade：

```diff
drone.mycomopany.com {
    proxy / localhost:8000 {
+       websocket
        transparent
    }
}
```

<!--You must configure the proxy to include `X-Forwarded` headers using the `transparent` directive:-->

您需要使用 `transparent` 语法来配置反向代理包含 `X-Forwarded` 头

```diff
drone.mycomopany.com {
    proxy / localhost:8000 {
        websocket
+       transparent
    }
}
```
