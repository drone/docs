+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Stop"
url = "api-build-stop"

[menu.api]
  weight = 2
  identifier = "api-build-stop"
  parent = "api_build"
+++

Stop the specified build.
Please note this api requires administrative privileges and the request parameter `{build}` is not the build id but the build number.

```text
DELETE /api/repos/{owner}/{repo}/builds/{build}
```
