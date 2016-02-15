+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Network Proxy"
weight = 10
menu = "installation"
toc = true
+++

# Overview

This document provides high-level instructions for configuring Drone to work with a corporate proxy server. This may be required when running Drone behind a **corporate firewall**.

# HTTP Proxy

The HTTP_PROXY environment variable holds the hostname or IP address of your proxy server. You can specify the HTTP_PROXY variables in your server configuration file or as an environment variable.

```bash
HTTPS_PROXY=https://proxy.example.com
HTTP_PROXY=http://proxy.example.com
```

These variables are propagated throughout your build environment, including build and plugin containers. To verify the environment variables are being set in your build container you can add the `env` command to your build script.

We also recommend you provide both uppercase and lowercase environment variables. We've found that certain common unix tools are case-sensitive:

```bash
HTTP_PROXY=http://proxy.example.com
http_proxy=http://proxy.example.com
```

# No Proxy

The `NO_PROXY` variable should contain a comma-separated list of domain extensions the proxy should not be used for. This typically includes resources inside your network, such as your GitHub Enterprise server.

```bash
NO_PROXY=.example.com, *.docker.example.com
```

You may also need to add your Docker daemon hostnames to the above list.
