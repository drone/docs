+++
date = "2015-12-05T16:00:21-08:00"
title = "drone repo"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++


# Overview

The repo subcommands are available to authenticated users. These commands can be used to manage active repositories and view repository information.

# Repo List

Get a list of all activated repositories:

```
$ drone repo ls

octocat/hello-world
octocat/spoon
```

# Repo Info

Get individual repository information:

```
$ drone repo info octocat/hello-world

Owner: octocat
Repo: hello-world
Type: git
Private: false
Remote: https://github.com/octocat/hello-world.git
```

# Repo Activate

Activates the repository and register post-commit hooks with the remote system.

```
$ drone repo add octocat/hello-world

Owner: octocat
Repo: hello-world
Type: git
Private: false
Remote: https://github.com/octocat/hello-world.git
```


# Repo Remove

Remove the repository and cleanup post-commit hooks registered with the remote system.

```
$ drone repo rm octocat/hello-world
```
