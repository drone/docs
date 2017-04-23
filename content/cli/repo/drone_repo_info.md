+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo info"
url = "cli-repository-info"

[menu.cli]
  weight = 5
  identifier = "cli-repository-info"
  parent = "repo"
+++

This subcommand prints the named repository details.

```text
{{< cat "content/cli/repo/data/drone_repo_info.out.txt" >}}
```

Example usage:

```text
$ drone repo info octocat/hello-world

Owner: octocat
Repo: hello-world
Type: git
Private: false
Remote: https://github.com/octocat/hello-world.git
```

Format the output using a custom Go template:

```text
$ drone repo info --format="{{ .Link }}" octocat/hello-world

https://github.com/octocat/hello-world.git
```
