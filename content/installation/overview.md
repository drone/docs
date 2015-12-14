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

# Running

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


The above example mounts a volume on the host machine. This is important because the default configuration uses a sqlite database which must be mounted on the host machine to avoid data loss.

```
--volume /var/lib/drone:/var/lib/drone
```

The above example sends your configuration to Docker using the `--env-file` flag:

```
--env-file /etc/drone/dronerc
```

Any changes to the above file require you to stop, remove and re-create the drone container.

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

This error indicates you did not setup your GitHub (or GitLab, Gogs, Bitbucket) integration prior to starting Drone. Please configure version control integration per the above installation instructions:

```
FATA[0000] unknown remote driver
```
