+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Build & Test"
weight = 7
menu = "usage"
toc = true
+++

# Overview

# Skip Commits

Instruct Drone to skip builds by adding `[ci skip]` to your commit message.

# Skip Branches

Instruct Drone to skip branches by including a branch white-list in your `.drone.yml`

```
branches:
  - master
  - develop
```
