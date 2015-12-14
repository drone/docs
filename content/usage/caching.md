+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Caching"
weight = 5
menu = "usage"
toc = true
+++

# Overview

Drone allows you to cache directories within the build workspace (inside the `/drone` volume). When a build successfully completes, the named directories are gzipped and stored on the host machine. When a new build starts, the named directories are restored from the gzipped files. This can be used to improve the performance of your builds.

Example configuration to cache the `.git` and the `node_modules` directory:

```yaml
---
cache:
  mount:
    - node_modules
    - .git
```

# Branches and Matrix

Drone keeps a separate cache for each Branch and Matrix combination. Let’s say, for example, you are using matrix builds to test node `0.11.x` and node `0.12.x` and you are caching `node_modules`. Drone will separately cache `node_modules` for each version of node.

# Pull Requests

Pull requests have read-only access to the cache. This means pull requests are not permitted to re-build the cache. This is done for security and stability purposes, to prevent a pull request from corrupting your cache.


# Deleting the Cache

There is currently no mechanism to automatically delete or flush the cache. This must be done manually, on each worker node. The cache is located in `/var/lib/drone/cache/`.

# Distributed Cache

This is outside the scope of Drone. You may, for example, use a distributed filesystem such as `ceph` or `gluster` mounted to `/var/lib/drone/cache/` to share the cache across nodes.

# Common Issues

The most common issue occurs when you try to cache files or folders that are not children of the `/drone` volume. The below example **will not work**:

```yaml
---
cache:
  mount:
    - /root/node_modules
    - /usr/lib/python3.2/site-packages
```

You should instead only cache directories in `/drone`:


```yaml
---
cache:
  mount:
    - /drone/node_modules
    - /drone/site-packages
```

To work around this limitation we can cache files or folders in `/drone` and copy they to the expected location at the start of the build, and then copy them back to `/drone` at the end fo the build:

```yaml
---
build:
  image: node
  commands:
    - cp -a /drone/node_modules /root
    - rm -rf /drone/node_modules
    - npm -g install
    - npm run build
    - npm run tests
    - cp -a /root/node_modules /drone
```
