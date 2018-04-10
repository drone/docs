+++
date = "2000-01-01T00:00:00+00:02"
title = "drone user add"
url = "cli-user-add"

[menu.cli]
  weight = 1
  identifier = "cli-user-add"
  parent = "cli_user"
+++

This subcommand registers a new user with the system. Please note this command requires administrative privileges.

```text
{{< cat "content/cli/user/data/drone_user_add.out.txt" >}}
```

Example usage, adds a user by username:

```text
$ drone user add octocat
```
