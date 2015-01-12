+++
weight = 3
date = "2014-11-08T12:54:06-08:00"
title = "Gitlab"

[menu.main]
parent = "Configure"
+++

You may configure Drone to integrate with GitLab. This can be configured in the `/etc/drone/drone.toml` configuration file:

```toml
[gitlab]
url = "http://gitlab.drone.io"
```

## Environment Variables

You may also configure Gitlab using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_GILAB_URL="http://gitlab.drone.io"
```
