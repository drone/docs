+++
date = "2015-12-05T16:00:21-08:00"
title = "CLI overview"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++

# Overview

Drone provides a comprehensive command line utility so you can easily work with Drone from your terminal or shell scripts. This section of the documents provides a brief overview to the command line utility, including installation and setup instructions.

# Install on Linux

Download and install the x64 linux binary:

```
curl http://downloads.drone.io/release/linux/amd64/drone.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

# Install on Mac

Download and install using Homebrew:

```
brew tap drone/drone
brew install drone
```

Or manually download and install the binary:

```
curl http://downloads.drone.io/release/darwin/amd64/drone.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

# Install on Windows

Download and install the x64 windows binary:

```
curl http://downloads.drone.io/release/windows/amd64/drone.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

# Authentication

Drone requires the server address and token for authorization purposes. Your token can be found in your Drone profile screen. You can export these values as environment variables and add to your `.profile` or `.bashrc` file:

```
export DRONE_TOKEN=...
export DRONE_SERVER=...
```
