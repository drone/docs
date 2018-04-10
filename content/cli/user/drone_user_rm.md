+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user rm"
url = "cli-user-rm"

[menu.cli]
  weight = 4
  identifier = "cli-user-rm"
  parent = "cli_user"
+++

This subcommand deletes a registered user from the system. Please note this command requires administrative privileges.

```text
{{< cat "content/cli/user/data/drone_user_rm.out.txt" >}}
```

Example usage:

```text
$ drone user rm octocat
```
