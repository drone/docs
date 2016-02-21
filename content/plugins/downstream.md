+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Trigger"
description = "Trigger downstream repository builds"
user = "drone-plugins"
repo = "drone-downstream"
image = "plugins/drone-downstream"
tags = ["trigger", "drone"]
categories = "clone"
draft = false
date = 2016-02-21T08:35:06Z
menu = ""
weight = 1

+++

Use this plugin to trigger builds for a list of downstream repositories. This is useful when updates to a repository have downstream impacts that should also be tested. These are the configuration options:

* `repositories` - list of repositories to trigger
* `token` - drone API token from your user setttings

The following is a sample configuration in your .drone.yml file:

```yaml
notify:
  downstream:
    repositories:
      - octocat/Hello-World
      - octocat/Spoon-Knife
    token: e3b0c44298fc1c149afbf4
    when:
      event: push
      branch: master
      success: true
```

In some cases you may want to trigger specific branches:

```
notify:
  downstream:
    repositories:
      - octocat/Hello-World@develop
      - octocat/Spoon-Knife@master
```

Please be sure to include the `when` section in your `.drone.yml` to avoid triggering builds for pull requests, tags and failed builds.

You can also trigger a new build for a repository using `fork: true` :

```yaml
notify:
  downstream:
    repositories:
      - octocat/Hello-World
      - octocat/Spoon-Knife
    token: e3b0c44298fc1c149afbf4
    fork: true
    when:
      event: push
      branch: master
      success: true
```

If you want to trigger a drone build on a custom server, you can use the `server` argument. Do not forget to include the protocol for it to work.

```yaml
notify:
  downstream:
    repositories:
      - octocat/Hello-World
      - octocat/Spoon-Knife
    token: e3b0c44298fc1c149afbf4
    fork: true
    server: https://ci.exmaple.com
    when:
      event: push
      branch: master
      success: true
```

