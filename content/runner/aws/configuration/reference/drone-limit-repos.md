---
date: 2000-01-01T00:00:00+00:00
title: DRONE_LIMIT_REPOS
author: tphoney
weight: 1
---

Optional comma-separated string value. Configures the runner to only process matching repositories. This provides an extra layer of security and can stop untrusted repositories from executing pipelines with this runner.

```bash
DRONE_LIMIT_REPOS=octocat/hello-world,spaceghost/*
```
