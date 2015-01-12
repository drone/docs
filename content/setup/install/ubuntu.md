+++
weight = 1
date = "2014-11-08T12:53:08-08:00"
title = "Ubuntu"

[menu.main]
parent = "Install"
+++

These are the instructions for running Drone on Ubuntu . We recommend using Ubuntu `14.04`, the latest stable distribution. We also highly recommend using `AUFS` as the default file system driver for Docker.

## System Requirements

The default Drone installation uses a SQLite3 database for persistence. Please ensure you have `libsqlite3-dev` installed:

```bash
sudo apt-get update
sudo apt-get install libsqlite3-dev
```

The default Drone installation also assumes Docker is installed locally on the host machine. To install Docker on Ubuntu:

```bash
sudo apt-get update
sudo apt-get install docker.io
```

## Installation

Once the environment is prepared you can install Drone from a debian file. Drone will automatically start on port `80`.

```bash
wget downloads.drone.io/master/drone.deb
sudo install -t drone.deb
```

## Start & Stop

The Drone service is managed by upstart and can be started, stopped or restarted using the following commands:

```bash
sudo start drone
sudo stop drone
sudo restart drone
```

## Logging

The Drone service logs are written to the `/var/log/upstart/drone` directory.