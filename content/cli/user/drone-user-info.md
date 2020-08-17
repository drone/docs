---
date: 2000-01-01T00:00:00+00:00
title: drone user info
author: bradrydzewski
weight: 5
aliases:
- /cli-user-info
- /cli/user/drone-user-info
toc: false
---

This subcommand prints information about the named registered user. Please note this command requires administrative privileges.

```
NAME:
   drone user info - show user details

USAGE:
   drone user info [command options] [arguments...]

OPTIONS:
   --format value  format output (default: "User: {{ .Login }}\nEmail: {{ .Email }}")
``` 

Example usage:

```
$ drone user info octocat

User: octocat
Email: octocat@github.com
```

Format the output using a custom Go template:

```
$ drone user info --format="{{ .Login }} <{{ .Email }}> octocat"

octocat <octocat@github.com>
```