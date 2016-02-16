+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Bitbucket Cloud"
weight = 1
menu = "remotes"
toc = true
+++

Drone comes with built-in support for Bitbucket Cloud. Bitbucket Server is not yet supported. To enable Bitbucket Cloud you should configure the Bitbucket driver using the following environment variables:

```bash
REMOTE_DRIVER=bitbucket
REMOTE_CONFIG=https://bitbucket.org?client_id=${client_id}&client_secret=${client_secret}
```

# Bitbucket configuration

The following is the standard URI connection scheme:

```
scheme://host[?options]
```

The components of this string are:

* `scheme` server protocol `http` or `https`.
* `host` server address to connect to.
* `?options` connection specific options.

# Bitbucket options

This section lists all connection options used in the connection string format. Connection options are pairs in the following form: `name=value`. The value is always case sensitive. Separate options with the ampersand (i.e. &) character:

* `client_id` oauth client id for registered application.
* `client_secret` oauth client secret for registered application.
* `open=false` allows users to self-register. Defaults to false.
* `orgs=drone&orgs=docker` restricts access to these Bitbucket organizations. **Optional**

# Bitbucket registration

You must register your application with Bitbucket in order to generate a Client and Secret. Navigate to your account settings and choose OAuth from the menu, and click Add Consumer.

Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL. You will also need to check the following permissions:

* Account:Email
* Account:Read
* Team Membership:Read
* Repositories:Read
* Webhooks:Read and Write

# Remote Driver Feature Chart

Drone currently only supports BitBucket Cloud (hosted BitBucket). 
Contributions for BitBucket Server are welcomed.
 
Also, we only support git on BitBucket Cloud. While we do have a Mercurial 
Drone Plugin, the webhook output is slightly different for Hg projects. We'd 
love to see pull requests for this as well.

| Feature/Remote            | BitBucket Cloud      |
|---------------------------|----------------------|
| Supported version         | BitBucket Cloud Only |
| VCS                       | **git ( only )**     |
| Auth method               | oauth2               |
| Push events               | yes                  |
| Push tags events          | yes                  |
| Merge requests            | **partially**        |
| Commit statuses           | yes                  |
| Restrict by organizations | yes                  |

# Next Steps

Once you have configured your Remote Driver, it's time to [Select and 
Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
