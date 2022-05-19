---
date: 2000-01-01T00:00:00+00:00
title: Installation
author: bradrydzewski
weight: 1
aliases:
- /cli/install
- /cli-installation/
- /cli/drone-fmt
---

The Drone command line tools are used to interact with the Drone from the command line, and provide important utilities for managing users and repository settings.

# Binary Downloads

Download and install the raw binaries by platform:

Platform	| Download | Verify
------------|----------|-------
Linux x64	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_linux_amd64.tar.gz)   | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)
Linux arm64	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_linux_arm64.tar.gz)   | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)
Linux arm	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_linux_arm.tar.gz)     | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)
Windows x64	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_windows_amd64.tar.gz) | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)
Darwin x64	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_darwin_amd64.tar.gz)  | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)
Darwin arm64	| [tarball](https://github.com/harness/drone-cli/releases/latest/download/drone_darwin_arm64.tar.gz)  | [checksum](https://github.com/harness/drone-cli/releases/latest/download/drone_checksums.txt)

# Install on Linux

* Download and install on Linux:

  ```
  $ curl -L https://github.com/harness/drone-cli/releases/latest/download/drone_linux_amd64.tar.gz | tar zx
  $ sudo install -t /usr/local/bin drone
  ```

# Install on OSX

* Download and install on OSX:

  ```
  $ curl -L https://github.com/harness/drone-cli/releases/latest/download/drone_darwin_amd64.tar.gz | tar zx
  $ sudo cp drone /usr/local/bin
  ```

* Or download and install on OSX using Homebrew:

  ```
  $ brew tap drone/drone
  $ brew install drone
  ```

# Install on Windows

* Download and install on Windows using [scoop](https://scoop.sh/):

  ```
  $ scoop install drone
  ```
