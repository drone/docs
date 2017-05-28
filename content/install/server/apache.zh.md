+++
date = "2017-04-15T14:39:04+02:00"
title = "Apache 配置"
url = "zh/setup-with-apache"

[menu.install]
  identifier = "setup-with-apache-zh"
  parent = "install_server"
  weight = 5
+++

<!--This guide provides a brief overview for installing Drone server behind the Apache2 webserver. This is an example configuration:-->

这个指南概述了将 Drone 安装在 Apache2 服务器上。下面是一个示例配置：

```nohighlight
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

ProxyPass /ws/ ws://localhost:8000/ws/
ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

<!--You must have the below Apache modules installed.-->

您需要安装下面的 Apache 模块：

```nohighlight
a2enmod proxy
a2enmod proxy_http
a2enmod proxy_wstunnel
```

<!--You must configure Apache to set `X-Forwarded-Proto` when using https.-->

如果需要使用 https，您需要设置  `X-Forwarded-Proto`

```diff
ProxyPreserveHost On

+RequestHeader set X-Forwarded-Proto "https"

ProxyPass /ws/ ws://localhost:8000/ws/
ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```

<!--You must configure Apache to enable websocket upgrades.-->

您需要配置 Apache 启用 websocket upgrade 。

```diff
ProxyPreserveHost On

RequestHeader set X-Forwarded-Proto "https"

+ProxyPass /ws/ ws://localhost:8000/ws/
+ProxyPassReverse /ws/ ws://localhost:8000/ws/

ProxyPass / http://127.0.0.1:8000/
ProxyPassReverse / http://127.0.0.1:8000/
```
