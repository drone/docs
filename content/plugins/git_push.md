+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Git Push"
description = "Deploy or update a project via Git"
user = "drone-plugins"
repo = "drone-git-push"
image = "plugins/drone-git-push"
tags = ["git"]
categories = "deploy"
draft = false
date = 2016-02-13T08:58:57Z
menu = ""
weight = 1

+++

> This plugin has not been fully tested. Proceed with caution.

Use this plugin for deplying an application via `git push`. You can override
the default configuration with the following parameters:

* `remote` - Target remote repository
* `branch` - Target remote branch, defaults to master
* `force` - Force push using the `--force` flag, defaults to false
* `skip_verify` - Skip verification of HTTPS certs, defaults to false

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  git_push:
    branch: master
    remote: git@git.heroku.com:falling-wind-1624.git
    force: false
```

