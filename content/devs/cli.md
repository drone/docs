+++
title = "CLI Reference"
menu = "developers"
weight = 4
toc = true
+++

# Install

This tool requires Docker 1.6 or higher. If you are using Windows or Mac we recommend installing the [Docker Toolbox](https://www.docker.com/docker-toolbox).

__Linux__ binary installation instructions:

```bash
curl http://downloads.drone.io/drone-cli/drone_linux_amd64.tar.gz | tar zx
sudo install -t /usr/local/bin drone
```

__OSX__ binary installation instructions:

```bash
curl http://downloads.drone.io/drone-cli/drone_darwin_amd64.tar.gz | tar zx
sudo cp drone /usr/local/bin
```

# Setup

You will need to provide the Drone command line tools with the location of your Drone installation as well as an access token for authentication. You can retrieve a Drone personal access token from your user profile screen.

```bash
export DRONE_SERVER=<http://>
export DRONE_TOKEN=<token>
```

# Repo Commands

The `drone repo` commands help you manage repositories, including activation and de-activation. Run `drone build help` for the usage documentation:


```
NAME:
   drone repo - manage repos

USAGE:
   drone repo [global options] command [command options] [arguments...]

COMMANDS:
   ls		lists repositories
   info		show repository details
   add		add a repository
   rm		remove a repository
   help, h	Shows a list of commands or help for one command
```

Activates a repository:

```
drone repo add hello-world/octocat
```

Remove a repository:

```
drone repo rm hello-world/octocat
```

List all activated repositories:

```
drone repo ls
```

# Build Commands

The `drone build` commands help you manage builds, including starting, stopping and inspecting running builds. This could be used with CRON to schedule nighty builds, for example. Run `drone build help` for the usage documentation:

```
NAME:
   drone build - manage builds

USAGE:
   drone build [global options] command [command options] [arguments...]

COMMANDS:
   info		show build information
   list		list recent builds
   last		last build (push only)
   logs		show build logs
   start	start a stopped build
   stop		stop a running build job
   help, h	Shows a list of commands or help for one command
```

List recent builds for a repository:

```
drone build list octocat/hello-world
```

Get the last build for a repository:

```
drone build last octocat/hello-world
```

Restart a previously executed build by build number:

```
drone build start octocat/hello-world 35
```

# User Commands

The `drone user` commands help the system administrator manage registered users, including adding and removing user access. Run `drone user help` for the usage documentation:

```
NAME:
   drone user - manage users

USAGE:
   drone user [global options] command [command options] [arguments...]

COMMANDS:
   ls		list all users
   info		show user details
   add		adds a user
   rm		remove a user
   self		show the current user details
   help, h	Shows a list of commands or help for one command
```

Setup a new user by login:

```
drone user add octocat
```

Remove a user by login:

```
drone user rm octocat
```

List all registered users:

```
drone user ls
```

# Worker Commands

TODO

# Secret Management

TODO

# Local Testing

Drone gives you the ability to run your build locally, on your personal computer, without a Drone server. Local builds use Docker for build isolation and container orchestration. The local build process is almost identical to the server build process, making it very convenient for setup or troubleshooting failed builds.

First ensure you have a running Docker instance. Navigate to the root of your git repository, where the `.drone.yml` file is located, and run the following command:

```
drone exec
```

Run your build with the SSH key injected into your build environment. This may be useful if your build script fetches private repositories or connects to remote servers:

```
drone exec -i ~/.ssh/id_rsa
```

Run your build with the trusted flag to enable privileged mode and volume mounting:

```
drone exec --trusted
```
