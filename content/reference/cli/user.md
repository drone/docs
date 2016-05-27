+++
date = "2015-12-05T16:00:21-08:00"
title = "drone user"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++

# Overview

The user subcommands are available to Drone administrators only. These commands can be used to manage registered users, including granting and revoking access to the system.

# User List

Get a list of all registered users:

```
$ drone user ls

jonhsmith
janedoe
octocat
```

# User Info

Get individual account information:

```
$ drone user info octocat

User: octocat
Email: octocat@github.com
```

# User Creation

Add a new user to the system:

```
$ drone user add octocat

User: octocat
Email: octocat@github.com
```


# User Removal

Remove an existing user from the system:

```
$ drone user rm octocat
```
