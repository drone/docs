+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Overview"
weight = 1
menu = "installation"
toc = true
+++

# Download

Drone ships as a single binary file and is distributed as a minimalist 20 MB Docker image. Download the official Drone image from DockerHub:

```
sudo docker pull drone/drone:0.4
```

For a full list of available tags, see the [drone/drone](https://hub.docker.com/r/drone/drone/) Docker Hub repo.

# Configure

Create a `/etc/drone/dronerc` file to store configuration variables in `KEY=VALUE` format. Docker will use this file to load environment variables from disk. Please note these variables should never be quoted. For example:

```
REMOTE_DRIVER=github
REMOTE_CONFIG=https://github.com?client_id=....&client_secret=....
```

We'll be adding to this file as we progress through the setup process. 

# Select and Configure a Remote Driver

With your `dronerc` file created, it's time to configure your [Remote Drivers] ({{< relref "remotes.md" >}}). We currently have support for [GitHub]({{< relref "github.md" >}}), [Bitbucket]({{< relref "bitbucket.md" >}}), [GitLab]({{< relref "gitlab.md" >}}), and [Gogs]({{< relref "gogs.md" >}}).

# Select and Configure a Database

Now that you have configured your Remote Driver, you are ready to set up the Database that Drone will use for storing state. See the [Databases]({{< relref "database.md" >}}) article for a walkthrough.

# Create and Run a Drone Container

Create and run your container:

```
sudo docker run \
	--volume /var/lib/drone:/var/lib/drone \
	--volume /var/run/docker.sock:/var/run/docker.sock \
	--env-file /etc/drone/dronerc \
	--restart=always \
	--publish=80:8000 \
	--detach=true \
	--name=drone \
	drone/drone:0.4
```

Please note configuration parameters are provided to Docker using the `--env-file` flag. Any changes to the `dronerc` file require you to stop, remove and re-create the container:

```
sudo docker stop drone
sudo docker rm drone
sudo docker run ...
```

At this point, you should have a running Drone instance. If you are running into issues, read the following *Troubleshooting* and *Common Errors* sections. If that fails, see [Getting Help]({{< relref "community/overview.md#getting-help" >}}).

Assuming everything went well, you are ready to look at our [Usage]({{< relref "usage/overview.md" >}}) documentation and get some builds going.

# Advanced Configuration

* For a full list of configuration options for Drone itself, see [Settings Reference]({{< relref "settings.md" >}}).
* If your Drone instance needs to go through a proxy server, see [Network Proxy]({{< relref "proxy_network.md" >}}).
* If you'd like to park Drone behind a reverse proxy, see [Reverse Proxy]({{< relref "proxy_reverse.md" >}}).

# Troubleshooting

Please use the Drone logs troubleshoot issues:

```
sudo docker logs drone
```

For more verbose logging add `DEBUG` to your `dronerc` file:

```
DEBUG=true
```

# Common Errors

The below error indicates you did not setup your GitHub (or GitLab, Gogs, Bitbucket) integration prior to starting Drone. Please configure version control integration per the above installation instructions.

```
FATA[0000] unknown remote driver
```

The below error indicates you forgot to run Drone with `--volume /var/lib/drone:/var/lib/drone` per the above installation guide. This is important because you need to mount your database on the host machine.

```
INFO[0000] using database driver sqlite3
INFO[0000] using database config /var/lib/drone/drone.sqlite
ERRO[0000] unable to open database file
FATA[0000] migration failed
```
