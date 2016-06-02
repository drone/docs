+++
date = "2015-12-05T16:00:21-08:00"
title = "drone secret"
weight = 1
toc = true

[menu.main]
	parent="cli"
+++

# Overview

Drone provides the ability to store sensitive information such as passwords, tokens and keys in a central store so that you don't need to include them in your Yaml file. This section of the documentation provides usage information for the secret subcommand.

# Secret Add

Register the password with your repository for a specific images or plugins:

```
$ drone secret add \
    --image=s3 \
    octocat/hello-world AWS_ACCESS_KEY_ID AKIAIOSFODNN7EXAMPLE
```

Register the password for multiple images or plugins:

```
$ drone secret add \
    --image=golang \
    --image=node:* \
    --image=s3 \
    octocat/hello-world AWS_ACCESS_KEY_ID AKIAIOSFODNN7EXAMPLE
```

Register the password for pull requests:

```
$ drone secret add \
    --image=docker \
    --event pull_request \
    --event push \
    --event tag \
    --event deployment \
    octocat/hello-world AWS_ACCESS_KEY_ID AKIAIOSFODNN7EXAMPLE
```

# Secret Files

In some cases you may need to upload an SSH key or token that is stored in a file. You can use the `@` notation, similar to curl, to provide the path to the secret file:

```
$ drone secret add \
    --image ssh \
    octocat/hello-world SSH_KEY @/root/.ssh/id_rsa
```

# Secret Remove

Remove the password from your repository:

```
$ drone secret remove octocat/hello-world DOCKER_PASSWORD
```
