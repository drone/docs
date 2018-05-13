+++
date = "2000-01-01T00:00:00+00:05"
title = "CLI Installation"
url = "cli-installation"

[menu.cli]
  weight = 1
  identifier = "cli-installation"
  parent = "cli_overview"
+++

The Drone command line tools are used to interact with the Drone from the command line, and provides important utilities for managing users and repository settings.

# Install Binaries

Download and install the raw binaries by platform:

Platform    | Download
------------|---------
Linux x64   | [tarball](https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_linux_amd64.tar.gz)
Linux arm64 | [tarball](https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_linux_arm64.tar.gz)
Linux arm   | [tarball](https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_linux_arm.tar.gz)
Windows x64 | [tarball](https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_windows_amd64.tar.gz)
Darwin x64  | [tarball](https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_darwin_amd64.tar.gz)

# Install on Windows

Download and install on Windows using [scoop](http://scoop.sh):

```nohighlight
scoop install drone
```

# Install on Linux

Download and install on Linux:

```nohighlight
curl -L https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_linux_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

# Install on OSX

Download and install on OSX:

```nohighlight
curl -L https://github.com/drone/drone-cli/releases/download/v0.8.5/drone_darwin_amd64.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

Download and install on OSX using Homebrew:

```nohighlight
brew tap drone/drone
brew install drone
```
