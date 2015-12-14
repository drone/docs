+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Notifications"
weight = 9
menu = "usage"
toc = true
+++

# Overview

Notifications are triggered at the end of your entire build. If you are running a matrix build the notification is sent after all matrix jobs complete.

Example Slack notification using the Slack plugin:

```yaml
---
notify:
  slack:
    webhook_url: https://hooks.slack.com/services/...
    channel: dev
    username: drone
```

Please note notifications are executed after a build is already terminated and will not write any output to your build logs. For troubleshooting failed notifications please contact your system administrator to research errors in the Drone server logs.

# Plugins

Drone does not have any builtin notification capabilities. Drone instead offers a large library of plugins for sending [Slack](#), [Hipchat](#), [Email](#) notifications and more. Please see the [plugin documentation](#) for a complete list.

# Conditions

Use the `when` block to limit notification step execution:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      success: false
      failure: false
      change: true
```

Execute a notification step if the branch is `master`:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      branch: master
```

Execute a notification step if the branch is `master` or `develop`:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      branch: [master, develop]
```

Execute a notification step if the branch is _not_ `develop`:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      branch: "!develop"
```

Execute a notification step if the branch is starts with `prefix/*`:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      branch: prefix/*
```

Execute a notification step when the commit is a `pull_request`, `push` or `tag`:

```yaml
---
notify:
  slack:
    channel: foo
    when:
      event: pull_request
```

Define the same notification step multiple times, using different configuration based on branch:

```yaml
---
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
