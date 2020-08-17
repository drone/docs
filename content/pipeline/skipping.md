---
date: 2000-01-01T00:00:00+00:00
title: Skipping
author: bradrydzewski
weight: 1001

description: |
  Control whether a pipeline is skipped or executed.
---

You can skip pipeline execution for an individual commit by adding the skip directive to your commit message. The skip directive is ignored for tags and promotion events, and when manually triggering pipelines.

Example commit message with skip directive:

```
git commit -m "update readme [CI SKIP]"
git push origin master
```

If you need to skip pipelines or individual pipeline steps based on commit or hook metadata, such as branch or hook event, please see the [Triggers]({{< relref "triggers.md" >}}) and [Conditions]({{< relref "conditions.md" >}}) sections of the documentation.