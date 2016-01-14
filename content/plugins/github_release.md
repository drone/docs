+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "GitHub Release"
description = "Publishs files and artifacts to GitHub Releases"
user = "drone-plugins"
repo = "drone-github-release"
image = "plugins/drone-github-release"
tags = ["github", "release"]
categories = "publish"
draft = false
date = 2016-01-14T18:45:49Z
menu = ""
weight = 1

+++

Use this  plugin for publishing files and artifacts to GitHub releases. You
can override the default configuration with the following parameters:

* `api_key` - GitHub oauth token with public_repo or repo permission
* `files` - Files to upload to GitHub Release, globs are allowed
* `base_url` - GitHub base URL, only required for GHE
* `upload_url` - GitHub upload URL, only required for GHE

Sample configuration:

```yaml
publish:
  github_release:
    api_key: my_github_api_key
    files:
      - dist/*
```

