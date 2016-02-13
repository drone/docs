+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Google Container Registry"
description = "Build and publish images to Google Container Registry"
user = "drone-plugins"
repo = "drone-gcr"
image = "plugins/drone-gcr"
tags = ["docker", "image", "container", "google"]
categories = "publish"
draft = false
date = 2016-02-13T08:59:33Z
menu = ""
weight = 1

+++

Use this plugin to build and push Docker images to the Google Container Registry (GCR). Please read the GCR [documentation](https://cloud.google.com/container-registry/) before you begin. You will need to generate a [JSON token](https://developers.google.com/console/help/new/#serviceaccounts) to authenticate to the registry and push images.

The following parameters are used to configure this plugin:

* `registry` - authenticates to this registry (defaults to `gcr.io`)
* `token` - json token
* `repo` - repository name for the image
* `tag` - repository tag for the image (defaults to `latest`)
* `storage_driver` - use `aufs`, `devicemapper`, `btrfs` or `overlay` driver

Sample configuration:

```yaml
publish:
  gcr:
    repo: foo/bar
    token: >
      {
        "private_key_id": "...",
        "private_key": "...",
        "client_email": "...",
        "client_id": "...",
        "type": "..."
      }
```

Sample configuration using multiple tags:

```
publish:
  gcr:
    repo: foo/bar
    tag:
      - latest
      - "1.0.1"
      - "1.0"
    token: >
      {
        "private_key_id": "...",
        "private_key": "...",
        "client_email": "...",
        "client_id": "...",
        "type": "..."
      }
```

## JSON Token

When setting your token in the `.drone.yml` file you must use [folded block scalars](http://www.yaml.org/spec/1.2/spec.html#id2796251) to avoid parsing errors:

```
publish:
  gcr:
    token: >
      {
        "private_key_id": "...",
        "private_key": "...",
        "client_email": "...",
        "client_id": "...",
        "type": "..."
      }
```

When injecting secrets you must also use a folded block scalar:

```
publish:
  gcr:
    token: >
      $$GOOGLE_KEY
```

When defining secrets in a `.drone.sec` file you must also use folded block scalars. Please note that you should either place your JSON string on a single line:

```
environment:
  GOOGLE_KEY: >
    { "private_key_id": "...", "private_key": "..." ... }

```

Or you must ensure all lines have the exact same indentation:

```
environment:
  GOOGLE_KEY: >
    {
    "private_key_id": "...",
    "private_key": "...",
    "client_email": "...",
    "client_id": "...",
    "type": "..."
    }

```

## Troubleshooting

For detailed output you can set the `DOCKER_LAUNCH_DEBUG` environment variable in your plugin configuration. This starts Docker with verbose logging enabled.

```
publish:
  gcr:
    environment:
      - DOCKER_LAUNCH_DEBUG=true
```

## Known Issues

There are known issues when attempting to run this plugin on CentOS, RedHat, and Linux installations that do not have a supported storage driver installed. You can check by running `docker info | grep 'Storage Driver:'` on your host machine. If the storage driver is not `aufs` or `overlay` you will need to re-configure your host machine.

This error occurs when trying to use the default `aufs` storage Driver but aufs is not installed:

```
level=fatal msg="Error starting daemon: error initializing graphdriver: driver not supported
```

This error occurs when trying to use the `overlay` storage Driver but overlay is not installed:

```
level=error msg="'overlay' not found as a supported filesystem on this host.
Please ensure kernel is new enough and has overlay support loaded." 
level=fatal msg="Error starting daemon: error initializing graphdriver: driver not supported"
```

This error occurs when using CentOS or RedHat which default to the `devicemapper` storage driver:

```
level=error msg="There are no more loopback devices available." 
level=fatal msg="Error starting daemon: error initializing graphdriver: loopback mounting failed" 
Cannot connect to the Docker daemon. Is 'docker -d' running on this host?
```

The above issue can be resolved by setting `storage_driver: vfs` in the `.drone.yml` file. This may work, but will have very poor performance as discussed [here](https://github.com/rancher/docker-from-scratch/issues/20).

