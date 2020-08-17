---
date: 2000-01-01T00:00:00+00:00
title: drone repo repair
author: bradrydzewski
weight: 5
aliases:
- /cli-repository-repair
- /cli/user/drone-repo-repair
toc: false
---

This subcommand re-creates webhooks for your repository in your version control system (e.g GitHub). This can be used if you accidentally delete your webhooks.

```
NAME:
   drone repo repair - repair repository webhooks

USAGE:
   drone repo repair [arguments...]
```

Example usage:

```
$ drone repo repair octocat/hello-world
```