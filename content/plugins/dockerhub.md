+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "DockerHub"
description = "Trigger a DockerHub remote build"
user = "drone-plugins"
repo = "drone-dockerhub"
image = "plugins/drone-dockerhub"
tags = ["docker"]
categories = "publish"
draft = false
date = 2016-02-13T08:59:01Z
menu = ""
weight = 1

+++

Use the DockerHub plugin to trigger a remote build after your build successfully completes.
The following parameters are used to configuration this plugin:

* **token** - json payloads are sent here
* **repo** - messages sent to the above webhook are posted here

The following is a sample DockerHub configuration in your .drone.yml file:

```yaml
publish:
  dockerhub:
    token: be579c82-7c0e-11e4-81c4-0242ac110020
    repo: foo/bar
```

