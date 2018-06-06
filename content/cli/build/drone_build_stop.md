+++
date = "2000-01-01T00:00:00+00:02"
title = "drone build stop"
url = "cli-build-stop"

[menu.cli]
  weight = 9
  identifier = "cli-build-stop"
  parent = "cli_build"
+++

Stop the specified build. Please note this command requires administrative privileges and the request parameter `{build}` is not the build id but the build number.

```text
{{< cat "content/cli/build/data/drone_build_stop.out.txt" >}}
```
