+++
date = "2015-12-05T16:00:21-08:00"
title = "Server"
weight = 1
menu = "installation"
toc = true
+++

# Overview

Drone ships as a single binary file inside a minimalist 20 MB Docker image. Docker is the only dependency. We recommend using Drone with the latest stable version of Docker and officially support Docker 1.N-1 (latest minor release minus one).

# Installation

Get started by downloading the Docker image:

```
docker pull drone/drone:0.5
```

Create an `/etc/dronerc` file to hold your configuration parameters. Please read the documentation to configure your remote system (ie GitHub) before completing this step. Please also note that parameter values must not be quoted.

```
DRONE_GITHUB_CLIENT=...
DRONE_GITHUB_SECRET=...
```

Create a secret passcode used to generate authentication tokens for your builds agents. We recommend using a [random password generator](http://correcthorsebatterystaple.net/) to create this secret passcode.

```
DRONE_AGENT_SECRET=...
```

Create and run your container:

```
docker run \
	--volume /var/lib/drone:/var/lib/drone \
	--env-file /etc/dronerc \
	--restart=always \
	--publish=80:8000 \
	--detach=true \
	--name=drone \
	drone/drone:0.5
```

Note the above example mounts a volume on the host machine. The default configuration uses a sqlite database and should therefore be mounted on the host machine as a volume to avoid data loss.

```
--volume /var/lib/drone:/var/lib/drone
```

Note the above example uses the `--env-file` flag to provide Docker with configuration parameters. Changes to this file require you to stop, remove and re-create the container:

```
sudo docker stop drone
sudo docker rm drone
sudo docker run ...
```

# Troubleshooting

Please use the Drone logs troubleshoot issues:

```
docker logs drone
```

For more verbose logging enable debugging in your `dronerc` file:

```
DRONE_DEBUG=true
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
