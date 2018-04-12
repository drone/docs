+++
date = "2000-01-01T00:00:00+00:02"
title = "drone exec"
url = "cli-exec"

[menu.cli]
  weight = 2
  identifier = "cli-exec"
  parent = "cli_misc"
+++

This subcommand executes a local build.

```text
{{< cat "content/cli/misc/data/drone_exec.out.txt" >}}
```

If your pipeline uses secrets, these can be injected locally simply by passing environment variables:

```text
$ MY_SECRET=mybigsecret drone exec
```
