+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "GitHub Pages"
description = "Copies a directory to GitHub Pages"
user = "drone-plugins"
repo = "drone-gh-pages"
image = "plugins/drone-gh-pages"
tags = ["github", "pages"]
categories = "publish"
draft = false
date = 2016-07-30T23:50:02Z
menu = ""
weight = 1

+++

Use the gh-pages plugin to build and push a directory from your build to your repo's GitHub Pages.
The following parameters are used to configure this plugin:

* `remote` - name of the remote repository to push changes to (default origin)
* `source` - the directory of your build that should become your pages content (default docs)
* `temp` - the temporary directory used to store the gh-pages clone (default .tmp)
* `branch` - the branch to publish to (default gh-pages)

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  gh_pages:
    source: pages
    remote: origin
```

## Known Issues

At this time, this plugin will not allow you to create the gh-pages branch if it does not exist.

