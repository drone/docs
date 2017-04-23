+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Registry Credentials"
url = "release-0.6.0-registry-credentials"
+++

Registry credentials for downloading private pipeline images are now managed separately from secrets, and are no longer specified in the yaml configuration file.

Example legacy configuration:

```diff
pipeline:
  build:
    image: bradrydzewski/private-image
-   auth_config:
-     username: bradrydzewski
-     password: password
    commands:
      - go build
      - go test
```

Registry credentials are adding to drone using the command line utility:

```text
drone registry add \
  --repository=<repo> \
  --hostname=docker.io \
  --username=<username> \
  --password=<password>
```

Example command:

```text
drone registry add \
  --repository=octocat/hello-world \
  --hostname=docker.io \
  --username=bradrydzewski \
  --password=password
```

Drone matches the registry hostname to each image in your yaml. If the hostnames match the registry credentials are used to authenticate to your registry and pull the image. Note that registry credentials are used by the Drone agent and are never exposed to your build containers.

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
