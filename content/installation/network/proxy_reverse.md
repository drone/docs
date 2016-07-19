+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Reverse Proxy"
weight = 31
toc = true


[menu.main]
	parent="proxy"
+++

# Overview

This section of the documentation provides sample configurations for using Drone with reverse proxy servers such as nginx. When using a reverse proxy it is imperative you set the `X-Forwarded-Proto` and `X-Forwarded-For` header variables. You will also need to run the Drone container using an alternate exposed port, for example `--port-8000:8000`.

# Nginx

Example [nginx](http://nginx.org) reverse proxy configuration:

```
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}

location / {
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
    proxy_set_header Origin "";
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;

    proxy_pass http://127.0.0.1:8000;
    proxy_http_version 1.1;
    proxy_redirect off;
    proxy_buffering off;

    chunked_transfer_encoding off;
}
```

# Caddy

Example [caddy server](https://caddyserver.com/) reverse proxy configuration:

```
drone.mycomopany.com {
    proxy / localhost:8000 {
        websocket
        proxy_header X-Forwarded-Proto {scheme}
        proxy_header X-Forwarded-For {host}
        proxy_header Host {host}
    }
}
```


# Apache

Example [apache](http://httpd.apache.org) reverse proxy configuration:

```
  ProxyPreserveHost On
  RequestHeader set X-Forwarded-Proto "https"

  # Websocket endpoints

  ProxyPass /ws/ ws://localhost:8000/ws/
  ProxyPassReverse /ws/ ws://localhost:8000/ws/

  RewriteEngine On
  RewriteCond %{REQUEST_METHOD} =GET
  RewriteRule ^/api/queue/logs/(.*) ws://localhost:8000/api/queue/logs/$1 [P]
  ProxyPassReverse /api/queue/logs ws://localhost:8000/api/queue/logs

  ProxyPass /api/queue/pull ws://127.0.0.1:8000/api/queue/pull
  ProxyPassReverse /api/queue/pull ws://127.0.0.1:8000/api/queue/pull

  ProxyPass /api/stream ws://127.0.0.1:8000/api/stream
  ProxyPassReverse /api/stream ws://127.0.0.1:8000/api/stream

  # All the rest (not websockets)

  ProxyPass / http://127.0.0.1:8000/
  ProxyPassReverse / http://127.0.0.1:8000/
```

Do not forget to to install proxy modules:
```
a2enmod proxy
a2enmod proxy_http
a2enmod proxy_wstunnel
```
