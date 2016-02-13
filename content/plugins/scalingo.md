+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Scalingo"
description = "Deploy or update a project on Scalingo"
user = "drone-plugins"
repo = "drone-scalingo"
image = "plugins/drone-scalingo"
tags = ["scalingo"]
categories = "deploy"
draft = false
date = 2016-02-13T09:01:36Z
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

