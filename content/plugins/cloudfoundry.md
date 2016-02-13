+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Cloud Foundry"
description = "Deploy or update a project on Cloud Foundry"
user = "drone-plugins"
repo = "drone-cloudfoundry"
image = "plugins/drone-cloudfoundry"
tags = ["paas"]
categories = "deploy"
draft = false
date = 2016-02-13T08:59:58Z
menu = ""
weight = 1

+++

Use this plugin for deplying an application to Cloud Foundry. You can override
the default configuration with the following parameters:

## Example

```yaml
deploy:
  cloudfoundry:
    api: api.run.pivotal.io
    user: $$USER
    password: $$PASSWORD
    org: $$USER
    space: production
    when:
      branch: master
```

### Options

| Option | Description | Example |
| ------ | ----------- | ---- |
| `api` (**required**) | Target API | `api.run.pivotal.io` |
| `org` (**required**) | Target Org | `xyz-org` |
| `space` (**required**) | Target Space | `development` |
| `user` (**required**) | Auth username | `john@doe.com` |
| `password` (**required**) | Auth password | `supersecure` |
| `name` | Override application name | `app-canary` |
| `manifest` | Path to manifest | `test.manifest.yml` |
| `path` | App path | `build/assets` |
| `command` | Startup command | `start-script.sh` |
| `buildpack` | Custom buildpack | `https://....` |
| `disk` | Disk limit | `256M` |
| `memory` | Memory limit | `1G` |
| `instances` | Number of instances | `4` |
| `hostname` | Hostname | `my-subdomain` |
| `random-route` | Create a random route for this app | `false` |
| `domain` | Domain | `example.com` |
| `no-route` | Do not map a route to this app and remove routes from previous pushes of this app. | `false` |
| `no-start` | Do not start an app after pushing | `false` |
| `no-hostname` | Map the root domain to this app | `false` |
| `no-manifest` | Ignore manifest file | `false` |
| `skip-ssl-validation` | For login, do not validate ssl (for self-signed cert) | `false` |

