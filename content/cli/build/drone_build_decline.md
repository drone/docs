+++
date = "2000-01-01T00:00:00+00:02"
title = "drone build decline"
url = "cli-build-decline"

[menu.cli]
  weight = 2
  identifier = "cli-build-decline"
  parent = "cli_build"
+++

Declines a blocked build. Please note this command requires write access to the repository, and the request parameter `{build}` is not the build id but the build number.

```text
{{< cat "content/cli/build/data/drone_build_decline.out.txt" >}}
```
