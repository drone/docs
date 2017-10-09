---
title: Manage Registry Credentials
url: manage-registry-credentials

menu:
  usage:
    weight: 1
    identifier: manage_registry_credentials
    parent: usage_secrets
---

Drone provides the ability to store registry credentials. These credentials can be used to pull private pipeline images defined in your Yaml configuration file.

{{% alert info %}}
These credentials are never exposed to your pipeline, which means they cannot be used to push, and are safe to use with pull requests, for example. Pushing to a registry still require setting credentials for the appropriate plugin.
{{% /alert %}}

# Image Caching Behavior

{{% alert error %}}
Pull private images with caution.
{{% /alert %}}

All images (including private images) are pulled and cached by the Docker daemon.
Neither Docker nor Drone restrict the use of cached images.
An image already in the local cache can be used by any pipeline.

It is possible for one repository to have credentials and pull a private image that is cached by Docker, and used by another repository that does not have registry credentials configured.

Keep this in mind when running Drone and jobs in a shared or public environment.

# Configuration

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
  --repository <repository> \
  --hostname <image> \
  --username <name> \
  --password <value>
```

Example command:

```diff
drone registry add \
  --repository octocat/hello-world \
  --hostname gcr.io \
  --username <name> \
  --password <value>
```

Example commands loads the password from file:

```diff
drone registry add \
  --repository octocat/hello-world \
  --hostname gcr.io \
  --username _json_key \
  --password @/absolute/path/to/keyfile.json
```

Please note that in the above examples the `--repository` flag should be set to your version control repository name (e.g. your github repository name).

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
