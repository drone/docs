+++
date = "2000-01-01T00:00:00+00:02"
title = "drone build approve"
url = "cli-build-approve"

[menu.cli]
  weight = 1
  identifier = "cli-build-approve"
  parent = "cli_build"
+++

Approves a blocked build. Please note this command requires write access to the repository, and the request parameter `{build}` is not the build id but the build number.

```text
{{< cat "content/cli/build/data/drone_build_approve.out.txt" >}}
```
