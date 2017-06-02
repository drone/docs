---
date: 2017-04-15T14:39:04+02:00
title: "Error: image not found"
url: error-image-not-found
---

This error message comes from the Docker daemon when an image isn’t specified, or when the specified image cannot be downloaded. To further diagnose this error we recommend looking at the Docker daemon logs, since this is not a Drone issue.

```nohighlight
Trying to pull foo/bar from https://registry-1.docker.io v2
Trying to pull foo/bar from https://index.docker.io v1
[registry] Calling GET https://index.docker.io/v1/repositories/foo/bar/images
Not continuing with pull after error: Error: image foo/bar not found
```

Example failure when image does not exist, or has a typo:

```diff
pipeline:
  build:
-   image: go
+   image: golang
```

Example failure when image tag does not exist, or has a typo:

```diff
pipeline:
  build:
-   image: golang:latdst
+   image: golang:latest
```

Example failure when image is not declared:

```diff
pipeline:
  build:
-                    
+   image: golang:1.5
```

Example failure when image is private, but no registry credentials provided:


```diff
pipeline:
  build:
    image: index.company.com/octocat/hello-image
+   auth_config:
+     username: octocat
+     password: password
```
