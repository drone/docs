+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret info"
url = "cli-secret-info"

[menu.cli]
  weight = 2
  identifier = "cli-secret-info"
  parent = "cli_secret"
+++

This subcommand prints the named secret details. Please note this command requires administrative privilege to the repository.

```text
{{< cat "content/cli/secrets/data/drone_secret_info.out.txt" >}}
```

Example command:

```text
$ drone secret info \
  -repository octocat/hello-world \
  -name docker_password

docker_password
Events: push, tag, deployment
Images: plugins/docker
```
