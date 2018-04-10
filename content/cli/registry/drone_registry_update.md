+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry update"
url = "cli-registry-update"

[menu.cli]
  weight = 5
  identifier = "cli-registry-update"
  parent = "cli_registry"
+++

This subcommand updates the named registry credentials. Please note this command requires authentication and administrative privilege to the repository.

```text
{{< cat "content/cli/registry/data/drone_registry_update.out.txt" >}}
```

Example usage:

```text
$ drone registry update \
  -repository octocat/hello-world \
  -hostname docker.io \
  -password ...
```
