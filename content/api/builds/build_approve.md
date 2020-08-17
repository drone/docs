+++
date = "2000-01-01T00:00:00+00:02"
title = "Build Approve"

description = "Endpoint to approve a build"
+++

Approves a blocked build.
Please note this api requires write access to the repository,
and the request parameter `{build}` is not the build id but the build number.

```
POST /api/repos/{owner}/{repo}/builds/{build}/approve
```
