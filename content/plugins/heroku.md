+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Heroku"
description = "Deploy or update a project on Heroku"
user = "drone-plugins"
repo = "drone-heroku"
image = "plugins/drone-heroku"
tags = ["heroku"]
categories = "deploy"
draft = false
date = 2016-02-13T08:58:53Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Heroku. You can override the
default configuration with the following parameters:

* `app` - Application name, defaults to repo name
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  heroku:
    app: awesomeness
    force: true
```

