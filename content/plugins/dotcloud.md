+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "dotCloud"
description = "Deploy or update a project on dotCloud"
user = "drone-plugins"
repo = "drone-dotcloud"
image = "plugins/drone-dotcloud"
tags = ["dotcloud"]
categories = "deploy"
draft = false
date = 2016-02-13T09:00:02Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to dotCloud. You can override
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
  dotcloud:
    app: helloworld
    deployment: default
    email: octocat@github.com
    password: my_password
```

