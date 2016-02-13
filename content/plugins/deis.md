+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Deis"
description = "Deploy or update a project on Deis"
user = "drone-plugins"
repo = "drone-deis"
image = "plugins/drone-deis"
tags = ["deis"]
categories = "deploy"
draft = false
date = 2016-02-13T09:00:18Z
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

