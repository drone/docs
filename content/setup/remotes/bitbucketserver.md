+++
date = "2016-04-20T16:00:21-08:00"
draft = false
title = "Bitbucket Server(Stash)"
weight = 1
menu = "remotes"
toc = true
+++

Drone comes with built-in support for Bitbucket Server(stash) To enable Bitbucket Server you should configure the Bitbucket driver using the following environment variables:

```bash
REMOTE_DRIVER=bitbucketserver
REMOTE_CONFIG=https://{servername}?consumer_key={consumer_key}&git_username={git_username}&git_password={git_password}&consumer_rsa={path_private_key}&open={unused}
```

# Bitbucket Server configuration

The following is the standard URI connection scheme:

```
scheme://host[?options]
```

The components of this string are:

* `scheme` server protocol `http` or `https`.
* `host` server address to connect to.
* `?options` connection specific options.

# Bitbucket Server options

This section lists all connection options used in the connection string format. Connection options are pairs in the following form: `name=value`. The value is always case sensitive. Separate options with the ampersand (i.e. &) character:

* `consumer_key` oauth client id for registered application (must add in application links incoming auth in Bitbucket server).
* `git_username` Service account that can http clone the projects.
* `git_password` password for the service account that can http clone the projects.
* `consumer_rsa` Path to the private key for oauth1 (will need to copy into a customer docker image or map the volume to the correct location).
* `open=false` not currently used but will be added later

# Bitbucket Server registration

You must register your application with Bitbucket in order to allow for auth1. Navigate to admin choose application links and add a new application link

Please use `http://drone.mycompany.com/authorize` as the Authorization callback URL in the incoming auth(at the bottom):

* Add the url to drone on the application path
* You'll also need to add the public key (public key to the private key in consumer_rsa_key file)



# Remote Driver Feature Chart

Drone currently only supports BitBucket Cloud (hosted BitBucket).
Contributions for BitBucket Server are welcomed.

Also, we only support git on BitBucket Cloud. While we do have a Mercurial
Drone Plugin, the webhook output is slightly different for Hg projects. We'd
love to see pull requests for this as well.

| Feature/Remote            | BitBucket Server     |
|---------------------------|----------------------|
| Supported version         | 4.5 (ymwv on others) |
| VCS                       | **git ( only )**     |
| Auth method               | oauth1               |
| Push events               | yes                  |
| Push tags events          | no                   |
| Merge requests            | no                   |
| Commit statuses           | yes                  |
| Restrict by organizations | no                   |

# Next Steps

Once you have configured your Remote Driver, it's time to [Select and
Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
