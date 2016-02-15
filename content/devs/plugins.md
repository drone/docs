+++
title = "Plugins"
menu = "developers"
weight = 2
toc = true
+++

# Overview

Plugins are Docker containers that attach to your build at runtime and perform custom operations. These operations may include deploying code, publishing artifacts and sending notifications. __BONUS__ because plugins are Docker containers they can be written in any language!

Plugins are declared in the `.drone.yml` file. Drone automatically downloads plugins from the registry and executes. This is an example of the [slack](../../plugins/slack) plugin configuration:

```yaml
---
notify:
  slack:
    webhook_url: https://hooks.slack.com/services/...
    username: captain_freedom
    channel: ics
```

# Plugin Input

Plugins receive a JSON payload as a command line argument that includes repository and build details. Plugin configuration parameters from the YAML file are stored in the `vargs` section of the document.

Example plugin input:

```js
{
    "system": {
        "link": "http://drone.mycompany.com"
    },
    "repo": {
        "owner": "octocat",
        "name": "hello-world",
        "full_name": "octocat/hello-world",
        "link_url": "https://github.com/octocat/hello-world"
        "clone_url": "https://github.com/octocat/hello-world.git",
    },
    "build": {
        "number": 1,
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master",
        "author": "octocat",
        "author_email": "octocat@github.com"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/octocat/hello-world",
        "keys": {
            "private": "-----BEGIN RSA PRIVATE KEY-----\nMIICXAIBAAKBgQC..."
        }
    },
    "vargs": { ... }
}
```

# Plugin Output

Plugins cannot send structured data back to Drone. Plugins can, however, write information to stdout that appears in your build logs. Plugins can fail a build by exiting with a non-zero status code.

Drone does not provide a storage mechanism for metrics or build artifacts. Instead you should write plugins that integrate with third party metrics services, like [Hound](https://houndci.com) or [Coveralls](https://coveralls.io), or storage services like [S3](../../s3).

# Disk Access

Plugins share the `/drone` container volume with your build environment. This means plugins have full access to your source code as well as any assets generated in your build workspace.

Please note plugins run in __isolated__ Docker container. This means, aside from the `/drone` volume, plugins do not share any resources with your build container.

# Official Plugins

Official plugins are found in the [drone-plugins](https://github.com/drone-plugins) GitHub organization. These plugins are vetted by the community and are considered trusted. Requirements of official plugins include:

* Plugin is useful to the broader community
* Plugin is documented
* Plugin is written in Go, Node or Python
* Plugin has manifest file
* Plugin uses Apache2 license
* Plugin uses `alpine` base image (unless technical limitations prohibit)

# Custom Plugins

Plugins can be easily forked or created from scratch. Use the `image` attribute to provide Drone with your custom plugin's fully qualified image name:

```yaml
---
deploy:
  heroku:
    image: bradrydzewski/custom-heroku
    app: hello-world
```

Custom plugins also must be white-listed. The plugin white-list is provided to the Drone server using the `PLUGIN_FILTER` environment variable. This variables is a space-separated list of patterns for matching trusted plugins:

```bash
PLUGIN_FILTER=plugins/* bradrydzewski/*
```

# Plugin Docs

Plugin documentation should be stored in a `DOCS.md` file in the root of your plugin's GitHub repository. Please use the [slack](https://github.com/drone-plugins/drone-slack/blob/master/DOCS.md) documentation for reference.

# Plugin Logo

Plugin logos should be stored in a `logo.svg` file in the root of your plugin's GitHub repository. This logo file is used by the Drone plugin website.

# Plugin Manifest

Plugin details should be stored in the `.drone.yml` file. These plugin details are used by the Drone plugin website. Please use the [slack](https://github.com/drone-plugins/drone-slack/blob/master/.drone.yml) yaml file for reference.

# Plugin Tutorials

> This section will link to plugin creation Guides in the [drone-demos](https://github.com/drone-demos) GitHub organization. Such guides still need to be created.
