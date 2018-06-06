+++
date = "2000-01-01T00:00:00+00:02"
title = "drone build start"
url = "cli-build-start"

[menu.cli]
  weight = 8
  identifier = "cli-build-start"
  parent = "cli_build"
+++

Restart the specified build. Please note this command requires read and write access to the repository and the request parameter `{build}` is not the build id but the build number.

```text
{{< cat "content/cli/build/data/drone_build_start.out.txt" >}}
```
