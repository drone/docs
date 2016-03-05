+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "GitHub"
weight = 2
menu = "remotes"
toc = true
+++

Drone comes with built-in support for GitHub and GitHub Enterprise. To enable GitHub you should configure the GitHub driver using the following environment variables:

```bash
REMOTE_DRIVER=github
REMOTE_CONFIG=https://github.com?client_id=${client_id}&client_secret=${client_secret}
```

# GitHub configuration

The following is the standard URI connection scheme:

```
scheme://host[:port][?options]
```

The components of this string are:

* `scheme` server protocol `http` or `https`.
* `host` server address to connect to. The default value is github.com if not specified.
* `:port` optional. The default value is :80 if not specified.
* `?options` connection specific options.

# GitHub options

This section lists all connection options used in the connection string format. Connection options are pairs in the following form: `name=value`. The value is always case sensitive. Separate options with the ampersand (i.e. &) character:

* `client_id` oauth client id for registered application.
* `client_secret` oauth client secret for registered application.
* `open=false` allows users to self-register. Defaults to false..
* `orgs=drone&orgs=docker` restricts access to these GitHub organizations. **Optional**
* `private_mode=false` indicates GitHub Enterprise is running in private mode.
* `skip_verify=false` skip ca verification if self-signed certificate. Defaults to false.
* `merge_ref=merge` configure pull requests to use head vs merge ref. Defaults to merge.

# GitHub registration

You must register your application with GitHub in order to generate a Client and Secret. Navigate to your account settings and choose Applications from the menu, and click Register new application.

Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL.

# Remote Driver Feature Chart

Since the majority of our users (including Drone itself) use GitHub, the 
GitHub remote is in great shape. At any given time, Drone should work well 
with both GitHub and GitHub Enterprise.

| Feature/Remote            | GitHub                       |
|---------------------------|------------------------------|
| Supported versions        | GitHub and GitHub Enterprise |
| VCS                       | git                          |
| Auth method               | ouath2                       |
| Push events               | yes                          |
| Push tags events          | yes                          |
| Merge requests            | yes                          |
| Commit statuses           | yes                          |
| Restrict by organizations | yes                          |

# Next Steps

Once you have configured your Remote Driver, it's time to [Select and 
Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
