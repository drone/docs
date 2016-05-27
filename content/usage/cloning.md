+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Cloning"
weight = 21
toc = true

[menu.main]
	parent="usage"
+++

# Overview

Drone will automatically clone your repository during the build lifecycle. You can customize this behavior by adding a `clone` step to the pipeline section of your Yaml file. This is entirely optional.

This is an example Yaml configuration:

```yaml
pipeline:
  clone:
    depth: 50
    recursive: true
```

Which results in the following command:

```
git clone --depth=50 --recusive=true https://github.com/drone/drone.git
```

# Authentication

Drone uses `git+https` combined with a `.netrc` file to authenticate and clone private repositories. This allows Drone to clone private repository submodules that also use git+https.

# Customization

Limit the amount of git history that is fetched from the remote repository:

```yaml
pipeline:
  clone:
    depth: 50
```

Fetch all tags from the remote repository:

```yaml
pipeline:
  clone:
    tags: true
```

Disable tls verification when cloning the remote repository:

```yaml
pipeline:
  clone:
    skip_verify: true
```

Recursively clone all submodules:

```yaml
pipeline:
  clone:
    recursive: true
```

Recursively clone and update all submodules:

```yaml
pipeline:
  clone:
    recursive: true
    submodule_update_remote: true
```

Use your own custom clone plugin:

```yaml
pipeline:
  clone:
    image: custom/git
```

# Private Submodules

Private submodules may encounter the following error:

```
Warning: Permanently added 'github.com' (RSA) to the list of known hosts.
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
pipeline:
  clone:
    recursive: true
    submodule_override:
      hello-world: https://github.com/octocat/hello-world.git
```

This gives us the ability to authenticate using the `.netrc` file when cloning private submodules.
