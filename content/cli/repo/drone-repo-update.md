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
   drone repo update [command options] <repo/name>

OPTIONS:
   --trusted                    repository is trusted
   --protected                  repository is protected
   --timeout value              repository timeout (default: 0s)
   --visibility value           repository visibility
   --ignore-forks               ignore forks
   --ignore-pull-requests       ignore pull requests
   --auto-cancel-pull-requests  automatically cancel pending pull request builds
   --auto-cancel-pushes         automatically cancel pending push builds
   --auto-cancel-running        automatically cancel running builds if newer commit pushed
   --config value               repository configuration path (e.g. .drone.yml)
   --build-counter value        repository starting build number (default: 0)
   --unsafe                     validate updating the build-counter is unsafe
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

Example command updates the auto cancellation flag:

```
$ drone repo update octocat/hello-world --auto-cancel-pushes=true
$ drone repo update octocat/hello-world --auto-cancel-pull-requests=true
```
