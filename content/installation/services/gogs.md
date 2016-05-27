+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Gogs"
weight = 21
toc = true

[menu.main]
	parent="remotes"
+++

# Overview

Drone comes with built-in support for Gogs the latest stable version of Gogs. To enable Gogs you should configure the driver using following required environment variables:

```
DRONE_GOGS=true
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

Pull Request and Tag events are not supported. We are interested in patches to include this functionality. If you are interested in contributing to Drone and submitting a patch please [contact us](https://gitter.im/drone/drone).

# Common Issues

When running Drone and Gogs inside separate Docker containers you may experience inter-container network and communication issues. You should configure Drone and Gogs to use fully qualified domain names and not custom hostnames.
