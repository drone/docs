+++
date = "2000-01-01T00:00:00+00:05"
title = "CLI Installation"
url = "cli-installation"

[menu.cli]
  weight = 1
  identifier = "cli-installation"
  parent = "overview"
+++

The Drone command line tools are used to interact with the Drone from the command line, and provides important utilities for managing users and repository settings.

# Install Binaries

Download and install the raw binaries by platform:

Platform    | Download
------------|---------
Linux x64   | [tarball](http://downloads.drone.io/0.6.0/release/linux/amd64/drone.tar.gz), [sha256](http://downloads.drone.io/0.6.0/release/linux/amd64/drone.sha256)
Linux arm64 | [tarball](http://downloads.drone.io/0.6.0/release/linux/arm64/drone.tar.gz), [sha256](http://downloads.drone.io/0.6.0/release/linux/arm64/drone.sha256)
Linux arm   | [tarball](http://downloads.drone.io/0.6.0/release/linux/arm/drone.tar.gz), [sha256](http://downloads.drone.io/0.6.0/release/linux/arm/drone.sha256)
Windows x64 | [tarball](http://downloads.drone.io/0.6.0/release/windows/amd64/drone.tar.gz), [sha256](http://downloads.drone.io/0.6.0/release/windows/amd64/drone.sha256)
Darwin x64  | [tarball](http://downloads.drone.io/0.6.0/release/darwin/amd64/drone.tar.gz), [sha256](http://downloads.drone.io/0.6.0/release/darwin/amd64/drone.sha256)

# Install on Linux

Download and install on Linux:

```nohighlight
curl http://downloads.drone.io/0.6.0/release/linux/amd64/drone.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

# Install on OSX

Download and install on OSX:

```nohighlight
curl http://downloads.drone.io/0.6.0/release/darwin/amd64/drone.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

Download and install on OSX using Homebrew:

```nohighlight
brew tap drone/drone
brew install drone --devel
```
