+++
date = "2000-01-01T00:00:00+00:02"
title = "drone build info"
url = "cli-build-info"

[menu.cli]
  weight = 3
  identifier = "cli-build-info"
  parent = "cli_build"
+++

Show the specified repository build. Please note this command requires read access to the repository and the request parameter `{build}` is not the build id but the build number.

```text
{{< cat "content/cli/build/data/drone_build_info.out.txt" >}}
```
