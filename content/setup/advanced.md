+++
date = "2016-02-05T18:20:44-06:00"
draft = false
title = "Advanced Settings"
weight = 10
menu = "installation"
toc = true
+++

Advanced settings allow for highly customized deployments.

# Public Mode

Drone restricts the build history to authenticated users with pull access to
the source repository. In some deployments it may be desirable to make the
build history visible to all users on the network. Public mode permits
unauthenticated users to access the build history for every repository,
_including private repositories_.

Add the following configuration to `/etc/drone/dronerc` to enable public
mode:

```
PUBLIC_MODE=true
```

Enabling this feature on a public network will leak sensitive information
in the build logs.
