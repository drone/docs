+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Notifications"
weight = 9
menu = "usage"
toc = true
+++

# Overview

# Conditions

Use the `when` block to limit notification step execution. This first example executes a notification step when the branch is `master`:

```
notify:
  slack:
    channel: foo
    when:
      branch: master
```

Execute a notification step only the branch is `master` or `develop`:

```
notify:
  slack:
    channel: foo
    when:
      branch: [master, develop]
```

Execute a notification step if the branch is _not_ `develop`:

```
notify:
  slack:
    channel: foo
    when:
      branch: "!develop"
```

Execute a notification step if the branch is starts with `prefix/*`:

```
notify:
  slack:
    channel: foo
    when:
      branch: prefix/*
```

Execute a notification step when the commit is a `pull_request`, `push` or `tag`:

```
notify:
  slack:
    channel: foo
    when:
      event: pull_request
```

Define the same notification step multiple times, using different configuration based on branch:

```
notify:
  slack:
    channel: foo
    when:
      branch: develop

  slack:
    channel: bar
    when:
      branch: master
```
