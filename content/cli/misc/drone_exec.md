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

Running a full matrix build is not supported at this time, however you can execute a single matrix axis by passing environment variables to the `drone exec` command:

```text
$ GO_VERSION=1.4 drone exec
```
