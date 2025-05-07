---
date: 2000-01-01T00:00:00+00:00
title: drone build create
author: maxlever
weight: 1
aliases:
- /cli-build-create/
---

This subcommand triggers a new build.

```
NAME:
   drone build create - create a build

USAGE:
   drone build create [command options] <repo/name>

OPTIONS:
   --commit value           source commit
   --branch value           source branch
   --param value, -p value  custom parameters to be injected into the job environment. Format: KEY=value
   --format value           format output (default: "Number: {{ .Number }}\nStatus: {{ .Status }}\nEvent: {{ .Event }}\nCommit: {{ .After }}\nBranch: {{ .Target }}\nRef: {{ .Ref }}\nAuthor: {{ .Author }} {{ if .AuthorEmail }}<{{.AuthorEmail}}>{{ end }}\nMessage: {{ .Message }}\n")
```

Example usage:

```
$ drone build create octocat/hello-world --branch=main --param=COSTUME=pirate-octocat
```
