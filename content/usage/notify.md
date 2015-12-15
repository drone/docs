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

Drone does not have any builtin notification capabilities. Drone instead offers a large library of plugins for sending [Slack](/plugins/slack/), [Hipchat](/plugins/hipchat/), [Email](/plugins/email/) notifications and more. Please see the [plugin documentation](/plugins/) for a complete list.

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

# Templates

Many of the Drone notifications support [handlebars](http://handlebarsjs.com/) templates for customizing notification messages. This example demonstrates a custom hipchat notification message:

```yaml
---
notify:
  hipchat:
    template: |
      {{ repo.full_name }} finished build {{ build.number }}
      with a status of {{ build.status }}
```

Example template that generates a different message based on build status:

```yaml
---
notify:
  hipchat:
    template: |
      {{#success build.status}}
        {{ build.author }} successfully pushed to {{ build.branch}}
      {{else}}
        {{ build.author }} broke the build. Hang your head in shame.
      {{/success}}
```

# Template Data

This is an example data structure passed to the template engine. Any of the below variables may be referenced in your template:

```js
{
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "event": "push",
        "commit": "7fd1a60",
        "branch": "master",
        "message": "Update README",
        "author": "octocat",
        "author_email": "octocat@github.com",
        "link_url": "https://github.com/octocat/Hello-World/commit/7fd1a60"
    },
    "repo": {
        "owner": "octocat",
        "name": "hello-world",
        "full_name": "octocat/hello-world",
        "link_url": "https://github.com/octocat/hello-world",
        "clone_url": "https://github.com/octocat/hello-world.git"
    },
    "system": {
        "link_url": "https://drone.mycompany.com"
    }
}
```

# Template Builtins

Drone includes a number of builtin handlebars functions to help format your messages. This is an example builtin function that converts a string to all uppercase characters:

```handlebars
{{uppercase build.status}}
```

Converts a string to all lowercase characters:

```handlebars
{{uppercase build.author}}
```

Truncates a string to N characters:

```handlebars
{{ truncate build.commit 8 }}
```

Calculates a duration and returns a human readable string:

```handlebars
{{ duration build.started_at build.finished_at }}
```

Converts a timestamp to a human readable string:

```handlebars
finished at {{ datetime build.finished_at "3:04PM" "UTC" }}
```

Returns true if the build is successful:

```handlebars
{{#success build.status}}
  It works
{{/success}}
```

Returns true if the build failed:

```handlebars
{{#failure build.status}}
  Something is busted
{{/failure}}
```
