+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Approve"
url = "api-build-approve"

[menu.api]
  weight = 7
  identifier = "api-build-approve"
  parent = "api_build"
+++

Approves a blocked build.
Please note this api requires write access to the repository,
and the request parameter `{build}` is not the build id but the build number.

```text
POST /api/repos/{owner}/{repo}/builds/{build}/approve
```
