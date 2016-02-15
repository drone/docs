+++
date = "2016-02-05T18:20:44-06:00"
draft = false
title = "Settings Reference"
weight = 4
menu = "installation"
toc = true
+++

Unless otherwise specified, the settings described below are all environment variables.

# Remote Driver Settings

## REMOTE_DRIVER

Set this to whichever remote driver you're using. Drone will use your remote for authentication, and will add webhooks to your projects to facilitate the build process. See [Remote Drivers]({{< relref "remotes.md" >}}) for possible values.

## REMOTE_CONFIG

The value of this setting is highly dependent upon the remote driver chosen. See [Remote Drivers]({{< relref "remotes.md" >}}) for more details.

# Database Settings

See the [Database]({{< relref "database.md" >}}) page for more details on selecting and configuring a database engine for Drone.

## DATABASE_DRIVER

One of the following values:

* `sqlite3`
* `postgres`
* `mysql`

## DATABASE_CONFIG

Differs based on the selected database engine. See See the [Database]({{< relref "database.md" >}}) for a list of possible values for each DB.

# Plugin Settings

## PLUGIN_FILTER

Default: `plugins/*`

This setting contains a space-separated list of patterns for determining which plugins can be pulled an ran. By disallowing open season on plugins, we prevent some classes of malicious plugins sneaking into builds.

To use an example, let's say you've created an organization on Docker
Hub called `myorg`. Here's a REMOTE_DRIVER setting value that would allow the standard Drone plugins plus your org:

```
PLUGIN_FILTER=plugins/* myorg/*
```

Or perhaps your org is on Quay instead of Docker Hub:

```
PLUGIN_FILTER=plugins/* quay.io/myorg/*
```

## PLUGIN_PARAMS

This setting may be used to define some global parameters to pass into all plugins on all repositories. 

> Be careful what you put in here, since it will be trivial for a malicious developer to obtain the values found in this setting.

Key/value pairs are separated by spaces:

```
PLUGIN_PARAMS=MYSETTING=value1 ANOTHER=value2
```

# General Drone Settings

## PUBLIC_MODE

Drone restricts the build history to authenticated users with pull access to the source repository. In some deployments it may be desirable to make the build history visible to all users on the network. Public mode permits unauthenticated users to access the build history for every repository, _including private repositories_.

Add the following configuration to `/etc/drone/dronerc` to enable public mode:

```
PUBLIC_MODE=true
```

Enabling this feature on a public network will leak sensitive information in the build logs.
