+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Command Line Tools"
weight = 4
menu = "main"
toc = true
+++

# Installation

Use the command line tools to manage `MAINTAINERS` files from the terminal. This is especially useful if you would like to manage `MAINTAINERS` files in batch.

Install on OSX:

```
$ curl -L https://github.com/lgtmco/lgtm-cli/releases/download/v1.0.0/lgtm_darwin_amd64.tar.gz | tar zx
$ sudo cp lgtm /usr/local/bin
```

Install on Linux:

```
$ curl -L https://github.com/lgtmco/lgtm-cli/releases/download/v1.0.0/lgtm_linux_amd64.tar.gz | tar zx
$ sudo install -t /usr/local/bin drone
```

# Authentication

You will need to provide the command line client with an access token.

```
$ export LGTM_TOKEN=f0e4c2f76c58916ec258f246851bea
```

You can exchange a [GitHub personal token](https://github.com/settings/tokens) for an access token:

```
$ export GITHUB_TOKEN=f0e4c2f76c58916ec258f246851bea
$ export LGTM_TOKEN=$(lgtm token $GITHUB_TOKEN)
```

# Basic Usage

Activate or deactivate a repository:

```
$ lgtm add octocat/hello-world
$ lgtm rm octocat/hello-world
```

List all activated repositories:

```
$ lgtm ls
```

List all inactive repositories:

```
$ lgtm ls --inactive
```

Get an active repository's `MAINTAINERS` file:

```
$ lgtm get octocat/hello-world
```

# Batch Usage

Activate all repositories for an organization:

```
$ lgtm ls --inactive | grep octocat/ | xargs lgtm add
```

Copy a `MAINTAINERS` file from one repository to another:

```
$ export GITHUB_TOKEN=f0e4c2f76c58916ec258f246851bea
$ lgtm get octocat/helloworld plugins | lgtm push octocat/Spoon-Knife
```

Push a central `MAINTAINERS` file on your local machine to multiple repositories.

```
$ export GITHUB_TOKEN=f0e4c2f76c58916ec258f246851bea

$ lgtm get octocat/hello-world > /tmp/MAINTAINERS
$ lgtm ls --exclude=octocat/hello-world | xargs -I{} lgtm push {} /tmp/MAINTAINERS
```
