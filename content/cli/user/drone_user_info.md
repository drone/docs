+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user info"
url = "cli-user-info"
weight = 2

[menu.cli]
  weight = 41
  identifier = "cli-user-info"
  parent = "user"
+++

This subcommand prints information about the named registered user. Please note this command requires administrative privileges.

```text
{{< cat "content/cli/user/data/drone_user_info.out.txt" >}}
```

Example usage:

```text
$ drone user info octocat

User: octocat
Email: octocat@github.com
```

Format the output using a custom Go template:

```text
$ drone user info --format="{{ .Login }} <{{ .Email }}> octocat"

octocat <octocat@github.com>
```
