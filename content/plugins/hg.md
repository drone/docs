+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Mercurial"
description = "Clones a mercurial repository"
user = "drone-plugins"
repo = "drone-hg"
image = "plugins/drone-hg"
tags = ["scm", "vcs", "hg"]
categories = "clone"
draft = false
date = 2016-02-13T08:58:26Z
menu = ""
weight = 1

+++

Use the Mercurial plugin to clone a mercurial repository. Note that Drone uses this plugin
by default for all Bitbucket mercurial repositories, without any configuration required. You can override
the default configuration with the following parameters:

* `path` relative path inside /drone/src where the repository is cloned

The following is a sample Git clone configuration in your .drone.yml file:

```yaml
clone:
  image: hg
  path: bitbucket.org/foo/bar
```

