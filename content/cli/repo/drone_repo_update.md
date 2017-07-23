+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo update"
url = "cli-repository-update"

[menu.cli]
  weight = 4
  identifier = "cli-repository-update"
  parent = "cli_repo"
+++

This subcommand updates a named repository. Please note this command requires write access to the repository.


```text
{{< cat "content/cli/repo/data/drone_repo_update.out.txt" >}}
```

Example command updates the trusted flag:

```text
$ drone repo update octocat/hello-world --trusted=true
```

Example command updates the gated flag:

```text
$ drone repo update octocat/hello-world --gated=true
```

Example command updates the timeout value:

```text
$ drone repo update octocat/hello-world --timeout=90m
```

Example command updates the drone.yml file path

```text
$ drone repo update octocat/hello-world --config=.github/.drone.yml
```
