+++
date = "2015-12-05T16:00:21-08:00"
title = "Notifications"
weight = 28
toc = true
draft = false

[menu.main]
	parent="usage"
+++

# Overview

Drone does not have built-in notification capabilities. Instead notifications are delegated to plugins. This section of the documentation provides a high level overview of notifications and should be used in conjunction with the respective plugin documentation.

Example Slack notification using the Slack plugin:

```yaml
pipeline:
  build:
    image: node
    commands:
      - npm install
      - npm run build
      - npm run tests

  slack:
    channel: dev
    username: drone
```

# Templates

Many notification plugins support [handlebars](http://handlebarsjs.com/) templates for customizing notification messages. This example demonstrates a custom slack message:

```yaml
pipeline:
  slack:
    channel: dev
    username: drone
    template: |
      {{ repo.name }} finished build {{ build.number }}
      with a status of {{ build.status }}
```

Example template that generates a different message based on build status:

```yaml
pipeline:
  slack:
    channel: dev
    username: drone
    template: |
      {{#success build.status}}
        {{ build.author }} successfully pushed to {{ build.branch}}
      {{else}}
        {{ build.author }} broke the build. Hang your head in shame.
      {{/success}}
```

# Template Variables

Variables available to your notiication templates containers:

NAME                         | DESC
-----------------------------|--------------------------------------------------
`repo.owner`                 | repository owner
`repo.name`                  | repository name
`repo.link`                  | repository link
`repo.avatar`                | repository avatar
`repo.branch`                | repository default branch (master)
`repo.clone`                 | repository clone url
`commit.sha`                 | commit sha
`commit.ref`                 | commit ref
`commit.branch`              | commit branch
`commit.link`                | commit link in remote
`commit.message`             | commit message
`commit.author.name`         | commit author username
`commit.author.email`        | commit author email address
`commit.author.avatar`       | commit author avatar
`build.number`               | build number
`build.event`                | build event (push, pull_request, tag)
`build.status`               | build status (success, failure)
`build.link`                 | build result link
`build.created`              | build created unix timestamp
`build.started`              | build started unix timestamp
`build.finished`             | build finished unix timestamp

# Template Builtins

Drone includes a number of builtin handlebars functions to help format your messages. This is an example builtin function that converts a string to all uppercase characters:

```handlebars
{{uppercase build.status}}
```

Converts a string to all lowercase characters:

```handlebars
{{lowercase build.author.name}}
```

Truncates a string to N characters:

```handlebars
{{ truncate build.commit 8 }}
```

Calculates a duration and returns a human readable string:

```handlebars
{{ duration build.started build.finished }}
```

Converts a timestamp to a human readable string:

```handlebars
finished at {{ datetime build.finished "3:04PM" "UTC" }}
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
