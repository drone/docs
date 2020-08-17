---
date: 2000-01-01T00:00:00+00:00
title: drone repo update
author: bradrydzewski
weight: 5
aliases:
- /cli-repository-update
- /cli/user/drone-repo-update
toc: false
---

This subcommand updates a named repository. Please note this command requires write access to the repository.

```
NAME:
   drone repo update - update a repository

USAGE:
   drone repo update [command options] [arguments...]

OPTIONS:
   --trusted              repository is trusted
   --protected            repository is protected
   --timeout value        repository timeout (default: 0s)
   --visibility value     repository visibility
   --config value         repository configuration path (e.g. .drone.yml)
   --build-counter value  repository starting build number (default: 0)
   --unsafe               validate updating the build-counter is unsafe
```   

Example command updates the trusted flag:

```
$ drone repo update octocat/hello-world --trusted=true
```

Example command updates the protected flag:

```
$ drone repo update octocat/hello-world --protected=true
```

Example command updates the timeout value:

```
$ drone repo update octocat/hello-world --timeout=90m
```

Example command updates the drone.yml file path

```
$ drone repo update octocat/hello-world --config=.github/.drone.yml
```

Example command updates the current build number

```
$ drone repo update octocat/hello-world --build-counter=10 --unsafe
```
