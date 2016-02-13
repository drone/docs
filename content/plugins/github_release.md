+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "GitHub Release"
description = "Publish files and artifacts to GitHub Releases"
user = "drone-plugins"
repo = "drone-github-release"
image = "plugins/drone-github-release"
tags = ["github", "release"]
categories = "publish"
draft = false
date = 2016-02-13T09:00:08Z
menu = ""
weight = 1

+++

Use this  plugin for publishing files and artifacts to GitHub releases. You
can override the default configuration with the following parameters:

* `api_key` - GitHub oauth token with public_repo or repo permission
* `files` - Files to upload to GitHub Release, globs are allowed
* `file_exists` - What to do if an file asset already exists, supported values: **overwrite** (default), **skip** and **fail**
* `checksum` - Checksum takes hash methods to include in your GitHub release for the files specified. Supported hash methods include md5, sha1, sha256, sha512, adler32, and crc32.
* `draft` - create a draft release if set to true
* `base_url` - GitHub base URL, only required for GHE
* `upload_url` - GitHub upload URL, only required for GHE

Sample configuration:

```yaml
publish:
  github_release:
    api_key: my_github_api_key
    files: dist/*
    checksum: sha1
```

or

```yaml
publish:
  github_release:
    api_key: my_github_api_key
    files:
      - dist/*
      - bin/binary.exe
    checksum:
      - md5
      - sha1
      - sha256
      - sha512
      - adler32
      - crc32
```

