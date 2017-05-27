+++
date = "2000-01-01T00:00:00+00:05"
title = "命令行 安装"
url = "zh/cli-installation"

[menu.cli]
  weight = 1
  identifier = "cli-installation-zh"
  parent = "cli_overview"
+++

<!--The Drone command line tools are used to interact with the Drone from the command line, and provides important utilities for managing users and repository settings.-->

使用 Drone 命令行工具来管理 Drone 用户和仓库设置等多项功能。

<!--# Install Binaries-->

# 安装二进制包

下载和安装这些平台的二进制包：

<!--Download and install the raw binaries by platform:-->

Platform    | Download
------------|---------
Linux x64   | [tarball](https://github.com/drone/drone-cli/releases/)
Linux arm64 | [tarball](https://github.com/drone/drone-cli/releases/)
Linux arm   | [tarball](https://github.com/drone/drone-cli/releases/)
Windows x64 | [tarball](https://github.com/drone/drone-cli/releases/)
Darwin x64  | [tarball](https://github.com/drone/drone-cli/releases/)

<!--# Install on Linux-->

# Linux 安装

<!--Download and install on Linux:-->

访问 https://github.com/drone/drone-cli/releases/ 了解最新版本号，可替换 v0.7.0 为最新版本。

```nohighlight
curl -L https://github.com/drone/drone-cli/releases/download/v0.7.0/drone_linux_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

<!--# Install on OSX-->

# OSX 安装

<!--Download and install on OSX:-->

访问 https://github.com/drone/drone-cli/releases/ 了解最新版本号，可替换 v0.7.0 为最新版本。

```nohighlight
curl -L https://github.com/drone/drone-cli/releases/download/v0.7.0/drone_darwin_amd64.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

<!--Download and install on OSX using Homebrew:-->

使用 Homebrew 安装

```nohighlight
brew tap drone/drone
brew install drone
```
