+++
date = "2017-04-15T14:39:04+02:00"
title = "Setup with Nginx"
url = "setup-with-nginx"

[menu.install]
  identifier = "setup-with-nginx"
  parent = "install_server"
  weight = 3
+++

This guide provides a basic overview for installing Drone server behind the nginx webserver. For more advanced configuration options please consult the official nginx [documentation](https://www.nginx.com/resources/admin-guide/).

Example configuration:

```nginx
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
}
```

You must configure the proxy to set `X-Forwarded` proxy headers:

```diff
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
}
```
