+++
weight = 1
date = "2014-11-08T12:54:12-08:00"
title = "Settings"

[menu.main]
parent = "Configure"
+++

The Drone configuration settings are declared in `/etc/drone/drone.toml`. Alternatively, you may declare Drone configuration settings as environment variables, which is useful when running Drone in Docker.

In order to run Drone you must setup at least one of the following integration points:

* [GitHub](../github/)
* [GitHub Enterprise](../github)
* [Gitlab](../gitlab)
* [Gogs](../gogs)
* [Bitbucket](../bitbucket)

All other settings are options.

---

## Default Port and SSL

By default, Drone runs on port `:80` and uses `http`. You can modify this behavior with the following configuration variables:

```toml
[server]
port=":443"

[server.ssl]
key="/path/to/key.pem"
cert="/path/to/cert.pem"
```

Or by using environment variables:

```bash
DRONE_SERVER_PORT=":443"
DRONE_SERVER_SSL_KEY="/path/to/key.pem"
DRONE_SERVER_SSL_CERT="/path/to/cert.pem"
```

---

## Session Expiration

Session timeout is 72 hours by default. Session tokens are signed with a secret string that is randomly generated when Drone starts. In some cases you may not want Drone to use a random secret. You can modify this behavior with the following configuration variables:

```toml
[session]
secret="70038eaa535c0ad1ae468c9b5ade50834086a585"
expires="72h"
```

Or by using environment variables:

```bash
DRONE_SESSION_SECRET="70038eaa535c0ad1ae468c9b5ade50834086a585"
DRONE_SESSION_EXPIRES="72h"
```

---
