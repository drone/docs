+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Artifacts"
weight = 11
menu = "usage"
toc = true
+++


# Overview

Publishing files or build artifacts is triggered at the end of a successful build. Note that publish steps are not executed for pull requests or failed builds.

Example to publish a Node package to the [npmjs.com](https://npmjs.com) registry using the NPM plugin:

```yaml
---
publish:
  npm:
    username: octocat
    password: password
    email: octocat@github.com
```

# Plugins

Drone does not have any builtin publishing or artifact capabilities. Drone instead offers a large library of plugins for publish to [DockerHub](#), [S3](#), [Goolge Storage](#) and more. Please see the [plugin documentation](#) for a complete list.

# Docker

Building and publishing a Docker image to the registry is handled by plugins. We have multiple plugins for building and publishing Docker images to [DockerHub](#), [Google Container Registry](#), and more.

# Conditions

Use the `when` block to limit publish step execution:

```yaml
---
publish:
  heroku:
    app: mycompany.com
    when:
      branch: master
```

Execute a publish step if the branch is `master` or `release`:

```yaml
---
publish:
  rubygems:
    api_key: e123f83113f79eb17308bbf1ca8885bb
    when:
      branch: [master, release]
```

Execute a publish step when the commit is a `tag`:

```yaml
---
publish:
  rubygems:
    api_key: e123f83113f79eb17308bbf1ca8885bb
    when:
      event: tag
```

Execute a publish step for a single matrix permutation:

```yaml
---
publish:
  rubygems:
    api_key: e123f83113f79eb17308bbf1ca8885bb
    when:
      matrix:
        RUBY_VERSION: 2.3
        REDIS_VERSION: 2.8
```
