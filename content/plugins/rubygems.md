+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Rubygems"
description = "Publish ruby gems to a Rubygems server"
user = "drone-plugins"
repo = "drone-rubygems"
image = "plugins/drone-rubygems"
tags = ["ruby", "gem", "rubygems"]
categories = "publish"
draft = false
date = 2016-02-13T09:00:59Z
menu = ""
weight = 1

+++

Use this plugin for publishing gems to a Rubygems server. You can override the
default configuration with the following parameters:

* `api_key` - Rubygems API Key, get it [here](https://rubygems.org/profile/edit)
* `username` - Username, only required without API key
* `password` - Password, only required without API key
* `name` - Name of the gem, defaults to repo name
* `gemspec` - Path to gemspec, defaults to repo name with `.gemspec` suffix
* `host` - Rubygems host, only required for a selfhosted gem server
* `skip_cleanup` - Flag to deploy from the current file state

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  rubygems:
    api_key: e123f83113f79eb17308bbf1ca8885bb
    when:
      branch: master
```

