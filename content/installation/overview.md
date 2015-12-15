+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Overview"
weight = 1
menu = "installation"
toc = true
+++

# Installation

Drone ships as a single binary file and is distributed as a minimalist 20 MB Docker image. Download the official Drone image from DockerHub:

```
sudo docker pull drone/drone:0.4
```

# Configuration

Create a `/etc/lgtm/dronerc` file to hold your configuration parameters in `KEY=VALUE` format. These variables should not be quoted:

```
REMOTE_DRIVER=github
REMOTE_CONFIG=https://github.com?client_id=....&client_secret=....
```

Please configure your [GitHub](../github) integration (or GitLab, Bitbucket, Gogs) before your proceed. Drone will not start until you have configured your version control integration.

# Create and Run

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

Notice we send your Drone configuration to Docker as an environment variable file using the `--env-file` flag. Any changes to the `dronerc` file require you to stop, remove and re-create the drone container.

```
--env-file /etc/drone/dronerc
```

# Troubleshooting

Please use the Drone logs troubleshoot issues:

```
sudo docker logs drone
```

For more verbose logging add `DEBUG` to your `dronerc` file:

```
DEBUG=true
```

# Common Issues

The below error indicates you did not setup your GitHub (or GitLab, Gogs, Bitbucket) integration prior to starting Drone. Please configure version control integration per the above installation instructions.

```
FATA[0000] unknown remote driver
```

The below error indicates you forgot run Drone with `--volume /var/lib/drone:/var/lib/drone` per the above installation guide. This is important because you need to mount your database on the host machine. 

```
INFO[0000] using database driver sqlite3
INFO[0000] using database config /var/lib/drone/drone.sqlite
ERRO[0000] unable to open database file
FATA[0000] migration failed
```
