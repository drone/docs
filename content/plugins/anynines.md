+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Anynines"
description = "Deploy or update a project on Anynines"
user = "drone-plugins"
repo = "drone-anynines"
image = "plugins/drone-anynines"
tags = ["anynines", "paas"]
categories = "deploy"
draft = false
date = 2016-02-13T08:59:49Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Anynines. You can override the
default configuration with the following parameters:

* `username` - Username for Anynines
* `password` - Password for Anynines
* `organization` - Target organization for Anynines
* `space` - Target space for Anynines
* `skip_cleanup` - Flag to deploy from the current file state

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  anynines:
    username: octocat@github.com
    password: password
    organization: octocat_github_com
    space: production
    when:
      branch: master
```

