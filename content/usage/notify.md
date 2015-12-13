+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Notifications"
weight = 9
menu = "usage"
toc = true
+++

# Overview

Notifications are triggered at the end of your entire build. If you are running a matrix build the notification is sent after all matrix jobs complete. Please note notifications are executed after your build has terminated and will not write any output to your build logs.

Example Slack notification using the Slack plugin:

```
notify:
  slack:
    webhook_url: https://hooks.slack.com/services/...
    channel: dev
    username: drone
```

# Conditions

Use the `when` block to limit notification step execution:

```
notify:
  slack:
    channel: foo
    when:
      success: false
      failure: false
      change: true
```

Execute a notification step if the branch is `master`:

```
notify:
  slack:
    channel: foo
    when:
      branch: master
```

Execute a notification step if the branch is `master` or `develop`:

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
