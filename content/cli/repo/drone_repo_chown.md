+++
date = "2000-01-01T00:00:00+00:02"
title = "drone repo chown"
url = "cli-repository-chown"

[menu.cli]
  weight = 2
  identifier = "cli-repository-chown"
  parent = "cli_repo"
+++

This subcommand lets a user assume ownership of a named repository. Please note this command requires administrative access to the repository.

```text
{{< cat "content/cli/repo/data/drone_repo_chown.out.txt" >}}
```

Below command assumes ownership to the named repository. Note that the individual executing this command becomes the owner.

```text
$ drone repo chown octocat/hello-world
```
