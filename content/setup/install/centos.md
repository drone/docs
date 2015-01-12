+++
weight = 3
date = "2014-11-08T12:50:08-08:00"
title = "CentOS"

[menu.main]
parent = "Install"
+++


These are the instructions for running Drone on CentOS / RedHat Linux. Please note that Drone is tested primarily on Ubuntu. CentOS support is still considered experimental.

## Installation

You can install Drone from a pre-built rpm file. Drone will automatically start on port `80`.

```bash
wget downloads.drone.io/master/drone.rpm
sudo yum localinstall drone.rpm
```
