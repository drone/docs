+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Stop"
description = "Endpoint to stop a build"
+++

Stop the specified build.
Please note this api requires administrative privileges and the request parameter `{build}` is not the build id but the build number.

```
DELETE /api/repos/{owner}/{repo}/builds/{build}
```
