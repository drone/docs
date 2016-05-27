+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Caching"
weight = 22
toc = true


[menu.main]
	parent="usage"
+++

# Overview

Drone does not have built-in caching capabilities. Instead caching is delegated to plugins. This section of the documentation provides a high level overview of caching and should be used in conjunction with the respective plugin documentation.

Example configuration to cache the `.git` and the `node_modules` directory:

```yaml
pipeline:

  # restore the cache from an sftp server
  sftp_cache:
    server: 1.2.3.4:22
    restore: true
    mount: [ node_modules, .git ]

  build:
    image: node
    commans:
      - npm install
      - npm run build
      - npm run tests

  # rebuild the cache on the sftp server
  sftp_cache:
    server: 1.2.3.4:22
    rebuild: true
    mount: [ node_modules, .git ]
```

# Limitations

Drone cannot cache files or folders that are outside of the base directory in your workspace, which defaults to `/drone`. Each build step executes in a separate docker container and writes to ephemeral disk, which means cache plugins only have access to the shared workspace.

# Common Issues

The most common issue occurs when you try to cache files or folders that are located in your workspace. The below example **will not work**:

```yaml
pipeline:
  sftp_cache:
    server: 1.2.3.4
    restore: true
    mount: /root/node_modules
```

You should instead only cache directories in your workspace:

```yaml
workspace:
  path: /drone/src/github.com/octocat/hello-world
  base: /drone

pipeline:
  sftp_cache:
    server: 1.2.3.4
    restore: true
    mount: /drone/node_modules
```

To work around this limitation, we can cache files or folders in our workspace and copy they to the expected location at the start of the build, and then copy them back to workspace at the end of the build:

```yaml
pipeline:
  sftp_cache:
    server: 1.2.3.4
    restore: true
    mount: /drone/node_modules

  build:
    image: node
    commands:
      - cp -a /drone/node_modules /root
      - rm -rf /drone/node_modules
      - npm -g install
      - npm run build
      - npm run tests
      - cp -a /root/node_modules /drone

  sftp_cache:
    server: 1.2.3.4
    rebuild: true
    mount: /drone/node_modules
```
