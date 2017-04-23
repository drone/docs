+++
date = "2000-01-01T00:00:00+00:02"
title = "drone secret add"
url = "cli-secret-add"

[menu.cli]
  weight = 10
  identifier = "cli-secret-add"
  parent = "cli_secret"
+++

This subcommand adds a secret to your repository secret store. Please note this command requires administrative privilege to the repository.

```text
{{< cat "content/cli/secrets/data/drone_secret_add.out.txt" >}}
```

Example usage:

```text
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
  -value ...
```

Example usage limits the secret to a specific image:

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
+ -image plugins/docker \
  -value ...
```

Example usage limits the secret to specific hook events:

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name docker_password \
  -image plugins/docker \
+ -event push \
+ -event pull_request \
+ -event tag \
+ -event deployment \
  -value ...
```

Example usage adds the secret from a file:

```diff
$ drone secret add \
  -repository octocat/hello-world \
  -name ssh_key \
+ -value @/root/.ssh/id_rsa
```
