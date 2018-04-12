+++
date = "2000-01-01T00:00:00+00:02"
title = "drone registry add"
url = "cli-registry-add"

[menu.cli]
  weight = 1
  identifier = "cli-registry-add"
  parent = "cli_registry"
+++

This subcommand adds registry credentials to your repository. Please note this command requires authentication and administrative privilege to the repository.

```text
{{< cat "content/cli/registry/data/drone_registry_add.out.txt" >}}
```

Example usage:

```text
$ drone registry add \
  -repository octocat/hello-world \
  -username octocat \
  -password password
```

Example loads the repository password from file:

```text
$ drone registry add \
  -repository octocat/hello-world \
  -hostname gcr.io \
  -username octocat \
  -password @/path/to/token.js
```
