+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "BitBalloon"
description = "Deploy or update a project on BitBalloon"
user = "drone-plugins"
repo = "drone-bitballoon"
image = "plugins/drone-bitballoon"
tags = ["bitballoon", "html5"]
categories = "deploy"
draft = false
date = 2016-02-13T09:00:29Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Bitballoon. You can override the
default configuration with the following parameters:

* `access_token` - BitBalloon API access token
* `site` - Site name, should be the full domain name
* `path` - Relative path to the folder or ZIP file, defaults to repo root
* `draft` - Deploy as a draft version, defaults to `false`

## Example

```yaml
deploy:
  bitballoon:
    access_token: 6f34aed19d59420e3f026ab0f04fb3824710e36d26df4dc8aac8d56dc82a06ac
    site: happy-page.bitballoon.com
    path: dist/
    draft: true
```

