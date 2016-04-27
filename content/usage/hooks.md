+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Hooks"
weight = 5
menu = "usage"
toc = true
break = true
+++

# Overview

When you activate your repository, hooks are automatically registered with your version control system (ie GitHub). These hooks are used to automatically trigger builds for specific repository events. This section of the documents describes available hook events and configurations.

# Hook Events

Drone triggers builds for the following event types:

EVENT             | DESC
------------------|-------------------------------------------------------------
`push`            | triggered when code is pushed to your repository
`tag`             | triggered when a tag is pushed to your repository
`pull_request`    | triggered when a pull request is opened or synchronized
`deployment`      | triggered when a deployment is created

 The tag and deployment events are disabled by default and need to be enabled in your Drone repository settings. Please note support for event types may vary based on version control hosting provider.

# Skip Events

You can enable or disable hooks for event types using the toggles in your repository settings. The deployment and tag hooks are disabled and skipped by default, and need to be manually enabled.


# Skip Branches

You can skip specific branches using the `branches` section in your `.drone.yml` file. When a hook is received the branch is compared to the below list. If the branch matches the hook is processed, otherwise ignored.

```
branches:
  - master
  - feature/*
```

You can exclude specific branches or patterns:

```
branches:
  exclude:
    - master
    - develop
```

You can combine include and exclude:

```
branches:
  include: feature/*
  exclude:
    - master
    - develop
```

# Skip Commits

You can skip individual hooks by including `[CI SKIP]` in the commit message.
