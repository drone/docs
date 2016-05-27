+++
date = "2015-12-05T16:00:21-08:00"
title = "drone agents"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++

# Overview

The agents subcommands are available to Drone administrators only. These commands can be used to manage build agents, including monitoring usage and health.

# Agent List

```
$ drone agents

1.2.3.5
Platform: linux/amd64
Capacity: 2 concurrent builds
Pinged: 5m ago
Update: 37d

5.6.7.8
Platform: linux/amd64
Capacity: 2 concurrent builds
Pinged: 2m ago
Update: 10d
```

For scripting purposes you can customize the command output using a Go template:

```
$ drone agents --format "{{ .Address }}"

1.2.3.4
5.6.7.8
```
