+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Git Push"
description = "Deploys or updates a project via Git"
weight = 1
menu = "plugins"
repo = "drone-git"
categories = ["publish", "deploy"]
tags = "git"
+++

Use this plugin for deploying an application via `git push`. You can override
the default configuration with the following parameters:

* `remote` - Target remote repository
* `branch` - Target remote branch, defaults to master
* `force` - Force push using the `--force` flag, defaults to false
* `skip_verify` - Skip verification of HTTPS certs, defaults to false

# Example

The following is a sample configuration in your .drone.yml file:

```yaml
---
deploy:
  git_push:
    branch: master
    remote: git@git.heroku.com:falling-wind-1624.git
    force: false
```
