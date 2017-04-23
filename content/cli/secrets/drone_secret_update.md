+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret update"
url = "cli-secret-update"

[menu.cli]
  weight = 11
  identifier = "cli-secret-update"
  parent = "secret"
+++

This subcommand updates a secret in your repository secret store. Please note this command requires administrative privilege to the repository.

```text
{{< cat "content/cli/secrets/data/drone_secret_update.out.txt" >}}
```

Example usage:

```text
$ drone secret update \
  -repository octocat/hello-world \
  -name docker_username \
  -value ...
```
