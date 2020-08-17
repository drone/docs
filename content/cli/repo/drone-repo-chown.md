---
date: 2000-01-01T00:00:00+00:00
title: drone repo chown
author: bradrydzewski
weight: 5
aliases:
- /cli-repository-chown
- /cli/user/drone-repo-chown
toc: false
---

This subcommand lets a user assume ownership of a named repository. Please note this command requires administrative access to the repository.

```
NAME:
   drone repo chown - assume ownership of a repository

USAGE:
   drone repo chown [arguments...]
```

The below command assumes ownership to the named repository. The individual executing this command becomes the owner.

```
$ drone repo chown octocat/hello-world
```
