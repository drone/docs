+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Git"
description = "Clones a git repository"
user = "drone-plugins"
repo = "drone-git"
image = "plugins/drone-git"
tags = ["scm", "vcs", "git"]
categories = "clone"
draft = false
date = 2016-02-13T08:58:07Z
menu = ""
weight = 1

+++

Use the Git plugin to clone a git repository. Note that Drone uses the Git plugin
by default for all repositories, without any configuration required. You can override
the default configuration with the following parameters:

* `depth` - creates a shallow clone with truncated history
* `recursive` - recursively clones git submodules
* `path` - relative path inside `/drone/src` where the repository is cloned
* `skip_verify` - disables ssl verification when set to `true`
* `tags` - fetches tags when set to `true`
* `submodule_override` - override submodule urls
* `submodule_update_remote` - pass the `--remote` flag to git submodule update (useful when tracking a branch instead of a commit in a submodule)

Sample configuration:

```yaml
clone:
  depth: 50
  recursive: false
  tags: false
```

Sample configuration to override the relative clone path:

```
# clones to /drone/src/foo/bar

clone:
  path: foo/bar
```

Sample configuration to clone submodules:

```
clone:
  recursive: true
```

## Private Submodules

Private submodules may encounter the following error:

```
Warning: Permanently added 'github.com,192.30.252.130' (RSA) to the list of known hosts.
ERROR: Repository not found.
fatal: Could not read from remote repository.
```

This happens when a private submodule uses a `git+ssh` url:

```
[submodule "hello-world"]
    path = hello-world
    url = git@github.com:octocat/hello-world.git
```

This can be mitigated by overriding the submodule url to use `git+https`:

```
clone:
  recursive: true
  submodule_override:
    hello-world: https://github.com/octocat/hello-world.git
```

Overriding the submodule url to force `git+https` allows us to take advantage of the `netrc` file and automatically authenticate to the submodule repository.

