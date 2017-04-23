+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user ls"
url = "cli-user-ls"
weight = 4

[menu.cli]
  weight = 43
  identifier = "cli-user-ls"
  parent = "user"
+++

This subcommand prints a list of all registered users. Please note this command requires administrative privileges.

```text
{{< cat "content/cli/user/data/drone_user_ls.out.txt" >}}
```

Example usage:

```text
$ drone user ls

jonhsmith
janedoe
octocat
```

Format the output using a custom Go template:

```text
$ drone user ls --format="{{ .Login }} <{{ .Email }}>"

jonhsmith <jonh.smith@github.com>
janedoe   <jane.doe@github.com>
octocat   <octocat@github.com>
```
