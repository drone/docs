+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Decline"
url = "api-build-decline"

[menu.api]
  weight = 8
  identifier = "api-build-decline"
  parent = "api_build"
+++

Declines a blocked build.
Please note this api requires write access to the repository,
and the request parameter `{build}` is not the build id but the build number.

```text
POST /api/repos/{owner}/{repo}/builds/{build}/decline
```
