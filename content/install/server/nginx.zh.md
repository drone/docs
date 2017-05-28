+++
date = "2017-04-15T14:39:04+02:00"
title = "Setup with Nginx"
url = "zh/setup-with-nginx"

[menu.install]
  identifier = "setup-with-nginx-zh"
  parent = "install_server"
  weight = 3
+++

<!--This guide provides a basic overview for installing Drone server behind the nginx webserver. For more advanced configuration options please consult the official nginx [documentation](https://www.nginx.com/resources/admin-guide/).-->

这个指南概述了将 Drone 安装在 Nginx 上。访问 [Nginx 官方文档](https://www.nginx.com/resources/admin-guide/) 来了解更多高级配置。下面是一个示例配置：

```nginx
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}

server {
    listen 80;
    server_name drone.example.com;

    location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;

        proxy_pass http://127.0.0.1:8000;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }

    location ~* /ws {
        proxy_pass http://127.0.0.1:8000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 86400;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;
    }
}
```

<!--You must configure the proxy to set `X-Forwarded` proxy headers:-->

您需要配置方向代理使用 `X-Forwarded` 头。

```diff
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}

server {
    listen 80;
    server_name drone.example.com;

    location / {
+       proxy_set_header X-Forwarded-For $remote_addr;
+       proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass http://127.0.0.1:8000;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }

    location ~* /ws {
        proxy_pass http://127.0.0.1:8000;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 86400;
+       proxy_set_header X-Forwarded-For $remote_addr;
+       proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

<!--You must configure Nginx to upgrade websockets.-->

您需要配置 Nginx websockets upgrade 。

```diff
+map $http_upgrade $connection_upgrade {
+    default upgrade;
+    ''      close;
+}

server {
    listen 80;
    server_name drone.example.com;

    location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;

        proxy_pass http://127.0.0.1:8000;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }

    location ~* /ws {
        proxy_pass http://127.0.0.1:8000;
        proxy_http_version 1.1;
+       proxy_set_header Upgrade $http_upgrade;
+       proxy_set_header Connection "upgrade";
        proxy_read_timeout 86400;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```
