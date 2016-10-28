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

If you opt to use a different exposed port, be sure to adjust `upstream drone` accordingly.

If SSL is not to be used, change the listening port from `443` to `80`
and comment out the SSL configuration block.

Example [nginx](http://nginx.org) reverse proxy configuration:

```
# Handle connection upgrading.
map $http_upgrade $connection_upgrade {
    default upgrade;
    ''      close;
}

# Establish the upstream to Drone.
upstream drone {
    server 127.0.0.1:8000;
}

server {
    listen 443 ssl;
    server_name drone.example.com;

    # Proxy all requests to the Drone application.
    location / {
        access_log off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header Host $http_host;

        proxy_pass http://drone;
        proxy_redirect off;
        proxy_http_version 1.1;
        proxy_buffering off;

        chunked_transfer_encoding off;
    }

    # Handle WebSockets by catching all /ws (case-insensitive) and upgrade the connection.
    location ~* /ws {
        access_log off;
        proxy_pass http://drone;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection "upgrade";
        proxy_read_timeout 86400;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header X-Forwarded-Proto $scheme;
    }

	# SSL Configuration
    ssl on;
    ssl_certificate /etc/nginx/ssl/example.com.crt;
    ssl_certificate_key /etc/nginx/ssl/example.com.key;
    ssl_trusted_certificate /etc/nginx/ssl/ca-certs.pem;
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

  ProxyPass /api/queue/logs ws://127.0.0.1:8000/api/queue/logs
  ProxyPassReverse /api/queue/logs ws://127.0.0.1:8000/api/queue/logs

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
