+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Gogs"
weight = 4
menu = "remotes"
toc = true
+++

Drone comes with built-in support for Gogs version `0.6.16.1022` and higher. Please ensure you are running the latest version of Gogs to avoid compatibility issues. To enable Gogs you should configure the Gogs driver using the following environment variables:

```bash
REMOTE_DRIVER=gogs
REMOTE_CONFIG=https://gogs.hooli.com?open=false
```

# Gogs configuration

The following is the standard URI connection scheme:

```
scheme://host[:port][?options]
```

The components of this string are:

* `scheme` server protocol `http` or `https`.
* `host` server address to connect to. The default value is github.com if not specified.
* `:port` optional. The default value is :80 if not specified.
* `?options` connection specific options.

# Gogs options

This section lists all connection options used in the connection string format. Connection options are pairs in the following form: `name=value`. The value is always case sensitive. Separate options with the ampersand (i.e. &) character:

* `open=false` allows users to self-register. Defaults to false for security reasons.
* `skip_verify=false` skip ca verification if self-signed certificate. Defaults to false for security reasons.

# Remote Driver Feature Chart

While Gogs and Drone's support for Gogs are still developing, this is an 
excellent option for teams with very simple hosting needs. 

| Feature/Remote            | Gogs         |
|---------------------------|--------------|
| Supported version         | 0.6.16.1022+ |
| VCS                       | git          |
| Auth method               | manual       |
| Push events               | yes          |
| Push tags events          | **no**       |
| Merge requests            | **no**       |
| Commit statuses           | **no**       |
| Restrict by organizations | **no**       |

# Next Steps

Once you have configured your Remote Driver, it's time to [Select and 
Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
