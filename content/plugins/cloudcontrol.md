+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "cloudControl"
description = "Deploy or update a project on cloudControl"
user = "drone-plugins"
repo = "drone-cloudcontrol"
image = "plugins/drone-cloudcontrol"
tags = ["cloud", "control"]
categories = "deploy"
draft = false
date = 2016-02-13T09:00:48Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to cloudControl. You can override
the default configuration with the following parameters:

* `app` - Application name, defaults to repo name
* `deployment` - Deployment name, defaults to `default`
* `email` - Email address to authenticate
* `password` - Password to authenticate
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

```yaml
deploy:
  cloudcontrol:
    app: helloworld
    deployment: default
    email: octocat@github.com
    password: my_password
```

