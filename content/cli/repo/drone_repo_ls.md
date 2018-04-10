+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo ls"
url = "cli-repository-list"

[menu.cli]
  weight = 4
  identifier = "cli-repository-list"
  parent = "cli_repo"
+++

This subcommand returns a full list of user repositories.

```text
{{< cat "content/cli/repo/data/drone_repo_ls.out.txt" >}}
```

Example usage:

```text
$ drone repo ls

octocat/hello-world
octocat/fork-spoon
github/gitignore
github/hubot
github/hub
```

Filter the output by organization:

```text
$ drone repo ls --org=github

github/gitignore
github/hubot
github/hub
```

Format the output using a custom Go template:

```text
$ drone repo ls --format="{{ .Link }}"

https://github.com/octocat/hello-world
https://github.com/octocat/fork-spoon
https://github.com/github/gitignore
https://github.com/github/hubot
https://github.com/github/hub
```
