+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Azure Web Apps"
description = "Deploy or update a project on Azure Web Apps"
user = "drone-plugins"
repo = "drone-azure-web-apps"
image = "plugins/drone-azure-web-apps"
tags = ["azure", "web", "apps"]
categories = "deploy"
draft = false
date = 2016-02-13T08:59:54Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Azure Web Apps. You can override
the default configuration with the following parameters:

* `username` - Username, required for authentication
* `password` - Password, required for authentication
* `site` - Site name, defaults to repo name
* `slot` - Slot name, defaults to value of `site`
* `force` - Force a push, defaults to `false`
* `commit` - Commit local changes, defaults to `false`

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  azure_web_apps:
    username: octocat
    password: my_password
    site: awesome
    force: true
```

