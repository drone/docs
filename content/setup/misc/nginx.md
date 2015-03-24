+++
weight = 1
date = "2014-11-08T12:50:08-08:00"
title = "Nginx"

[menu.main]
parent = "Misc"
+++

Using a proxy server is not really necessary. Drone serves most static content from a CDN and uses the Go standard library’s high-performance net/http package to serve dynamic content.

If using Nginx to proxy traffic to Drone, please ensure you have version 1.3.13 or greater. You also need to configure nginx to proxy websocket connections:

```nginx
location / {
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
    proxy_set_header Origin "";

    proxy_pass http://127.0.0.1:8000;
    proxy_redirect off;
}

location /api/stream {
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
    proxy_set_header Origin "";

    proxy_pass http://127.0.0.1:8000;
    proxy_redirect off;

    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```

You may also want to change Drone's default port when proxying traffic. You can change the port in the `/etc/drone/drone.toml` configuration file:

```toml
[server]
port = ":8000"
```
