+++
title = "Scalingo"
description = "Deploys or updates a project on Scalingo"
user = "drone-plugins"
repo = "drone-scalingo"
image = "plugins/drone-scalingo"
tags = ["scalingo"]
categories = "deploy"
draft = false
date = 2015-12-15T06:53:30Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Scalingo. You can override the
default configuration with the following parameters:

* `app` - Application name, defaults to repo name
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  scalingo:
    app: awesomeness
    force: true
```

