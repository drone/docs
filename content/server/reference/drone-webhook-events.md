---
date: 2000-01-01T00:00:00+00:00
title: DRONE_WEBHOOK_EVENTS
author: bradrydzewski
---

Optional string value provides a comma-separated list of events and actions that trigger webhooks. If unset all events and actions trigger webhooks.

Limit by event type:

```
DRONE_WEBHOOK_EVENTS=user,build
```

Limit by event type and action:

```
DRONE_WEBHOOK_EVENTS=user:created,repo
```

List of available events:

```
user:created
user:updated
user:deleted
repo:enabled
repo:disabled
repo:updated
build:created
build:updated
```