+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Remote Drivers"
weight = 2
menu = "installation"
toc = true
+++

Drone uses Remote Drivers to integrate and execute builds against your code hosting system of choice. Some things that the remote driver handles include:

* Deploy key or credential management (to allow Drone to pull code)
* Webhook configuration (so Drone knows when pushes happen)
* Authenticating users against the remote driver (typically through oauth2)

# Features support

While the majority of user and developer scrutiny has been directed at our GitHub driver, the experience with our other supported services is solid. Here is where each of the remote drivers currently stands:

| Feature/Remote            | GitHub  | GitLab  | BitBucket Cloud     | Gogs    |
|---------------------------|---------|---------|---------------------|---------|
| Supported version         | --/--   | 8.2+    | --/--               | --/--   |
| VCS                       | **git** | **git** | **git ( only )**    | **git** |
| Auth method               | ouath2  | oauth2  | oauth2              | manual  |
| Push events               | **yes** | **yes** | **yes**             | **yes** |
| Push tags events          | **yes** | **yes** | **yes**             | no      |
| Merge requests            | **yes** | **yes** | **partially**       | no      |
| Commit statuses           | **yes** | **yes** | **yes**             | no      |
| Restrict by organizations | **yes** | no      | **yes**             | no      |

*partially* - Drone can fetch merge requests direct from source repository, from BitBucket, but if user who activates drone on target repo has no access to source repo, git clone fails

# Configuring a Remote Driver

Depending on which service or system you are hosting your code on, you'll want to refer to the corresponding Remote Driver sub-page for details on configuration:

* [Bitbucket]({{< relref "bitbucket.md" >}})
* [GitHub]({{< relref "github.md" >}})
* [GitLab]({{< relref "gitlab.md" >}})
* [Gogs]({{< relref "gogs.md" >}})

Once you have configured your remote, it's time to [Select and Configure a Database]({{< relref "setup/overview.md#select-and-configure-a-database" >}}).
