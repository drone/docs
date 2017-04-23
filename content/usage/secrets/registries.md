+++
date = "2017-04-15T14:39:04+02:00"
title = "Manage Registry Credentials"
url = "manage-registry-credentials"

[menu.usage]
  weight = 1
  parent = "usage_secrets"
  identifier = "manage-registry-credentials"
+++


Drone provides the ability to store registry credentials. These credentials can be used to download private pipeline images defined in your Yaml configuration file.

Example configuration using a private image:

```diff
pipeline:
  build:
+   image: gcr.io/custom/golang
    commands:
      - go build
      - go test
```

Registries are added to your repository using the command line utility:

```diff
drone registry add \
  --repository <registry> \
  --hostname <image> \
  --username <name> \
  --password <value>
```

Example command:

```diff
drone registry add \
  --repository <registry> \
  --hostname gcr.io \
  --username <name> \
  --password <value>
```

# Matching

Drone matches the registry hostname to each image in your yaml. If the hostnames match, the registry credentials are used to authenticate to your registry and pull the image. Note that registry credentials are used by the Drone agent and are never exposed to your build containers.

Example registry hostnames:

* Image `gcr.io/foo/bar` has hostname `gcr.io`
* Image `foo/bar` has hostname `docker.io`
* Image `qux.com:8000/foo/bar` has hostname `qux.com:8000`

Example registry hostname matching logic:

* Hostname `gcr.io` matches image `gcr.io/foo/bar`
* Hostname `docker.io` matches `golang`
* Hostname `docker.io` matches `library/golang`
* Hostname `docker.io` matches `bradyrydzewski/golang`
* Hostname `docker.io` matches `bradyrydzewski/golang:latest`
