---
date: 2000-01-01T00:00:00+00:00
title: drone user ls
author: bradrydzewski
weight: 5
aliases:
- /cli-user-ls
- /cli/user/drone-user-ls
toc: false
---

This subcommand prints a list of all registered users. Please note this command requires administrative privileges.

```
NAME:
   drone user ls - list all users

USAGE:
   drone user ls [command options] [arguments...]

OPTIONS:
   --format value  format output (default: "{{ .Login }}")
```   

Example usage:

```
$ drone user ls

jonhsmith
janedoe
octocat
```

Format the output using a custom Go template:

```
$ drone user ls --format="{{ .Login }} <{{ .Email }}>"

jonhsmith <jonh.smith@github.com>
janedoe   <jane.doe@github.com>
octocat   <octocat@github.com>
```