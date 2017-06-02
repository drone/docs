---
date: 2017-04-15T14:39:04+02:00
title: "Drone Is Broken"
url: drone-is-broken
draft: true
---



# Yaml File

If your build is failing please provide the full yaml configuration file. It will not be possible to effectively debug issues if we cannot see the commands and configuration being used.

# Build Logs

If your build is failing please provide a sample of the logs and include the error message. It will not be possible to effectively debug issues if we cannot see the error message.

# Drone Logs



```
docker logs <drone server>
docker logs <drone agent>
```

# Drone Version

Please provide your Drone version number.


```
$ docker images | grep drone
drone/drone   0.6    de54e7f05ec6   2 months ago   6.61 MB
```
