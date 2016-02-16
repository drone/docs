+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Remote Drivers"
weight = 2
menu = "installation"
#toc = true
+++

Drone uses Remote Drivers to integrate and execute builds against your code hosting system of choice. Some things that the remote driver handles include:

* Deploy key or credential management (to allow Drone to pull code)
* Webhook configuration (so Drone knows when pushes happen)
* Authenticating users against the remote driver (typically through oauth2)

# Configuring a Remote Driver

Depending on which service or system you are hosting your code on, you'll want to refer to the corresponding Remote Driver sub-page for details on configuration:

* [Bitbucket]({{< relref "bitbucket.md" >}})
* [GitHub]({{< relref "github.md" >}})
* [GitLab]({{< relref "gitlab.md" >}})
* [Gogs]({{< relref "gogs.md" >}})

# Next Steps

Once you have configured your Remote Driver, it's time to [Select and 
Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
