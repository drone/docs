+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret ls"
url = "cli-secret-ls"

[menu.cli]
  weight = 14
  identifier = "cli-secret-ls"
  parent = "cli_secret"
+++

This subcommand returns a list of secrets for the named repository. Please note this command requires administrative privilege to the repository.

```text
{{< cat "content/cli/secrets/data/drone_secret_ls.out.txt" >}}
```

Example usage:

```text
$ drone repo ls -repository octocat/hello-world

docker_username
Events: push, tag, deployment
Images: plugins/docker

docker_password
Events: push, tag, deployment
Images: plugins/docker
```
