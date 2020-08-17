---
date: 2000-01-01T00:00:00+00:00
title: drone repo ls
author: bradrydzewski
weight: 5
aliases:
- /cli-repository-list
- /cli/user/drone-repo-ls
toc: false
---

This subcommand returns a full list of user repositories.

```
NAME:
   drone repo ls - list all repos

USAGE:
   drone repo ls [command options] [arguments...]

OPTIONS:
   --format value  format output
   --org value     filter by organization
```

Example usage:

```
$ drone repo ls

octocat/hello-world
octocat/fork-spoon
github/gitignore
github/hubot
github/hub
```

Filter the output by organization:

```
$ drone repo ls --org=github

github/gitignore
github/hubot
github/hub
```

Format the output using a custom Go template:

```
$ drone repo ls --format="{{ .Link }}"

https://github.com/octocat/hello-world
https://github.com/octocat/fork-spoon
https://github.com/github/gitignore
https://github.com/github/hubot
https://github.com/github/hub
```