+++
weight = 5
date = "2014-11-08T12:54:06-08:00"
title = "Gogs"

[menu.main]
parent = "Configure"
+++

You may configure Drone to integrate with Gogs. This can be configured in the `/etc/drone/drone.toml` configuration file:

```toml
[gogs]
url = "http://gogs.drone.io"
secret = "c0aaff74c060ff4a950d"
```

## Environment Variables

You may also configure Gogs using environment variables. This is useful when running Drone inside Docker containers, for example.

```bash
DRONE_GOGS_URL="http://gogs.drone.io"
DRONE_GOGS_SECRET="c0aaff74c060ff4a950d"
```

## Known Issues

Private repositories are not supported. The [Gogs API](https://github.com/gogits/go-gogs-client) does not provide a means to upload a Public Key to a repository for cloning purposes. If you are interested in implementing this feature please contact us on [Gitter](https://gitter.im/drone/drone).