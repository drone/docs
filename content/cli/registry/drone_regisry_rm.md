+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry rm"
url = "cli-registry-rm"

[menu.cli]
  weight = 23
  identifier = "cli-registry-rm"
  parent = "cli_registry"
+++

This subcommand deletes a named registry credentials. Please note this command requires authentication and administrative privilege to the repository.

```text
{{< cat "content/cli/registry/data/drone_registry_rm.out.txt" >}}
```

Example usage:

```text
$ drone registry rm -repository octocat/hello-world
```

Example usage with custom docker registry:

```text
$ drone registry rm \
  -repository octocat/hello-world \
  -hostname gcr.io
```
