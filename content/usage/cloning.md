+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Cloning"
weight = 4
menu = "usage"
toc = true
+++

# Overview

Drone will automatically clone your repository during the build lifecycle. You can customize this behavior in the **clone** section of the yaml file. This section is entirely optional.

This is an example yaml configuration:

```yaml
---
clone:
  depth: 50
  recursive: true
  path: github.com/drone/drone
```

Which results in the following command:

```
git clone --depth=50 --recusive=true \
    https://github.com/drone/drone.git \
    /drone/src/github.com/drone/drone
```

# Authentication

Drone uses `git+https` combined with a `.netrc` file to authenticate and clone private repositories. This allows Drone to clone private Git submodules that also use `git+https`.

# Customization

Limit the amount of git history that is fetched from the remote repository:

```yaml
---
clone:
  depth: 50
```

Fetch all tags from the remote repository:

```yaml
---
clone:
  tags: true
```

Disable tls verification when cloning the remote repository:

```yaml
---
clone:
  skip_verify: true
```

Recursively clone all submodules:

```yaml
---
clone:
  recursive: true
```

Recursively clone and update all submodules:

```yaml
---
clone:
  recursive: true
  submodule_update_remote: true
```

# Private Submodules

Private submodules may encounter the following error:

```
Warning: Permanently added 'github.com,192.30.252.130' (RSA) to the list of known hosts.
ERROR: Repository not found.
fatal: Could not read from remote repository.
```

This happens when a private submodule uses a git+ssh url:

```git
[submodule "hello-world"]
    path = hello-world
    url = git@github.com:octocat/hello-world.git
```

This can be mitigated by overriding the submodule url to use git+https:

```yaml
---
clone:
  recursive: true
  submodule_override:
    hello-world: https://github.com/octocat/hello-world.git
```

Overriding the submodule url to force `git+https` allows us to take advantage of the `netrc` file and automatically authenticate to the submodule repository.
