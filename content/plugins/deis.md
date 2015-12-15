+++
title = "Deis"
description = "Deploys or updates a project on Deis"
user = "drone-plugins"
repo = "drone-deis"
image = "plugins/drone-deis"
tags = ["deis"]
categories = "deploy"
draft = false
date = 2015-12-15T06:52:33Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Deis. You can override the
default configuration with the following parameters:

* `controller` - Deis controller, e.g. deis.deisapps.com
* `app` - Application name, defaults to repo name
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  deis:
    controller: deis.deisapps.com
    app: awesomeness
    force: true
```

