+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret rm"
url = "cli-secret-rm"

[menu.cli]
  weight = 13
  identifier = "cli-secret-rm"
  parent = "cli_secret"
+++

This subcommand deletes a named repository secret. Please note this command requires administrative privilege to the repository.

```text
{{< cat "content/cli/secrets/data/drone_secret_rm.out.txt" >}}
```

Example usage:


```text
$ drone secret rm -repository octocat/hello-world -name docker_password
```
