+++
date = "2018-06-19T14:39:04+02:00"
title = "Network Proxy"
url = "network-proxy"

[menu.install]
  identifier = "network-proxy"
  parent = "install_server"
  weight = 10
+++

This is a brief guide on how to set a network proxy for drone.
This can be useful when deploying drone behind a corporate firewall

```diff
services:
  drone-server:
    image: drone/drone:0.8

    ports:
      - 80:8000
      - 9000
    volumes:
      - /var/lib/drone:/var/lib/drone/
    restart: always
    environment:
      - DRONE_OPEN=true
      - DRONE_HOST=${DRONE_HOST}
      - DRONE_SECRET=${DRONE_SECRET}
+     - HTTP_PROXY=http://my-proxy-server:1234
+     - HTTPS_PROXY=http://my-proxy-server:1234
+     - NO_PROXY=localhost,127.0.0.1,my-docker-registry,my-git-server,drone-agent

  drone-agent:
    image: drone/agent:0.8

    command: agent
    restart: always
    depends_on:
      - drone-server
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    environment:
      - DRONE_SERVER=drone-server:9000
      - DRONE_SECRET=${DRONE_SECRET}
+     - HTTP_PROXY=http://my-proxy-server:1234
+     - HTTPS_PROXY=http://my-proxy-server:1234
+     - NO_PROXY=localhost,127.0.0.1,my-docker-registry,my-git-server,drone-agent
```
