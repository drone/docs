+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Gogs"
weight = 21
menu = "installation"
toc = true
+++

# Overview

Drone comes with built-in support for Gogs the latest stable version of Gogs. To enable Gogs you should configure the driver using following required environment variables:

```
DRONE_GOGS_URL=http://gogs.mycompany.com
```

Once configured you can login to Drone using your Gogs username and password. Gogs does not currently support oauth which is why username and password are required.

# Configuration

This is a full list of configuration options. Please note that many of these options use default configuration value that should work for the majority of installations.

NAME                        | DESC
----------------------------|--------------------------------------------------------
`DRONE_GOGS_URL`            | gogs server address
`DRONE_GOGS_PRIVATE_MODE`   | gogs is running in private mode
`DRONE_GOGS_SKIP_VERIFY`    | gogs certificate is self-signed

# Missing Features

Please note that Drone does not currently support Pull Request or Tag events. We are absolutely open to pull requests to include this functionality. If you are interested in contributing to Drone and adding these capabilities please [get in touch](https://gitter.im/drone/drone).

# Common Issues

When running Drone and Gogs inside separate Docker containers you may exerpience inter-container network and communication issues. You should configure Drone and Gogs to use fully qualified domain names and not custom hostnames.
